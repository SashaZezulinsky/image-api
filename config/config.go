package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

// Config is an app config struct
type Config struct {
	Server  ServerConfig
	MongoDB MongoDB
	//Metrics  Metrics
	Logger Logger
}

// ServerConfig is a config struct
type ServerConfig struct {
	AppVersion        string
	Port              string
	Mode              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	CtxDefaultTimeout time.Duration
	Debug             bool
}

// Logger config
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// MongoDB config
type MongoDB struct {
	MongoURI string

	Database           string
	ImageCollection    string
	MetadataCollection string
}

// Metrics config
type Metrics struct {
	URL         string
	ServiceName string
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
