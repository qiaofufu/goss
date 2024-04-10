package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"master/internal/biz"
	"time"
)

type Block struct {
	ID        int64          `gorm:"primaryKey; column:block_id"`
	ObjectID  int64          `gorm:"column:object_id"`
	Size      int64          `gorm:"column:size"`
	Num       int32          `gorm:"column:number"`
	Checksum  int32          `gorm:"column:checksum"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

type blockRepo struct {
	data  *Data
	log   *log.Helper
	idGen UniqueIDGenerator
}

func NewBlockRepo(data *Data, logger log.Logger, idGen UniqueIDGenerator) biz.BlockRepo {
	return &blockRepo{
		data:  data,
		log:   log.NewHelper(logger),
		idGen: idGen,
	}
}

func (r *blockRepo) ListBlock(objectID int64) ([]*biz.Block, int64, error) {
	var blocks []Block
	err := r.data.db.Where("object_id = ?", objectID).Find(&blocks).Error
	if err != nil {
		return nil, 0, err
	}
	var res []*biz.Block
	for _, block := range blocks {
		res = append(res, &biz.Block{
			ID:        block.ID,
			ObjectID:  block.ObjectID,
			Size:      block.Size,
			Num:       int(block.Num),
			Checksum:  int(block.Checksum),
			CreatedAt: block.CreatedAt.UnixMilli(),
		})
	}
	return res, int64(len(res)), nil
}

func (r *blockRepo) CreateBlock(b *biz.Block) (*biz.Block, error) {
	block := &Block{
		ID:        r.idGen.GenerateID(),
		ObjectID:  b.ObjectID,
		Size:      b.Size,
		Num:       int32(b.Num),
		Checksum:  int32(b.Checksum),
		CreatedAt: time.Now(),
	}
	err := r.data.db.Create(block).Error
	if err != nil {
		return nil, err
	}
	return &biz.Block{
		ID:        block.ID,
		ObjectID:  block.ObjectID,
		Size:      block.Size,
		Num:       int(block.Num),
		Checksum:  int(block.Checksum),
		CreatedAt: block.CreatedAt.UnixMilli(),
	}, nil
}

func (r *blockRepo) DeleteBlock(id int64) error {
	err := r.data.db.Model(Block{}).Delete("object_id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}
