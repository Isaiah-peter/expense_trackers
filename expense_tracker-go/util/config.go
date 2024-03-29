package util

import "github.com/spf13/viper"

type Config struct {
	DBdriver string `mapstructure:"DB_DRIVER"`
	DBsource string `mapstructure:"DB_SOURCE"`
	Serverddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}