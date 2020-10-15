package utils

import (
	"encoding/json"
	"net/http"
)

func GetEnv(env, fallback string) string {
	curr :=  (env)
	if curr != "" {
		return curr
	}
	return fallback
}

func JSONResponse(w http.ResponseWriter, code int, output interface{}) {
	// Convert our interface to JSON
	response, _ := json.Marshal(output)
	// Set the content type to json for browsers
	w.Header().Set("Content-Type", "application/json")
	// Our response code
	w.WriteHeader(code)

	w.Write(response)
}