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

// Send response with custom data stucture
func Send(w http.ResponseWriter, statusCode int, data interface{}) {
	// response header settings should be set before status code
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(data)
}

// SendData response with success data
func SendData(w http.ResponseWriter, data interface{}) {
	success := Success{
		Status: StatusSuccess,
		Data:   data,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(success)
}

// SendError response with custom status code and error message
func SendError(w http.ResponseWriter, code int, message string) {
	error := Error{
		Status:  StatusError,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusUnauthorized)

	json.NewEncoder(w).Encode(error)
}

// ReqParamError notifies an error of request param form
func ReqParamError(w http.ResponseWriter) {
	error := Error{
		Status:  StatusError,
		Message: "Request content type error! Your request's Content-Type is not set to application/x-www-form-urlencoded. Set it and try again.",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(error)
}
