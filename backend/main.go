package main

import (
	"flag"
	"log"

	"github.com/tetsuzawa/web-spat/interfaces/server"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 1991, "Web server port")
	flag.Parse()

	s := server.NewServer()
	if err := s.Init(); err != nil {
		log.Fatal(err)
	}
	s.Run(port)
}
