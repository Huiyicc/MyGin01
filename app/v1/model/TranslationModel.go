package model

import "github.com/jinzhu/gorm"

type Translation struct {
	gorm.Model
	Article Article `gorm:"association_foreignkey:uid"`
	Apid int `gorm:"association_foreignkey:Apid"`
	Stype int
	money int
	supplies string
}
