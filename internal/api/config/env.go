package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/rs/zerolog"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/gateway_provider_enum"
)

type DBConfig struct {
	User     string
	Port     string
	Password string
	Address  string
	Host     string
	Name     string
}

type JWTConfig struct {
	Secret              string
	ExpirationInSeconds int64
}

type LoggerConfig struct {
	LogLevel   int
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	Output     *os.File
}

type StripeConfig struct {
	Key        string
	ProviderID int32
}

type GatewaysConfig struct {
	Stripe StripeConfig
}

type ProvidersConfig struct {
	Gateways GatewaysConfig
}

type KafkaConfig struct {
	Host      string
	Port      int
	Host_Port string
}

type Config struct {
	ApiToken   string
	AppEnv     string
	PublicHost string
	CookieName string
	Port       string
	DB         DBConfig
	JWT        JWTConfig
	Logger     LoggerConfig
	Providers  ProvidersConfig
	Kafka      KafkaConfig
}

var (
	Envs     Config
	initOnce sync.Once
)

func Init() {
	initOnce.Do(func() {
		Envs = Config{
			ApiToken:   getEnv("API_TOKEN", "ToYaaRUiza7cYAMzD+Pk2ha9N2Xn3rwMpuhd2JVEQ/Usdbte6kFaIOoIWm6qXgOXt0qYZo3uHTvecySPo4p5zQ=="),
			AppEnv:     getEnv("APP_ENV", "development"),
			PublicHost: fmt.Sprintf("%s:%s", getEnv("PUBLIC_HOST", "http://localhost"), getEnv("PORT", "8080")),
			Port:       getEnv("PORT", "8080"),
			CookieName: getEnv("COOKIE_NAME", "ticket_sale_jwt"),
			DB: DBConfig{
				User:     getEnv("DB_USER", "root"),
				Password: getEnv("DB_PASSWORD", "root"),
				Port:     getEnv("DB_PORT", "5432"),
				Name:     getEnv("DB_NAME", "ticket_sale"),
				Host:     getEnv("DB_HOST", "localhost"),
				Address:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "5432")),
			},
			JWT: JWTConfig{
				Secret:              getEnv("JWT_SECRET", "tvoxJ4l0kbR8jIdLcH+ywmrb0wXa7+Ob7z+m6pKmnpME3ZOF4A0Ma57JUgeceG4o0eSOKOnY4fJPtuoaGfy1tw=="),
				ExpirationInSeconds: getEnvAsInt64("JWT_EXPIRATION", 3600*24),
			},
			Logger: LoggerConfig{
				LogLevel:   getEnvAsInt("LOG_LEVEL", int(zerolog.InfoLevel)),
				Filename:   getEnv("LOG_FILE_NAME", "internal/api/log_files/api.log"),
				MaxSize:    getEnvAsInt("LOG_MAX_SIZE", 5),
				MaxBackups: getEnvAsInt("LOG_MAX_BACKUPS", 5),
				MaxAge:     getEnvAsInt("LOG_MAX_AGE", 30),
				Compress:   getEnvAsBool("LOG_COMPRESS", true),
				Output:     os.Stdout,
			},
			Providers: ProvidersConfig{
				Gateways: GatewaysConfig{
					Stripe: StripeConfig{
						Key:        getEnv("STRIPE_SECRET_KEY", "sk_test_Ho24N7La5CVDtbmpjc377lJI"),
						ProviderID: gateway_provider_enum.STRIPE,
					},
				},
			},
			Kafka: KafkaConfig{
				Host:      getEnv("KAFKA_HOST", "localhost"),
				Port:      getEnvAsInt("KAFKA_PORT", 9092),
				Host_Port: fmt.Sprintf("%s:%d", getEnv("KAFKA_HOST", "localhost"), getEnvAsInt("KAFKA_PORT", 9092)),
			},
		}
	})
}

func getEnv(key string, callback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return callback
}

func getEnvAsInt64(key string, callback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return callback
		}

		return i
	}
	return callback
}

func getEnvAsInt(key string, callback int) int {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return callback
		}

		return i
	}
	return callback
}

func getEnvAsBool(key string, callback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		b, err := strconv.ParseBool(value)
		if err != nil {
			return callback
		}
		return b
	}
	return callback
}
