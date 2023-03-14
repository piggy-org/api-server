package main

import (
	"github.com/Vincent18zt/piggy-market-backend/http-server"
)

func main() {
	// start a gin http-server
	server := http_server.NewGinServer()
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
