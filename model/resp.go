package model

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func OperateSuccess(c *gin.Context) {
	var res struct {
		Code int64  `json:"code"`
		Msg  string `json:"msg"`
	}

	res.Code = ok
	res.Msg = "success"
	c.JSON(http.StatusOK, res)
}
