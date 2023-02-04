package main

import (
	"fmt"
	"github.com/adrg/xdg"
	"gopkg.in/yaml.v3"
	"os"
	"path"
)

type Config struct {
	Token string `yaml:"token"`

	OutputFormat string `yaml:"output_format"`
	TableFormat string `yaml:"table_format"`
}

const ConfigFileName = "config.yaml"

func getConfigFile() (*os.File, error) {
	configFilePath, err := xdg.ConfigFile(path.Join(AppName, ConfigFileName))
	if err != nil {
		return nil, fmt.Errorf("unable to get config file path: %v", err)
	}

	// Check if file exists
	_, err = os.Stat(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Config file doesn't exist, creating it
			return createDefaultConfig(configFilePath)
		}
		return nil, fmt.Errorf("unable to stat file: %v", err)
	}

	return os.Open(configFilePath)
}

func createDefaultConfig(filePath string) (*os.File, error) {
	f, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to create file: %v", err)
	}

	config := Config{}
	enc := yaml.NewEncoder(f)
	err = enc.Encode(&config)

	f.Close()
	return os.Open(filePath)
}

func ParseConfig() (*Config, error) {
	f, err := getConfigFile()
	if err != nil {
		return nil, err
	}

	var config Config
	dec := yaml.NewDecoder(f)
	err = dec.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("unable to decode YAML: %v", err)
	}
	return &config, nil
}