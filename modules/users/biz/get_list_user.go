package biz

import (
	"context"
	"first-app/common"
	"first-app/modules/users/models"
)

type GetListUserStorage interface {
	GetListUser(ctx context.Context, paging *common.Paging, moreKey ...string) ([]models.User, error)
}

type GetListUserBiz struct {
	store GetListUserStorage
}

func NewGetListUserBiz(store GetListUserStorage) *GetListUserBiz {
	return &GetListUserBiz{store: store}
}

func (biz *GetListUserBiz) GetNewListUser(ctx context.Context, paging *common.Paging) ([]models.User, error) {
	data, err := biz.store.GetListUser(ctx, paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}