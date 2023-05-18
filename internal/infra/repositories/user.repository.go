package repository

import (
	mapper "github.com/nuttchai/go-ddd/common/infra/data-mappers"
	repository "github.com/nuttchai/go-ddd/common/infra/repositories"
	entity "github.com/nuttchai/go-ddd/internal/domain/entities"
	irepository "github.com/nuttchai/go-ddd/internal/domain/repositories"
	model "github.com/nuttchai/go-ddd/internal/infra/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	*repository.Repository[entity.User, model.User]
	queryAdapter *gorm.DB
	dataMapper   mapper.IDataMapper[entity.User, model.User]
}

func NewUserRepository(queryAdapter *gorm.DB, dataMapper mapper.IDataMapper[entity.User, model.User]) irepository.IUserRepository {
	return &UserRepository{
		Repository:   repository.NewRepository(queryAdapter, dataMapper),
		queryAdapter: queryAdapter,
		dataMapper:   dataMapper,
	}
}

func (r *UserRepository) UpdateUser(entity *entity.User) error {
	item := new(model.User)
	if dbResult := r.PreloadAll().Where("id = ?", entity.ID).First(item); dbResult.Error != nil {
		return dbResult.Error
	}

	dalEntity := r.dataMapper.ToDalEntity(entity)
	dalEntity.Address.ID = item.AddressID
	dbResult := r.queryAdapter.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&dalEntity)
	return dbResult.Error
}

func (r *UserRepository) FindOneByEmail(email string) (*entity.User, error) {
	userDAL := &model.User{Email: email}
	if dbResult := r.queryAdapter.Where("email = ?", userDAL.Email).First(&userDAL); dbResult.Error != nil {
		if dbResult.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, dbResult.Error
	}
	return r.dataMapper.ToDomainEntity(userDAL), nil
}
