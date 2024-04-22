package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"master/internal/biz"

	pb "master/api/location/v1"
)

type LocationService struct {
	pb.UnimplementedLocationServer
	uc  *biz.LocationUsecase
	log *log.Helper
}

func NewLocationService(uc *biz.LocationUsecase, logger log.Logger) *LocationService {
	return &LocationService{uc: uc, log: log.NewHelper(logger)}
}

func (s *LocationService) CreateLocation(ctx context.Context, req *pb.CreateLocationRequest) (*pb.CreateLocationReply, error) {
	err := s.uc.CreateLocation(ctx, req.Type, req.ReferID, req.NodeID)
	if err != nil {
		s.log.Errorf("create location error: %v", err)
		return &pb.CreateLocationReply{
			Status: 1,
			Msg:    "create location error",
		}, nil
	}
	return &pb.CreateLocationReply{
		Status: 0,
		Msg:    "success",
	}, nil
}
func (s *LocationService) DeleteLocation(ctx context.Context, req *pb.DeleteLocationRequest) (*pb.DeleteLocationReply, error) {
	err := s.uc.DeleteLocation(ctx, req.Type, req.ReferID)
	if err != nil {
		s.log.Errorf("delete location error: %v", err)
		return &pb.DeleteLocationReply{
			Status: 1,
			Msg:    "delete location error",
		}, nil
	}
	return &pb.DeleteLocationReply{
		Status: 0,
		Msg:    "success",
	}, nil
}
func (s *LocationService) ListLocation(ctx context.Context, req *pb.ListLocationRequest) (*pb.ListLocationReply, error) {
	location, err := s.uc.ListLocation(ctx, req.Type, req.ReferID)
	if err != nil {
		s.log.Errorf("list location error: %v", err)
		return &pb.ListLocationReply{
			Status: 1,
			Msg:    "list location error",
		}, nil
	}
	var locations []*pb.LocationInfo
	for _, v := range location {
		locations = append(locations, &pb.LocationInfo{
			Type:    v.Type,
			ReferID: v.ReferID,
			NodeID:  v.NodeID,
		})
	}
	return &pb.ListLocationReply{
		Status:    0,
		Msg:       "success",
		Locations: locations,
	}, nil
}
