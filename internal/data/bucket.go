package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"master/internal/biz"
	"time"
)

type Bucket struct {
	ID          int64          `gorm:"primary_key; column:bucket_id"`
	Name        string         `gorm:"column:name"`
	Description string         `gorm:"column:description"`
	OwnerID     int64          `gorm:"column:owner_id"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

type bucketRepo struct {
	data  *Data
	log   *log.Helper
	idGen UniqueIDGenerator
}

func NewBucketRepo(data *Data, logger log.Logger, idGen UniqueIDGenerator) biz.BucketRepo {
	return &bucketRepo{
		data:  data,
		log:   log.NewHelper(logger),
		idGen: idGen,
	}
}

func (b *bucketRepo) CreateBucket(ctx context.Context, name string, description string, ownerID int64) (*biz.Bucket, error) {
	curTime := time.Now()
	model := Bucket{
		ID:          b.idGen.GenerateID(),
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
		CreatedAt:   curTime,
		UpdatedAt:   curTime,
	}
	err := b.data.db.Create(&model).Error
	if err != nil {
		return nil, err
	}
	return &biz.Bucket{
		ID:          model.ID,
		Description: model.Description,
		OwnerID:     model.OwnerID,
		CreatedAt:   model.CreatedAt.UnixMilli(),
		UpdatedAt:   model.UpdatedAt.UnixMilli(),
	}, nil
}

func (b *bucketRepo) UpdateBucket(ctx context.Context, bucketID int64, newName string, newDescription string) error {
	err := b.data.db.Model(&Bucket{}).Where("id = ?", bucketID).Updates(map[string]interface{}{
		"name":        newName,
		"description": newDescription,
		"updated_at":  time.Now().UnixMilli(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (b *bucketRepo) GetBucketByID(ctx context.Context, id int64) (*biz.Bucket, error) {
	var model Bucket
	err := b.data.db.Where("id = ?", id).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &biz.Bucket{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		OwnerID:     model.OwnerID,
		CreatedAt:   model.CreatedAt.UnixMilli(),
		UpdatedAt:   model.UpdatedAt.UnixMilli(),
	}, nil
}

func (b *bucketRepo) ListBucket(ctx context.Context, ownerID int64, limit, offset int64) ([]*biz.Bucket, error) {
	var models []Bucket
	err := b.data.db.Where("owner_id = ?", ownerID).Limit(int(limit)).Offset(int(offset)).Find(&models).Error
	if err != nil {
		return nil, err
	}
	var buckets []*biz.Bucket
	for _, model := range models {
		buckets = append(buckets, &biz.Bucket{
			ID:          model.ID,
			Name:        model.Name,
			Description: model.Description,
			OwnerID:     model.OwnerID,
			CreatedAt:   model.CreatedAt.UnixMilli(),
			UpdatedAt:   model.UpdatedAt.UnixMilli(),
		})
	}
	return buckets, nil
}

func (b *bucketRepo) DeleteBucket(ctx context.Context, id int64) error {
	return b.data.db.Where("id = ?", id).Delete(&Bucket{}).Error
}
