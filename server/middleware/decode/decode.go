// Package decode is for specifying decoders for XML, JSON, and YAML
// requests as as well as for decoding proxy configuration files
package decode

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type XMLDecoder struct {
}

type JsonDecoder struct {
}

// The engine for interfacing with YAML Config files
type YamlDecoder struct {
}

// Creates a json decoder for the incoming request body
func requestBodyDecoder(request *http.Request) (*json.Decoder, error) {
	// Read the body to a buffer
	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("Error decoding payload body: %v", err)
		return nil, err
	}

	request.Body = io.NopCloser(bytes.NewBuffer(body))

	return json.NewDecoder(io.NopCloser(bytes.NewBuffer(body))), nil
}

// Parses the request body using the specified decoder
func parseRequestBody(req *http.Request) requestPayload {
	decoder, err := requestBodyDecoder(req)
	if err != nil {
		log.Panic("Error creating json body decoder: ", err)
	}
	var requestPayload requestPayload

	decoder.Decode(&requestPayload)

	return requestPayload
}