package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Block struct {
	ID        int64
	ObjectID  int64
	Num       int
	CreatedAt int64
	Size      int64
	Checksum  int
}

type BlockRepo interface {
	ListBlock(objectID int64) ([]*Block, int64, error)
	CreateBlock(b *Block) (*Block, error)
	DeleteBlock(id int64) error
}

type BlockUsecase struct {
	repo BlockRepo
	log  *log.Helper
}

func NewBlockUsecase(repo BlockRepo, logger log.Logger) *BlockUsecase {
	return &BlockUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *BlockUsecase) ListBlock(ctx context.Context, objectID int64) ([]*Block, int64, error) {
	return uc.repo.ListBlock(objectID)
}

func (uc *BlockUsecase) CreateBlock(ctx context.Context, b *Block) (*Block, error) {
	return uc.repo.CreateBlock(b)
}

func (uc *BlockUsecase) DeleteBlock(ctx context.Context, objectID int64) error {
	return uc.repo.DeleteBlock(objectID)
}
