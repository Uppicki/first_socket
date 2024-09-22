package main

import (
	"first_socket/cmd"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	server := cmd.NewApp()

	server.Run()
}
