package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppName string `mapstructure:"APP_NAME"`
	Env     string `mapstructure:"ENV"`
	Port    string `mapstructure:"PORT"`
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
