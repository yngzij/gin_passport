package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type RegisterRouter interface {
	RegisterRouter(engine *gin.Engine)
}

func main() {
	engine := gin.New()

	log.Fatal(engine.Run(":8888"))
}
