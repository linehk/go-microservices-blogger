package api

import (
	"github.com/Unknwon/com"
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
	return
}
