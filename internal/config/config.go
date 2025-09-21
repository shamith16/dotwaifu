package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DetectedShell   string `yaml:"detected_shell"`
	PreferredEditor string `yaml:"preferred_editor"`
	InitBasic       bool   `yaml:"init_basic"`
	CreateExamples  bool   `yaml:"create_examples"`
}

func GetConfigDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "dotwaifu")
}

func GetConfigPath() string {
	return filepath.Join(GetConfigDir(), "config.yaml")
}

func Load() (*Config, error) {
	configPath := GetConfigPath()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &Config{}, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	return &config, err
}

func (c *Config) Save() error {
	configDir := GetConfigDir()
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(GetConfigPath(), data, 0644)
}