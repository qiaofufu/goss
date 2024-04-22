package data

import (
	"context"
	"fmt"
	"master/internal/biz"
)

type Location struct {
	Type    int32 `gorm:"column:type"`
	ReferID int64 `gorm:"column:refer_id"`
	NodeID  int64 `gorm:"column:node_id"`
}

func (Location) TableName() string {
	return "location"
}

type locationRepo struct {
	data *Data
}

func NewLocationRepo(data *Data) biz.LocationRepo {
	return &locationRepo{data: data}
}

func (r *locationRepo) CreateLocation(ctx context.Context, t int32, referID, nodeID int64) error {
	err := r.data.db.Create(Location{
		Type:    t,
		ReferID: referID,
		NodeID:  nodeID,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *locationRepo) DeleteLocation(ctx context.Context, t int32, referID int64) error {
	err := r.data.db.Where("type = ? and refer_id = ?", t, referID).Delete(Location{}).Error
	if err != nil {
		return fmt.Errorf("data: %w", err)
	}
	return nil
}

func (r *locationRepo) ListLocation(ctx context.Context, t int32, referID int64) ([]*biz.Location, error) {
	var locations []*Location
	err := r.data.db.Where("type = ? and refer_id = ?", t, referID).Find(&locations).Error
	if err != nil {
		return nil, fmt.Errorf("data: %w", err)
	}
	var res []*biz.Location
	for _, location := range locations {
		res = append(res, &biz.Location{
			Type:    location.Type,
			ReferID: location.ReferID,
			NodeID:  location.NodeID,
		})
	}
	return res, nil
}
