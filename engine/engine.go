// Package engine is the engine for running the Proxy server. Routes defined in the
// are implemented in this package.
package engine

import (
	"context"
	"io"
)


type Engine struct {
	// The driver (Handler) that will perform the given action
	Driver Driver

	// The Logger to pass request/response information to
	//Logger log.Logger

	// The Context
	//Ctx 	context.Context
}

// Type Driver defines the minimum requirements for 
// creating a new driver to interface with the Engine.
type Driver struct {
	// Handler is an undefined function containing the
	// action to be performed by the driver.
	Handler string
	// Defines the information to be included in the
	// standard log Entry format for the driver
	//Log(*Log Entry)
	// return/add to context
	//Context(ctx context.Context)
}

// function to declare a new engine 
func NewEngine(d *Driver) *Engine {
	
	return &Engine{
		Driver: "TestDriver",
	},
}