package dto

import "github.com/bitwormhole/web-api-descriptor/app/data/dxo"

// API ...
type API struct {
	ID dxo.APIID `json:"id"`

	Base

	Name        string `json:"name"`
	Version     string `json:"version"`
	Title       string `json:"title"`
	Description string `json:"description"`

	Type   string `json:"type"`
	Path   string `json:"path"`
	Method string `json:"method"`

	Params []string `json:"params"`
}
