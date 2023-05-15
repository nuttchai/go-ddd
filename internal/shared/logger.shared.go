package app

import (
	"github.com/nuttchai/go-ddd/common/logger"
	constant "github.com/nuttchai/go-ddd/internal/http/client/constants"
)

var Logger logger.ILogger

func init() {
	Logger = logger.NewLogger(constant.UserDomain)
}
