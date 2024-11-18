package biz

import (
	"context"
	"first-app/modules/tables/models"
)

type UpdateTableByIdStorage interface {
	GetTableById(ctx context.Context, cond map[string]interface{}) (*models.Table, error)
	UpdateTableById(ctx context.Context, cond map[string]interface{}, dataUpdate *models.Table) error
}

type UpdateTableByIdBiz struct {
	store UpdateTableByIdStorage
}

func NewUpdateTableByIdBiz(store UpdateTableByIdStorage) *UpdateTableByIdBiz {
	return &UpdateTableByIdBiz{store: store}
}

func (biz *UpdateTableByIdBiz) UpdateNewTableById(ctx context.Context, id int, dataUpdate *models.Table) error {

	data, err := biz.store.GetTableById(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if data.Status == 10 {
		return models.ErrTableDeleted
	}

	if err := biz.store.UpdateTableById(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}
	return nil
}
