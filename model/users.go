package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
}
