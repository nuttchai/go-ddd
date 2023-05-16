package model

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	ID      string `gorm:"primaryKey;column:id;" json:"id"`
	Street  string `gorm:"type:varchar(30);size:30;column:street;" json:"street"`
	City    string `gorm:"type:varchar(30);size:30;column:city;" json:"city"`
	State   string `gorm:"type:varchar(5);size:5;column:state;" json:"state"`
	ZipCode string `gorm:"type:varchar(5);size:5;column:zip_code;" json:"zip_code"`
}
