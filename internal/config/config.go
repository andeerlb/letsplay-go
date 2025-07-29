package config

import (
	"fmt"
	_ "os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Env           string `mapstructure:"-"`
	ServerPort    string `mapstructure:"server_port"`
	AuthServerUrl string `mapstructure:"auth_server_url"`
	LogLevel      string `mapstructure:"log_level"`
	JwtSecret     []byte `mapstructure:"-"`
	DB            struct {
		Host       string `mapstructure:"host"`
		Port       int    `mapstructure:"port"`
		User       string `mapstructure:"user"`
		Password   string `mapstructure:"password"`
		Name       string `mapstructure:"name"`
		SSLMode    string `mapstructure:"sslmode"`
		SearchPath string `mapstructure:"search_path"`
	} `mapstructure:"db"`
}

func LoadConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")

	files, err := filepath.Glob("configs/*.yaml")
	if err != nil || len(files) == 0 {
		return nil, fmt.Errorf("config file not found in ./configs")
	}

	configPath := files[0]
	envName := strings.TrimSuffix(filepath.Base(configPath), ".yaml")

	v.SetConfigFile(configPath)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	v.SetEnvPrefix("LETSPLAY")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}

	jwtSecret := v.GetString("jwt_secret")
	if jwtSecret == "" {
		return nil, fmt.Errorf("jwt_secret must be set")
	}
	cfg.JwtSecret = []byte(jwtSecret)
	cfg.Env = envName

	if cfg.ServerPort == "" {
		return nil, fmt.Errorf("server_port must be set")
	}
	if cfg.AuthServerUrl == "" {
		return nil, fmt.Errorf("auth_server_url must be set")
	}
	if cfg.LogLevel == "" {
		return nil, fmt.Errorf("log_level must be set")
	}

	return &cfg, nil
}
