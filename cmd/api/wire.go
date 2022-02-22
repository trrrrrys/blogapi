//go:build wireinject
// +build wireinject

package main

import (
	"blog-api/application/usecase"
	"blog-api/config"
	"blog-api/infrastructure"
	"blog-api/interface/api/graphql/handler"
	"blog-api/interface/api/middleware"
	"blog-api/interface/api/rest"
	"blog-api/interface/api/router"
	"fmt"
	"net/http"
	"time"

	"github.com/google/wire"
)

var initSet = wire.NewSet(
	config.NewConfig,
)

var infraSet = wire.NewSet(
	infrastructure.NewContentRepository,
	infrastructure.NewUserRepository,
)

var usecaseSet = wire.NewSet(
	usecase.NewContentUsecase,
	usecase.NewUserUsecase,
)

func ProvideServer(h http.Handler, c *config.Config) *http.Server {
	// set middleware
	handler := middleware.SetMiddleware(h)
	return &http.Server{
		Addr:              fmt.Sprintf(":%v", c.Port),
		Handler:           handler,
		ReadHeaderTimeout: 30 * time.Second,
	}
}

var handlerSet = wire.NewSet(
	handler.NewContentHandler,
	handler.NewUserHandler,
	handler.NewHandler,
	rest.NewRestHandler,
	rest.NewContentHandler,
	router.Route,
	ProvideServer,
)

func InitializeServer() (*http.Server, func(), error) {
	wire.Build(
		initSet,
		infraSet,
		usecaseSet,
		handlerSet,
	)
	return nil, nil, nil
}
