//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	di "github.com/nuttchai/go-ddd/internal/di"
)

func InitializeHTTPServer() (*di.HTTPServer, func(), error) {
	wire.Build(di.HTTPServerProviderSet)
	return &di.HTTPServer{}, nil, nil
}
