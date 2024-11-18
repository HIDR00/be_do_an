package biz

import (
	"context"
	"first-app/common"
	"first-app/modules/categories/models"
)

type GetListCategoriesStorage interface {
	GetListCategories(ctx context.Context, paging *common.Paging, moreKey ...string) ([]models.Categories, error)
}

type GetListCategoriesBiz struct {
	store GetListCategoriesStorage
}

func NewGetListCategoriesBiz(store GetListCategoriesStorage) *GetListCategoriesBiz {
	return &GetListCategoriesBiz{store: store}
}

func (biz *GetListCategoriesBiz) GetNewListCategories(ctx context.Context,paging *common.Paging) ([]models.Categories, error) {

	data, err := biz.store.GetListCategories(ctx,paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}
