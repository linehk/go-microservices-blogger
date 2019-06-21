package api

import (
	"log"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/linehk/gin-blog/errno"
)

var PageSize = 10

func PageNum(c *gin.Context) int {
	count := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		count = (page - 1) * 10
	}
	return count
}

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Response(c *gin.Context, httpCode, errCode int, data interface{}) {
	c.JSON(httpCode, Resp{
		Code: httpCode,
		Msg:  errno.Msg[errCode],
		Data: data,
	})
}

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, errno.INVALID_PARAMS
	}
	valid := validation.Validation{}
	ok, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, errno.ERROR
	}
	if !ok {
		for _, err := range valid.Errors {
			log.Print(err.Key, err.Message)
		}
		return http.StatusBadRequest, errno.INVALID_PARAMS
	}
	return http.StatusOK, errno.SUCCESS
}

func LogErrors(errors []*validation.Error) {
	for _, err := range errors {
		log.Print(err.Key, err.Message)
	}
}
