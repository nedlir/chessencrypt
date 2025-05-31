package json

import (
	"chessencryption/files"
	"encoding/json"
	"fmt"
)

type StringSet map[string]bool

type SetMap map[string]StringSet

func LoadToMapFromFile(filepath string) (SetMap, error) {
	data, err := files.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var jsonData map[string][]string
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, fmt.Errorf("failed to parse JSON from %s: %w", filepath, err)
	}

	result := make(SetMap, len(jsonData))
	for key, values := range jsonData {
		// make a Set for each key
		set := make(StringSet, len(values))
		for _, value := range values {
			set[value] = true
		}
		result[key] = set
	}

	return result, nil
}
