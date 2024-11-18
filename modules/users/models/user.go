package models

import (
	"errors"
	"first-app/common"
)

var (
	ErrTitleIsBlank = errors.New("email or password cannot be blank")
	ErrItemDeleted  = errors.New("title cannot be deleted")
)

type User struct {
	common.SQLModel
	Name     string `json:"name" gorm:"column:name;"`
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"password" gorm:"column:password;"`
	RoleType int    `json:"role_type" gorm:"column:role_type;"`
}

type UserAuthen struct {
	common.SQLModel
	Name     string `json:"name" gorm:"column:name;"`
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"password" gorm:"column:password;"`
	RoleType int    `json:"role_type" gorm:"column:role_type;"`
}
