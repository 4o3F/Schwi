package database

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var (
	RecaptchaKey string
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

	fmt.Println("Init Config")
}
