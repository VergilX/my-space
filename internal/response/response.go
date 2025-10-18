package response

import (
	"encoding/json"
	"maps"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data any) error {
	return JSONWithHeaders(w, status, data, nil)
}

func JSONWithHeaders(w http.ResponseWriter, status int, data any, headers http.Header) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// adding stuff for terminal
	jsonData = append(jsonData, '\n')

	// adding headers (map)
	maps.Copy(w.Header(), headers)
	// for key, value := range headers {
	// 	w.Header()[key] = value
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	w.Write(jsonData)

	return nil
}
