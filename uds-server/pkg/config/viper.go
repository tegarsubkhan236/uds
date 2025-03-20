package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	AppName       = "apps.name"
	AppHost       = "apps.host"
	AppPort       = "apps.port"
	AppProduction = "apps.production"
	AppPrefork    = "apps.prefork"

	DbUser      = "database.username"
	DbPassword  = "database.password"
	DbHost      = "database.host"
	DbPort      = "database.port"
	DbName      = "database.name"
	DbParsetime = "database.parsetime"

	JwtSecret = "jwt.secret"
)

func NewViper() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return nil
}
