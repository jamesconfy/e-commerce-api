package utils

import "github.com/spf13/viper"

// import "github.com/spf13/viper"

type Config struct {
	DEVELOPMENT_DATABASE string `mapstructure:"DEVELOPMENT_DATABASE"`
	PRODUCTION_DATABASE  string `mapstructure:"PRODUCTION_DATABASE"`
	ADDR                 string `mapstructure:"ADDR"`
	SECRET_KEY_TOKEN     string `mapstructure:"SECRET_KEY_TOKEN"`
	MODE                 string `mapstructure:"MODE"`
}

var AppConfig Config

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		panic(err)
	}
}
