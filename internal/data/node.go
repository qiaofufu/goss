package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"master/internal/biz"
)

type Node struct {
	ID     int64  `gorm:"primary_key; column:node_id"`
	Addr   string `gorm:"unique;not null;column:ip"`
	Status int    `gorm:"not null;column:status"`
}

func NewNodeRepo(data *Data, logger log.Logger, idGen UniqueIDGenerator) biz.NodeRepo {
	return &nodeRepo{
		data:  data,
		log:   log.NewHelper(logger),
		idGen: idGen,
	}
}

type nodeRepo struct {
	data  *Data
	log   *log.Helper
	idGen UniqueIDGenerator
}

func (n *nodeRepo) GetNodeByAddr(addr string) (*biz.Node, error) {
	var node Node
	err := n.data.db.Where("ip = ?", addr).First(node).Error
	if err != nil {
		return nil, err
	}
	return &biz.Node{
		ID:     node.ID,
		Addr:   node.Addr,
		Status: node.Status,
	}, nil
}

func (n *nodeRepo) ExistNodeByAddr(addr string) (bool, error) {
	var count int64
	err := n.data.db.Where("ip = ?", addr).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (n *nodeRepo) CreateNode(addr string) (*biz.Node, error) {
	node := Node{
		ID:     n.idGen.GenerateID(),
		Addr:   addr,
		Status: 0,
	}
	if err := n.data.db.Create(&node).Error; err != nil {
		return nil, err
	}
	return &biz.Node{
		ID:     node.ID,
		Addr:   node.Addr,
		Status: node.Status,
	}, nil
}

func (n *nodeRepo) DeleteNode(id int64) error {
	err := n.data.db.Delete(&Node{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (n *nodeRepo) GetNode(id int64) (*biz.Node, error) {
	var node Node
	err := n.data.db.First(&node, id).Error
	if err != nil {
		return nil, err
	}
	return &biz.Node{
		ID:     node.ID,
		Addr:   node.Addr,
		Status: node.Status,
	}, nil
}

func (n *nodeRepo) ListNode(limit int64, offset int64) ([]*biz.Node, error) {
	var nodes []Node
	err := n.data.db.Limit(int(limit)).Offset(int(offset)).Find(&nodes).Error
	if err != nil {
		return nil, err
	}
	var res []*biz.Node
	for _, node := range nodes {
		res = append(res, &biz.Node{
			ID:     node.ID,
			Addr:   node.Addr,
			Status: node.Status,
		})
	}
	return res, nil
}

func (n *nodeRepo) Count() int64 {
	var count int64
	n.data.db.Model(&Node{}).Count(&count)
	return count
}
