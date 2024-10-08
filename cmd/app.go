package cmd

import (
	"first_socket/internal/handlers"
	"first_socket/internal/repositories"
	"first_socket/internal/router"
	"first_socket/internal/services"
	"first_socket/internal/store"
	"fmt"
)

type App struct {
	store  store.IStore
	router *router.Router
}

func (app *App) Setup() {
	app.store.Migrate()

	userRepository := repositories.NewUserRepository(app.store)

	tokenService := services.NewTokenService()

	authHandler := handlers.NewAuthHandler(
		userRepository,
		tokenService,
	)

	app.router.SetupRoutes(
		tokenService,
		authHandler,
	)
}

func (app *App) Run() {
	app.router.Run()
}

func NewApp() *App {

	str, errStr := store.NewStore("postgres")
	if errStr != nil {
		fmt.Println(errStr)
		return nil
	}

	router := router.NewRouter()

	return &App{
		store:  str,
		router: router,
	}
}
