package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func (u UserController) RegisterRouter(engine *gin.Engine) {
	user := engine.Group("/user")
	user.POST("/login", u.login)
	user.POST("/logout", u.logout)
}

func (u UserController) login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"": "",
		})
	}
	// TODO return token
}

func (u UserController) logout(c *gin.Context) {

}
