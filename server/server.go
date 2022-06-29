// Package server defines a customizable server for configuration. It allows the use of a standard
// or custom server, router, and engine as well as customization and integration
// of middleware packages.
package server

import (
	"net/http"
	"time"

	"github.com/ganymede-security/san/server/config"
	"github.com/ganymede-security/san/server/engine"
)

// Struct specifies the requirements for a new server
// when not specified.
type CustomServer struct {
	Server config.ServerConfig
	Router config.RouterConfig
	Engine config.EngineConfig
}

// Function New sets up a new Proxy Server with default
// configuration settings specified in the Config package.
func NewServer(h http.Handler, c *config.EngineConfig) *CustomServer {
	//c = config.NewDefaultServer
	//server := &ProxyServer{
	//	proxyEngine: ,
	//}
	return c
}

// Function to create a new server with default configuration settings
func Default(c *CustomServer) /*ServerConfig*/ {
	router := &http.ServeMux{}
	// Server defaults
	srv := &http.Server{
		Addr:         "localhost:80",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
		//ErrorLog: 		netlog.ProxyLogger,
	}
	engine := &engine.Engine{}
	return &EngineConfig{
		Server: srv,
		Router: http.ServeMux{},
	}
}

/*
// Wrapper function passes the information to the Driver ListenAndServe Interface from the Engine package
func (srv *ProxyServer) ListenAndServe(listenAddr string, certificate string, key string, h http.Handler) error {
	//return srv.proxyEngine == engine.Engine.ListenAndServe()
}

// Function to add a NewRoute to the Proxy server
func (srv *ProxyServer) NewRoute() {

}

func (srv *ProxyServer) GracefullyShutdown() {

}

//function to get the URL based on the input condition. Takes an input of a URL as a string
// and outputs the URL
func getProxyUrl(proxyCondition string) string {
	switch {
	case proxyCondition == "A":
		return aUrl
	case proxyCondition == "B":
		return bUrl
	default:
		return defaultUrl
	}
}

*/
