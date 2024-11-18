package biz

import (
	"context"
	"first-app/modules/tables/models"
)

type CreateTableStorage interface {
	CreateTable(ctx context.Context,data *models.Table) error
}

type CreateTableBiz struct {
	store CreateTableStorage
}

func NewCreateTableBiz(store CreateTableStorage) *CreateTableBiz {
	return &CreateTableBiz{store: store}
}

func (biz *CreateTableBiz) CreateNewTable(ctx context.Context,data *models.Table) error {

	if err := biz.store.CreateTable(ctx,data); err != nil {
		return err
	}

	return nil
}
