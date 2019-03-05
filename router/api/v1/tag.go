package v1

import (
	"net/http"

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
		PageSize: 10,
	}

	tags, err := vmTag.GetAll()
	if err != nil {
		api.Response(c, http.StatusInternalServerError,
			11, nil)
		return
	}

	count, err := vmTag.Count()
	if err != nil {
		api.Response(c, http.StatusInternalServerError,
			22, nil)
		return
	}

	api.Response(c, http.StatusOK, 33, map[string]interface{}{
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
