package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	URLs                []string `json:"urls"`
	IntervalSeconds     int      `json:"intervalSeconds"`
	NotificationWebhook string   `json:"notificationWebhook,omitempty"`
}

func LoadConfig(filePath string) (*Config, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("Configuration file not found at %s. Please create it.\n", filePath)
		}
		return nil, fmt.Errorf("failed to read config file %s: %w", filePath, err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file %s: %w", filePath, err)
	}

	if cfg.IntervalSeconds <= 0 {
		cfg.IntervalSeconds = 60
	}

	return &cfg, nil
}
