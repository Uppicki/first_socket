package handlers

import (
	"first_socket/internal/payload/responses"
	"first_socket/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrashHandler struct {
	repository *repositories.UserRepository
}

func (handler *TrashHandler) GetAllUsers(ctx *gin.Context) {
	users := handler.repository.GetAllUsers()

	ctx.JSON(http.StatusOK, responses.UserListResponse{
		Users: users,
	})
}

func NewTrashHandler(repository *repositories.UserRepository) *TrashHandler {
	return &TrashHandler{
		repository: repository,
	}
}
