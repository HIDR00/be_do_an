package biz

import (
	"context"
	"first-app/common"
	"first-app/modules/menus/models"
)

type GetListMenuStorage interface {
	GetListMenu(ctx context.Context, paging *common.Paging, moreKey ...string) ([]models.Menu, error)
}

type GetListMenuBiz struct {
	store GetListMenuStorage
}

func NewGetListMenuBiz(store GetListMenuStorage) *GetListMenuBiz {
	return &GetListMenuBiz{store: store}
}

func (biz *GetListMenuBiz) GetNewListMenu(ctx context.Context, paging *common.Paging) ([]models.Menu, error) {

	data, err := biz.store.GetListMenu(ctx,paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}
