package app

import service "github.com/nuttchai/go-ddd/user/domain/services"

type UserApp struct {
	userService service.IUserService
}

func NewUserApp(userService service.IUserService) *UserApp {
	return &UserApp{
		userService: userService,
	}
}

func (a *UserApp) FindOneById(id string) (*User, error) {
	user, err := a.userService.FindOneById(id)
	if err != nil {
		return nil, err
	}
	return NewUserFromEntity(user), nil
}

func (a *UserApp) CreateUser(user *User) error {
	return a.userService.CreateUser(user.ToEntity())
}

func (a *UserApp) UpdateUser(user *User) error {
	return a.userService.UpdateUser(user.ToEntity())
}

func (a *UserApp) UpdateFirstNameIfIdExist(id, firstName string) error {
	return a.userService.UpdateFirstNameIfIdExist(id, firstName)
}

func (a *UserApp) FindOneByEmail(email string) (*User, error) {
	user, err := a.userService.FindOneByEmail(email)
	if err != nil {
		return nil, err
	}
	return NewUserFromEntity(user), nil
}
