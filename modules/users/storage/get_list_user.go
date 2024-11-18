package storage

import (
	"context"
	"first-app/common"
	"first-app/modules/users/models"
)

func (s *sqlStore) GetListUser(ctx context.Context, paging *common.Paging, moreKey ...string) ([]models.User, error){
	var result []models.User

	db := s.db

	if err := s.db.Table(models.User{}.TableName()).Count(&paging.Total).Error; err != nil {
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
