package biz


import (
	"context"
	"first-app/modules/menus/models"
)

type DeleteMenuByIdStorage interface {
	GetMenuById(ctx context.Context,cond map[string]interface{}) (*models.Menu,error)
	DeleteMenuById(ctx context.Context,cond map[string]interface{}) error
}

type DeleteMenuByIdBiz struct {
	store DeleteMenuByIdStorage
}

func NewDeleteMenuByIdBiz(store DeleteMenuByIdStorage) *DeleteMenuByIdBiz {
	return &DeleteMenuByIdBiz{store: store}
}

func (biz *DeleteMenuByIdBiz) DeleteMenuById(ctx context.Context,id int) error {

	data, err := biz.store.GetMenuById(ctx,map[string]interface{}{"id":id}); 

	if err != nil {
		return err
	}

	if data.Name == "" {
		return models.ErrMenuDeleted
	}

	if err := biz.store.DeleteMenuById(ctx,map[string]interface{}{"id": id}); err != nil {
		return err
	}
	return nil
}
