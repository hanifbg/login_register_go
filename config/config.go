package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var defaultConfig *viper.Viper

func init() {
	defaultConfig = readViperConfig()
}

func readViperConfig() *viper.Viper {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName("app.config")
	v.SetConfigType("json")
	err := v.ReadInConfig()
	if err == nil {
		fmt.Printf("Using config file: %s \n\n", v.ConfigFileUsed())
	}

	return v
}

func ProviderConfig() *viper.Viper {
	return defaultConfig
}
