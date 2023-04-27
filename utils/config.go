package utils

import "github.com/spf13/viper"

// import "github.com/spf13/viper"

type Config struct {
	// Default settings
	ADDR             string `mapstructure:"ADDR"`
	SECRET_KEY_TOKEN string `mapstructure:"SECRET_KEY_TOKEN"`
	MODE             string `mapstructure:"MODE"`
	ENABLE_CACHE     bool   `mapstructure:"ENABLE_CACHE"`

	// Development Postgres Database
	DEVELOPMENT_POSTGRES_HOST     string `mapstructure:"DEVELOPMENT_POSTGRES_HOST"`
	DEVELOPMENT_POSTGRES_USERNAME string `mapstructure:"DEVELOPMENT_POSTGRES_USERNAME"`
	DEVELOPMENT_POSTGRES_PASSWORD string `mapstructure:"DEVELOPMENT_POSTGRES_PASSWORD"`
	DEVELOPMENT_POSTGRES_DBNAME   string `mapstructure:"DEVELOPMENT_POSTGRES_DBNAME"`

	// Development Redis Database
	DEVELOPMENT_REDIS_DATABASE_USERNAME string `mapstructure:"DEVELOPMENT_REDIS_DATABASE_USERNAME"`
	DEVELOPMENT_REDIS_DATABASE_HOST     string `mapstructure:"DEVELOPMENT_REDIS_DATABASE_HOST"`
	DEVELOPMENT_REDIS_DATABASE_PASSWORD string `mapstructure:"DEVELOPMENT_REDIS_DATABASE_PASSWORD"`

	// Production Postgres Database
	PRODUCTION_POSTGRES_HOST     string `mapstructure:"PRODUCTION_POSTGRES_HOST"`
	PRODUCTION_POSTGRES_USERNAME string `mapstructure:"PRODUCTION_POSTGRES_USERNAME"`
	PRODUCTION_POSTGRES_PASSWORD string `mapstructure:"PRODUCTION_POSTGRES_PASSWORD"`
	PRODUCTION_POSTGRES_DBNAME   string `mapstructure:"PRODUCTION_POSTGRES_DBNAME"`

	// Production Redis Database
	PRODUCTION_REDIS_DATABASE_HOST     string `mapstructure:"PRODUCTION_REDIS_DATABASE_HOST"`
	PRODUCTION_REDIS_DATABASE_USERNAME string `mapstructure:"PRODUCTION_REDIS_DATABASE_USERNAME"`
	PRODUCTION_REDIS_DATABASE_PASSWORD string `mapstructure:"PRODUCTION_REDIS_DATABASE_PASSWORD"`
	MYSQL_PRODUCTION_DATABASE          string `mapstructure:"MYSQL_PRODUCTION_DATABASE"`
}

var AppConfig Config

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

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
