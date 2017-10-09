package mimetypes

type MimeTypes []MimeType

func (a MimeTypes) Len() int           { return len(a) }
func (a MimeTypes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a MimeTypes) Less(i, j int) bool { return a[i].Name < a[j].Name }

type MimeType struct {
	Name         string
	Compressible bool
	Extensions   []string
}

type mimeTypeDetails struct {
	Source       string   `json:"source"`
	Compressible bool     `json:"compressible"`
	Extensions   []string `json:"extensions"`
}
