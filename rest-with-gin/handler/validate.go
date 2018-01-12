package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Check(flag bool, err string, c *gin.Context) (ret bool) {
	ret = flag
	if ret {
		Logger.Error(err)
		c.JSON(http.StatusOK, APIError{Code: BADREQUEST_ERROR, Msg: err})
	}
	return
}

func CheckErr(err error, c *gin.Context) (ret bool) {
	ret = err != nil
	if ret {
		Logger.Error(err.Error())
		c.JSON(http.StatusOK, APIError{Code: BADREQUEST_ERROR, Msg: err.Error()})
	}
	return
}
