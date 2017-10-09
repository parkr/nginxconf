package mimetypes

import (
	"encoding/json"
	"net/http"
	"sort"
)

const MimeTypeDBURL = "https://raw.githubusercontent.com/jshttp/mime-db/master/db.json"

func Fetch() ([]MimeType, error) {
	resp, err := http.Get(MimeTypeDBURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	db := map[string]mimeTypeDetails{}
	err = json.NewDecoder(resp.Body).Decode(&db)
	if err != nil {
		return nil, err
	}

	var types []MimeType
	for mimeType, details := range db {
		if len(details.Extensions) > 0 {
			types = append(types, MimeType{
				Name:         mimeType,
				Compressible: details.Compressible,
				Extensions:   details.Extensions,
			})
		}
	}
	sort.Stable(MimeTypes(types))

	return types, nil
}
