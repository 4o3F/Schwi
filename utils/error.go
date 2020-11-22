package utils

import "github.com/getsentry/sentry-go"

func ErrorHandler(err error) bool {
	if err != nil {
		sentry.CaptureException(err)
		return true
	}
	return false
}
