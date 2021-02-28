package server

import (
	"fmt"
	"log"

	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tetsuzawa/web-spat/config"
	"github.com/tetsuzawa/web-spat/infrastructure/persistence"
	"github.com/tetsuzawa/web-spat/interfaces/server/handler"
	"github.com/tetsuzawa/web-spat/interfaces/server/openapi"
	"github.com/tetsuzawa/web-spat/usecase"
)

func Run(port int) {
	// connect to DB
	db, err := config.NewDBConnection()
	defer db.Close()
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to init DB connection -> %w", err))
	}

	e := echo.New()

	if buildEnv := config.GetEnvWithDefault("BUILD_ENV", "Release"); buildEnv == "Debug" {
		log.Println("server is running on Debug mode...")
		e.Debug = true
	}

	// routing
	h := handler.NewIntegratedHandler(
		*handler.NewExperimentsHandler(usecase.NewExperimentUseCase(persistence.NewExperimentRepository(db))),
		*handler.NewUtilHandler(),
	)

	openapi.RegisterHandlersWithBaseURL(e, h, "/v1")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	swagger, err := openapi.GetSwagger()
	if err != nil {
		log.Fatalf("Error loading swagger spec\n: %s", err)
	}

	e.Use(oapimiddleware.OapiRequestValidator(swagger))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
