package dto

// Schema 表示一套 API 方案
type Schema struct {
	URL         string `json:"url"`
	Version     string `json:"version"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
