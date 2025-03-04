package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	Info info `json:"info"`
}

type info struct {
	Version string `json:"productVersion"`
}

//go:embed wails.json
var configData []byte

var appInfo info

func init() {
	var cfg config

	err := json.Unmarshal(configData, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	appInfo = cfg.Info
}

type Config struct {
	LastVaultPath string `json:"lastVaultPath"`
}

func getConfigPath() (string, error) {
	// Get user's config directory
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	// Create app config directory if it doesn't exist
	appConfigDir := filepath.Join(configDir, "avdu")
	if err := os.MkdirAll(appConfigDir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(appConfigDir, "config.json"), nil
}

func loadConfig() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	// If config file doesn't exist, return default config
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &Config{}, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func saveConfig(config *Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}
