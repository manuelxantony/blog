package main

import (
	"errors"

	"github.com/manuelxantony/blog/config"
	"github.com/spf13/viper"
)

func LoadConfig() *config.Configuration {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		panic(errors.New("failed to read configuration"))
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		panic(errors.New("unable to decode configuration into struct"))
	}

	return &configuration

}
