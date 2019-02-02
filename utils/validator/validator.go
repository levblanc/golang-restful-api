package validator

import "net/http"

// ValidContentType checks whether the request content type is"application/x-www-form-urlencoded"
func ValidContentType(req *http.Request) bool {
	return req.Header.Get("Content-Type") == "application/x-www-form-urlencoded"
}
