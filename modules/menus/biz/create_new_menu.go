package biz

import (
	"context"
	"first-app/modules/menus/models"
	"strings"
)

type CreateMenuStorage interface {
	CreateMenu(ctx context.Context,data *models.Menu) error
}

type CreateMenuBiz struct {
	store CreateMenuStorage
}

func NewCreateMenuBiz(store CreateMenuStorage) *CreateMenuBiz {
	return &CreateMenuBiz{store: store}
}

func (biz *CreateMenuBiz) CreateNewMenu(ctx context.Context,data *models.Menu) error {
	title := strings.TrimSpace(data.Name)

	if title == "" {
		return models.ErrTitleIsBlank
	}

	if err := biz.store.CreateMenu(ctx,data); err != nil {
		return err
	}

	return nil
}
