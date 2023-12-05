package dto

import "github.com/bitwormhole/web-api-descriptor/app/data/dxo"

// Service ...
type Service struct {
	ID dxo.ServiceID `json:"id"`

	Base

	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Schema      string `json:"schema"`   // API 方案 URL
	Protocol    string `json:"protocol"` // 服务协议名称（例如：https）
	Profiles    string `json:"profiles"` // like: "release,debug,dev,stable"
	Host        string `json:"host"`
	Port        int    `json:"port"`
}
