package app

import (
	"github.com/nuttchai/go-ddd/common/logger"
	constant "github.com/nuttchai/go-ddd/user/config/constants"
)

var Logger logger.ILogger

func init() {
	Logger = logger.NewLogger(constant.UserDomain)
}
