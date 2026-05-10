package di

import (
	"errors"
	"flag"
	"fmt"

	"github.com/labstack/echo"
	config "github.com/nuttchai/go-ddd/common/config"
	cmapper "github.com/nuttchai/go-ddd/common/infra/data-mappers"
	middleware "github.com/nuttchai/go-ddd/common/middlewares"
	application "github.com/nuttchai/go-ddd/internal/app"
	entity "github.com/nuttchai/go-ddd/internal/domain/entities"
	irepository "github.com/nuttchai/go-ddd/internal/domain/repositories"
	iservice "github.com/nuttchai/go-ddd/internal/domain/services"
	route "github.com/nuttchai/go-ddd/internal/http/client/routers"
	controller "github.com/nuttchai/go-ddd/internal/http/controllers"
	model "github.com/nuttchai/go-ddd/internal/infra/models"
	repository "github.com/nuttchai/go-ddd/internal/infra/repositories"
	dto "github.com/nuttchai/go-ddd/internal/shared/dtos"
	types "github.com/nuttchai/go-ddd/internal/shared/types"
	context "github.com/nuttchai/go-ddd/utils/context"
	env "github.com/nuttchai/go-ddd/utils/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	initConnectionTimeout = 5

	invalidUserRepositoryTypeAssertion = "invalid_user_repository_type_assertion"
	invalidUserServiceTypeAssertion    = "invalid_user_service_type_assertion"
	invalidUserAppTypeAssertion        = "invalid_user_application_type_assertion"
	invalidUserControllerTypeAssertion = "invalid_user_controller_type_assertion"
)

type HTTPServer struct {
	Echo      *echo.Echo
	AppConfig *config.AppConfig
}

type userRoutesReady struct{}

func ProvideAppConfig() (*config.AppConfig, error) {
	appEnv := env.GetEnv("APP_ENV", env.Local.Name)
	envDefaultDir, err := env.GetDefaultEnvFileDirectoryPath(appEnv)
	if err != nil {
		return nil, err
	}

	envDir := env.GetEnv("ENV_PATH", envDefaultDir)
	_ = env.LoadEnvFile(envDir)

	dbType := env.GetEnv("DB_TYPE", "postgres")
	dbUser := env.GetEnv("APP_DB_USER", "postgres")
	dbPass := env.GetEnv("APP_DB_PASS", "postgres")
	dbHost := env.GetEnv("DB_HOST", "localhost")
	dbPort := env.GetEnv("DB_PORT", "5432")
	dbName := env.GetEnv("APP_DB_NAME", "userdb")
	dbDriver := env.GetEnv("DB_DRIVER", "postgres")
	port := env.GetEnv("APP_PORT", "8000")
	dbConnStr := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
		dbType,
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	var envArg, serverPort, dsn, driver string
	flag.StringVar(&envArg, "env", appEnv, "Application Environment")
	flag.StringVar(&serverPort, "port", port, "Server Listening Port")
	flag.StringVar(&dsn, "dsn", dbConnStr, "Data Source Name")
	flag.StringVar(&driver, "driver", dbDriver, "Database Driver")
	flag.Parse()

	appConfig := types.AppConfig
	appConfig.SetENV(envArg)
	appConfig.SetRESTConfig(serverPort)
	appConfig.SetDBMetaData(dsn, driver)

	return appConfig, nil
}

func ProvideDB(appConfig *config.AppConfig) (*gorm.DB, func(), error) {
	dbConfig := &gorm.Config{}
	if appConfig.GetENV() == env.Production.Name {
		dbConfig.Logger = logger.Default.LogMode(logger.Silent)
	}

	db, err := gorm.Open(postgres.Open(appConfig.GetDBConfig().GetDSN()), dbConfig)
	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(initConnectionTimeout)
	defer cancel()
	if err = sqlDB.PingContext(ctx); err != nil {
		return nil, nil, err
	}

	appConfig.SetDBInstance(db)
	cleanup := func() {
		_ = sqlDB.Close()
	}

	return db, cleanup, nil
}

func ProvideEcho() *echo.Echo {
	e := echo.New()
	middleware.EnableCORS(e)
	return e
}

func ProvideUserRepository(queryAdapter *gorm.DB, dataMapper cmapper.IDataMapper[entity.User, model.User]) (*repository.UserRepository, error) {
	userRepository := repository.NewUserRepository(queryAdapter, dataMapper)
	repo, ok := userRepository.(*repository.UserRepository)
	if !ok {
		return nil, errors.New(invalidUserRepositoryTypeAssertion)
	}
	return repo, nil
}

func ProvideUserService(userRepo irepository.IUserRepository) (*iservice.UserService, error) {
	userService := iservice.NewUserService(userRepo)
	service, ok := userService.(*iservice.UserService)
	if !ok {
		return nil, errors.New(invalidUserServiceTypeAssertion)
	}
	return service, nil
}

func ProvideUserApplicationService(userService iservice.IUserService, userReqDataMapper cmapper.IDataMapper[entity.User, dto.UserDTO]) (*application.UserApplicationService, error) {
	userApplicationService := application.NewUserApplicationService(userService, userReqDataMapper)
	appService, ok := userApplicationService.(*application.UserApplicationService)
	if !ok {
		return nil, errors.New(invalidUserAppTypeAssertion)
	}
	return appService, nil
}

func ProvideUserController(userAppService application.IUserApplicationService) (*controller.UserController, error) {
	userController := controller.NewUserController(userAppService)
	httpController, ok := userController.(*controller.UserController)
	if !ok {
		return nil, errors.New(invalidUserControllerTypeAssertion)
	}
	return httpController, nil
}

func RegisterUserRoutes(e *echo.Echo, r *route.Router, userController controller.IUserController) userRoutesReady {
	r.InitUserRouter(e, userController)
	return userRoutesReady{}
}

func ProvideHTTPServer(e *echo.Echo, appConfig *config.AppConfig, _ userRoutesReady) *HTTPServer {
	return &HTTPServer{
		Echo:      e,
		AppConfig: appConfig,
	}
}
