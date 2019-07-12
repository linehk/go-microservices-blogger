package config

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

// 服务器设置
type ServerConfig struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var Server = &ServerConfig{}

// 数据库设置
type DatabaseConfig struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string // 数据库名
	TablePrefix string // 表前缀
}

var Database = &DatabaseConfig{}

var cfg *ini.File

func Setup() {
	var err error
	configFilePath := "config/config.ini"
	cfg, err = ini.Load(configFilePath)
	if err != nil {
		log.Fatalf("can't load file: %s, err: %v", configFilePath, err)
	}

	// 把 .ini 文件的内容映射到结构体里
	mapTo("server", Server)
	mapTo("database", Database)

	// 转换数据类型
	Server.ReadTimeout *= time.Second
	Server.WriteTimeout *= time.Second
}

// 映射到对应的结构体
func mapTo(s string, v interface{}) {
	if err := cfg.Section(s).MapTo(v); err != nil {
		log.Fatalf("section: %s can't map to: %v, err: %v", s, v, err)
	}
}
