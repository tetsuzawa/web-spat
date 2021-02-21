package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tetsuzawa/web-spat/config"
	"github.com/tetsuzawa/web-spat/interfaces/server"
	"github.com/tetsuzawa/web-spat/interfaces/server/handler"
	openapi "github.com/tetsuzawa/web-spat/interfaces/server/openapi"
)

func init() {
	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)
}

func main() {
	var port int
	flag.IntVar(&port, "port", 1991, "Web server port")
	flag.Parse()

	// connect to DB
	db, err := config.NewDBConnection()
	defer db.Close()
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to init DB connection -> %w", err))
	}

	// routing
	e := server.NewRouter()
	h := handler.NewIntegratedHandler(
		*handler.NewExperimentsHandler(),
		*handler.NewUtilHandler(),
	)
	openapi.RegisterHandlersWithBaseURL(e, h, "/v1")

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
