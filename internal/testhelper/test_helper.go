package testhelper

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func LoadJsonFromFile(filePath string) ([]byte, error) {
	basePath := "./"
	filePath = filepath.Join(basePath, filepath.Clean(filePath))

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open test cases file: %v", err)
	}
	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read test cases file: %v", err)
	}
	return jsonData, nil
}
