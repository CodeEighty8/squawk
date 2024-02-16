package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Default_AppConfig(t *testing.T) {
	appConfig := &AppConfig{}
	baseConfig, err := GetConfig(appConfig)

	assert.Nil(t, err)
	assert.Equal(t, &AppConfig{Name: "chat-app"}, appConfig)
	assert.Equal(t, ServerConfig{Env: "dev", Port: 5000}, baseConfig.Server)
}

func Test_Custom_AppConfig(t *testing.T) {
	type AppConfig struct {
		Key string `json:"Key"`
	}

	appConfig := &AppConfig{}
	baseConfig, err := GetConfig(appConfig, "./mock/test_config.json")

	assert.Nil(t, err)
	assert.Equal(t, &AppConfig{Key: "value"}, appConfig)
	assert.Equal(t, ServerConfig{Env: "dev", Port: 5000}, baseConfig.Server)
}

func Test_Custom_AppConfig_Error(t *testing.T) {
	type AppConfig struct {
		Key string `json:"Key"`
	}

	appConfig := &AppConfig{}
	baseConfig, err := GetConfig(appConfig, "./mock/no_file.json")

	assert.Nil(t, baseConfig)
	assert.NotNil(t, err)
}
