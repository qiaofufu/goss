package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"sync"
	"time"
)

type Node struct {
	ID            int64
	Addr          string
	Status        int
	LastHeartbeat int64
}

type NodeRepo interface {
	CreateNode(addr string) (*Node, error)
	DeleteNode(id int64) error
	GetNode(id int64) (*Node, error)
	GetNodeByAddr(addr string) (*Node, error)
	ListNode(limit int64, offset int64) ([]*Node, error)
	Count() int64
	ExistNodeByAddr(addr string) (bool, error)
}

type NodeUsecase struct {
	repo           NodeRepo
	log            *log.Helper
	consistentHash *ConsistentHash
	nodeMap        map[int64]*Node
	mu             sync.RWMutex
	errChan        chan error
}

func NewNodeUsecase(repo NodeRepo, logger log.Logger, hash *ConsistentHash) *NodeUsecase {
	usecase := &NodeUsecase{repo: repo, log: log.NewHelper(logger), consistentHash: hash, errChan: make(chan error), nodeMap: make(map[int64]*Node)}
	usecase.adjustNodeMap()
	return usecase
}

func (u *NodeUsecase) CreateNode(ctx context.Context, addr string) (*Node, error) {
	u.log.WithContext(ctx).Infof("create node: %s", addr)
	if addr == "" {
		return nil, fmt.Errorf("addr is empty")
	}
	if exist, err := u.repo.ExistNodeByAddr(addr); err != nil {
		return nil, fmt.Errorf("check node exist error: %v", err)
	} else if exist {
		return nil, fmt.Errorf("node already exists")
	}

	node, err := u.repo.CreateNode(addr)
	if err != nil {
		return nil, err
	}
	u.consistentHash.Add(node.ID)
	u.mu.Lock()
	u.nodeMap[node.ID] = node
	u.mu.Unlock()

	return node, nil
}

func (u *NodeUsecase) DeleteNode(ctx context.Context, id int64) error {
	u.log.WithContext(ctx).Infof("delete node: %d", id)
	err := u.repo.DeleteNode(id)
	if err != nil {
		return err
	}
	u.consistentHash.Remove(id)
	u.mu.Lock()
	delete(u.nodeMap, id)
	u.mu.Unlock()
	return nil
}

func (u *NodeUsecase) GetNode(ctx context.Context, key int64, replicas int) ([]*Node, error) {
	u.log.WithContext(ctx).Infof("get node: %d", key)

	if replicas == 0 {
		return nil, fmt.Errorf("replicas must be greater than 0")
	}
	nodeIDs, err := u.consistentHash.GetN(key, replicas)
	if err != nil {
		return nil, err
	}
	u.mu.RLock()
	defer u.mu.RUnlock()

	nodes := make([]*Node, 0, len(nodeIDs))
	for _, id := range nodeIDs {
		var tmpnode Node
		mnode, ok := u.nodeMap[id]
		if !ok {
			dnode, err := u.repo.GetNode(id)
			if err != nil {
				return nil, err
			}
			tmpnode = *dnode
		} else {
			tmpnode = *mnode
		}
		nodes = append(nodes, &tmpnode)
	}

	return nodes, nil
}

func (u *NodeUsecase) ListNode(ctx context.Context, limit, offset int64) ([]*Node, int64, error) {
	u.log.WithContext(ctx).Infof("list node")
	node, err := u.repo.ListNode(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	total := u.repo.Count()
	return node, total, nil
}

func (u *NodeUsecase) adjustNodeMap() {
	// initial node map
	count := u.repo.Count()
	nodes, err := u.repo.ListNode(count, 0)
	if err != nil {
		u.log.Errorf("adjust node map error: %v", err)
		return
	}
	for _, v := range nodes {
		u.nodeMap[v.ID] = v
		u.consistentHash.Add(v.ID)
	}
	// adjust node map
	go func() {
		time.Sleep(time.Second * 5)
		u.mu.Lock()
		nodes := make([]int64, 0)
		for _, v := range u.nodeMap {
			if time.Now().UnixMilli()-v.LastHeartbeat > time.Second.Milliseconds()*60 {
				nodes = append(nodes, v.ID)
			}
		}
		u.mu.Unlock()

		for _, v := range nodes {
			err := u.DeleteNode(context.Background(), v)
			if err != nil {
				u.errChan <- err
			}
		}
	}()
	// error log
	go func() {
		for err := range u.errChan {
			u.log.Errorf("adjust node map error: %v", err)
		}
	}()
}
