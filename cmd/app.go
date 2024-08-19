package cmd

import (
	"first_socket/internal/handlers"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func (app *App) Run() {
	app.router.Run(":5555")
}

func NewApp() *App {
	router := gin.Default()

	router.GET("/", handlers.HelloHandler)

	return &App{
		router: router,
	}
}
