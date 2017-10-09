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
	var extensionsSeen map[string]bool
	for mimeType, details := range db {
		var uniqueExtensions []string
		for _, ext := range details.Extensions {
			if _, ok := extensionsSeen[ext]; ok {
				extensionsSeen[ext] = true
				uniqueExtensions = append(uniqueExtensions, ext)
			}
		}
		if len(uniqueExtensions) > 0 {
			types = append(types, MimeType{
				Name:         mimeType,
				Compressible: details.Compressible,
				Extensions:   uniqueExtensions,
			})

		}
	}
	sort.Stable(MimeTypes(types))

	return types, nil
}
