package utils

import "log"

// LogFatal logs a fatal error
func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
