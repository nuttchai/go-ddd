package repository

import (
	"errors"

	constant "github.com/nuttchai/go-ddd/common/constants"
	mapper "github.com/nuttchai/go-ddd/common/infrastructure/data-mappers"
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
	return r.dataMapper.ToDomainEntity(item)
}

func (r *Repository[TDomainEntity, TDalEntity]) Save(entity *TDomainEntity) error {
	isExisted, err := r.IsExisted(entity)
	if err != nil {
		return err
	}
	if isExisted {
		return errors.New(constant.ItemAlreadyExisted)
	}

	dalEntity, err := r.dataMapper.ToDalEntity(entity)
	if err != nil {
		return err
	}

	return r.queryAdapter.Save(dalEntity).Error
}

func (r *Repository[TDomainEntity, TDalEntity]) Update(entity *TDomainEntity) error {
	return nil
}

func (r *Repository[TDomainEntity, TDalEntity]) Delete(id string) error {
	return nil
}

func (r *Repository[TDomainEntity, TDalEntity]) IsExisted(entity *TDomainEntity) (bool, error) {
	dalEntity, err := r.dataMapper.ToDalEntity(entity)
	if err != nil {
		return false, err
	}

	if dbResult := r.queryAdapter.First(dalEntity); dbResult.RowsAffected > 0 {
		return false, nil
	}

	return true, nil
}
