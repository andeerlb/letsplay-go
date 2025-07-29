package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Env           string
	ServerPort    string
	AuthServerUrl string
	LogLevel      string
	JwtSecret     []byte
}

func LoadConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigType("env")

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	v.SetConfigName(env)
	v.AddConfigPath("./configs")

	// Default values
	v.SetDefault("SERVER_PORT", "8080")
	v.SetDefault("LOG_LEVEL", "info")
	v.SetDefault("AUTH_SERVER_URL", "http://localhost:9000")

	// Try read config file, ignore error if not found
	_ = v.ReadInConfig()

	// Read from environment variables with prefix LETSPLAY_
	v.SetEnvPrefix("LETSPLAY")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	jwtSecret := v.GetString("JWT_SECRET")

	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET cannot be empty")
	}

	cfg := &Config{
		Env:           env,
		ServerPort:    v.GetString("SERVER_PORT"),
		LogLevel:      v.GetString("LOG_LEVEL"),
		AuthServerUrl: v.GetString("AUTH_SERVER_URL"),
		JwtSecret:     []byte(jwtSecret),
	}

	// Basic validation
	if cfg.ServerPort == "" {
		return nil, fmt.Errorf("SERVER_PORT cannot be empty")
	}
	if cfg.LogLevel == "" {
		cfg.LogLevel = "info"
	}

	return cfg, nil
}
