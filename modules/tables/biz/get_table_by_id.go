package biz

import (
	"context"
	"first-app/modules/tables/models"
)

type GetTableByIdStorage interface {
	GetTableById(ctx context.Context, cond map[string]interface{}) (*models.Table, error)
}

type GetTableByIdBiz struct {
	store GetTableByIdStorage
}

func NewGetTableByIdBiz(store GetTableByIdStorage) *GetTableByIdBiz {
	return &GetTableByIdBiz{store: store}
}

func (biz *GetTableByIdBiz) GetNewTableById(ctx context.Context, id int) (*models.Table, error) {

	data, err := biz.store.GetTableById(ctx, map[string]interface{}{"id": id, "deleted_at": nil})
	if err != nil {
		return nil, err
	}

	return data, nil
}
