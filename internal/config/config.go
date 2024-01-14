package config

import (
	"github.com/chapa-ai/application-gaspar/internal/util"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port string `envconfig:"PORT" default:":9997"`
}

func GetConfig() (*Config, error) {
	var _config Config
	if err := envconfig.Process("", &_config); err != nil {
		util.LogFatalf("failed load configuration -- %s", err.Error())
		return nil, err
	}
	return &_config, nil
}
