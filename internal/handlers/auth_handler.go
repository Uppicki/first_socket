package handlers

import (
	"first_socket/internal/domain"
	"first_socket/internal/payload/requests"
	"first_socket/internal/payload/responses"
	"first_socket/internal/repositories"
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	repo repositories.IUserRepository
}

func (handler *AuthHandler) AvailableLogin(ctx *gin.Context) {
	login := ctx.Query("login")

	isValidLength := utf8.RuneCountInString(login) >= 5 && utf8.RuneCountInString(login) <= 32

	isAvailable := isValidLength && handler.repo.IsLoginExsist(login)

	ctx.JSON(http.StatusOK, responses.LoginAvailableResponse{
		IsAvailable: isAvailable,
	})
}

func (handler *AuthHandler) RegistrUser(ctx *gin.Context) {
	var request requests.RegRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		14,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user := domain.User{
		Login:    request.Login,
		Password: string(hashPassword),
	}

	ierr := handler.repo.CreateUser(user)

	if ierr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Reg error"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User successfully registered"})
}

func NewAuthHandler(repository repositories.IUserRepository) *AuthHandler {
	return &AuthHandler{
		repo: repository,
	}
}
