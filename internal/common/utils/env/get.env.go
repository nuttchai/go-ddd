package env

import (
	"os"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}

	return value
}

func GetDefaultEnvFileDirectoryPath(appEnv string) (string, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	envFileName := Local.File
	if appEnv == Production.Name {
		envFileName = Production.File
	}

	envFilePath := rootDir + "/" + envFileName
	return envFilePath, nil
}
