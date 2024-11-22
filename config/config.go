package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const (
	CONFIG_FILE = "gnt-config.yml"
)

type Configuration struct {
	BaseConfiguration BaseConfiguration `yaml:"base"`
}

type BaseConfigurationRoot struct {
	BaseConfiguration BaseConfiguration `yaml:"base"`
}

type BaseConfiguration struct {
	ProjectName string `yaml:"project_name"`
	BaseURL     string `yaml:"base_url"`
	RootDir     string `yaml:"root_dir"`
	SourceDir   string `yaml:"source_dir"`
}

func New() (Configuration, error) {
	yamlData, err := os.ReadFile(CONFIG_FILE)
	if err != nil {
		return Configuration{}, err
	}

	var config Configuration
	if err := yaml.Unmarshal(yamlData, &config); err != nil {
		return Configuration{}, err
	}

	return config, nil
}
