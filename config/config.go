package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const (
	CONFIG_FILE = "bob.yml"
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
}

func New() (*Configuration, bool, error) {
	configFileExists := validateExistence()

	if configFileExists {
		cfg, err := parseConfigFile()

		return cfg, configFileExists, err
	}

	return nil, configFileExists, nil
}

func validateExistence() bool {
	_, err := os.Stat(CONFIG_FILE)

	isConfigFileInvalid := os.IsNotExist(err)

	return !isConfigFileInvalid
}

func parseConfigFile() (*Configuration, error) {
	yamlData, err := os.ReadFile(CONFIG_FILE)
	if err != nil {
		return &Configuration{}, err
	}

	var config Configuration
	if err := yaml.Unmarshal(yamlData, &config); err != nil {
		return &Configuration{}, err
	}

	return &config, nil
}
