package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress 	string `mapstructure:"SERVER_ADDRESS"`
}

var config Config

func LoadConfig(path string) (c Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()

	if err = viper.ReadInConfig(); err != nil {
		log.Fatal("env Read Error : &w", err)
	}

	if err = viper.Unmarshal(&config); err != nil {
		log.Fatal("env Marshal Error : &w", err)
	}

	return config, err
}