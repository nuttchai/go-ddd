package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        string    `gorm:"primaryKey;column:id;"`
	FirstName string    `gorm:"type:varchar(100);size:100;not null;default:null;column:first_name;"`
	LastName  string    `gorm:"type:varchar(100);size:100;not null;default:null;column:last_name;"`
	Email     string    `gorm:"type:varchar(50);size:50;not null;default:null;column:email;"`
	CreatedAt time.Time `gorm:"autoCreateTime:true;column:created_at;"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:true;column:updated_at;"`

	Address   Address `gorm:"foreignkey:AddressId"`
	AddressId string  `gorm:"column:address_id;"`
}
