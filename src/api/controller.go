package api

import (
	"fmt"
	"go-db/src/filesystem"
	"go-db/src/models"
	"log"
	"net/http"
)

func HandleCreateIndex(w http.ResponseWriter, r *http.Request, Config models.Config) {
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

func HandleInsertKeyValue(w http.ResponseWriter, r *http.Request, Config models.Config) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fileName := r.URL.Query().Get("fileName")
	if fileName == "" {
		http.Error(w, "fileName is required", http.StatusBadRequest)
		return
	}

	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	if key == "" || value == "" {
		http.Error(w, "key and value are required", http.StatusBadRequest)
		return
	}

	fullPath := Config.IndexesPath + "/" + fileName

	err := filesystem.InsertKeyValue(fullPath, key, value)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to insert key-value pair: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "key-value pair (%s, %s) inserted successfully\n", key, value)
	log.Default().Println("key-value pair inserted successfully:", key, value)
}
