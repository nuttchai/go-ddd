package types

import config "github.com/nuttchai/go-ddd/common/config"

var AppConfig *config.AppConfig

func init() {
	AppConfig = &config.AppConfig{}
}
