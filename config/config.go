package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ServiceIp      string        `yaml:"SERVICE_IP"`
	ServicePort    int           `yaml:"SERVICE_PORT"`
	IgnoreIp       string        `yaml:"IGNORE_IP"`
	MaxRequest     int           `yaml:"MAX_REQUEST"`
	ExpirationTime time.Duration `yaml:"EXPIRATION_TIME"`
	AutKey         string        `yaml:"AUT_KEY"`
	AutPass        string        `yaml:"AUT_PASS"`
}

// LoadConfig function to load configuration from a YAML file
func LoadConfig(filename string) (*Config, error) {
	// Default values
	defaultConfig := Config{
		ServiceIp:      "localhost",
		ServicePort:    3000,
		IgnoreIp:       "127.0.0.1",
		MaxRequest:     10,
		ExpirationTime: 30 * time.Second,
		AutKey:         "X-API-Key",
		AutPass:        "12345",
	}

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file doesn't exist, return the default configuration
			log.Println("YAML file not found, using default configuration.")
			return &defaultConfig, nil
		}
		return nil, err // Return error for any other issues
	}
	defer file.Close()

	// Decode YAML into a Config struct
	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	// Merge missing fields with default values
	if config.ServiceIp == "" {
		config.ServiceIp = defaultConfig.ServiceIp
	}
	if config.ServicePort == 0 {
		config.ServicePort = defaultConfig.ServicePort
	}
	if config.IgnoreIp == "" {
		config.IgnoreIp = defaultConfig.IgnoreIp
	}
	if config.MaxRequest == 0 {
		config.MaxRequest = defaultConfig.MaxRequest
	}
	if config.ExpirationTime == 0 {
		config.ExpirationTime = defaultConfig.ExpirationTime
	}
	if config.AutKey == "" {
		config.AutKey = defaultConfig.AutKey
	}
	if config.AutPass == "" {
		config.AutPass = defaultConfig.AutPass
	}

	return &config, nil
}
