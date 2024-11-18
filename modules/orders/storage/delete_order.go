package storage

import (
	"context"
	"first-app/common"
	"first-app/modules/orders/models"
	"time"
)

func (s *sqlStore) DeleteOrder(ctx context.Context, cond map[string]interface{}) error {
	var data []models.OrderCreate

	if err := s.db.Where(cond).Find(&data).Error; err != nil {
		return err
	}
	da := time.Now()
	if err := s.db.Model(&data).Select("DeletedAt").Updates(models.OrderCreate{
		SQLModel: common.SQLModel{DeletedAt: &da},
	}).Error; err != nil {
		return err
	}
	return nil
}
