package jsonhandler

import (
	"encoding/json"
	"fmt"

	"github.com/nedlir/chessencrypt/utils/fileshandler"
)

type StringSet map[string]bool

type SetMap map[string]StringSet

func LoadToMapFromFile(filepath string) (SetMap, error) {
	data, err := fileshandler.ReadFile(filepath)
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
