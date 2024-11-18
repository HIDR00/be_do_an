package models

import (
	"errors"
	"first-app/common"
)

var (
	ErrTitleIsBlank= errors.New("title cannot be blank")
	ErrItemIsExit= errors.New("Device token is exit")
)

type Devices struct {
	common.SQLModel
	DeviceToken       string      `json:"device_token" gorm:"column:device_token;"`
}