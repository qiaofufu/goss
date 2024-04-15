package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"master/internal/biz"
	"time"
)

type Acl struct {
	ID           int64          `gorm:"primary_key;column:acl_id"`
	ResourceId   int64          `gorm:"column:resource_id"`
	ResourceType int32          `gorm:"column:resource_type"`
	TargetId     int64          `gorm:"column:target_id"`
	TargetType   int32          `gorm:"column:target_type"`
	Permission   int32          `gorm:"column:permission"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
}

type aclRepo struct {
	data  *Data
	log   *log.Helper
	idGen UniqueIDGenerator
}

func NewAclRepo(data *Data, logger log.Logger, idGen UniqueIDGenerator) biz.AclRepo {
	return &aclRepo{
		data:  data,
		log:   log.NewHelper(log.With(logger, "module", "data/acl")),
		idGen: idGen,
	}
}

func (a *aclRepo) CreateAcl(ctx context.Context, acl *biz.Acl) (*biz.Acl, error) {
	a.log.WithContext(ctx).Infof("CreateAcl: %v", acl)
	if exist, err := a.Exist(acl.ResourceId, acl.ResourceType, acl.TargetType, acl.TargetId, acl.Permission); exist {
		return nil, fmt.Errorf("acl already exists")
	} else if err != nil {
		return nil, err
	}
	aclModel := Acl{
		ID:           a.idGen.GenerateID(),
		ResourceId:   acl.ResourceId,
		ResourceType: acl.ResourceType,
		TargetId:     acl.TargetId,
		TargetType:   acl.TargetType,
		Permission:   acl.Permission,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err := a.data.db.Create(&aclModel).Error
	if err != nil {
		return nil, err
	}
	return &biz.Acl{
		Id:           aclModel.ID,
		ResourceId:   aclModel.ResourceId,
		ResourceType: aclModel.ResourceType,
		TargetId:     aclModel.TargetId,
		TargetType:   aclModel.TargetType,
		Permission:   aclModel.Permission,
	}, nil
}

func (a *aclRepo) UpdateAcl(ctx context.Context, aclId int64, permission int32) (*biz.Acl, error) {
	a.log.WithContext(ctx).Infof("UpdateAcl: %v", aclId)

	err := a.data.db.Model(&Acl{}).Where("acl_id = ?", aclId).Update("permission", permission).Error
	if err != nil {
		return nil, err
	}
	var aclModel Acl
	err = a.data.db.Model(&Acl{}).Where("acl_id = ?", aclId).First(&aclModel).Error
	if err != nil {
		return nil, err
	}
	return &biz.Acl{
		Id:           aclModel.ID,
		ResourceId:   aclModel.ResourceId,
		ResourceType: aclModel.ResourceType,
		TargetId:     aclModel.TargetId,
		TargetType:   aclModel.TargetType,
		Permission:   aclModel.Permission,
		CreatedAt:    aclModel.CreatedAt.UnixMilli(),
		UpdatedAt:    aclModel.UpdatedAt.UnixMilli(),
	}, nil
}

func (a *aclRepo) GetAcl(ctx context.Context, resourceId int64, resourceType int32, targetId int64, targetType int32) (*biz.Acl, error) {
	a.log.WithContext(ctx).Infof("GetAcl: resourceId: %v, resourceType: %v, targetId: %v, targetType: %v", resourceId, resourceType, targetId, targetType)
	var aclModel Acl
	err := a.data.db.Model(&Acl{}).Where("resource_id = ? AND resource_type = ? AND target_id = ? AND target_type = ?",
		resourceId, resourceType, targetId, targetType).First(&aclModel).Error
	if err != nil {
		return nil, err
	}
	return &biz.Acl{
		Id:           aclModel.ID,
		ResourceId:   aclModel.ResourceId,
		ResourceType: aclModel.ResourceType,
		TargetId:     aclModel.TargetId,
		TargetType:   aclModel.TargetType,
		Permission:   aclModel.Permission,
		CreatedAt:    aclModel.CreatedAt.UnixMilli(),
		UpdatedAt:    aclModel.UpdatedAt.UnixMilli(),
	}, nil
}

func (a *aclRepo) DeleteAcl(ctx context.Context, aclId int64) error {
	a.log.WithContext(ctx).Infof("DeleteAcl: %v", aclId)
	err := a.data.db.Model(&Acl{}).Where("acl_id = ?", aclId).Delete(&Acl{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *aclRepo) Exist(resourceId int64, resourceType, targetType int32, targetId int64, permission int32) (bool, error) {
	var (
		cnt int64
	)
	err := a.data.db.Model(Acl{}).Where("resource_id = ? AND resource_type = ? AND target_id = ? AND target_type = ? AND permission = ?",
		resourceId, resourceType, targetId, targetType, permission).Count(&cnt).Error
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (a *aclRepo) ListAcl(ctx context.Context, resourceId int64, resourceType int32, limit, offset int32) ([]*biz.Acl, int64, error) {
	a.log.WithContext(ctx).Infof("ListAcl: resourceId: %v, limit: %v, offset: %v", resourceId, limit, offset)
	var aclModels []Acl
	err := a.data.db.Model(&Acl{}).Where("resource_id = ? and resource_type = ?", resourceId, resourceType).Limit(int(limit)).Offset(int(offset)).Find(&aclModels).Error
	if err != nil {
		return nil, 0, err
	}

	acls := make([]*biz.Acl, 0)
	for _, aclModel := range aclModels {
		acls = append(acls, &biz.Acl{
			Id:           aclModel.ID,
			ResourceId:   aclModel.ResourceId,
			ResourceType: aclModel.ResourceType,
			TargetId:     aclModel.TargetId,
			TargetType:   aclModel.TargetType,
			Permission:   aclModel.Permission,
			CreatedAt:    aclModel.CreatedAt.UnixMilli(),
			UpdatedAt:    aclModel.UpdatedAt.UnixMilli(),
		})
	}
	var cnt int64
	a.data.db.Model(&Acl{}).Where("resource_id = ? and resource_type = ?", resourceId, resourceType).Count(&cnt)

	return acls, cnt, nil
}
