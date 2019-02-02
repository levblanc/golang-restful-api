package response

import (
	"encoding/json"
	"net/http"
)

// status texts
const (
	StatusSuccess = "success"
	StatusError   = "error"
)

// Success is the data structure of success responses
type Success struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// Error is the data structure of error responses
type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Send writes encoded json data to the response writer
func Send(w http.ResponseWriter, statusCode int, data interface{}) {
	// response header settings should be set before status code
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(data)
}
