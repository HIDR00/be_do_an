package storage

import (
	"context"
	"first-app/modules/users/models"
)

func (s *sqlStore) DeleteUserById(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table(models.User{}.TableName()).Where(cond).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
