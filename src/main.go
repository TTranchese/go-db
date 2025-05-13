package main

import (
	"fmt"
	"go-db/src/filesystem"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Staring GO DB...")

	http.HandleFunc("/create-index", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
			return
		}

		fileName := r.URL.Query().Get("fileName")
		if fileName == "" {
			http.Error(w, "fileName is required", http.StatusBadRequest)
			return
		}

		file, err := filesystem.CreateIndex(fileName)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to create index file: %v", err), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		fmt.Fprintf(w, "index file %s create successfully", fileName)
	})

	port := ":8080"
	fmt.Printf("server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
