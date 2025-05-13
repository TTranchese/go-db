package api

import (
	"fmt"
	"go-db/src/filesystem"
	"go-db/src/models"
	"log"
	"net/http"
)

func StartServer(Config models.Config) {
	fmt.Println("starting GO DB...")

	port := checkPortPrefix(Config.Port)

	http.HandleFunc("/create-index", func(w http.ResponseWriter, r *http.Request) {
		handleCreateIndex(w, r, Config)
	})

	log.Default().Println("server is running on port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleCreateIndex(w http.ResponseWriter, r *http.Request, Config models.Config) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fileName := r.URL.Query().Get("fileName")
	if fileName == "" {
		http.Error(w, "fileName is required", http.StatusBadRequest)
		return
	}

	fullPath := Config.IndexesPath + "/" + fileName

	file, err := filesystem.CreateIndex(fullPath)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create index file: %v", err), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	fmt.Fprintf(w, "index file %s created successfully\n", fileName)
	log.Default().Println("index file created successfully:", fileName)
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
