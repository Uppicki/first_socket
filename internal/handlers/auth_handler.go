package handlers

import (
	"first_socket/internal/payload/requests"
	"first_socket/internal/payload/responses"
	"first_socket/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	repository *repositories.UserRepository
}

func (handler *AuthHandler) LoginHandler(ctx *gin.Context) {
	var request requests.LoginRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	handler.repository.AddUser(request.Name)

	ctx.JSON(http.StatusOK, responses.LoginResponse{
		Login: request.Name,
	})
}

func NewAuthHandler(repository *repositories.UserRepository) *AuthHandler {
	return &AuthHandler{
		repository: repository,
	}
}
