package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	ServerPort string `envconfig:"SERVER_PORT"`
	DBString   string `envconfig:"DB_STRING"`
}

func ReadConfig() (*Config, error) {
	var config Config

	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
