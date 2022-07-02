package router

import (
	"context"
	"log"
	"net/http"

	"github.com/ganymede-security/san/engine"
)

// Wrapper around a ServeMux
type Router struct {
	serveMux 	*http.ServeMux
	logger 		*log.Logger
	ctx 		context.Context
}

func NewRouter() *Router {
	router := &Router{}
	return router
}

// Function to add a new route
func AddHandleRoute(string, h engine.Driver) {

}

// Function to add a route to be used by a different engine.
func AddEngineRoute() {

}