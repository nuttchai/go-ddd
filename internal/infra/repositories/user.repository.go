package repository

import (
	cmapper "github.com/nuttchai/go-ddd/common/infra/data-mappers"
	crepo "github.com/nuttchai/go-ddd/common/infra/repositories"
	entity "github.com/nuttchai/go-ddd/internal/domain/entities"
	irepo "github.com/nuttchai/go-ddd/internal/domain/repositories"
	model "github.com/nuttchai/go-ddd/internal/infra/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	*crepo.Repository[entity.User, model.User]
	queryAdapter *gorm.DB
	dataMapper   cmapper.IDataMapper[entity.User, model.User]
}

func NewUserRepository(queryAdapter *gorm.DB, dataMapper cmapper.IDataMapper[entity.User, model.User]) irepo.IUserRepository {
	return &UserRepository{
		Repository:   crepo.NewRepository(queryAdapter, dataMapper),
		queryAdapter: queryAdapter,
		dataMapper:   dataMapper,
	}
}

func (r *UserRepository) UpdateUser(entity *entity.User) error {
	item := new(model.User)
	if dbResult := r.PreloadAll().Where("id = ?", entity.ID).First(item); dbResult.Error != nil {
		return dbResult.Error
	}
	if item.AddressID == "" {
		return r.Save(entity)
	}

	dalEntity := r.dataMapper.ToDalEntity(entity)
	dalEntity.AddressID = item.AddressID
	dalAddress := dalEntity.Address
	dalEntity.Address = model.Address{}
	if updateUserResult := r.queryAdapter.Model(&dalEntity).Updates(&dalEntity); updateUserResult.Error != nil {
		return updateUserResult.Error
	}

	updateAddressResult := r.queryAdapter.Model(&dalAddress).Where("id = ?", item.AddressID).Updates(&dalAddress)
	return updateAddressResult.Error
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
