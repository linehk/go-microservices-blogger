package model

type Article struct {
	Model

	TagID int `gorm:"index" json:"tag_id"`
	Tag   Tag `json:"tag"`

	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
