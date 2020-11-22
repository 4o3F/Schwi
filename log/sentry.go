package log

import (
	"fmt"
	"github.com/CardinalDevLab/Schwi-Backend/database"
	"github.com/getsentry/sentry-go"
)

func SentryInit()  {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: database.SentryDSN,
	})
	if err != nil {
		fmt.Println("Sentry Init: %s", err)
	} else {
		fmt.Println("Init Sentry")
	}
}