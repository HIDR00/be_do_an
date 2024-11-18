package biz

import (
	"context"
	"first-app/common"
	"first-app/modules/devices/models"
)

type GetListDevicesStorage interface {
	GetListDevices(ctx context.Context, paging *common.Paging, moreKey ...string) ([]models.Devices, error)
}

type GetListDevicesBiz struct {
	store GetListDevicesStorage
}

func NewGetListDevicesBiz(store GetListDevicesStorage) *GetListDevicesBiz {
	return &GetListDevicesBiz{store: store}
}

func (biz *GetListDevicesBiz) GetNewListDevices(ctx context.Context,paging *common.Paging) ([]models.Devices, error) {

	data, err := biz.store.GetListDevices(ctx,paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}
