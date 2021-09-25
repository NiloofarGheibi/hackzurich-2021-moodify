package config

import (
	"os"

	"github.com/jinzhu/configor"
)

// Config contains all the configuration info for running the service
type Config struct {
	Environment string
	Database    *Database
	API         *API
	Spotify     *Spotify
}

// Database contains all of the api specific configuration
type API struct {
	Host string
}

// Database contains all of the database specific configuration
type Database struct {
	URL      string
	UserName string
	Password string
}

// Spotify contains all of the spotify specific configuration
type Spotify struct {
	ClientID     string
	ClientSecret string
	TokenURL     string
	Host         string
}

// Load will load the config from file and environment variables
func Load() (*Config, error) {
	file := os.Getenv("CONFIG_FILE")
	if file == "" {
		file = "config.development.yaml"
	}

	var conf Config

	err := configor.Load(&conf, file)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
