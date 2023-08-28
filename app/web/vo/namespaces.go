package vo

import "github.com/bitwormhole/web-api-descriptor/app/web/dto"

// Namespaces ...
type Namespaces struct {
	Base

	Namespaces []*dto.Namespace `json:"namespaces"`
}
