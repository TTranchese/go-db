package api

import (
	"fmt"
	"go-db/src/filesystem"
	"go-db/src/models"
	"log"
	"net/http"
)

func StartServer(Config models.Config, Port string) {
	fmt.Println("starting GO DB...")

	http.HandleFunc("/create-index", func(w http.ResponseWriter, r *http.Request) {
		handleCreateIndex(w, r, Config)
	})

	log.Default().Println("server is running on port ", Port)
	log.Fatal(http.ListenAndServe(Port, nil))
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
