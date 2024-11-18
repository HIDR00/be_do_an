package storage

import (
	"context"
	"first-app/modules/orders/models"
)

func (s *sqlStore) UpdateTableStatusPayById(ctx context.Context,cond map[string]interface{},dataUpdate *models.TableCreate) error {
	var tables models.Table

	db := s.db
    if err := db.Where(cond).First(&tables).Error; err != nil {
        return err
    }

	if err := s.db.Model(&tables).Select("Status", "PayType").Updates(models.TableCreate{Status: dataUpdate.Status,PayType: dataUpdate.PayType}).Error; err != nil {
		return err
	}
	return nil
}