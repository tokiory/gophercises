package urlshort

import (
	"encoding/json"
	"net/http"
)

func JSONHandler(jsonData []byte, fallback http.HandlerFunc) (http.HandlerFunc, error) {
	var paths PathList
	err := json.Unmarshal(jsonData, &paths)
	if err != nil {
		return nil, err
	}

	return MapHandler(paths.Hashmap(), fallback), nil
}
