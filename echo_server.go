package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
	"gitlab.com/k1476/scaffolding/config"
	"gitlab.com/k1476/scaffolding/middleware"

	authDeliveryPackage "gitlab.com/k1476/scaffolding/src/auth/v1/delivery"
	authRepositoryPackage "gitlab.com/k1476/scaffolding/src/auth/v1/repository"
	authUsecasePackage "gitlab.com/k1476/scaffolding/src/auth/v1/usecase"
	"gitlab.com/k1476/scaffolding/src/shared/external"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// EchoServer structure
type EchoServer struct {

	// auth
	authEchoHandler *authDeliveryPackage.EchoHandler
}

// Run main function for serving echo http server
func (s *EchoServer) Run() {
	basicAuthConfig := middleware.BasicAuthConfig{Username: config.BasicAuthUsername, Password: config.BasicAuthPassword}

	e := echo.New()
	e.Use(middleware.ServerHeader, middleware.Logger)
	//e.Use(mid.Recover())

	if config.Development == "1" {
		e.Debug = true
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Up and running !!")
	})

	e.Static("/static", "static")

	// Serve Coverage result
	coverGroup := e.Group("/cover")
	coverGroup.Use(middleware.BasicAuth(basicAuthConfig))
	coverGroup.GET("", func(c echo.Context) error {
		return c.File("coverages/index.html")
	})

	// auth v1 route
	authGroupV1 := e.Group("/v1/auth")
	authGroupV1.Use(middleware.BasicAuth(basicAuthConfig))
	s.authEchoHandler.Mount(authGroupV1)

	listenerPort := fmt.Sprintf(":%d", config.HTTPPort)
	e.Logger.Fatal(e.Start(listenerPort))
}

// NewEchoServer function
func NewEchoServer(writeDb, readDb *gorm.DB, redisWrite *redis.Client) (*EchoServer, error) {
	// trxservice
	externalService, err := external.NewNotifWhatsapp()
	if err != nil {
		return nil, err
	}

	// auth
	authRepositoryWrite := authRepositoryPackage.NewAuthRepositoryGorm(writeDb)
	authUsecase := authUsecasePackage.NewAuthUsecaseImpl(authRepositoryWrite, externalService)
	authEchoHandler := authDeliveryPackage.NewEchoHandler(authUsecase)

	return &EchoServer{

		authEchoHandler: authEchoHandler,
	}, nil
}
