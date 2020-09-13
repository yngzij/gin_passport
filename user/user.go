package user

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (u UserController) RegisterRouter(engine *gin.Engine) {
	user := engine.Group("/user")
	user.POST("/login", u.login)
	user.POST("/logout", u.logout)
}

func (u UserController) login(c *gin.Context) {

}

func (u UserController) logout(c *gin.Context) {

}
