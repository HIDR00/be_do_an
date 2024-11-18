package biz

import (
	"context"
	"first-app/modules/orders/models"
)

type CreateOrderStorage interface {
	UpdateTableStatusPayById(ctx context.Context,cond map[string]interface{},dataUpdate *models.TableCreate) error
	CreateOrder(ctx context.Context,data *[]models.OrderCreate,cond map[string]interface{}) error
}

type CreateOrderBiz struct {
	store CreateOrderStorage
}

func NewCreateOrderBiz(store CreateOrderStorage) *CreateOrderBiz {
	return &CreateOrderBiz{store: store}
}

func (biz *CreateOrderBiz) CreateNewOrder(ctx context.Context,data *models.TableCreate,id int) error {

	if err := biz.store.UpdateTableStatusPayById(ctx,map[string]interface{}{"id":id},data); err != nil {
		return err
	}

	if err := biz.store.CreateOrder(ctx,&data.Orders,map[string]interface{}{"id":id}); err != nil {
		return err
	}

	return nil
}
