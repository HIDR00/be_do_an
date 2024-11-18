package biz

import (
	"context"
	"first-app/modules/orders/models"
)

type GetOrderByIdStorage interface {
	GetOrderById(ctx context.Context,cond map[string]interface{}) (*models.TableGet,error)
}

type GetOrderByIdBiz struct {
	store GetOrderByIdStorage
}

func NewGetOrderByIdBiz(store GetOrderByIdStorage) *GetOrderByIdBiz {
	return &GetOrderByIdBiz{store: store}
}

func (biz *GetOrderByIdBiz) GetNewOrderById(ctx context.Context,id int) (*models.TableGet,error) {

	data, err := biz.store.GetOrderById(
		ctx,
		map[string]interface{}{"id":id,"deleted_at": nil},
	); 
	if err != nil {
		return nil,err
	}

	return data,nil
}
