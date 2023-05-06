package config

import (
	types "github.com/nuttchai/go-ddd/types"
)

var AppConfig *types.AppConfig

func init() {
	AppConfig = &types.AppConfig{}
}
