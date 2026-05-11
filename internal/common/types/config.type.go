package types

import "github.com/nuttchai/go-ddd/internal/client/config"

var AppConfig *config.AppConfig

func init() {
	AppConfig = &config.AppConfig{}
}
