package main

import (
	"go-db/src/api"
	"go-db/src/filesystem"
)

func main() {
	config, err := filesystem.LoadConfig()
	if err != nil {
		panic(err)
	}

	api.StartServer(config, ":8080")
}
