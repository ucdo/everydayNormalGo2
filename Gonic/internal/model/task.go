package model

import "gorm.io/gorm"

type Task struct {
	Name    string `gorm:"type:varchar(255)"`
	Context string `gorm:"type:text"`
	Status  uint   `gorm:"type:tinyint(1)"`
	gorm.Model
}
