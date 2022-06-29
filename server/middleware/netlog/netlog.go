// Package netlog specifies a log configuration
// for passing to a logger package and formats
// logs before passing to a logger
package netlog

import (
	"net"
	"net/http"
	"time"
	//"github.com/ganymede-security/janus/log"
)

// NetLogger is a wrapper for the Log method.
// The Log method structures a standard log entry
// for the proxy server
type NetLogger interface {
	Log(*ServerLogEntry)
}

// Struct representing a log entry, contains information from the
// ProxyRequest and ProxyResponse structs along with additional information
type ServerLogEntry struct {
	Request *proxyRequest
	// A pointer to the incoming http request
	/* 	Includes information on:
		- HttpMethod
		- URL
		- Transport Protocol (TCP/UDP)
		- Proto Major/Minor
		- HttpHeaders
	HttpBody Information on:
		- Content Length
		- Destination Host:Port
		- Destination URL
		- Remote Address
		- Request Uri
		- TLS Connection State
	*/
	Response *proxyResponse
	// Time of the incoming request
	RequestTime time.Time
	// TO DO: Implement Context
	//Context string
	Latency time.Duration
	// TO DO: Size of the incoming packet
	//BytesReceived int64
	// TO DO: Size of the outgoing packet
	//BytesSent int64
	// TO DO: The service being requested
	//Service string
}

// The handler struct ensures wraps a standard
// httpHandler to include a logger
type Handler struct {
	log     NetLogger
	handler http.Handler
}

// Function NewProxyHandler creates a new Http Handler that will output
// to a log. It requires an input of a log of type ProxyLogger
// and an HttpHandler.
func NewProxyHandler(log NetLogger, handler http.Handler) *Handler {
	return &Handler{
		log:     log,
		handler: handler,
	}
}

// Specifies the handler then logs the incoming request.
func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	timeOfReq := time.Now()

	// request information
	reqContext := getContext(req)

	hSize := int64(len(req.Header))

	remIp := getRemoteIp(req)
	// Get the local Address if Available
	localAddr := getLocalIp(req)

	requestInfo := &proxyRequest{
		Request:    reqContext,
		Method:     req.Method,
		Url:        req.URL.String(),
		HeaderSize: hSize,
		Agent:      req.UserAgent(),
		Referer:   req.Referer(),
		Protocol:   req.Proto,
		RemoteIp:   remIp,
		LocalIp:    localAddr,
	}

	writer2 := &proxyResponse{w: w}

	latency := time.Since(timeOfReq)

	// After all information has been gathered we combine the
	// information into a single log entry
	logEntry := &ServerLogEntry{
		Request:     requestInfo,
		Response:    writer2,
		RequestTime: timeOfReq,
		Latency:     latency,
	}

	h.log.Log(logEntry)
	// After Logging is complete continue serving Http
	newRequest := new(http.Request)
	
	h.handler.ServeHTTP(writer2, newRequest)

}

// A Struct representing the http Request received
// by the proxy server to be logged
type proxyRequest struct {
	Request    *http.Request
	Method     string
	Url        string
	HeaderSize int64
	Agent      string
	Referer   string
	Protocol   string
	RemoteIp   string
	LocalIp    string
}

// Helper function to copy the context from an Http Request
// and strip the body
func getContext(req *http.Request) *http.Request {
	clonedReq := req.Clone(req.Context())
	// response body is not needed
	clonedReq.Body = nil

	return clonedReq
}

// Helper function to get the IP address of the remote server
// from the Request.RemoteAddr field
func getRemoteIp(req *http.Request) string {
	// port is not needed
	addr, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return ""
	}
	// format the results if not already formatted
	if len(addr) > 0 && addr[0] == '[' {
		return addr[1 : len(addr)-1]
	}
	return addr
}

// Helper function to get the IP address of the local server
// from the Request context field. If the local server was not in the
// context an empty string will be returned
func getLocalIp(req *http.Request) string {
	addrKey := req.Context().
	Value(http.LocalAddrContextKey).
	(net.Addr).String()

	// port is not needed
	addr, _, err := net.SplitHostPort(addrKey)
	if err != nil {
		return ""
	}
	// format the results if not already formatted
	if len(addr) > 0 && addr[0] == '[' {
		return addr[1 : len(addr)-1]
	}
	return addr

}

// A struct representing the http Response information
// that will be logged
type proxyResponse struct {
	w 				http.ResponseWriter
	headerSize 		int64
	resCode  		int
}

func (p *proxyResponse) Header() http.Header {
	return p.w.Header()
}

// Helper function to pass the header from the wrapper to the 
// httprequestwriter
func (p *proxyResponse) WriteHeader(statusCode int) {
	if p.resCode != 0 {
		return
	}
	lenHeader := int64(len(p.w.Header()))

	p.headerSize = lenHeader
	p.w.WriteHeader(statusCode)
	p.resCode = statusCode
}

func (p *proxyResponse) Write(val []byte) (out int, err error) {
	if p.resCode == 0 {
		p.WriteHeader(http.StatusOK)
	}
	out, err = p.w.Write(val)
	return
}