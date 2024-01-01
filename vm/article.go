package vm

import (
	"github.com/linehk/gin-blog/model"
)

// Article 用于传递数据
type Article struct {
	ID      int
	TagID   int
	Title   string
	Desc    string
	Content string

	State      int
	CreatedBy  string
	ModifiedBy string

	PageNum  int
	PageSize int
}

func (a *Article) Add() error {
	article := map[string]interface{}{
		"tag_id":     a.TagID,
		"title":      a.Title,
		"desc":       a.Desc,
		"content":    a.Content,
		"created_by": a.CreatedBy,
		"state":      a.State,
	}
	if err := model.AddArticle(article); err != nil {
		return err
	}
	return nil
}

func (a *Article) Edit() error {
	return model.EditArticle(a.ID, map[string]interface{}{
		"tag_id":      a.TagID,
		"title":       a.Title,
		"desc":        a.Desc,
		"content":     a.Content,
		"state":       a.State,
		"modified_by": a.ModifiedBy,
	})
}

func (a *Article) Get() (*model.Article, error) {
	article, err := model.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a *Article) GetAll() ([]*model.Article, error) {
	var articles []*model.Article
	articles, err := model.GetArticles(a.PageNum, a.PageSize, a.toMap())
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *Article) Delete() error {
	return model.DeleteArticle(a.ID)
}

func (a *Article) HasID() (bool, error) {
	return model.HasArticleByID(a.ID)
}

func (a *Article) Count() (int64, error) {
	return model.GetArticlesCount(a.toMap())
}

func (a *Article) toMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["deleted_on"] = 0
	if a.State != -1 {
		m["state"] = a.State
	}
	if a.TagID != -1 {
		m["tag_id"] = a.TagID
	}
	return m
}
