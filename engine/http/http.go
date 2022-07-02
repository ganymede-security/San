package http

import (
	"crypto/tls"
	"net/http"

	"github.com/ganymede-security/san/engine"
)

// The Engine interface is fulfilled by the ListenAndHandle and
// GracefullyShutdown Functions
type HTTPEngine struct {
	// Specifies TLS Configuration
	TLSConfig *tls.Config
	// The driver to use
	Driver engine.Driver
}

type Driver interface {
	// ListenAndServe listens on the specified TCP port and responds
	// with the specified handler. TLS settings can be specified with the certificate
	// and key options. If key and certificate are nil TLS is not used.
	ListenAndHandle()

	GracefullyShutdown()
}

type Handler interface {
	ServeHTTP(w http.ResponseWriter, req *http.Request)
	LogEntry()
}

func () New() *HTTPEngine {
	srv := http.Server
}

// Function StartServer starts the Proxy server. Requires a string defining the port to run on.
func StartEngine(port string) {
	
	http.Handler

	router := http.DefaultServeMux

	router.Handle("/ping")
	// Start Server
	http.HandleFunc("/")

	http.ListenAndServe(port, nil)

}