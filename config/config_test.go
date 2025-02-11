package config

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Test case: YAML file does not exist
	t.Run("FileNotFound", func(t *testing.T) {
		config, err := LoadConfig("nonexistent.yml")
		assert.NoError(t, err)
		assert.NotNil(t, config)
		assert.Equal(t, "localhost", config.ServiceIp)
		assert.Equal(t, 3000, config.ServicePort)
		assert.Equal(t, "127.0.0.1", config.IgnoreIp)
		assert.Equal(t, 10, config.MaxRequest)
		assert.Equal(t, 30*time.Second, config.ExpirationTime)
		assert.Equal(t, "X-API-Key", config.AutKey)
		assert.Equal(t, "12345", config.AutPass)
	})

	// Test case: YAML file with complete configuration
	t.Run("CompleteConfig", func(t *testing.T) {
		yamlContent := `
SERVICE_IP: "192.168.1.1"
SERVICE_PORT: 8080
IGNORE_IP: "192.168.0.1"
MAX_REQUEST: 20
EXPIRATION_TIME: 6s
AUT_KEY: "Custom-API-Key"
AUT_PASS: "custom123"
`
		filename := "test_config.yml"
		err := ioutil.WriteFile(filename, []byte(yamlContent), 0644)
		assert.NoError(t, err)
		defer os.Remove(filename)

		config, err := LoadConfig(filename)
		assert.NoError(t, err)
		assert.NotNil(t, config)
		assert.Equal(t, "192.168.1.1", config.ServiceIp)
		assert.Equal(t, 8080, config.ServicePort)
		assert.Equal(t, "192.168.0.1", config.IgnoreIp)
		assert.Equal(t, 20, config.MaxRequest)
		assert.Equal(t, 6*time.Second, config.ExpirationTime)
		assert.Equal(t, "Custom-API-Key", config.AutKey)
		assert.Equal(t, "custom123", config.AutPass)
	})

	// Test case: YAML file with partial configuration
	t.Run("PartialConfig", func(t *testing.T) {
		yamlContent := `
SERVICE_IP: "192.168.1.2"
MAX_REQUEST: 15
`
		filename := "partial_config.yml"
		err := ioutil.WriteFile(filename, []byte(yamlContent), 0644)
		assert.NoError(t, err)
		defer os.Remove(filename)

		config, err := LoadConfig(filename)
		assert.NoError(t, err)
		assert.NotNil(t, config)
		assert.Equal(t, "192.168.1.2", config.ServiceIp)       // Custom value
		assert.Equal(t, 3000, config.ServicePort)              // Default value
		assert.Equal(t, "127.0.0.1", config.IgnoreIp)          // Default value
		assert.Equal(t, 15, config.MaxRequest)                 // Custom value
		assert.Equal(t, 30*time.Second, config.ExpirationTime) // Default value
		assert.Equal(t, "X-API-Key", config.AutKey)            // Default value
		assert.Equal(t, "12345", config.AutPass)               // Default value
	})
}
