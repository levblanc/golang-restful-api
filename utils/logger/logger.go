package logger

import "log"

// Fatal logs fatal error
func Fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
