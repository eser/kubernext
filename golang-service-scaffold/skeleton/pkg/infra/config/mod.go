package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Env string `mapstructure:"ENV"`

	AppName           string `mapstructure:"APP_NAME"`
	Port              string `mapstructure:"PORT"`
	CorsOrigin        string `mapstructure:"CORS_ORIGIN"`
	CorsStrictHeaders bool   `mapstructure:"CORS_STRICT_HEADERS"`
	RedisAddr         string `mapstructure:"REDIS_ADDR"`
	RedisPwd          string `mapstructure:"REDIS_PWD"`
	RedisConnCheck    bool   `mapstructure:"REDIS_CONN_CHECK"`
	MongoAddr         string `mapstructure:"MONGO_ADDR"`
	MongoMaxPoolSize  uint64 `mapstructure:"MONGO_MAX_POOL_SIZE"`
	MongoConnCheck    bool   `mapstructure:"MONGO_CONN_CHECK"`
	MongoConnTimeout  uint64 `mapstructure:"MONGO_CONN_TIMEOUT"`
}

func loadConfig(filename string) error {
	viper.SetConfigName(filename)
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	err := viper.MergeInConfig()
	return err
}

func loadEnv() {
	env := strings.ToLower(os.Getenv("ENV"))
	if env == "" {
		env = "development"
	}

	viper.AutomaticEnv()

	loadConfig(".env")

	loadConfig(".env." + env)

	if env != "test" {
		loadConfig(".env.local")
	}

	loadConfig(".env." + env + ".local")

	viper.Set("ENV", env)
}

func NewConfig() (*Config, error) {
	config := &Config{}

	viper.SetDefault("APP_NAME", "go-service")
	viper.SetDefault("PORT", "8080")

	loadEnv()

	err := viper.Unmarshal(&config)

	return config, err
}
