package utils

import "github.com/spf13/viper"

type Config struct {
	DATA_SOURCE_NAME string `mapstructure:"DATA_SOURCE_NAME"`
	ADDR             string `mapstructure:"ADDR"`
	SECRET_KEY_TOKEN string `mapstructure:"SECRET_KEY_TOKEN"`
	HOST             string `mapstructure:"HOST"`
	PORT             string `mapstructure:"PORT"`
	PASSWD           string `mapstructure:"PASSWD"`
	EMAIL            string `mapstructure:"EMAIL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
