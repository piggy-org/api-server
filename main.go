package main

import "github.com/piggy-org/api-server/server"

func main() {
	// start a gin server
	httpServer := server.NewGinServer()
	err := httpServer.Run()
	if err != nil {
		panic(err)
	}
}
