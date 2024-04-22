package biz

import (
	"context"
	"fmt"
)

type Location struct {
	Type    int32 `json:"type"`
	ReferID int64 `json:"refer_id"`
	NodeID  int64 `json:"node_id"`
}

type LocationRepo interface {
	CreateLocation(ctx context.Context, t int32, referID, nodeID int64) error
	DeleteLocation(ctx context.Context, t int32, referID int64) error
	ListLocation(ctx context.Context, t int32, referID int64) ([]*Location, error)
}

type LocationUsecase struct {
	repo LocationRepo
}

func NewLocationUsecase(repo LocationRepo) *LocationUsecase {
	return &LocationUsecase{repo: repo}
}

func (uc *LocationUsecase) CreateLocation(ctx context.Context, t int32, referID, nodeID int64) error {
	err := uc.repo.CreateLocation(ctx, t, referID, nodeID)
	if err != nil {
		return fmt.Errorf("biz: %w", err)
	}
	return nil
}

func (uc *LocationUsecase) DeleteLocation(ctx context.Context, t int32, referID int64) error {
	err := uc.repo.DeleteLocation(ctx, t, referID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *LocationUsecase) ListLocation(ctx context.Context, t int32, referID int64) ([]*Location, error) {
	locations, err := uc.repo.ListLocation(ctx, t, referID)
	if err != nil {
		return nil, err
	}
	return locations, nil
}
