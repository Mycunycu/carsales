package config

import (
	"strings"

	"github.com/spf13/viper"
)

type (
	//Config ...
	Config struct {
		HTTP  HTTPConfig
		Mongo MongoConfig
	}
	// HTTPConfig ...
	HTTPConfig struct {
		Host string
		Port string
	}
	// MongoConfig ...
	MongoConfig struct {
		ConnString string
		DbName     string
	}
)

// Init ...
func Init(path string) (*Config, error) {
	if err := readConfig(path); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func readConfig(path string) error {
	pathParts := strings.Split(path, "/")

	viper.AddConfigPath(pathParts[0]) // folder
	viper.SetConfigName(pathParts[1]) // file name

	return viper.ReadInConfig()
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("mongo", &cfg.Mongo); err != nil {
		return err
	}

	return nil
}
