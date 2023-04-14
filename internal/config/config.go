package config

import (
	"os"
)

type Config struct {
	ApiHost      string
	ApiPort      int
	KafkaHost    string
	KafkaGroupID string
}

func GetConfig() *Config {
	return &Config{
		ApiHost:      os.Getenv("API_HOST"),
		ApiPort:      1323,
		KafkaHost:    os.Getenv("KAFKA_HOST"),
		KafkaGroupID: os.Getenv("KAFKA_GROUP_ID"),
	}
}
