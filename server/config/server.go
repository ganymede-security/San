// Package config is where the server environment variables are defined
package config

import (
	"time"
)

// Struct for configuring Environment variables and setting
// server configuration options
type ServerConfig struct {
	// The port the server should listen on. The zero value
	// is http port 80.
	ListenPort string

	ReadTimeout time.Time

	WriteTimeout time.Time

	IdleTimeout time.Time

	TLSSupported bool
}

// RouterConfig is where the router options are set.
type RouterConfig struct {
	// Specifies which router to use by default. Zero value is an
	// http.ServeMux router.
	DefaultRouter string
}

// EngineConfig sets the Engine configuration options
type EngineConfig struct {
	// Specifies the engine to be used.
	// TODO: Define engines to use (?HTTP/DNS/SNMP?)

	// The only currently supported Engine is HTTP.
	EngineType string
}

// TODO Create separate prod and dev environments
/* Function ConfigureServer takes an argument of a string specifying the environment to use
// when starting the server and a type of ServerConfig specifying the environment variables to set.
// Current Environments include "dev" and "prod"
func ConfigureServerEnv(env string) *EngineConfig {
	environ := strings.ToLower(env)

	switch environ {
	case "dev":
		// TO DO: Allow different log files to be added
		// Create a log in the current working directory
		logPath, err := log.NewLogFile()
		if err != nil {
			panic(err)
		}
		// Set Logger as MultiWriter
		logger := log.New(true)
		logger.Log().Msgf("Log file created at path: %s", logPath)

		// Set env variables
		var c EngineConfig

		return &c

	case "prod":
		// TO DO: Allow different log files to be added
		// Create a log in the current working directory
		logPath, err := log.NewLogFile()
		if err != nil {
			panic(err)
		}
		// Set Logger to write to file
		logger := log.New(true)
		logger.Log().Msgf("Log file created at path: %s", logPath)

		// Set env variables
		var c EngineConfig

		return &c

	default:
		return nil
	}
}
*/
