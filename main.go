package main

import (
	"log"

	"passport/controller"

	"github.com/gin-gonic/gin"
)

type RegisterRouter interface {
	RegisterRouter(engine *gin.RouterGroup)
}

var Routers []RegisterRouter

func main() {
	engine := gin.New()

	Routers = append(Routers, controller.UsersController{})

	g := engine.Group("/api/v1")
	for _, r := range Routers {
		r.RegisterRouter(g)
	}

	log.Fatal(engine.Run(":8888"))
}
