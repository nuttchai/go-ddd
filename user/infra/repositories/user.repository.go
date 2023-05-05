package repository

import (
	"errors"

	cmapper "github.com/nuttchai/go-ddd/common/infra/data-mappers"
	crepo "github.com/nuttchai/go-ddd/common/infra/repositories"
	constant "github.com/nuttchai/go-ddd/user/domain/constants"
	entity "github.com/nuttchai/go-ddd/user/domain/entities"
	model "github.com/nuttchai/go-ddd/user/infra/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	*crepo.Repository[entity.User, model.User]
	queryAdapter *gorm.DB
	dataMapper   cmapper.IDataMapper[entity.User, model.User]
}

func NewUserRepository(queryAdapter *gorm.DB, dataMapper cmapper.IDataMapper[entity.User, model.User]) *UserRepository {
	return &UserRepository{
		Repository:   crepo.NewRepository(queryAdapter, dataMapper),
		queryAdapter: queryAdapter,
		dataMapper:   dataMapper,
	}
}

func (r *UserRepository) CreateUser(user *entity.User) error {
	userDb, err := r.FindOneByEmail(user.Email)
	if err != nil {
		return err
	}
	if userDb != nil {
		return errors.New(constant.EmailAlreadyExisted)
	}

	return r.Repository.Save(user)
}

func (r *UserRepository) UpdateFirstNameIfIdExist(id, firstName string) error {
	user, err := r.Repository.FindOneById(id)
	if err != nil {
		return err
	}

	userDAL, err := r.dataMapper.ToDalEntity(user)
	if err != nil {
		return err
	}

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

	return r.dataMapper.ToDomainEntity(&userDAL)
}
