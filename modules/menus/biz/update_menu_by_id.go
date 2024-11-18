package biz

import (
	"context"
	"first-app/modules/menus/models"
)

type UpdateMenuByIdStorage interface {
	GetMenuById(ctx context.Context, cond map[string]interface{}) (*models.Menu, error)
	UpdateMenuById(ctx context.Context, cond map[string]interface{}, dataUpdate *models.Menu) error
}

type UpdateMenuByIdBiz struct {
	store UpdateMenuByIdStorage
}

func NewUpdateMenuByIdBiz(store UpdateMenuByIdStorage) *UpdateMenuByIdBiz {
	return &UpdateMenuByIdBiz{store: store}
}

func (biz *UpdateMenuByIdBiz) UpdateNewMenuById(ctx context.Context, id int, dataUpdate *models.Menu) error {

	data, err := biz.store.GetMenuById(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if data.Name == "" {
		return models.ErrMenuDeleted
	}

	if err := biz.store.UpdateMenuById(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}
	return nil
}
