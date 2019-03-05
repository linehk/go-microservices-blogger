package config

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type ServerConfig struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var Server = &ServerConfig{}

type DatabaseConfig struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var Database = &DatabaseConfig{}

var cfg *ini.File

func Init() {
	var err error
	configFilePath := "config/config.ini"
	cfg, err = ini.Load(configFilePath)
	if err != nil {
		log.Fatalf("can't load file: %s, err: %v", configFilePath, err)
	}
	mapTo("server", Server)
	mapTo("database", Database)

	Server.ReadTimeout *= time.Second
	Server.WriteTimeout *= time.Second
}

func mapTo(s string, v interface{}) {
	if err := cfg.Section(s).MapTo(v); err != nil {
		log.Fatalf("section: %s can't map to: %v, err: %v", s, v, err)
	}
}
