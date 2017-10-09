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

	// Sort the MIME types
	var mimeTypeKeys []string
	for mimeType, details := range db {
		if len(details.Extensions) > 0 {
			mimeTypeKeys = append(mimeTypeKeys, mimeType)
		}
	}
	sort.Strings(mimeTypeKeys)

	// One by one in sorted order, add each eligible MIME type.
	var types []MimeType
	extensionsSeen := map[string]bool{}
	for _, mimeType := range mimeTypeKeys {
		var uniqueExtensions []string
		for _, ext := range db[mimeType].Extensions {
			if _, ok := extensionsSeen[ext]; !ok {
				extensionsSeen[ext] = true
				uniqueExtensions = append(uniqueExtensions, ext)
			}
		}
		if len(uniqueExtensions) > 0 {
			types = append(types, MimeType{
				Name:         mimeType,
				Compressible: db[mimeType].Compressible,
				Extensions:   uniqueExtensions,
			})
		}
	}

	return types, nil
}
