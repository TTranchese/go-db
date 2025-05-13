package api

import (
	"fmt"
	"go-db/src/models"
	"log"
	"net/http"
)

func StartServer(Config models.Config) {
	fmt.Println("starting GO DB...")

	port := checkPortPrefix(Config.Port)

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		HandleCreateIndex(w, r, Config)
	})
	http.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		HandleInsertKeyValue(w, r, Config)
	})
	log.Default().Println("server is running on port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func checkPortPrefix(port string) string {
	if port == "" {
		return ":8080"
	}

	if port[0] != ':' {
		port = ":" + port
	}

	return port
}
