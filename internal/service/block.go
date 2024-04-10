package service

import (
	"context"
	"master/internal/biz"

	pb "master/api/block/v1"
)

type BlockService struct {
	pb.UnimplementedBlockServer
	uc *biz.BlockUsecase
}

func NewBlockService(uc *biz.BlockUsecase) *BlockService {
	return &BlockService{uc: uc}
}

func (s *BlockService) CreateBlock(ctx context.Context, req *pb.CreateBlockRequest) (*pb.CreateBlockReply, error) {
	block, err := s.uc.CreateBlock(ctx, &biz.Block{
		ObjectID: req.ObjectId,
		Num:      int(req.Num),
		Size:     req.Size,
		Checksum: int(req.Checksum),
	})
	if err != nil {
		return &pb.CreateBlockReply{
			Status: 1,
			Msg:    err.Error(),
		}, err
	}
	return &pb.CreateBlockReply{
		BlockId: block.ID,
		Status:  0,
		Msg:     "Success",
	}, nil
}

func (s *BlockService) DeleteBlock(ctx context.Context, req *pb.DeleteBlockRequest) (*pb.DeleteBlockReply, error) {
	err := s.uc.DeleteBlock(ctx, req.ObjectId)
	if err != nil {
		return &pb.DeleteBlockReply{
			Status: 1,
			Msg:    err.Error(),
		}, err
	}
	return &pb.DeleteBlockReply{
		Status: 0,
		Msg:    "Success",
	}, nil
}

func (s *BlockService) ListBlock(ctx context.Context, req *pb.ListBlockRequest) (*pb.ListBlockReply, error) {
	blocks, total, err := s.uc.ListBlock(ctx, req.ObjectId)
	if err != nil {
		return &pb.ListBlockReply{
			Status: 1,
			Msg:    err.Error(),
		}, err
	}
	resp := &pb.ListBlockReply{}
	for _, block := range blocks {
		resp.Blocks = append(resp.Blocks, &pb.BlockInfo{
			BlockId:   block.ID,
			ObjectId:  block.ObjectID,
			Num:       int32(block.Num),
			Size:      block.Size,
			Checksum:  int32(block.Checksum),
			CreatedAt: block.CreatedAt,
		})
	}
	resp.Total = total
	return resp, nil
}
