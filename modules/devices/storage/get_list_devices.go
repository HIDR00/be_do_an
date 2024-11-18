package storage

import (
	"context"
	"first-app/common"
	"first-app/modules/devices/models"
)

func (s *sqlStore) GetListDevices(ctx context.Context, paging *common.Paging, moreKey ...string) ([]models.Devices, error) {
	var result []models.Devices

	db := s.db

	if err := s.db.Table(models.Devices{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("id asc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
