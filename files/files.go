package files

import (
	"fmt"
	"os"
)

func ReadFile(filepath string) ([]byte, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filepath, err)
	}
	return data, nil
}

// func WriteFile() { } // will be implemented later
