package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type config struct {
	Server   server
	Database database
}

type server struct {
	Mode         string
	Addr         string
	ReadTimeout  int
	WriteTimeout int
}

type database struct {
	Dialect  string
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	TimeZone string
}

var Cfg config

// init 初始化 Cfg 全局变量。
func init() {
	// ../config.toml 用于各个子目录。
	if _, err := toml.DecodeFile("../config.toml", &Cfg); err != nil {
		// ./config.toml 用于 main.go。
		if _, err := toml.DecodeFile("./config.toml", &Cfg); err != nil {
			log.Fatal(err)
		}
	}
}
