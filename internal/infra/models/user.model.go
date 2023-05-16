package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        string  `gorm:"primaryKey;column:id;"`
	FirstName string  `gorm:"type:varchar(100);size:100;not null;default:null;column:first_name;"`
	LastName  string  `gorm:"type:varchar(100);size:100;not null;default:null;column:last_name;"`
	Email     string  `gorm:"type:varchar(50);size:50;not null;default:null;column:email;"`
	Address   Address `gorm:"foreignkey:AddressID"`
	AddressID string  `gorm:"column:address_id;"`
}
