package utils

import (
	"github.com/getsentry/sentry-go"
	"log"
)

func ErrorHandler(err error) bool {
	if err != nil {
		sentry.CaptureException(err)
		log.Println(err)
		return true
	}
	return false
}
