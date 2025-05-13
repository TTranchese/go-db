package filesystem

import (
	"fmt"
	"os"
)

func CreateIndex(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to create index file: %w", err)
	}
	return file, nil
}
