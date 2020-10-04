package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"passport/middleware"
	"passport/model"
)

type UsersController struct {
}

func (u UsersController) RegisterRouter(engine *gin.RouterGroup) {
	users := engine.Group("/users")
	users.POST("/login", u.login)
	users.POST("/logout", u.logout)
}

func (u UsersController) login(c *gin.Context) {
	var req struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := middleware.Validator(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 201,
			"msg":  err.Error(),
		})
		return
	}

	model.OperateSuccess(c)
}

func (u UsersController) logout(c *gin.Context) {

}
