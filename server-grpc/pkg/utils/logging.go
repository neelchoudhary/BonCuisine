package utils

import (
	"log"
)

// LogIfFatalAndExit if error, logs and exits
func LogIfFatalAndExit(err error, msgs ...string) {
	if err != nil {
		var message string
		for _, msg := range msgs {
			message += msg + " "
		}
		log.Fatal(message + ": " + err.Error())
	}
}
