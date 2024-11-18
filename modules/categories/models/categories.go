package models

import (
	"errors"
	"first-app/common"
)

var (
	ErrTitleIsBlank= errors.New("title cannot be blank")
	ErrItemDeleted= errors.New("title cannot be deleted")
)

type Categories struct {
	common.SQLModel
	Name       string      `json:"name" gorm:"column:name;"`
}