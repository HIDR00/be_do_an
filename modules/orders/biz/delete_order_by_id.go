package biz

import (
	"context"
	"first-app/modules/orders/models"
)

type DeleteOrderByIdStorage interface {
	UpdateTableStatusPayById(ctx context.Context,cond map[string]interface{},dataUpdate *models.TableCreate) error
	DeleteOrder(ctx context.Context, cond map[string]interface{}) error
}

type DeleteOrderByIdBiz struct {
	store DeleteOrderByIdStorage
}

func NewDeleteOrderByIdBiz(store DeleteOrderByIdStorage) *DeleteOrderByIdBiz {
	return &DeleteOrderByIdBiz{store: store}
}

func (biz *DeleteOrderByIdBiz) DeleteNewOrderById(ctx context.Context, id int, dataDelete *models.TableCreate) error {

	if err := biz.store.UpdateTableStatusPayById(ctx,map[string]interface{}{"id":id},dataDelete); err != nil {
		return err
	}

	if err := biz.store.DeleteOrder(ctx, map[string]interface{}{"table_id": id}); err != nil {
		return err
	}
	return nil
}