package utils

import "github.com/spf13/viper"

type Config struct {
	// Default settings
	ADDR             string `mapstructure:"ADDR"`
	SECRET_KEY_TOKEN string `mapstructure:"SECRET_KEY_TOKEN"`
	MODE             string `mapstructure:"MODE"`
	ENABLE_CACHE     bool   `mapstructure:"ENABLE_CACHE"`
	EXPIRES_AT       int    `mapstructure:"EXPIRES_AT"`

	// Postgres Database
	POSTGRES_HOST     string `mapstructure:"POSTGRES_HOST"`
	POSTGRES_USER     string `mapstructure:"POSTGRES_USER"`
	POSTGRES_PASSWORD string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRES_DB       string `mapstructure:"POSTGRES_DB"`

	// Redis Database
	REDIS_USERNAME string `mapstructure:"REDIS_USERNAME"`
	REDIS_HOST     string `mapstructure:"REDIS_HOST"`
	REDIS_PASSWORD string `mapstructure:"REDIS_PASSWORD"`
}

var AppConfig *Config

func init() {
	viper.AutomaticEnv()

	AppConfig = &Config{
		ADDR:             viper.GetString("ADDR"),
		SECRET_KEY_TOKEN: viper.GetString("SECRET_KEY_TOKEN"),
		MODE:             viper.GetString("MODE"),
		ENABLE_CACHE:     viper.GetBool("ENABLE_CACHE"),
		EXPIRES_AT:       viper.GetInt("EXPIRES_AT"),

		// Docker Postgres Database
		POSTGRES_HOST:     viper.GetString("POSTGRES_HOST"),
		POSTGRES_USER:     viper.GetString("POSTGRES_USER"),
		POSTGRES_PASSWORD: viper.GetString("POSTGRES_PASSWORD"),
		POSTGRES_DB:       viper.GetString("POSTGRES_DB"),
		REDIS_USERNAME:    viper.GetString("REDIS_USERNAME"),
		REDIS_HOST:        viper.GetString("REDIS_HOST"),
		REDIS_PASSWORD:    viper.GetString("REDIS_PASSWORD"),
	}
}
