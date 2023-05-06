package config

import (
	"flag"
	"fmt"

	env "github.com/nuttchai/go-ddd/utils/env"
)

func initEnv() error {
	// Load Environment Variables
	appEnv := env.GetEnv("APP_ENV", "development")
	envDefaultDir, err := env.GetDefaultEnvFileDirectoryPath(appEnv)
	if err != nil {
		return err
	}

	envDir := env.GetEnv("USER_ENV_PATH", envDefaultDir)
	env.LoadEnv(envDir)

	dbType := env.GetEnv("USER_DB_TYPE", "postgres")
	dbUser := env.GetEnv("USER_APP_DB_USER", "postgres")
	dbPass := env.GetEnv("USER_APP_DB_PASS", "postgres")
	dbHost := env.GetEnv("USER_DB_HOST", "localhost")
	dbPort := env.GetEnv("USER_DB_PORT", "5432")
	dbName := env.GetEnv("USER_APP_DB_NAME", "test")
	dbDriver := env.GetEnv("USER_DB_DRIVER", "postgres")
	port := env.GetEnv("USER_APP_PORT", "8000")
	dbConnStr := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
		dbType,
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	// Get Value from Command Line if Exist
	var env, serverPort, dsn, driver string
	flag.StringVar(&env, "userEnv", appEnv, "Application Environment")
	flag.StringVar(&serverPort, "userPort", port, "Server Listening Port")
	flag.StringVar(&dsn, "userDsn", dbConnStr, "Data Source Name")
	flag.StringVar(&driver, "userDriver", dbDriver, "Database Driver")
	flag.Parse()

	// Set Value to AppConfig
	AppConfig.SetENV(env)
	AppConfig.SetRESTConfig(serverPort)
	AppConfig.SetDBMetaData(dsn, driver)

	return nil
}
