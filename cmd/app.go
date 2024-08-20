package cmd

import (
	"first_socket/internal/handlers"
	"first_socket/internal/repositories"
	"first_socket/internal/store"

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

	router := gin.Default()

	router.GET("/", handlers.HelloHandler)

	router.GET("/users", trashHandler.GetAllUsers)
	router.POST("/login", authHandler.LoginHandler)

	return &App{
		router: router,
	}
}
