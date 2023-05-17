package env

import (
	"github.com/joho/godotenv"
)

func LoadEnvFile(filename string) error {
	if err := godotenv.Load(filename); err != nil {
		return err
	}
	return nil
}
