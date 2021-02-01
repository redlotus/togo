package main

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		DBTodo      ConfigSQLDB
		Service     ConfigService
		TokenSecret string
	}

	ConfigService struct {
		Address string `yaml:"address" json:"address"`
	}

	ConfigSQLDB struct {
		DSN string
	}
)

var (
	ConfEnv = []string{
		"db_dsn",
		"service_addr",
	}
)

func LoadConfigFromEnv() Config {
	conf := Config{}
	viper.SetConfigType("env")
	for _, e := range ConfEnv {
		viper.BindEnv(e)
	}

	conf.DBTodo.DSN = viper.GetString("db_dsn")
	conf.Service.Address = viper.GetString("service_addr")

	return conf
}
