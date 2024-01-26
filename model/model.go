package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"

	"github.com/linehk/gin-blog/config"
)

type Model struct {
	ID         int    `gorm:"primary_key" json:"id"` // 主键，根据约定不需要
	CreatedOn  int    `json:"created_on"`
	ModifiedOn int    `json:"modified_on"`
	DeletedOn  int    `json:"deleted_on"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

var db *gorm.DB

var dbc = config.Cfg.Database

func Setup() {
	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbc.Host, dbc.User, dbc.Password, dbc.DBName, dbc.Port, dbc.SSLMode, dbc.TimeZone)

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
	})
	if err != nil {
		log.Fatalf("can't open database err: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("can't get sqlDB err: %v", err)
	}

	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)

	err = db.AutoMigrate(&Article{}, &Tag{})
	if err != nil {
		log.Fatalf("AutoMigrate err: %v", err)
	}
}
