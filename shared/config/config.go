package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type (
	// ImmutableConfig ...
	ImmutableConfig interface {
		GetDatabaseUserName() string
		GetDatabasePassword() string
		GetDatabaseHost() string
		GetDatabaseName() string
		GetDatabaseMaxConnection() int
		GetDatabaseMinConnection() int
		GetDatabaseDebugMode() bool
		GetPort() int
	}

	immutable struct {
		DatabaseUserName      string `mapstructure:"DB_USERNAME"`
		DatabasePassword      string `mapstructure:"DB_PASSWORD"`
		DatabaseHost          string `mapstructure:"DB_HOST"`
		DatabaseName          string `mapstructure:"DB_NAME"`
		DatabaseMaxConnection int    `mapstructure:"DB_MAX_CONNECTION"`
		DatabaseMinConnection int    `mapstructure:"DB_MIN_CONNECTION"`
		DatabaseDebugMode     bool   `mapstructure:"DB_DEBUG_MODE"`
		Port                  int    `mapstructure:"PORT"`
	}
)

func (i *immutable) GetDatabaseUserName() string {
	return i.DatabaseUserName
}

func (i *immutable) GetDatabasePassword() string {
	return i.DatabasePassword
}

func (i *immutable) GetDatabaseHost() string {
	return i.DatabaseHost
}

func (i *immutable) GetDatabaseName() string {
	return i.DatabaseName
}

func (i *immutable) GetDatabaseMaxConnection() int {
	return i.DatabaseMaxConnection
}

func (i *immutable) GetDatabaseMinConnection() int {
	return i.DatabaseMinConnection
}

func (i *immutable) GetDatabaseDebugMode() bool {
	return i.DatabaseDebugMode
}

func (i *immutable) GetPort() int {
	return i.Port
}

var (
	imOnce sync.Once
	im     immutable
)

// GetDefaultImmutableConfig ...
func GetDefaultImmutableConfig() ImmutableConfig {
	var outer error
	var success = true

	env := os.Getenv("APP_ENV")
	pwd := os.Getenv("APP_PWD")

	if env == "test" && pwd == "" {
		panic("APP_PWD env is required in test env")
	}

	imOnce.Do(func() {
		v := viper.New()
		v.SetConfigName("app.config")

		if env == "test" {
			v.AddConfigPath(pwd)
		} else {
			v.AddConfigPath(".")
		}

		v.SetEnvPrefix("vp")
		v.AutomaticEnv()

		if err := v.ReadInConfig(); err != nil {
			success = false
			outer = fmt.Errorf("failed to read app.config file")
		}

		v.Unmarshal(&im)
	})

	if !success {
		panic(outer)
	}

	return &im
}
