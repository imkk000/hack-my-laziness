package server

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Routes []RouteConfig `yaml:"routes"`
}

type RouteConfig struct {
	Method   string          `yaml:"method"`
	Path     string          `yaml:"path"`
	Response ResponseConfig  `yaml:"response"`
}

type ResponseConfig struct {
	Status  int               `yaml:"status"`
	Headers map[string]string `yaml:"headers"`
	Body    string            `yaml:"body"`
	File    string            `yaml:"file"`
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &cfg, nil
}
