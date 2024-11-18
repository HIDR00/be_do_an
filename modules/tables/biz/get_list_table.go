package biz

import (
	"context"
	"first-app/common"
	"first-app/modules/tables/models"
)

type GetListTableStorage interface {
	GetListTable(ctx context.Context, paging *common.Paging, moreKey ...string) ([]models.Table, error)
}

type GetListTableBiz struct {
	store GetListTableStorage
}

func NewGetListTableBiz(store GetListTableStorage) *GetListTableBiz {
	return &GetListTableBiz{store: store}
}

func (biz *GetListTableBiz) GetNewListTable(ctx context.Context, paging *common.Paging) ([]models.Table, error) {

	data, err := biz.store.GetListTable(ctx,paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}
