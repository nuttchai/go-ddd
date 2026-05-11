//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	cmapper "github.com/nuttchai/go-ddd/common/infra/data-mappers"
	irepository "github.com/nuttchai/go-ddd/internal/core/app/repositories"
	usecase "github.com/nuttchai/go-ddd/internal/core/app/usecases"
	entity "github.com/nuttchai/go-ddd/internal/core/domain/entities"
	controller "github.com/nuttchai/go-ddd/internal/core/http/controllers"
	route "github.com/nuttchai/go-ddd/internal/core/http/routers"
	mapper "github.com/nuttchai/go-ddd/internal/core/infra/data-mappers"
	model "github.com/nuttchai/go-ddd/internal/core/infra/models"
	repository "github.com/nuttchai/go-ddd/internal/core/infra/repositories"
)

var InfrastructureProviderSet = wire.NewSet(
	ProvideAppConfig,
	ProvideDB,
	ProvideEcho,
	ProvideEnvelopeBuilder,
	route.NewRouter,
)

var UserModuleProviderSet = wire.NewSet(
	mapper.NewUserDataMapper,
	wire.Bind(new(cmapper.IDataMapper[entity.User, model.User]), new(*mapper.UserDataMapper)),
	ProvideUserRepository,
	wire.Bind(new(irepository.IUserRepository), new(*repository.UserRepository)),
	ProvideUserUsecase,
	wire.Bind(new(usecase.IUserUsecase), new(*usecase.UserUsecase)),
	ProvideUserController,
	wire.Bind(new(controller.IUserController), new(*controller.UserController)),
	RegisterUserRoutes,
)

var HTTPServerProviderSet = wire.NewSet(
	InfrastructureProviderSet,
	UserModuleProviderSet,
	ProvideHTTPServer,
)
