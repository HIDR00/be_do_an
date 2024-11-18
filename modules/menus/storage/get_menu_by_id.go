package storage;

import (
	"context"
	"first-app/modules/menus/models"
)

func (s *sqlStore) GetMenuById(ctx context.Context,cond map[string]interface{}) (*models.Menu,error) {
	var data models.Menu

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil,err
	}
	return &data,nil
}