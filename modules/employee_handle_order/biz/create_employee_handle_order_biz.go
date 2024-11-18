package biz

import (
	"context"
	"first-app/modules/employee_handle_order/models"
)

type CreateEmployeeHandleOrderStorage interface {
	CreateEmployeeHandleOrder(ctx context.Context, data *models.EmployeeHandleOrder) error
}

type CreateEmployeeHandleOrderBiz struct {
	store CreateEmployeeHandleOrderStorage
}

func NewCreateEmployeeHandleOrderBiz(store CreateEmployeeHandleOrderStorage) *CreateEmployeeHandleOrderBiz {
	return &CreateEmployeeHandleOrderBiz{store: store}
}

func (biz *CreateEmployeeHandleOrderBiz) CreateNewEmployeeHandleOrder(ctx context.Context, data *models.EmployeeHandleOrder) error {
	if err := biz.store.CreateEmployeeHandleOrder(ctx, data); err != nil {
		return err
	}
	return nil
}
