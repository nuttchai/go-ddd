package repository

import "gorm.io/gorm"

type IRepository[TDomainEntity any, TDalEntity any] interface {
	PreloadAll() *gorm.DB
	FindOneById(id string) (*TDomainEntity, error)
	Save(entity *TDomainEntity) error
	Delete(entity *TDomainEntity) error
	IsExisted(id string) (bool, error)
}
