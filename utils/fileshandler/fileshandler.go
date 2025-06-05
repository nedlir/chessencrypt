package fileshandler

import (
	"fmt"
	"os"
)

const MaxFileSize = 1024 * 1024 // 1 MB

func ReadFile(filepath string) ([]byte, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to get file info for %s: %w", filepath, err)
	}

	if fileInfo.Size() > MaxFileSize {
		return nil, fmt.Errorf("file %s too large: %d bytes (max %d bytes)", filepath, fileInfo.Size(), MaxFileSize)
	}

	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filepath, err)
	}
	return data, nil
}

func GetFileSize(filepath string) (int64, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return 0, fmt.Errorf("failed to get file info for %s: %w", filepath, err)
	}
	return fileInfo.Size(), nil
}

func WriteFile(filepath string, data []byte) error {
	err := os.WriteFile(filepath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", filepath, err)
	}
	return nil
}
