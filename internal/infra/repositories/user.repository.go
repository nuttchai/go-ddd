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

func (r *UserRepository) UpdateFirstNameIfIdExist(id, firstName string) error {
	user, err := r.Repository.FindOneById(id)
	if err != nil {
		return err
	}

	userDAL := r.dataMapper.ToDalEntity(user)
	userDAL.FirstName = firstName
	return r.Repository.Save(user)
}

func (r *UserRepository) FindOneByEmail(email string) (*entity.User, error) {
	userDAL := model.User{
		Email: email,
	}
	if dbResult := r.queryAdapter.Find(&userDAL).First(&userDAL); dbResult.Error != nil {
		if dbResult.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, dbResult.Error
	}

	return r.dataMapper.ToDomainEntity(&userDAL), nil
}
