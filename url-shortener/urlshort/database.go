package urlshort

import (
	"net/http"
	"url-shortener/database"
)

func DBHandler(db *database.Database, fallback http.Handler) (http.HandlerFunc, error) {
	paths, err := db.List()
	if err != nil {
		return nil, err
	}

	return MapHandler(paths, fallback), nil
}
