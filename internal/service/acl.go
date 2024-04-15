package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"master/internal/biz"

	pb "master/api/acl/v1"
)

type AclService struct {
	pb.UnimplementedAclServer
	uc  *biz.AclUsecase
	log *log.Helper
}

func NewAclService(uc *biz.AclUsecase, logger log.Logger) *AclService {
	return &AclService{uc: uc, log: log.NewHelper(log.With(logger, "module", "service/acl"))}
}

func (s *AclService) CreateAcl(ctx context.Context, req *pb.CreateAclRequest) (*pb.CreateAclReply, error) {
	acl, err := s.uc.CreateAcl(ctx, req.ResourceId, req.ResourceType, req.TargetId, req.TargetType, req.Permission)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAclReply{
		Id:         acl.Id,
		ResourceId: acl.ResourceId,
		TargetId:   acl.TargetId,
		TargetType: acl.TargetType,
		Permission: acl.Permission,
		CreatedAt:  acl.CreatedAt,
		UpdatedAt:  acl.UpdatedAt,
		Status:     0,
		Msg:        "Success",
	}, nil
}
func (s *AclService) UpdateAcl(ctx context.Context, req *pb.UpdateAclRequest) (*pb.UpdateAclReply, error) {
	_, err := s.uc.UpdateAcl(ctx, req.Id, req.Permission)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAclReply{
		Status: 0,
		Msg:    "Success",
	}, nil
}
func (s *AclService) DeleteAcl(ctx context.Context, req *pb.DeleteAclRequest) (*pb.DeleteAclReply, error) {
	err := s.uc.DeleteAcl(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteAclReply{
		Status: 0,
		Msg:    "Success",
	}, nil
}
func (s *AclService) ListAcl(ctx context.Context, req *pb.ListAclRequest) (*pb.ListAclReply, error) {
	acls, cnt, err := s.uc.ListAcl(ctx, req.ResourceId, req.ResourceType, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	var list []*pb.AclInfo
	for _, acl := range acls {
		list = append(list, &pb.AclInfo{
			Id:         acl.Id,
			ResourceId: acl.ResourceId,
			TargetId:   acl.TargetId,
			TargetType: acl.TargetType,
			Permission: acl.Permission,
			CreatedAt:  acl.CreatedAt,
			UpdatedAt:  acl.UpdatedAt,
		})
	}
	return &pb.ListAclReply{
		AclList: list,
		Total:   cnt,
		Status:  0,
		Msg:     "Success",
	}, nil
}
