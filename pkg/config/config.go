package config

import (
	"github.com/spf13/viper"
	"log"
)

var LocalConfig *Config

type Config struct {
	DBUser string `mapstructure:"DBUSER"`
	DBPass string `mapstructure:"DBPASS"`
	DBIP   string `mapstructure:"DBIP"`
	DBName string `mapstructure:"DBNAME"`
	Port   string `mapstructure:"PORT"`
}

func initConfig() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file ", err)
	}
	var config *Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error reading env file", err)
	}
	return config
}

func SetConfig() {
	LocalConfig = initConfig()
}
