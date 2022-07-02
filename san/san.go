package main

import (
	"github.com/ganymede-security/san/engine"
	"github.com/ganymede-security/san/router"
	"github.com/ganymede-security/san/server"
)

// The SanServer struct combines the Server, Engine, and Router into a single server
type Server struct {
	server server.Server

	engine engine.Engine

	router router.Router
}

func main() {

}

func CombineServer(server.Server, engine.Engine, router.Router) {

}

func NewDefault(server.Server, engine.Engine, router.Router) {

}