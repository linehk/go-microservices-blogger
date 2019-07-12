package model

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model
	Name string `json:"name"`
}

func GetTags(offset, limit int, cond map[string]interface{}) ([]Tag, error) {
	var tags []Tag
	var err error
	// 从 offset 开始读取 limit 条
	if limit > 0 && offset > 0 {
		err = db.Where(cond).Find(&tags).Offset(offset).Limit(limit).Error
	} else {
		err = db.Where(cond).Find(&tags).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, nil
}

func GetTag(id int) (*Tag, error) {
	var tag Tag
	if err := db.Where("id = ? AND deleted_on = ?", id, 0).First(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func HasTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ? AND deleted_on = ?", name, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	// id 为正数时才表示存在
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

func HasTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	// id 为正数时才表示存在
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

func GetTagsCount(cond map[string]interface{}) (int, error) {
	var count int
	if err := db.Model(&Tag{}).Where(cond).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func AddTag(name, createdBy string, state int) error {
	// 根据参数构造 tag 结构体
	tag := Tag{Name: name, Model: Model{CreatedBy: createdBy, State: state}}
	// 插入记录
	if err := db.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

func EditTag(id int, data map[string]interface{}) error {
	if err := db.Model(&Tag{}).Where("id = ? AND deleted_on = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTag(id int) error {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTags() (bool, error) {
	// Unscoped 返回所有记录，包含软删除的记录
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{}).Error; err != nil {
		return false, err
	}
	return true, nil
}
