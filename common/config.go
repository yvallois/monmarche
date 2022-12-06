package common

import (
	"github.com/spf13/viper"
	"log"
)

var Conf *Config

type Config struct {
	DatabaseName       string `mapstructure:"DATABASE_NAME"`
	DatabaseUser       string `mapstructure:"DATABASE_USER"`
	DatabasePassword   string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseHost       string `mapstructure:"DATABASE_HOST"`
	DatabasePort       string `mapstructure:"DATABASE_PORT"`
	RabbitMQExchange   string `mapstructure:"RABBITMQ_EXCHANGE"`
	RabbitMQRoutingKey string `mapstructure:"RABBITMQ_ROUTING_KEY"`
	RabbitMQQueue      string `mapstructure:"RABBITMQ_QUEUE"`
	RabbitMQURI        string `mapstructure:"RABBITMQ_URI"`
}

func InitConf(fileName string) Config {
	viper.SetConfigFile(fileName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Could not environment: %v", err)
	}
	err = viper.Unmarshal(&Conf)
	return *Conf
}
