package handlers

import (
	"first_socket/internal/payload/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, responses.HelloResponse{
		Message: "Hello world",
	})
}
