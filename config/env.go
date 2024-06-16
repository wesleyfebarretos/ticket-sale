package config

import (
	"fmt"
	"os"
)

type Config struct {
	ApiToken   string
	PublicHost string
	Port       string
	DBUser     string
	DBPort     string
	DBPassword string
	DBAddress  string
	DBHost     string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		ApiToken:   getEnv("API_TOKEN", "ToYaaRUiza7cYAMzD+Pk2ha9N2Xn3rwMpuhd2JVEQ/Usdbte6kFaIOoIWm6qXgOXt0qYZo3uHTvecySPo4p5zQ=="),
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "root"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "ticket_sale"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "5432")),
	}
}

func getEnv(key string, callback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return callback
}
