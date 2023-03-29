package utils

import "github.com/spf13/viper"

// import "github.com/spf13/viper"

type Config struct {
	DATA_SOURCE_NAME string `mapstructure:"DATA_SOURCE_NAME"`
	ADDR             string `mapstructure:"ADDR"`
	SECRET_KEY_TOKEN string `mapstructure:"SECRET_KEY_TOKEN"`
	// HOST             string `mapstructure:"HOST"`
	// PORT             string `mapstructure:"PORT"`
	// PASSWD           string `mapstructure:"PASSWD"`
	// EMAIL            string `mapstructure:"EMAIL"`
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
	// return
}
