package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config") // config.yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("No config file found; using defaults.")
	}

	viper.SetDefault("app.theme", "dark")
}
