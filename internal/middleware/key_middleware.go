package middleware

import (
	"first_socket/internal/repositories"

	"github.com/gin-gonic/gin"
)

const apiKeyHeader = "Authorization"
const UserContextKey = "user"

func KeyMIddleware(
	repository *repositories.UserRepository,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/*
			key := ctx.GetHeader(apiKeyHeader)

			user, err := repository.

			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
				return
			}

			ctx.Set(UserContextKey, user)*/

		ctx.Next()
	}
}
