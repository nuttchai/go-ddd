package model

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Id        string    `gorm:"primaryKey" json:"id"`
	Street    string    `gorm:"type:varchar(30);size:30;column:street;" json:"street"`
	City      string    `gorm:"type:varchar(30);size:30;column:city;" json:"city"`
	State     string    `gorm:"type:varchar(5);size:5;column:state;" json:"state"`
	ZipCode   string    `gorm:"type:varchar(5);size:5;column:zip_code;" json:"zip_code"`
	CreatedAt time.Time `gorm:"autoCreateTime:true;column:created_at;" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:true;column:updated_at;" json:"updated_at"`
}
