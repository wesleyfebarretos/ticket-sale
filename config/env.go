package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Config struct {
	ApiToken               string
	JWTSecret              string
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPort                 string
	DBPassword             string
	DBAddress              string
	DBHost                 string
	DBName                 string
	CookieName             string
	JWTExpirationInSeconds int64
}

var (
	Envs     Config
	initOnce sync.Once
)

func Init() {
	initOnce.Do(func() {
		Envs = Config{
			ApiToken:               getEnv("API_TOKEN", "ToYaaRUiza7cYAMzD+Pk2ha9N2Xn3rwMpuhd2JVEQ/Usdbte6kFaIOoIWm6qXgOXt0qYZo3uHTvecySPo4p5zQ=="),
			JWTSecret:              getEnv("JWT_SECRET", "tvoxJ4l0kbR8jIdLcH+ywmrb0wXa7+Ob7z+m6pKmnpME3ZOF4A0Ma57JUgeceG4o0eSOKOnY4fJPtuoaGfy1tw=="),
			JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION", 3600*24),
			PublicHost:             fmt.Sprintf("%s:%s", getEnv("PUBLIC_HOST", "http://localhost"), getEnv("PORT", "8080")),
			Port:                   getEnv("PORT", "8080"),
			DBUser:                 getEnv("DB_USER", "root"),
			DBPassword:             getEnv("DB_PASSWORD", "root"),
			DBPort:                 getEnv("DB_PORT", "5432"),
			DBName:                 getEnv("DB_NAME", "ticket_sale"),
			DBHost:                 getEnv("DB_HOST", "localhost"),
			DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "5432")),
			CookieName:             getEnv("COOKIE_NAME", "ticket_sale_jwt"),
		}
	})
}

func getEnv(key string, callback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return callback
}

func getEnvAsInt(key string, callback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return callback
		}

		return i
	}
	return callback
}
