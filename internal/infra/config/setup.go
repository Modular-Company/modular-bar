package config

import (
	"time"

	"github.com/magiconair/properties"
)

const (
	appConfigFilename = "../../../resources/application.properties"
)

type AppConfig struct {
	Database Database `properties:"database"`
	Port     string   `properties:"port"`
}

type Database struct {
	Username        string        `properties:"username"`
	Password        string        `properties:"password"`
	Schema          string        `properties:"schema"`
	Server          string        `properties:"server"`
	ConnMaxLifetime time.Duration `properties:"connMaxLifetime"`
	MaxOpenConns    int           `properties:"maxOpenConns"`
	MaxIddleConns   int           `properties:"maxIddleConns"`
}

func Load() (*AppConfig, error) {
	props := properties.MustLoadFile(appConfigFilename, properties.UTF8)

	config := &AppConfig{}

	if err := props.Decode(config); err != nil {
		panic(err.Error())
	}

	return config, nil
}
