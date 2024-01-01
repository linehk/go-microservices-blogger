package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	Model // 自定义的 Model
	// 关系：Article 拥有 Tag，TagID 为外键
	TagID   int    `gorm:"index" json:"tag_id"` // 索引
	Tag     Tag    `json:"tag"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func GetArticles(offset, limit int, cond map[string]interface{}) ([]*Article, error) {
	var articles []*Article
	// 预加载
	// 根据 cond 多表查询
	err := db.Preload("Tag").Where(cond).Offset(offset).Limit(limit).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

func GetArticle(id int) (*Article, error) {
	var article Article
	// 预加载
	// 多表查询
	err := db.Preload("Tag").Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, nil
}

func HasArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	// id 为正数才表示存在
	if article.ID > 0 {
		return true, nil
	}
	return false, nil
}

func GetArticlesCount(cond map[string]interface{}) (int64, error) {
	var count int64
	if err := db.Model(&Article{}).Where(cond).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func AddArticle(data map[string]interface{}) error {
	// 根据 data 参数构造 article 结构体
	article := Article{
		TagID:   data["tag_id"].(int),
		Title:   data["title"].(string),
		Desc:    data["desc"].(string),
		Content: data["content"].(string),
		Model: Model{
			CreatedBy: data["created_by"].(string),
			State:     data["state"].(int),
		},
	}
	// 插入记录
	if err := db.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

func EditArticle(id int, data map[string]interface{}) error {
	if err := db.Model(&Article{}).Where("id = ? AND deleted_on = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteArticle(id int) error {
	if err := db.Where("id = ?", id).Delete(&Article{}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteArticles() error {
	// Unscoped 返回所有记录，包含软删除的记录
	if err := db.Unscoped().Where("deleted_on != ?", 0).Delete(&Article{}).Error; err != nil {
		return err
	}
	return nil
}
