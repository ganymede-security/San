// Package engine is the engine for running the Proxy server. Routes defined in the
// are implemented in this package.
package engine

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"

	//"github.com/ganymede-security/janus/log"
)

// Defines the methods for the Engine
type Engine interface {
	// ListenAndServe listens on the specified TCP port and responds
	// with the specified handler. TLS settings can be specified with the certificate
	// and key options. If key and certificate are nil TLS is not used.
	ListenAndServe(ctx context.Context, string, certificate string, key string, h http.Handler) error
	// Function for just listening but not responding
	Listen(ctx context.Context)
	// The method for shutting down without interrupting server processes
	GracefullyShutdown(ctx context.Context)
}

// Struct Driver Implements the Engine interface
type Driver struct {

}

// Struct representing the request payload
type requestPayload struct {
	urlRoute string
}

// Function StartServer starts the Proxy server. Requires a string defining the port to run on.
func StartServer(port string) {
	// Start Server
	http.HandleFunc("/", HandleAndRedirect)
	
	http.ListenAndServe(port, nil)

}

func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {

	// Parse the url
	url, err := url.Parse(target)
	if err != nil {
		//log.Print("Error parsing URL: ", err)
	}
	//log.Print(url)
	// Create the reverse Proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update headers for SSL Redirection
	req.URL.Host = url.Host
	//log.Print("Url Host is: ", req.URL.Host)

	req.URL.Scheme = url.Scheme
	//log.Print("Url Scheme is: ", req.URL.Scheme)

	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host
	//log.Print("Request Host is: ", req.Host)

	// ServeHttp is non blocking and uses a goroutine
	proxy.ServeHTTP(res, req)
}

// HandleFunc. Given a request send it to the appropriate url
func HandleAndRedirect(res http.ResponseWriter, req *http.Request) {
	requestPayload := parseRequestBody(req)

	logRequest(requestPayload, url)

	serveReverseProxy(url, res, req)
}
