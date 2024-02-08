package config

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type ServerConfig struct {
	Env  string `json:"env"`
	Port int    `json:"port"`
}

type AppConfig struct {
	Name string `json:"name"`
}

type BaseConfig struct {
	Server      ServerConfig `json:"Server"`
	Application interface{}  `json:"Application"`
}

func loadConfigFile(loader *koanf.Koanf, path string) error {

	log.Printf("reading config from %s", path)
	if err := loader.Load(file.Provider(path), json.Parser()); err != nil {
		return err
	}
	return nil
}

func GetConfig(configType interface{}, additionalConfigPaths ...string) (*BaseConfig, error) {

	config := &BaseConfig{
		Application: configType,
	}

	configLoader := koanf.New(".")

	// Load from default location
	_, currentFile, _, _ := runtime.Caller(0)
	dir := filepath.Dir(currentFile)
	absolutePath := filepath.Join(dir, "./config.json")
	if err := loadConfigFile(configLoader, absolutePath); err != nil {
		return nil, fmt.Errorf("error reading default configs: %v", err)
	}

	// Load additional configs from configsPaths
	for _, path := range additionalConfigPaths {
		if err := loadConfigFile(configLoader, path); err != nil {
			return nil, fmt.Errorf("error reading default configs: %v", err)
		}
	}

	if err := configLoader.Unmarshal("", config); err != nil {
		return nil, fmt.Errorf("error loading configs: %v", err)
	}

	return config, nil
}
