package v1

import (
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/linehk/gin-blog/errno"
	"github.com/linehk/gin-blog/router/api"
	"github.com/linehk/gin-blog/vm"
)

func GetArticles(c *gin.Context) {
	valid := validation.Validation{}
	state := -1
	if s := c.PostForm("state"); s != "" {
		state = com.StrTo(s).MustInt()
		valid.Range(state, 0, 1, "state")
	}
	tagId := -1
	if t := c.PostForm("tag_id"); t != "" {
		tagId = com.StrTo(t).MustInt()
		valid.Min(tagId, 1, "tag_id")
	}

	if valid.HasErrors() {
		api.LogErrors(valid.Errors)
		api.Response(c,
			http.StatusBadRequest,
			errno.INVALID_PARAMS, nil)
		return
	}

	vmArticle := vm.Article{
		TagID:    tagId,
		State:    state,
		PageNum:  api.PageNum(c),
		PageSize: api.PageSize,
	}

	count, err := vmArticle.Count()
	if err != nil {
		api.Response(c,
			http.StatusInternalServerError,
			errno.ERROR_COUNT_ARTICLE_FAIL, nil)
		return
	}

	articles, err := vmArticle.GetAll()
	if err != nil {
		api.Response(c,
			http.StatusInternalServerError,
			errno.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = articles
	data["count"] = count

	api.Response(c, http.StatusOK, errno.SUCCESS, data)
}

func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")
	if valid.HasErrors() {
		api.LogErrors(valid.Errors)
		api.Response(c,
			http.StatusBadRequest,
			errno.INVALID_PARAMS, nil)
		return
	}

	vmArticle := vm.Article{ID: id}
	exist, err := vmArticle.HasID()
	if err != nil {
		api.Response(c,
			http.StatusInternalServerError,
			errno.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exist {
		api.Response(c,
			http.StatusOK,
			errno.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	article, err := vmArticle.Get()
	if err != nil {
		api.Response(c,
			http.StatusInternalServerError,
			errno.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	api.Response(c, http.StatusOK, errno.SUCCESS, article)
}

type AddArticleForm struct {
	TagID     int    `form:"tag_id" valid:"Required;Min(1)"`
	Title     string `form:"title" valid:"Required;MaxSize(100)"`
	Desc      string `form:"desc" valid:"Required;MaxSize(255)"`
	Content   string `form:"content" valid:"Required;MaxSize(65535)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" valid:"Range(0,1)"`
}

func AddArticles(c *gin.Context) {
	var form AddArticleForm
	httpCode, errCode := api.BindAndValid(c, &form)
	if errCode != errno.SUCCESS {
		api.Response(c, httpCode, errCode, nil)
		return
	}
	vmTag := vm.Tag{ID: form.TagID}
	exist, err := vmTag.HasID()
	if err != nil {
		api.Response(c,
			http.StatusInternalServerError,
			errno.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if !exist {
		api.Response(c, http.StatusOK,
			errno.ERROR_NOT_EXIST_TAG, nil)
		return
	}
	vmArticle := vm.Article{
		TagID:   form.TagID,
		Title:   form.Title,
		Desc:    form.Desc,
		Content: form.Content,
		State:   form.State,
	}
	if err := vmArticle.Add(); err != nil {
		api.Response(c,
			http.StatusInternalServerError,
			errno.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}
	api.Response(c, http.StatusOK, errno.SUCCESS, nil)
}

type EditArticleForm struct {
	ID         int    `form:"id" valid:"Required;Min(1)"`
	TagID      int    `form:"tag_id" valid:"Required;Min(1)"`
	Title      string `form:"title" valid:"Required;MaxSize(100)"`
	Desc       string `form:"desc" valid:"Required;MaxSize(255)"`
	Content    string `form:"content" valid:"Required;MaxSize(65535)"`
	ModifiedBy string `form:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `form:"state" valid:"Range(0,1)"`
}

func EditArticle(c *gin.Context) {
	form := EditArticleForm{ID: com.StrTo(c.Param("id")).MustInt()}
	httpCode, errCode := api.BindAndValid(c, &form)
	if errCode != errno.SUCCESS {
		api.Response(c, httpCode, errCode, nil)
		return
	}
	vmArticle := vm.Article{
		ID:         form.ID,
		TagID:      form.TagID,
		Title:      form.Title,
		Desc:       form.Desc,
		Content:    form.Content,
		ModifiedBy: form.ModifiedBy,
		State:      form.State,
	}
	exist, err := vmArticle.HasID()
	if err != nil {
		api.Response(c,
			http.StatusInternalServerError,
			errno.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exist {
		api.Response(c,
			http.StatusOK,
			errno.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	vmTag := vm.Tag{ID: form.TagID}
	exist, err = vmTag.HasID()
	if err != nil {
		api.Response(c,
			http.StatusInternalServerError,
			errno.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if !exist {
		api.Response(c,
			http.StatusOK,
			errno.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = vmArticle.Edit()
	if err != nil {
		api.Response(c,
			http.StatusInternalServerError,
			errno.ERROR_EDIT_ARTICLE_FAIL, nil)
		return
	}

	api.Response(c, http.StatusOK, errno.SUCCESS, nil)
}

func DeleteArticle(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if valid.HasErrors() {
		api.LogErrors(valid.Errors)
		api.Response(c,
			http.StatusOK,
			errno.INVALID_PARAMS, nil)
		return
	}

	vmArticle := vm.Article{ID: id}
	exist, err := vmArticle.HasID()
	if err != nil {
		api.Response(c,
			http.StatusInternalServerError,
			errno.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exist {
		api.Response(c, http.StatusOK,
			errno.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	err = vmArticle.Delete()
	if err != nil {
		api.Response(c,
			http.StatusInternalServerError,
			errno.ERROR_DELETE_ARTICLE_FAIL, nil)
		return
	}

	api.Response(c, http.StatusOK, errno.SUCCESS, nil)
}
