package config

import (
	"fmt"

	app "github.com/nuttchai/go-ddd/internal/shared/app"
)

func InitServer() {
	app.Logger.Info("Initializing the Application...")
	server, cleanup, err := InitializeHTTPServer()
	if err != nil {
		app.Logger.Error("Error Initializing the Application (Error: %s)", err.Error())
		return
	}
	defer cleanup()

	app.Logger.Info("Starting Server...")
	serverPort := fmt.Sprintf(":%s", server.AppConfig.GetRESTPort())
	if err := server.Echo.Start(serverPort); err != nil {
		app.Logger.Error("Server Start Failed (Error: %s)", err.Error())
	}
}
