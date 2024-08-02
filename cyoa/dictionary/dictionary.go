package dictionary

import (
	"encoding/json"
	"fmt"
	"os"
)

func Parse(path string) (*Dictionary, error) {
	dict := make(Dictionary)
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("can't open file %s, %s", path, err.Error())
	}
	err = json.Unmarshal([]byte(content), &dict)
	if err != nil {
		return nil, fmt.Errorf("can't process dictionary from json, %s", err.Error())
	}

	return &dict, nil
}
