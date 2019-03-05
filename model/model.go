package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/linehk/gin-blog/config"
)

type Model struct {
	ID         int    `gorm:"primary_key" json:"id"`
	CreatedAt  int    `json:"created_at"`
	CreatedBy  string `json:"created_by"`
	ModifiedAt int    `json:"modified_at"`
	ModifiedBy string `json:"modified_by"`
	DeletedAt  int    `json:"deleted_at"`
}

var db *gorm.DB

func Init() {
	DSL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Name)
	var err error
	db, err = gorm.Open(config.Database.Type, DSL)
	if err != nil {
		log.Fatalf("can't open database err: %v", err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func Close() {
	defer db.Close()
}
