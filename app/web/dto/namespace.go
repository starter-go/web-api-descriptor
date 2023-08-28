package dto

import "github.com/bitwormhole/web-api-descriptor/app/data/dxo"

// Namespace ...
type Namespace struct {
	ID dxo.NamespaceID `json:"id"`

	Base

	Name        string `json:"name"`
	Version     string `json:"version"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URI         string `json:"uri"`

	Types []string `json:"types"`

	APIs []*API `json:"apis"`
}
