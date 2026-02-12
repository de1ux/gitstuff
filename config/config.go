package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	AlwaysIgnoreWorktrees bool `json:"alwaysIgnoreWorktrees"`
}

var globalConfig *Config

// Load reads the config file from ~/.config/gitstuff.config
func Load() (*Config, error) {
	if globalConfig != nil {
		return globalConfig, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(homeDir, ".config", "gitstuff.config")

	// Create config directory if it doesn't exist
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, err
	}

	// Initialize config file if it doesn't exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		defaultConfig := &Config{
			AlwaysIgnoreWorktrees: false,
		}
		if err := Save(configPath, defaultConfig); err != nil {
			return nil, err
		}
		globalConfig = defaultConfig
		return defaultConfig, nil
	}

	// Read existing config
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	globalConfig = &cfg
	return &cfg, nil
}

// Save writes the config to the specified path
func Save(path string, cfg *Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// Get returns the loaded global config
func Get() *Config {
	if globalConfig == nil {
		// Try to load if not already loaded
		cfg, err := Load()
		if err != nil {
			// Return default config if load fails
			return &Config{
				AlwaysIgnoreWorktrees: false,
			}
		}
		return cfg
	}
	return globalConfig
}
