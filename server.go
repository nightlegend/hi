package main

import (
	"log"

	"github.com/nightlegend/hi/core"
)

func main() {
	log.Println("starting socket server...")
	core.SocketServer()
}
