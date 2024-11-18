package storage

import (
	"context"
	"first-app/modules/menus/models"
)

func (s *sqlStore) UpdateMenuById(ctx context.Context,cond map[string]interface{},dataUpdate *models.Menu) error {
	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}