package server

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tetsuzawa/web-spat/config"
	"github.com/tetsuzawa/web-spat/interfaces/server/handler"
)

type Server struct {
	e        *echo.Echo
	db       *sqlx.DB
	teardown func()
}

func NewServer() *Server {
	return &Server{teardown: func() {}}
}

func (s *Server) Init() (err error) {
	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)

	s.db, err = config.NewDBConnection()
	if err != nil {
		return fmt.Errorf("failed to init DB connection -> %w", err)
	}

	s.e = s.Route()
	s.teardown = func() {
		if err := s.db.Close(); err != nil {
			log.Printf("failed to close DB -> %v", err)
		}
	}
	return nil
}

func (s *Server) Run(port int) {
	// Start server
	err := s.e.Start(fmt.Sprintf(":%d", port))
	s.teardown()
	s.e.Logger.Fatal(err)
}

func (s *Server) Route() *echo.Echo {
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

	// Routes
	v1 := e.Group("/v1")
	v1.GET("/ping", handler.Ping)

	return e
}
