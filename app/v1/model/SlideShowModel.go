package model

import "github.com/jinzhu/gorm"

type SildeShow struct {
	gorm.Model
	Src string `gorm:"type:varchar(255);not null"`
	Disable int `gorm:"type:int(1);not null;default:0"`
	Url string `gorm:"type:varchar(255);not null"`
}