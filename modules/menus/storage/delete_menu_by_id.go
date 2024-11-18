package storage

import (
	"context"
	"first-app/modules/menus/models"
)

func (s *sqlStore) DeleteMenuById(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table(models.Menu{}.TableName()).Where(cond).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
