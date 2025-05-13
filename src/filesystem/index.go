package filesystem

import (
	"fmt"
	"os"
	"strings"
)

func CreateIndex(fileName string) (*os.File, error) {
	lastSeparator := strings.LastIndex(fileName, string(os.PathSeparator))

	if lastSeparator != -1 {
		dir := fileName[:lastSeparator]
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("failed to create directory: %w", err)
		}
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to create index file: %w", err)
	}
	return file, nil
}
