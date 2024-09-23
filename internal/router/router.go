package router

import (
	"first_socket/internal/handlers"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

func (r *Router) SetupRoutes(
	authHandler *handlers.AuthHandler,
) {
	r.router.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))

	authRoutes := r.router.Group("/api/v1/auth")
	{
		authRoutes.GET("/availableLogin", authHandler.AvailableLogin)
		authRoutes.POST("/registr", authHandler.RegistrUser)
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
