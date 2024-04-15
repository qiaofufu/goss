package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Acl struct {
	Id           int64
	ResourceId   int64
	ResourceType int32
	TargetId     int64
	TargetType   int32
	Permission   int32
	CreatedAt    int64
	UpdatedAt    int64
}

type AclRepo interface {
	CreateAcl(context.Context, *Acl) (*Acl, error)
	UpdateAcl(ctx context.Context, aclId int64, permission int32) (*Acl, error)
	GetAcl(ctx context.Context, resourceId int64, resourceType int32, targetId int64, targetType int32) (*Acl, error)
	DeleteAcl(ctx context.Context, aclId int64) error
	ListAcl(ctx context.Context, resourceId int64, resourceType int32, limit, offset int32) ([]*Acl, int64, error)
}

type AclUsecase struct {
	repo AclRepo
	log  *log.Helper
}

func NewAclUsecase(repo AclRepo, logger log.Logger) *AclUsecase {
	return &AclUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AclUsecase) CreateAcl(ctx context.Context, resourceId int64, resourceType int32, targetId int64, targetType int32, permission int32) (*Acl, error) {
	acl, err := uc.repo.CreateAcl(ctx, &Acl{
		ResourceId: resourceId,
		TargetId:   targetId,
		TargetType: targetType,
		Permission: permission,
	})
	if err != nil {
		return nil, err
	}
	return acl, nil
}

func (uc *AclUsecase) UpdateAcl(ctx context.Context, id int64, permission int32) (*Acl, error) {
	acl, err := uc.repo.UpdateAcl(ctx, id, permission)
	if err != nil {
		return nil, err
	}
	return acl, nil
}

func (uc *AclUsecase) CheckPermission(ctx context.Context, resourceId int64, resourceType int32, targetId int64, targetType int32, permission int32) (bool, error) {
	acl, err := uc.repo.GetAcl(ctx, resourceId, resourceType, targetId, targetType)
	if err != nil {
		return false, err
	}
	if acl == nil {
		return false, nil
	}
	if acl.Permission&permission == permission {
		return true, nil
	}
	return false, nil
}

func (uc *AclUsecase) DeleteAcl(ctx context.Context, aclId int64) error {
	err := uc.repo.DeleteAcl(ctx, aclId)
	if err != nil {
		return err
	}

	return nil
}

func (uc *AclUsecase) ListAcl(ctx context.Context, resourceId int64, resourceType int32, page, pageSize int32) ([]*Acl, int64, error) {
	if pageSize <= 0 {
		pageSize = 10
	}
	acls, cnt, err := uc.repo.ListAcl(ctx, resourceId, resourceType, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}
	return acls, cnt, nil
}
