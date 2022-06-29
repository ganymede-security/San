package router

import "net/http"

// Wrapper around a ServeMux
type Router struct {
	serveMux http.ServeMux
}

