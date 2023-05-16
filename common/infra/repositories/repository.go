package repository

import (
	mapper "github.com/nuttchai/go-ddd/common/infra/data-mappers"
	"gorm.io/gorm"
)

type Repository[TDomainEntity any, TDalEntity any] struct {
	queryAdapter *gorm.DB
	dataMapper   mapper.IDataMapper[TDomainEntity, TDalEntity]
}

func NewRepository[TDomainEntity any, TDalEntity any](queryAdapter *gorm.DB, dataMapper mapper.IDataMapper[TDomainEntity, TDalEntity]) *Repository[TDomainEntity, TDalEntity] {
	return &Repository[TDomainEntity, TDalEntity]{
		queryAdapter: queryAdapter,
		dataMapper:   dataMapper,
	}
}

func (r *Repository[TDomainEntity, TDalEntity]) FindOneById(id string) (*TDomainEntity, error) {
	item := new(TDalEntity)
	if dbResult := r.queryAdapter.Where("id = ?", id).First(item); dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return r.dataMapper.ToDomainEntity(item), nil
}

func (r *Repository[TDomainEntity, TDalEntity]) Save(entity *TDomainEntity) error {
	isExisted := r.IsExisted(entity)
	dalEntity := r.dataMapper.ToDalEntity(entity)
	if isExisted {
		dbResult := r.queryAdapter.Model(&dalEntity).Updates(&dalEntity)
		return dbResult.Error
	}
	dbResult := r.queryAdapter.Save(&dalEntity)
	return dbResult.Error
}

func (r *Repository[TDomainEntity, TDalEntity]) Delete(id string) error {
	return nil
}

func (r *Repository[TDomainEntity, TDalEntity]) IsExisted(entity *TDomainEntity) bool {
	dalEntity := r.dataMapper.ToDalEntity(entity)
	return r.queryAdapter.First(&dalEntity).RowsAffected > 0
}
