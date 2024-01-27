package config

import (
	"log"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var Raw = koanf.New(".")

func init() {
	err := Raw.Load(file.Provider(".env"), dotenv.Parser())
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
}
