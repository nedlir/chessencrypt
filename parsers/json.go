package parsers

import (
	"encoding/json"
	"fmt"
	"os"
)

type JSONParser struct{}

func NewJSONParser() *JSONParser {
	return &JSONParser{}
}

func (jp *JSONParser) LoadToMapFromFile(filepath string) (map[string]map[string]bool, error) {
	data, err := jp.readFile(filepath)
	if err != nil {
		return nil, err
	}

	var jsonData map[string][]string
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, fmt.Errorf("failed to parse JSON from %s: %w", filepath, err)
	}

	result := make(map[string]map[string]bool)
	for key, values := range jsonData {
		result[key] = make(map[string]bool)
		for _, value := range values {
			result[key][value] = true
		}
	}

	return result, nil
}

func (jp *JSONParser) readFile(filepath string) ([]byte, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filepath, err)
	}
	return data, nil
}
