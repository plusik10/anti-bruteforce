package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config - .
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		PG   `yaml:"postgres"`
		BL   `yaml:"bucket_limit"`
	}
	// App - .
	App struct {
		Name    string `env-required:"false" yaml:"name" env:"APP_NAME"`
		Version string `env-required:"false" yaml:"version" env:"APP_VERSION"`
		Debug   bool   `env-required:"true" yaml:"debug" env:"debug"`
	}
	// HTTP - .
	HTTP struct {
		Port        string        `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		Timeout     time.Duration `yaml:"timeout" env-default:"4s" env:"HTTP_TIMEOUT"`
		IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s" env:"HTTP_IDLE_TIMEOUT"`
	}

	// LOG - .
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env-required:"true" yaml:"url"        env:"PG_URL"`
	}

	BL struct {
		LoginLimit    int `env-required:"true" env-default:"3" yaml:"login_limit" env:"LOGIN_LIMIT"`
		PasswordLimit int `env-required:"true" env-default:"10" yaml:"password_limit" env:"PASSWORD"`
		IPLimit       int `env-required:"true" env-default:"1000" yaml:"ip_limit" env:"IP_LIMIT"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
