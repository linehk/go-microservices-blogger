package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/linehk/gin-blog/errno"
	"github.com/linehk/gin-blog/router/api"
	"github.com/linehk/gin-blog/vm"
)

func GetArticles(c *gin.Context) {
	valid := validation.Validation{}
	state := -1
	// PostForm 从 urlencoded 表单或 multipart 表单返回指定 key 的值
	if s := c.PostForm("state"); s != "" {
		state = com.StrTo(s).MustInt()
		// 规则：0 <= state <= 1
		valid.Range(state, 0, 1, "state")
	}

	tagId := -1
	if t := c.PostForm("tag_id"); t != "" {
		tagId = com.StrTo(t).MustInt()
		// 规则：tag_id >= 1
		valid.Min(tagId, 1, "tag_id")
	}

	// 表单验证错误
	if valid.HasErrors() {
		api.LogErrors(valid.Errors)
		api.Response(c, http.StatusBadRequest, errno.InvalidParams, nil)
		return
	}

	// 构造结构体
	vmArticle := vm.Article{TagID: tagId, State: state, PageNum: api.PageNum(c), PageSize: api.PageSize}

	// 计数
	count, err := vmArticle.Count()
	if err != nil {
		api.Response(c, http.StatusInternalServerError, errno.CountArticleFail, nil)
		return
	}

	// 获得所有文章
	articles, err := vmArticle.GetAll()
	if err != nil {
		api.Response(c, http.StatusInternalServerError, errno.GetArticleListFail, nil)
		return
	}

	// 填充进 data
	data := map[string]interface{}{"lists": articles, "count": count}

	// 根据 data 响应
	api.Response(c, http.StatusOK, errno.Success, data)
}

func GetArticle(c *gin.Context) {
	// c.Param 返回 URL 参数的值
	// router.GET("/user/:id", func(c *gin.Context) {
	// 	// a GET request to /user/john
	// 	id := c.Param("id") // id == "john"
	// })
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	// 规则：id >= 1
	valid.Min(id, 1, "id")
	if valid.HasErrors() {
		api.LogErrors(valid.Errors)
		api.Response(c, http.StatusBadRequest, errno.InvalidParams, nil)
		return
	}

	// 构造结构体
	vmArticle := vm.Article{ID: id}
	exist, err := vmArticle.HasID()
	if err != nil {
		api.Response(c, http.StatusInternalServerError, errno.CheckArticleIsExistFail, nil)
		return
	}
	if !exist {
		api.Response(c, http.StatusOK, errno.ArticleIsNotExist, nil)
		return
	}
	article, err := vmArticle.Get()
	if err != nil {
		api.Response(c, http.StatusInternalServerError, errno.GetArticleFail, nil)
		return
	}
	api.Response(c, http.StatusOK, errno.Success, article)
}

// AddArticleForm 用于验证表单
type AddArticleForm struct {
	// form 标签用于 c.Bind(form)
	TagID     int    `form:"tag_id" valid:"Required;Min(1)"`
	Title     string `form:"title" valid:"Required;MaxSize(100)"`
	Desc      string `form:"desc" valid:"Required;MaxSize(255)"`
	Content   string `form:"content" valid:"Required;MaxSize(65535)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" valid:"Range(0,1)"`
}

func AddArticle(c *gin.Context) {
	var form AddArticleForm
	// 绑定并验证
	httpCode, errCode := api.BindAndValid(c, &form)
	if errCode != errno.Success {
		api.Response(c, httpCode, errCode, nil)
		return
	}

	// 检查 Tag 的 ID
	vmTag := vm.Tag{ID: form.TagID}
	exist, err := vmTag.HasID()
	if err != nil {
		api.Response(c, http.StatusInternalServerError, errno.GetExistedTagFail, nil)
		return
	}
	if !exist {
		api.Response(c, http.StatusOK, errno.TagIsNotExist, nil)
		return
	}

	// 构造结构体
	vmArticle := vm.Article{
		TagID:   form.TagID,
		Title:   form.Title,
		Desc:    form.Desc,
		Content: form.Content,
		State:   form.State,
	}
	if err := vmArticle.Add(); err != nil {
		api.Response(c, http.StatusInternalServerError, errno.AddArticleFail, nil)
		return
	}
	api.Response(c, http.StatusOK, errno.Success, nil)
}

// EditArticleForm 用于验证表单
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
	// 获取 id
	form := EditArticleForm{ID: com.StrTo(c.Param("id")).MustInt()}
	httpCode, errCode := api.BindAndValid(c, &form)
	if errCode != errno.Success {
		api.Response(c, httpCode, errCode, nil)
		return
	}

	// 构造结构体
	vmArticle := vm.Article{
		ID:         form.ID,
		TagID:      form.TagID,
		Title:      form.Title,
		Desc:       form.Desc,
		Content:    form.Content,
		ModifiedBy: form.ModifiedBy,
		State:      form.State,
	}

	// 检查 Article 的 ID
	exist, err := vmArticle.HasID()
	if err != nil {
		api.Response(c, http.StatusInternalServerError, errno.CheckArticleIsExistFail, nil)
		return
	}
	if !exist {
		api.Response(c, http.StatusOK, errno.ArticleIsNotExist, nil)
		return
	}

	// 检查 Tag 的 ID
	vmTag := vm.Tag{ID: form.TagID}
	exist, err = vmTag.HasID()
	if err != nil {
		api.Response(c, http.StatusInternalServerError, errno.GetExistedTagFail, nil)
		return
	}
	if !exist {
		api.Response(c, http.StatusOK, errno.TagIsNotExist, nil)
		return
	}

	err = vmArticle.Edit()
	if err != nil {
		api.Response(c, http.StatusInternalServerError, errno.EditArticleFail, nil)
		return
	}

	api.Response(c, http.StatusOK, errno.Success, nil)
}

func DeleteArticle(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if valid.HasErrors() {
		api.LogErrors(valid.Errors)
		api.Response(c, http.StatusOK, errno.InvalidParams, nil)
		return
	}

	// 检查 Article 的 ID
	vmArticle := vm.Article{ID: id}
	exist, err := vmArticle.HasID()
	if err != nil {
		api.Response(c, http.StatusInternalServerError, errno.CheckArticleIsExistFail, nil)
		return
	}
	if !exist {
		api.Response(c, http.StatusOK, errno.ArticleIsNotExist, nil)
		return
	}

	err = vmArticle.Delete()
	if err != nil {
		api.Response(c, http.StatusInternalServerError, errno.DeleteArticleFail, nil)
		return
	}

	api.Response(c, http.StatusOK, errno.Success, nil)
}
