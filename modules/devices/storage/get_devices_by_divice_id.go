package storage

import (
	"context"
	"first-app/modules/devices/models"
)

func (s *sqlStore) GetDevicesByDeviceId(ctx context.Context,cond map[string]interface{}) (*models.Devices,error) {
	var data models.Devices

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil,err
	}
	return &data,nil
}