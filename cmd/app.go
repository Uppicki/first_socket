package cmd

import (
	//"first_socket/internal/handlers"
	"first_socket/internal/store"
	"fmt"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type App struct {
	store  store.IStore
	router *gin.Engine
}

func (app *App) Run() {
	app.router.Run(":5556")
}

func NewApp() *App {

	str, errStr := store.NewStore("postgres")
	if errStr != nil {
		fmt.Println(errStr)
		return nil
	}

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))

	return &App{
		store:  str,
		router: router,
	}
}
