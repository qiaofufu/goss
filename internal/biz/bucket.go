package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Bucket struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerID     int64  `json:"owner_id"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type BucketRepo interface {
	CreateBucket(ctx context.Context, name string, description string, ownerID int64) (*Bucket, error)
	UpdateBucket(ctx context.Context, bucketID int64, newName string, newDescription string) error
	GetBucketByID(ctx context.Context, id int64) (*Bucket, error)
	ListBucket(ctx context.Context, ownerID int64, limit, offset int64) ([]*Bucket, error)
	DeleteBucket(ctx context.Context, id int64) error
}

type BucketUsecase struct {
	repo BucketRepo
	log  *log.Helper
}

func NewBucketUsecase(repo BucketRepo, logger log.Logger) *BucketUsecase {
	return &BucketUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *BucketUsecase) CreateBucket(ctx context.Context, name string, description string, ownerID int64) (*Bucket, error) {
	uc.log.WithContext(ctx).Infof("CreateBucket: %v", name)
	return uc.repo.CreateBucket(ctx, name, description, ownerID)
}

func (uc *BucketUsecase) UpdateBucket(ctx context.Context, bucketID int64, newName string, newDescription string) error {
	uc.log.WithContext(ctx).Infof("UpdateBucket: %v", bucketID)
	return uc.repo.UpdateBucket(ctx, bucketID, newName, newDescription)
}

func (uc *BucketUsecase) GetBucket(ctx context.Context, id int64) (*Bucket, error) {
	uc.log.WithContext(ctx).Infof("GetBucket: %v", id)
	return uc.repo.GetBucketByID(ctx, id)
}

func (uc *BucketUsecase) ListBucket(ctx context.Context, ownerID int64, limit, offset int64) ([]*Bucket, error) {
	uc.log.WithContext(ctx).Infof("ListBucket: %v", ownerID)
	return uc.repo.ListBucket(ctx, ownerID, limit, offset)
}

func (uc *BucketUsecase) DeleteBucket(ctx context.Context, id int64) error {
	uc.log.WithContext(ctx).Infof("DeleteBucket: %v", id)
	return uc.repo.DeleteBucket(ctx, id)
}
