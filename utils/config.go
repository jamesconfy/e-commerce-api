package utils

import "github.com/spf13/viper"

// import "github.com/spf13/viper"

type Config struct {
	// Default settings
	ADDR             string `mapstructure:"ADDR"`
	SECRET_KEY_TOKEN string `mapstructure:"SECRET_KEY_TOKEN"`
	MODE             string `mapstructure:"MODE"`
	ENABLE_CACHE     bool   `mapstructure:"ENABLE_CACHE"`

	// Development
	DEVELOPMENT_DATABASE                string `mapstructure:"DEVELOPMENT_DATABASE"`
	DEVELOPMENT_REDIS_DATABASE          string `mapstructure:"DEVELOPMENT_REDIS_DATABASE"`
	DEVELOPMENT_REDIS_DATABASE_USERNAME string `mapstructure:"DEVELOPMENT_REDIS_DATABASE_USERNAME"`
	DEVELOPMENT_REDIS_DATABASE_HOST     string `mapstructure:"DEVELOPMENT_REDIS_DATABASE_HOST"`
	DEVELOPMENT_REDIS_DATABASE_PASSWORD string `mapstructure:"DEVELOPMENT_REDIS_DATABASE_PASSWORD"`

	// Production
	PRODUCTION_REDIS_DATABASE_HOST     string `mapstructure:"PRODUCTION_REDIS_DATABASE_HOST"`
	PRODUCTION_REDIS_DATABASE_USERNAME string `mapstructure:"PRODUCTION_REDIS_DATABASE_USERNAME"`
	PRODUCTION_REDIS_DATABASE_PASSWORD string `mapstructure:"PRODUCTION_REDIS_DATABASE_PASSWORD"`
	PRODUCTION_DATABASE                string `mapstructure:"PRODUCTION_DATABASE"`
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
