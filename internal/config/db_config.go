package config

import (
	"errors"
	"github.com/BurntSushi/toml"
	"log"
)

type DatabaseConfig struct {
	URL string `toml:"mongo_url"`
}

func NewDatabaseConfig() (*DatabaseConfig, error) {
	var dbCong DatabaseConfig
	_, err := toml.DecodeFile("configs/database.toml", &dbCong)
	if err != nil {
		log.Printf("Error for reading local config: %v", err)

		return nil, errors.New("Error for reading local config")
	}

	return &dbCong, nil
}