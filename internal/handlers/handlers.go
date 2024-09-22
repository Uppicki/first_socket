package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HelloHandler(ctx *gin.Context) {
	fmt.Println("Asd")
	ctx.File("./frontend/dist/index.html")
}
