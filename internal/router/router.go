package router

import (
	"first_socket/internal/handlers"
	"first_socket/internal/middleware"
	"first_socket/internal/services"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

func (r *Router) SetupRoutes(
	tokenService services.ITokenService,
	authHandler *handlers.AuthHandler,
) {
	r.router.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))
	r.router.Use(middleware.SetupCORSMidleware())

	authRoutes := r.router.Group("/api/v1/auth")
	{
		authRoutes.GET("/availableLogin", authHandler.AvailableLogin)
		authRoutes.POST("/registr", authHandler.RegistrUser)
		authRoutes.POST("/login", authHandler.LoginUser)
	}

	protectedRoutes := r.router.Group(
		"/api/v1/protected",
		middleware.JWTMiddleware(tokenService),
	)
	{
		protectedRoutes.GET(
			"/", func(ctx *gin.Context) {

				ctx.JSON(200, gin.H{
					"asd": "AD",
				},
				)
			})
	}

}

func (r *Router) Run() {
	r.router.Run(":5556")
}

func NewRouter() *Router {
	r := gin.Default()

	return &Router{
		router: r,
	}
}
