package config

import (
	"log"

	"github.com/spf13/viper"
)

type RedisConfig struct {
	Addr     string `mapstructure:"REDIS_ADDRESS"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	DB       int    `mapstructure:"REDIS_DBNUMBER"`
}

type Config struct {
	Port        string `mapstructure:"PORT"`
	DBUrl       string `mapstructure:"DB_URL"`
	RedisConfig RedisConfig
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./config/env")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	return
}
