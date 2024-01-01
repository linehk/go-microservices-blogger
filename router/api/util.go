package api

import (
	"log"
	"net/http"

	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/linehk/gin-blog/errno"
)

// PageSize 每页 Article 的数量
var PageSize = 10

func PageNum(c *gin.Context) int {
	count := 0
	// c.Query("page") 取回 URL 中的参数，然后再转换成 Int
	// GET /path?id=1234&name=Manu&value=
	// c.Query("id") == "1234"
	// c.Query("name") == "Manu"
	// c.Query("value") == ""
	// c.Query("wtf") == ""
	page := com.StrTo(c.Query("page")).MustInt()
	// page <= 1 时，count 为 0
	if page > 0 {
		// page = 2 时，count = 10
		count = (page - 1) * 10
	}
	return count
}

// Resp 统一返回格式
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response 根据数据返回响应
func Response(c *gin.Context, httpCode, errCode int, data interface{}) {
	c.JSON(httpCode, Resp{Code: httpCode, Msg: errno.Msg[errCode], Data: data})
}

// BindAndValid 绑定并验证表单
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	// c.Bind(form) 会根据 Content-Type 选择 binding
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, errno.InvalidParams
	}
	valid := validation.Validation{}
	// 验证该表单，必须是结构体或结构体指针
	ok, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, errno.Error
	}
	// 验证失败
	if !ok {
		// LogErrors(valid.Errors)
		for _, err := range valid.Errors {
			log.Print(err.Key, err.Message)
		}
		return http.StatusBadRequest, errno.InvalidParams
	}
	return http.StatusOK, errno.Success
}

// LogErrors 把验证错误输出到日志
func LogErrors(errors []*validation.Error) {
	for _, err := range errors {
		log.Print(err.Key, err.Message)
	}
}
