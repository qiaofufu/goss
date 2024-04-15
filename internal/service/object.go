package service

import (
	"context"
	"master/internal/biz"

	pb "master/api/object/v1"
)

type ObjectService struct {
	pb.UnimplementedObjectServer
	uc *biz.ObjectUsecase
}

func NewObjectService(uc *biz.ObjectUsecase) *ObjectService {
	return &ObjectService{uc: uc}
}

func (s *ObjectService) CreateObject(ctx context.Context, req *pb.CreateObjectRequest) (*pb.CreateObjectReply, error) {

	object, err := s.uc.CreateObject(ctx, req.Name, req.Size, req.BucketId, req.Type, int(req.Replicas))
	if err != nil {
		return &pb.CreateObjectReply{
			Status: 1,
			Msg:    err.Error(),
		}, nil
	}
	return &pb.CreateObjectReply{
		ObjectId: object.ID,
	}, nil
}
func (s *ObjectService) UpdateObject(ctx context.Context, req *pb.UpdateObjectRequest) (*pb.UpdateObjectReply, error) {
	err := s.uc.UpdateObject(ctx, req.ObjectId, req.Name, req.BucketId, req.Type)
	if err != nil {
		return &pb.UpdateObjectReply{
			Status: 1,
			Msg:    err.Error(),
		}, nil
	}
	return &pb.UpdateObjectReply{
		Status: 0,
		Msg:    "Success",
	}, nil
}
func (s *ObjectService) DeleteObject(ctx context.Context, req *pb.DeleteObjectRequest) (*pb.DeleteObjectReply, error) {
	err := s.uc.DeleteObject(ctx, req.ObjectId)
	if err != nil {
		return &pb.DeleteObjectReply{
			Status: 1,
			Msg:    err.Error(),
		}, nil
	}
	return &pb.DeleteObjectReply{
		Status: 0,
		Msg:    "Success",
	}, nil
}
func (s *ObjectService) GetObject(ctx context.Context, req *pb.GetObjectRequest) (*pb.GetObjectReply, error) {
	object, err := s.uc.GetObject(ctx, req.ObjectId)
	if err != nil {
		return &pb.GetObjectReply{
			Status: 1,
			Msg:    err.Error(),
		}, nil
	}
	return &pb.GetObjectReply{
		Object: &pb.ObjectInfo{
			ObjectId: object.ID,
			Name:     object.Name,
			Size:     object.Size,
			BucketId: object.BucketID,
			Type:     object.Type,
			Replicas: int32(object.Replicas),
		},
	}, nil
}
func (s *ObjectService) ListObject(ctx context.Context, req *pb.ListObjectRequest) (*pb.ListObjectReply, error) {
	objects, total, err := s.uc.ListObject(ctx, req.BucketId, int(req.PageSize), int(req.Page*req.PageSize))
	if err != nil {
		return &pb.ListObjectReply{
			Status: 1,
			Msg:    err.Error(),
		}, nil
	}
	var objectList []*pb.ObjectInfo
	for _, object := range objects {
		objectList = append(objectList, &pb.ObjectInfo{
			ObjectId:  object.ID,
			Name:      object.Name,
			Size:      object.Size,
			BucketId:  object.BucketID,
			Type:      object.Type,
			Replicas:  int32(object.Replicas),
			UpdatedAt: object.UpdatedAt,
			CreatedAt: object.CreatedAt,
		})
	}
	return &pb.ListObjectReply{
		Objects: objectList,
		Total:   total,
		Status:  0,
		Msg:     "Success",
	}, nil
}
