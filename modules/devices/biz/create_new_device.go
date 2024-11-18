package biz

import (
	"context"
	"first-app/modules/devices/models"
)

type CreateDevicesStorage interface {
	CreateDevices(ctx context.Context, data *models.Devices) error
	GetDevicesByDeviceId(ctx context.Context, cond map[string]interface{}) (*models.Devices, error)
}

type CreateDevicesBiz struct {
	store CreateDevicesStorage
}

func NewCreateDevicesBiz(store CreateDevicesStorage) *CreateDevicesBiz {
	return &CreateDevicesBiz{store: store}
}

func (biz *CreateDevicesBiz) CreateNewDevices(
	ctx context.Context,
	data *models.Devices,
	deviceToken string,
) error {

	result, err := biz.store.GetDevicesByDeviceId(ctx, map[string]interface{}{"device_token": deviceToken})

	if err != nil {
		return err
	}

	if result != nil {
		return models.ErrItemIsExit
	}

	if err := biz.store.CreateDevices(ctx, data); err != nil {
		return err
	}

	return nil
}
