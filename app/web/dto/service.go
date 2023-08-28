package dto

import "github.com/bitwormhole/web-api-descriptor/app/data/dxo"

// Service ...
type Service struct {
	ID dxo.ServiceID `json:"id"`

	Base

	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Protocol    string `json:"protocol"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
}
