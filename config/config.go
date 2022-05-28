package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

// Config is an app config struct
type Config struct {
	Server  ServerConfig
	MongoDB MongoDB
}

// ServerConfig is a config struct
type ServerConfig struct {
	Port string
}

// MongoDB config
type MongoDB struct {
	MongoURI string
	Database string
}

// LoadConfig loads config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// ParseConfig parses config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
