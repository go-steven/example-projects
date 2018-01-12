package test

import (
	"github.com/gin-gonic/gin"
	. "github.com/go-steven/example-projects/rest-with-gin/handler"
	"net/http"
)

func HelloHandler(c *gin.Context) {
	token := c.Query("token")
	if Check(token == "", "missing token", c) {
		panic("xxx")
		return
	}

	c.JSON(http.StatusOK, token)
}
