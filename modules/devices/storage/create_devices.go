package storage

import (
	"context"
	"first-app/modules/devices/models"
)

func (s *sqlStore) CreateDevices(ctx context.Context, data *models.Devices) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}