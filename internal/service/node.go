package service

import (
	"context"
	"master/internal/biz"

	pb "master/api/node/v1"
)

type NodeService struct {
	pb.UnimplementedNodeServer
	uc *biz.NodeUsecase
}

func NewNodeService(uc *biz.NodeUsecase) *NodeService {
	return &NodeService{uc: uc}
}

func (s *NodeService) CreateNode(ctx context.Context, req *pb.CreateNodeRequest) (*pb.CreateNodeReply, error) {
	node, err := s.uc.CreateNode(ctx, req.GetAddr())
	if err != nil {
		return &pb.CreateNodeReply{
			Status: 1,
			Msg:    err.Error(),
		}, nil
	}

	return &pb.CreateNodeReply{
		Id:     node.ID,
		Addr:   node.Addr,
		Status: 0,
		Msg:    "Success",
	}, nil
}
func (s *NodeService) DeleteNode(ctx context.Context, req *pb.DeleteNodeRequest) (*pb.DeleteNodeReply, error) {
	err := s.uc.DeleteNode(ctx, req.GetId())
	if err != nil {
		return &pb.DeleteNodeReply{
			Status: 1,
			Msg:    err.Error(),
		}, nil
	}

	return &pb.DeleteNodeReply{
		Status: 0,
		Msg:    "Success",
	}, nil
}
func (s *NodeService) GetNode(ctx context.Context, req *pb.GetNodeRequest) (*pb.GetNodeReply, error) {
	nodes, err := s.uc.GetNode(ctx, req.GetKey(), int(req.GetReplicas()))
	if err != nil {
		return &pb.GetNodeReply{
			Status: 1,
			Msg:    err.Error(),
		}, nil
	}
	list := make([]*pb.NodeInfo, 0)
	for _, node := range nodes {
		list = append(list, &pb.NodeInfo{
			Id:            node.ID,
			Addr:          node.Addr,
			Status:        int32(node.Status),
			LastHeartbeat: node.LastHeartbeat,
		})
	}
	return &pb.GetNodeReply{
		Nodes: list,
	}, nil
}
func (s *NodeService) ListNode(ctx context.Context, req *pb.ListNodeRequest) (*pb.ListNodeReply, error) {
	nodes, total, err := s.uc.ListNode(ctx, int64(req.GetSize()), int64(req.GetPage()*req.GetSize()))
	if err != nil {
		return &pb.ListNodeReply{
			Status: 1,
			Msg:    err.Error(),
		}, nil
	}
	list := make([]*pb.NodeInfo, 0)
	for _, node := range nodes {
		list = append(list, &pb.NodeInfo{
			Id:            node.ID,
			Addr:          node.Addr,
			Status:        int32(node.Status),
			LastHeartbeat: node.LastHeartbeat,
		})
	}
	return &pb.ListNodeReply{
		Nodes:  list,
		Total:  total,
		Status: 0,
		Msg:    "Success",
	}, nil
}
func (s *NodeService) KeepAlive(ctx context.Context, req *pb.KeepAliveRequest) (*pb.KeepAliveReply, error) {
	return &pb.KeepAliveReply{}, nil
}
