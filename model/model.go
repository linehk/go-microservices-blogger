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
	ID         int    `gorm:"primary_key" json:"id"` // 主键，根据约定不需要
	CreatedOn  int    `json:"created_on"`
	ModifiedOn int    `json:"modified_on"`
	DeletedOn  int    `json:"deleted_on"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

var db *gorm.DB

func Setup() {
	// 构建 DSL
	DSL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Name)

	// 连接到数据库
	var err error
	db, err = gorm.Open(config.Database.Type, DSL)
	if err != nil {
		log.Fatalf("can't open database err: %v", err)
	}

	// 替换表名 Handler，设置表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.Database.TablePrefix + defaultTableName
	}

	// 注册回调函数
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	// 全局禁用复数表名
	db.SingularTable(true)

	// 最大链接数
	db.DB().SetMaxIdleConns(10)
	// 最大打开链接
	db.DB().SetMaxOpenConns(100)

	// 自动迁移
	db.AutoMigrate(&Article{}, &Tag{})
}

func Close() {
	defer db.Close()
}

// 注册 gorm 回调函数
// see https://github.com/jinzhu/gorm/blob/master/callback_create.go

// 创建数据时的回调函数
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		now := time.Now().Unix()
		// 如果存在该列
		if createdOn, ok := scope.FieldByName("CreatedOn"); ok {
			// 如果该列的值为空
			if createdOn.IsBlank {
				// 设置该列的值
				if err := createdOn.Set(now); err != nil {
					scope.Log("createdOn.Set() err: %v", err)
				}
			}
		}
		// 如果存在该列
		if modifiedOn, ok := scope.FieldByName("ModifiedOn"); ok {
			// 如果该列的值为空
			if modifiedOn.IsBlank {
				// 设置该列的值
				if err := modifiedOn.Set(now); err != nil {
					scope.Log("modifiedOn.Set() err: %v", err)
				}
			}
		}
	}
}

// 更新数据时的回调函数
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	// 查找设置有 update_column 标签的列，如果没有
	if _, ok := scope.Get("gorm:update_column"); !ok {
		// 则设置该列的值
		if err := scope.SetColumn("ModifiedOn", time.Now().Unix()); err != nil {
			scope.Log("SetColumn() err: %v", err)
		}
	}
}

// 删除数据时的回调函数
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		// 查找设置有 delete_option 标签的列，如果有就保存起来
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		deletedOn, ok := scope.FieldByName("DeletedOn")
		// 如果未调用过 Unscoped()，并且存在该列
		if !scope.Search.Unscoped && ok {
			// 则软删除
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				// 返回引用的表名
				scope.QuotedTableName(),
				// Quote() 使用引号包裹参数，deletedOn.DBName 返回列名
				scope.Quote(deletedOn.DBName),
				// 防止 SQL 注入
				scope.AddToVars(time.Now().Unix()),
				// 返回组合条件 SQL
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			// 否则直接删除
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// 在字符串前添加一个空格
func addExtraSpaceIfExist(s string) string {
	if s == "" {
		return ""
	}
	return " " + s
}
