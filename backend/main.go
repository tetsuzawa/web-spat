package main

import (
	"flag"
	"log"
	"os"

	"github.com/tetsuzawa/web-spat/interfaces/server"
)

func init() {
	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)
}

func main() {
	var port int
	flag.IntVar(&port, "port", 1991, "Web server port")
	flag.Parse()
	server.Run(port)
}
