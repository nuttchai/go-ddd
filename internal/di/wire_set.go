//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	cmapper "github.com/nuttchai/go-ddd/common/infra/data-mappers"
	application "github.com/nuttchai/go-ddd/internal/app"
	entity "github.com/nuttchai/go-ddd/internal/domain/entities"
	irepository "github.com/nuttchai/go-ddd/internal/domain/repositories"
	iservice "github.com/nuttchai/go-ddd/internal/domain/services"
	route "github.com/nuttchai/go-ddd/internal/http/client/routers"
	controller "github.com/nuttchai/go-ddd/internal/http/controllers"
	mapper "github.com/nuttchai/go-ddd/internal/infra/data-mappers"
	model "github.com/nuttchai/go-ddd/internal/infra/models"
	repository "github.com/nuttchai/go-ddd/internal/infra/repositories"
	dto "github.com/nuttchai/go-ddd/internal/shared/dtos"
)

var InfrastructureProviderSet = wire.NewSet(
	ProvideAppConfig,
	ProvideDB,
	ProvideEcho,
	route.NewRouter,
)

var UserModuleProviderSet = wire.NewSet(
	mapper.NewUserDataMapper,
	wire.Bind(new(cmapper.IDataMapper[entity.User, model.User]), new(*mapper.UserDataMapper)),
	mapper.NewUserRequestDataMapper,
	wire.Bind(new(cmapper.IDataMapper[entity.User, dto.UserDTO]), new(*mapper.UserRequestDataMapper)),
	ProvideUserRepository,
	wire.Bind(new(irepository.IUserRepository), new(*repository.UserRepository)),
	ProvideUserService,
	wire.Bind(new(iservice.IUserService), new(*iservice.UserService)),
	ProvideUserApplicationService,
	wire.Bind(new(application.IUserApplicationService), new(*application.UserApplicationService)),
	ProvideUserController,
	wire.Bind(new(controller.IUserController), new(*controller.UserController)),
	RegisterUserRoutes,
)

var HTTPServerProviderSet = wire.NewSet(
	InfrastructureProviderSet,
	UserModuleProviderSet,
	ProvideHTTPServer,
)
