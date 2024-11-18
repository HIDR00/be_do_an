package biz

import (
	"context"
	"first-app/common"
	"first-app/modules/employee_handle_order/models"
)

type GetListEmployeeHandleOrderStorage interface {
	GetListEmployeeHandleOrder(ctx context.Context, paging *common.Paging, moreKey ...string) ([]models.GetEmployeeHandleOrder, error)
}

type GetListEmployeeHandleOrderBiz struct {
	store GetListEmployeeHandleOrderStorage
}

func NewGetListEmployeeHandleOrderBiz(store GetListEmployeeHandleOrderStorage) *GetListEmployeeHandleOrderBiz {
	return &GetListEmployeeHandleOrderBiz{store: store}
}

func (biz *GetListEmployeeHandleOrderBiz) GetNewListEmployeeHandleOrder(ctx context.Context, paging *common.Paging) ([]models.GetEmployeeHandleOrder, error) {
	data, err := biz.store.GetListEmployeeHandleOrder(ctx, paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}