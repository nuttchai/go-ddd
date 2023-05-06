package main

import (
	"github.com/nuttchai/go-ddd/common/logger"
	user "github.com/nuttchai/go-ddd/internal/config"
)

func main() {
	logger := logger.NewLogger()
	logger.Log("Starting server...")

	user.InitServer()
}
