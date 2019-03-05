package model

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model
	Name string `json:"name"`
}

func AddTag(name, createdBy string, state int) error {
	tag := Tag{
		Name: name,
		Model: Model{
			CreatedBy: createdBy,
			State:     state,
		},
	}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

func HasTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").
		Where("name = ? AND deleted_on = ?", name, 0).
		First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

func HasTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").
		Where("id = ? AND deleted_on = ?", id, 0).
		First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

func GetTags(offset, limit int, cond map[string]interface{}) ([]Tag, error) {
	var tags []Tag
	var err error
	// 从 offset 开始读取 limit 条
	if limit > 0 && offset > 0 {
		err = db.Where(cond).
			Find(&tags).
			Offset(offset).
			Limit(limit).Error
	} else {
		err = db.Where(cond).Find(&tags).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, nil
}

func GetTagsCount(cond map[string]interface{}) (int, error) {
	var count int
	if err := db.Model(&Tag{}).
		Where(cond).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func DeleteTag(id int) error {
	if err := db.Where("id = ?", id).
		Delete(&Tag{}).Error; err != nil {
		return err
	}
	return nil
}

func EditTag(id int, data map[string]interface{}) error {
	if err := db.Model(&Tag{}).
		Where("id = ? AND deleted_on = ?", id, 0).
		Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTags() (bool, error) {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{}).Error; err != nil {
		return false, err
	}
	return true, nil
}
