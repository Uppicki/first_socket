package cmd

import (
	"first_socket/internal/handlers"
	"first_socket/internal/middleware"
	"first_socket/internal/repositories"
	"first_socket/internal/store"
	wsusermanager "first_socket/internal/ws_manager/user_manager"

	"github.com/gin-gonic/gin"
)

type App struct {
	store  *store.Store
	router *gin.Engine
}

func (app *App) Run() {
	app.router.Run(":5556")
}

func NewApp() *App {
	store := store.NewStore()

	userRep := repositories.NewUserRepository(store)

	authHandler := handlers.NewAuthHandler(userRep)
	trashHandler := handlers.NewTrashHandler(userRep)

	wsUserManager := wsusermanager.NewWSUserManager(
		userRep,
	)

	router := gin.Default()

	router.GET("/", handlers.HelloHandler)

	router.GET("/users", trashHandler.GetAllUsers)
	router.POST("/login", authHandler.LoginHandler)

	wsprotected := router.Group(
		"/ws",
		middleware.KeyMIddleware(userRep),
	)
	{
		wsprotected.GET("/", wsUserManager.ServeWS)
	}

	return &App{
		router: router,
	}
}
