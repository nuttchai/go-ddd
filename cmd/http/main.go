package main

import (
	"github.com/nuttchai/go-ddd/common/logger"
	user "github.com/nuttchai/go-ddd/user/http/server"
)

func main() {
	logger.App.Log("Start client...")
	user.InitServer()
}
