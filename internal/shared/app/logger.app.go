package app

import (
	"github.com/nuttchai/go-ddd/common/logger"
)

const (
	domain = "USER_SERVICE"
)

var Logger logger.ILogger

func init() {
	Logger = logger.NewLogger(domain)
}
