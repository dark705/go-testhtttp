package helpers

import (
	"log"
)

func FailOnError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}
}
