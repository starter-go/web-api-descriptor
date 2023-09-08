package vo

import "github.com/bitwormhole/web-api-descriptor/app/web/dto"

// Schema 是包含 API 方案的文档
type Schema struct {
	Base

	Schema     dto.Schema       `json:"schema"`
	Namespaces []*dto.Namespace `json:"namespaces"`
}
