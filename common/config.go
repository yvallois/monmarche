package common

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"reflect"
	"strings"
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

func InitConf() Config {
	_ = godotenv.Load()
	Conf = &Config{
		DatabaseName:       os.Getenv("DATABASE_NAME"),
		DatabaseUser:       os.Getenv("DATABASE_USER"),
		DatabasePassword:   os.Getenv("DATABASE_PASSWORD"),
		DatabaseHost:       os.Getenv("DATABASE_HOST"),
		DatabasePort:       os.Getenv("DATABASE_PORT"),
		RabbitMQExchange:   os.Getenv("RABBITMQ_EXCHANGE"),
		RabbitMQRoutingKey: os.Getenv("RABBITMQ_ROUTING_KEY"),
		RabbitMQQueue:      os.Getenv("RABBITMQ_QUEUE"),
		RabbitMQURI:        os.Getenv("RABBITMQ_URI"),
	}

	var emptyVariables []string
	v := reflect.ValueOf(Conf)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == "" {
			emptyVariables = append(emptyVariables, typeOfS.Field(i).Name)
		}
	}
	if len(emptyVariables) != 0 {
		log.Fatalf("the following environment variables are missing: %s", strings.Join(emptyVariables, ", "))
	}
	return *Conf
}
