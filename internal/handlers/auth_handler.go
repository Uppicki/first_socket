package handlers

import (
	"first_socket/internal/domain"
	"first_socket/internal/payload/requests"
	"first_socket/internal/payload/responses"
	"first_socket/internal/repositories"
	"first_socket/internal/services"

	"net/http"

	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	repo repositories.IUserRepository

	tokenService services.ITokenService
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

func (handler *AuthHandler) LoginUser(ctx *gin.Context) {
	var request requests.RegRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, uErr := handler.repo.GetUserByLogin(request.Login)

	if uErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": uErr.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	accessToken, refreshToken, err := handler.tokenService.CreateTokens(user.Login)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       "Could not generate tokens",
			"inner_error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (handler *AuthHandler) CurrentUser(ctx *gin.Context) {
	login, _ := ctx.Get("login")

	loginStr, ok := login.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid login type"})
		return
	}

	ctx.JSON(http.StatusOK, responses.UserInfoResponse{
		Login: loginStr,
	})
}

func NewAuthHandler(
	repository repositories.IUserRepository,
	tokenService services.ITokenService,
) *AuthHandler {
	return &AuthHandler{
		repo:         repository,
		tokenService: tokenService,
	}
}
