package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type EnvSource struct {
	Prefix string
}

func (es EnvSource) Load() (Config, error) {
	_ = godotenv.Load("secrets/.env")
	_ = godotenv.Load(".env")

	var c Config
	pErr := envconfig.Process(es.Prefix, &c)
	if pErr != nil {
		return c, fmt.Errorf("envconfig.Process return error: %s", pErr)
	}

	return c, nil
}
