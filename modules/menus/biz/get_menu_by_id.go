package biz;

import (
	"context"
	"first-app/modules/menus/models"
)

type GetMenuByIdStorage interface {
	GetMenuById(ctx context.Context,cond map[string]interface{}) (*models.Menu,error)
}

type GetMenuByIdBiz struct {
	store GetMenuByIdStorage
}

func NewGetMenuByIdBiz(store GetMenuByIdStorage) *GetMenuByIdBiz {
	return &GetMenuByIdBiz{store: store}
}

func (biz *GetMenuByIdBiz) GetNewMenuById(ctx context.Context,id int) (*models.Menu,error) {

	data, err := biz.store.GetMenuById(ctx,map[string]interface{}{"id":id}); 
	if err != nil {
		return nil,err
	}

	return data,nil
}
