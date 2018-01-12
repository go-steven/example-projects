package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-steven/example-projects/rest-with-gin/handler/test"
)

func test_router(r *gin.Engine) {
	g := r.Group("/test")
	{
		g.GET("/hello", test.HelloHandler)
	}
}
