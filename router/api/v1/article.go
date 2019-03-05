package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Articles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"test1": "test2",
	})
}
