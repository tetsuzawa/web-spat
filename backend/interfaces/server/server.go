package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tetsuzawa/web-spat/config"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	buildEnv := config.GetEnvWithDefault("BUILD_ENV", "Release")
	if buildEnv == "Debug" {
		log.Println("server is running on Debug mode...")
		e.Debug = true
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	return e
}
