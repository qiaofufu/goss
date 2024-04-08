package service

import (
	"context"
	pb "master/api/bucket/v1"
	"master/internal/biz"
)

type BucketService struct {
	pb.UnimplementedBucketServer
	uc *biz.BucketUsecase
}

func NewBucketService(uc *biz.BucketUsecase) *BucketService {
	return &BucketService{uc: uc}
}

func (s *BucketService) CreateBucket(ctx context.Context, req *pb.CreateBucketRequest) (*pb.CreateBucketReply, error) {
	bucket, err := s.uc.CreateBucket(ctx, req.Name, req.Description, req.OwnerID)
	if err != nil {
		return &pb.CreateBucketReply{
			Status: 1,
			Msg:    err.Error(),
		}, err
	}
	return &pb.CreateBucketReply{
		Id:          bucket.ID,
		Name:        bucket.Name,
		Description: bucket.Description,
		OwnerID:     bucket.OwnerID,
		CreatedAt:   bucket.CreatedAt,
		UpdatedAt:   bucket.UpdatedAt,
	}, nil
}
func (s *BucketService) UpdateBucket(ctx context.Context, req *pb.UpdateBucketRequest) (*pb.UpdateBucketReply, error) {
	err := s.uc.UpdateBucket(ctx, req.Id, req.Name, req.Description)
	if err != nil {
		return &pb.UpdateBucketReply{
			Status: 1,
			Msg:    err.Error(),
		}, err
	}
	return &pb.UpdateBucketReply{
		Status: 0,
		Msg:    "success",
	}, nil
}
func (s *BucketService) DeleteBucket(ctx context.Context, req *pb.DeleteBucketRequest) (*pb.DeleteBucketReply, error) {
	err := s.uc.DeleteBucket(ctx, req.Id)
	if err != nil {
		return &pb.DeleteBucketReply{
			Status: 1,
			Msg:    err.Error(),
		}, err
	}
	return &pb.DeleteBucketReply{
		Status: 0,
		Msg:    "success",
	}, nil
}
func (s *BucketService) GetBucket(ctx context.Context, req *pb.GetBucketRequest) (*pb.GetBucketReply, error) {
	bucket, err := s.uc.GetBucket(ctx, req.Id)
	if err != nil {
		return &pb.GetBucketReply{
			Status: 1,
			Msg:    err.Error(),
		}, err
	}
	return &pb.GetBucketReply{
		Id:          bucket.ID,
		Name:        bucket.Name,
		Description: bucket.Description,
		OwnerID:     bucket.OwnerID,
		CreatedAt:   bucket.CreatedAt,
		UpdatedAt:   bucket.UpdatedAt,
	}, nil
}
func (s *BucketService) ListBucket(ctx context.Context, req *pb.ListBucketRequest) (*pb.ListBucketReply, error) {
	buckets, err := s.uc.ListBucket(ctx, req.OwnerID, int64(req.PageSize), int64(req.Page*req.PageSize))
	if err != nil {
		return &pb.ListBucketReply{
			Status: 1,
			Msg:    err.Error(),
		}, err
	}
	var pbBuckets []*pb.BucketInfo
	for _, bucket := range buckets {
		pbBuckets = append(pbBuckets, &pb.BucketInfo{
			Id:          bucket.ID,
			Name:        bucket.Name,
			Description: bucket.Description,
			OwnerID:     bucket.OwnerID,
			CreatedAt:   bucket.CreatedAt,
			UpdatedAt:   bucket.UpdatedAt,
		})
	}
	return &pb.ListBucketReply{
		Buckets: pbBuckets,
	}, nil
}
