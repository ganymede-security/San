// Package for formatting network logs

package netlog

import (

)

// Standard log format is defined by the ServerLogEntry struct.
// Type JanusLogger wraps the Log with the Janus Logger for formatting
type Logger struct {
	Logger log.Logger
	Log    *NetLogger
}

func NewLogger(error) {
	return
}
