package main

import "first_socket/cmd"

func main() {
	server := cmd.NewApp()

	server.Run()
}
