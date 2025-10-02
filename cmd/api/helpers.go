package main

import (
	"encoding/json"
	"maps"
	"net/http"
)

var envelope map[string]any

func (app *application) decodeJSON(w http.ResponseWriter, r *http.Request, dst any, disallowUnknownFields bool) error {
	dec := json.NewDecoder(r.Body)

	// unknown fields
	if disallowUnknownFields {
		dec.DisallowUnknownFields()
	}

	err := dec.Decode(&dst)
	if err != nil {
		return err
	}

	return nil
}

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
