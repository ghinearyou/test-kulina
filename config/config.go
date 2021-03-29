package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

var (
	once sync.Once
	conf *Config
)

type Config struct {
	Port int `envconfig:"PORT" default:"3001"`
	DBUsername string `envconfig:"DB_Username" default:"root"`
	DBPass string `envconfig:"DB_Pass" default:""`
	DBHost string `envconfig:"DB_Host" default:"localhost"`
	DBPort string `envconfig:"DB_Port" default:"3306"`
	DBName string `envconfig:"DB_Name" default:"testers"`
}

func Init() {
	conf = new(Config)
	envconfig.MustProcess("", conf)
}

func Get() *Config {
	return conf
}