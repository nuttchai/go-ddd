package config

import (
	"fmt"

	"github.com/labstack/echo"
	middleware "github.com/nuttchai/go-ddd/common/middlewares"
	app "github.com/nuttchai/go-ddd/internal/shared/app"
	types "github.com/nuttchai/go-ddd/internal/shared/types"
)

func InitServer() {
	// Add the Configuration into AppConfig
	app.Logger.Info("Loading Configuration...")
	if err := initEnv(); err != nil {
		app.Logger.Error("Error Loading App Configuration (Error: %s)", err.Error())
	}

	// Establish Database Connection
	app.Logger.Info("Connecting to Database...")
	if err := initDB(); err != nil {
		app.Logger.Error("Error Connecting to Database (Error: %s)", err.Error())
	}

	// Initialize the Application
	app.Logger.Info("Initializing the Application...")
	e := echo.New()
	middleware.EnableCORS(e)
	if err := initApp(e); err != nil {
		app.Logger.Error("Error Initializing the Application (Error: %s)", err.Error())
	}

	// Defer Closing the Database Connection
	db, _ := getDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Start Server
	app.Logger.Info("Starting Server...")
	serverPort := fmt.Sprintf(":%s", types.AppConfig.GetRESTPort())
	if err := e.Start(serverPort); err != nil {
		app.Logger.Error("Server Start Failed (Error: %s)", err.Error())
	}
}
