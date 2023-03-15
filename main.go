package main

import "github.com/piggy-org/api-server/http-server"

func main() {
	// start a gin http-server
	server := http_server.NewGinServer()
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
