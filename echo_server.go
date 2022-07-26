package main

import (
	"fmt"
	"net/http"

	"github.com/Klinisia/backend-ksi/config"
	"github.com/Klinisia/backend-ksi/middleware"
	"github.com/go-redis/redis"

	authDeliveryPackage "github.com/Klinisia/backend-ksi/src/auth/v1/delivery"
	authRepositoryPackage "github.com/Klinisia/backend-ksi/src/auth/v1/repository"
	authUsecasePackage "github.com/Klinisia/backend-ksi/src/auth/v1/usecase"
	"github.com/Klinisia/backend-ksi/src/shared/external"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
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
	authGroupV1 := e.Group("/auth")
	// authGroupV1.Use(middleware.BasicAuth(basicAuthConfig))
	authGroupV1.Use()
	s.authEchoHandler.Mount(authGroupV1)

	listenerPort := fmt.Sprintf(":%d", config.HTTPPort)
	e.Logger.Fatal(e.Start(listenerPort))
}

// NewEchoServer function
func NewEchoServer(writeDb, readDb *gorm.DB, redisWrite *redis.Client) (*EchoServer, error) {

	extServiceSms, err := external.NewSmsAcs()
	if err != nil {
		return nil, err
	}

	// auth
	authRepositoryWrite := authRepositoryPackage.NewAuthRepositoryGorm(writeDb)
	authUsecase := authUsecasePackage.NewAuthUsecaseImpl(authRepositoryWrite, extServiceSms)
	authEchoHandler := authDeliveryPackage.NewEchoHandler(authUsecase)

	return &EchoServer{

		authEchoHandler: authEchoHandler,
	}, nil
}
