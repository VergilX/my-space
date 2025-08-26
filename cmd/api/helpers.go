package main

import (
	"encoding/json"
	"maps"
	"net/http"
)

// Mainly takes data to convert to JSON
func (app *application) writeToJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	stream, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	stream = append(stream, '\n')

	// Add headers to writer header
	maps.Copy(w.Header(), headers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(stream)

	return nil
}
