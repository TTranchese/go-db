package filesystem

import (
	"fmt"
	"go-db/src/models"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig() (models.Config, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		log.Default().Fatal("Error opening config file:", err)
		return models.Config{}, err
	}
	defer file.Close()

	var config models.Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Default().Fatal("Error decoding config file:", err)
		return config, nil
	}

	err = os.MkdirAll(config.IndexesPath, os.ModePerm)
	if err != nil {
		return models.Config{}, fmt.Errorf("failed to create indexes directory: %w", err)
	}

	log.Default().Println("IndexesPath is set to:", config.IndexesPath)
	return config, nil
}
