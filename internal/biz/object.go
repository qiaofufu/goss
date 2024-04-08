package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Object struct {
	ID       int64
	Name     string
	Size     int64
	Type     string
	BucketID int64
	Replicas int

	CreatedAt int64
	UpdatedAt int64
}

type ObjectRepo interface {
	CreateObject(
		ctx context.Context,
		name string,
		size int64,
		bucketID int64,
		types string,
		replicas int,
	) (*Object, error)
	GetObject(ctx context.Context, id int64) (*Object, error)
	ListObject(ctx context.Context, bucketID int64, limit, offset int) ([]*Object, error)
	DeleteObject(ctx context.Context, id int64) error
	UpdateObject(ctx context.Context, id int64, name string, bucketID int64, types string) error
	Count(ctx context.Context, id int64) (int64, error)
}

type ObjectUsecase struct {
	ObjectRepo ObjectRepo
	log        *log.Helper
}

func NewObjectUsecase(repo ObjectRepo, logger log.Logger) *ObjectUsecase {
	return &ObjectUsecase{ObjectRepo: repo, log: log.NewHelper(logger)}
}

func (uc *ObjectUsecase) CreateObject(
	ctx context.Context,
	name string,
	size int64,
	bucketID int64,
	types string,
	replicas int,
) (*Object, error) {
	uc.log.WithContext(ctx).Infof(
		"CreateObject name: %s, size: %d, bucketID: %d, types: %s, replicas: %d",
		name,
		size,
		bucketID,
		types,
		replicas,
	)
	return uc.ObjectRepo.CreateObject(ctx, name, size, bucketID, types, replicas)
}

func (uc *ObjectUsecase) GetObject(ctx context.Context, id int64) (*Object, error) {
	uc.log.WithContext(ctx).Infof("GetObject id: %d", id)
	return uc.ObjectRepo.GetObject(ctx, id)
}

func (uc *ObjectUsecase) ListObject(ctx context.Context, bucketID int64, limit, offset int) ([]*Object, int64, error) {
	uc.log.WithContext(ctx).Infof("ListObject bucketID: %d, limit: %d, offset: %d", bucketID, limit, offset)
	list, err := uc.ObjectRepo.ListObject(ctx, bucketID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	total, err := uc.ObjectRepo.Count(ctx, bucketID)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (uc *ObjectUsecase) DeleteObject(ctx context.Context, id int64) error {
	uc.log.WithContext(ctx).Infof("DeleteObject id: %d", id)
	return uc.ObjectRepo.DeleteObject(ctx, id)
}

func (uc *ObjectUsecase) UpdateObject(
	ctx context.Context,
	id int64,
	name string,
	bucketID int64,
	types string,
) error {
	uc.log.WithContext(ctx).Infof(
		"UpdateObject id: %d, name: %s, bucketID: %d, types: %s",
		id,
		name,
		bucketID,
		types,
	)
	return uc.ObjectRepo.UpdateObject(ctx, id, name, bucketID, types)
}
