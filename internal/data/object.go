package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"master/internal/biz"
	"time"
)

type Object struct {
	ID        int64          `gorm:"primaryKey; column:object_id"`
	Name      string         `gorm:"column:name"`
	Size      int64          `gorm:"column:size"`
	BucketID  int64          `gorm:"column:bucket_id"`
	Type      string         `gorm:"column:type"`
	Replicas  int            `gorm:"column:replicas"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

type objectRepo struct {
	data  *Data
	log   *log.Helper
	idGen UniqueIDGenerator
}

func (o *objectRepo) CreateObject(ctx context.Context, name string, size int64, bucketID int64, types string, replicas int) (*biz.Object, error) {
	var cnt int64
	err := o.data.db.Model(&Object{}).Where("name = ?", name).Count(&cnt).Error
	if err != nil {
		return nil, errors.New(1000, err.Error(), "internal error")
	}
	if cnt > 0 {
		return nil, errors.New(1001, "", "object name already exists")
	}
	err = o.data.db.Create(&Object{
		ID:        o.idGen.GenerateID(),
		Name:      name,
		Size:      size,
		BucketID:  bucketID,
		Type:      types,
		Replicas:  replicas,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error
	if err != nil {
		return nil, errors.New(1000, err.Error(), "internal error")
	}

	return &biz.Object{
		ID:        o.idGen.GenerateID(),
		Name:      name,
		Size:      size,
		BucketID:  bucketID,
		Type:      types,
		Replicas:  replicas,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}, nil
}

func (o *objectRepo) GetObject(ctx context.Context, id int64) (*biz.Object, error) {
	object := &Object{}
	err := o.data.db.Where("object_id = ?", id).First(object).Error
	if err != nil {
		return nil, errors.New(1000, err.Error(), "internal error")
	}
	return &biz.Object{
		ID:        object.ID,
		Name:      object.Name,
		Size:      object.Size,
		BucketID:  object.BucketID,
		Type:      object.Type,
		Replicas:  object.Replicas,
		CreatedAt: object.CreatedAt.UnixMilli(),
		UpdatedAt: object.UpdatedAt.UnixMilli(),
	}, nil
}

func (o *objectRepo) ListObject(ctx context.Context, bucketID int64, limit, offset int) ([]*biz.Object, error) {
	objects := make([]*Object, 0)
	err := o.data.db.Where("bucket_id = ?", bucketID).Limit(limit).Offset(offset).Find(&objects).Error
	if err != nil {
		return nil, err
	}
	objectsList := make([]*biz.Object, 0)
	for _, object := range objects {
		objectsList = append(objectsList, &biz.Object{
			ID:        object.ID,
			Name:      object.Name,
			Size:      object.Size,
			BucketID:  object.BucketID,
			Type:      object.Type,
			Replicas:  object.Replicas,
			CreatedAt: object.CreatedAt.UnixMilli(),
			UpdatedAt: object.UpdatedAt.UnixMilli(),
		})
	}
	return objectsList, nil
}

func (o *objectRepo) DeleteObject(ctx context.Context, id int64) error {
	err := o.data.db.Where("object_id = ?", id).Delete(&Object{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *objectRepo) UpdateObject(ctx context.Context, id int64, name string, bucketID int64, types string) error {
	err := o.data.db.Model(&Object{}).Where("object_id = ?", id).Updates(map[string]interface{}{
		"name":       name,
		"bucket_id":  bucketID,
		"type":       types,
		"updated_at": time.Now().Unix(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *objectRepo) Count(ctx context.Context, id int64) (int64, error) {
	var count int64
	err := o.data.db.Model(&Object{}).Where("bucket_id = ?", id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func NewObjectRepo(data *Data, logger log.Logger, idGen UniqueIDGenerator) biz.ObjectRepo {
	return &objectRepo{
		data:  data,
		log:   log.NewHelper(logger),
		idGen: idGen,
	}
}
