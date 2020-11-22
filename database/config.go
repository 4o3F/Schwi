package database

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var (
	RecaptchaKey string
	FrontendDomain string
	CORSDomain []string
	UseTLS bool
	CookieDomain string
	SentryDSN string
)

func LoadConfig()  {
	viper.SetConfigName("config")
	viper.AddConfigPath("./data/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
	RecaptchaKey = viper.GetString("recaptchakey")
	FrontendDomain = viper.GetString("frontenddomain")
	CORSDomain = viper.GetStringSlice("corsdomain")
	UseTLS = viper.GetBool("usetls")
	CookieDomain = viper.GetString("cookiedomain")
	SentryDSN = viper.GetString("sentrydsn")

	fmt.Println("Init Config")
}
