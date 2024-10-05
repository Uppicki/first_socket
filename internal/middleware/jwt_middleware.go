package middleware

import (
	stringsRes "first_socket/internal/res/strings"
	"first_socket/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware(
	tokenService services.ITokenService,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			ctx.JSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "Authorization header is required",
				},
			)
			ctx.Abort()
			return
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		login, err := tokenService.VerifyAccessToken(tokenString)

		if err != nil {
			ctx.JSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "Invalid token",
				},
			)
			ctx.Abort()
			return
		}

		ctx.Set(stringsRes.LoginHeaderKey, login)
		ctx.Next()
	}
}
