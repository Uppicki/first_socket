package main

import "first_socket/cmd"

func main() {
	srv := cmd.NewApp()

	srv.Run()

}
