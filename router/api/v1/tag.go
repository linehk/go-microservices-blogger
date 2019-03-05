package v1

import (
	"net/http"

	"github.com/linehk/gin-blog/errno"

	"github.com/linehk/gin-blog/router/api"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/linehk/gin-blog/vm"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")
	state := -1
	if s := c.Query("state"); s != "" {
		state = com.StrTo(s).MustInt()
	}

	vmTag := vm.Tag{
		Name:     name,
		State:    state,
		PageNum:  api.PageNum(c),
		PageSize: api.PageSize,
	}

	tags, err := vmTag.GetAll()
	if err != nil {
		api.Response(c, http.StatusInternalServerError,
			errno.ERROR_GET_TAGS_FAIL, nil)
		return
	}

	count, err := vmTag.Count()
	if err != nil {
		api.Response(c, http.StatusInternalServerError,
			errno.ERROR_COUNT_TAG_FAIL, nil)
		return
	}

	api.Response(c, http.StatusOK,
		errno.SUCCESS, map[string]interface{}{
			"lists": tags,
			"count": count,
		})
}

func GetTag(c *gin.Context) {
}

func AddTag(c *gin.Context) {
}

func EditTag(c *gin.Context) {
}

func DeleteTag(c *gin.Context) {
}
