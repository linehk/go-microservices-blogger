package vm

import (
	"github.com/linehk/gin-blog/model"
)

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int
	PageNum    int
	PageSize   int
}

func (t *Tag) HasName() (bool, error) {
	return model.HasTagByName(t.Name)
}

func (t *Tag) HasID() (bool, error) {
	return model.HasTagByID(t.ID)
}

func (t *Tag) Add() error {
	return model.AddTag(t.Name, t.CreatedBy, t.State)
}

func (t *Tag) Edit() error {
	data := map[string]interface{}{
		"modified_by": t.ModifiedBy,
		"name":        t.Name,
	}
	if t.State >= 0 {
		data["state"] = t.State
	}
	return model.EditTag(t.ID, data)
}

func (t *Tag) Delete() error {
	return model.DeleteTag(t.ID)
}

func (t *Tag) Count() (int, error) {
	return model.GetTagsCount(t.toMap())
}

func (t *Tag) GetAll() ([]model.Tag, error) {
	var tags []model.Tag
	tags, err := model.GetTags(t.PageNum, t.PageSize, t.toMap())
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (t *Tag) Get() (*model.Tag, error) {
	var tag *model.Tag
	tag, err := model.GetTag(t.ID)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func (t *Tag) toMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["deleted_on"] = 0
	if t.Name != "" {
		m["name"] = t.Name
	}
	if t.State >= 0 {
		m["state"] = t.State
	}
	return m
}
