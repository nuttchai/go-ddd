package main

import (
	"github.com/nuttchai/go-ddd/common/logger"
	app "github.com/nuttchai/go-ddd/internal/http/client/config"
)

func main() {
	logger := logger.NewLogger()
	logger.Log("Starting server...")

	app.InitServer()
}
