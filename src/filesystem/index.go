package filesystem

import (
	"bufio"
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

func InsertKeyValue(fileName, key, value string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open index file: %w", err)
	}
	defer file.Close()

	var updatedContent strings.Builder
	keyFound := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, key+":") {
			updatedContent.WriteString(fmt.Sprintf("%s:%s\n", key, value))
			keyFound = true
		} else {
			updatedContent.WriteString(line + "\n")
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read index file: %w", err)
	}

	if !keyFound {
		updatedContent.WriteString(fmt.Sprintf("%s:%s\n", key, value))
	}

	err = os.WriteFile(fileName, []byte(updatedContent.String()), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to index file: %w", err)
	}

	return nil
}
