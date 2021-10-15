package config

import (
	"errors"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type DatabaseConfig struct {
	URL string `toml:"mongo_url"`
}

func NewDatabaseConfig() (*DatabaseConfig, error) {
	var dbCong DatabaseConfig
	hello := "test"
	fmt.Print(hello)
	_, err := toml.DecodeFile("configs/database.toml", &dbCong)
	if err != nil {
		log.Printf("Error for reading local config: %v", err)

		return nil, errors.New("Error for reading local config")
	}

	return &dbCong, nil
}
