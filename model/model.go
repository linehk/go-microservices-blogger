package model

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/linehk/gin-blog/config"
)

type Model struct {
	ID         int    `gorm:"primary_key" json:"id"`
	CreatedOn  int    `json:"created_on"`
	ModifiedOn int    `json:"modified_on"`
	DeletedOn  int    `json:"deleted_on"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
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

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.Database.TablePrefix + defaultTableName
	}

	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	db.SingularTable(true)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func Close() {
	defer db.Close()
}

// gorm callback function replace
// see https://github.com/jinzhu/gorm/blob/master/callback_create.go
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		now := time.Now().Unix()
		if createdOn, ok := scope.FieldByName("CreatedOn"); ok {
			if createdOn.IsBlank {
				if err := createdOn.Set(now); err != nil {
					scope.Log("updateTimeStampForCreateCallback createdOn.Set() err: %v", err)
				}
			}
		}
		if modifiedOn, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifiedOn.IsBlank {
				if err := modifiedOn.Set(now); err != nil {
					scope.Log("updateTimeStampForCreateCallback modifiedOn.Set() err: %v", err)
				}
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		if err := scope.SetColumn("ModifiedOn", time.Now().Unix()); err != nil {
			scope.Log("updateTimeStampForUpdateCallback SetColumn() err: %v", err)
		}
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		deletedOn, ok := scope.FieldByName("DeletedOn")
		if !scope.Search.Unscoped && ok {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOn.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(s string) string {
	if s == "" {
		return ""
	}
	return " " + s
}
