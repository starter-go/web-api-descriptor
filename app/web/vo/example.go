package vo

import "github.com/bitwormhole/web-api-descriptor/app/web/dto"

// Example ...
type Example struct {
	Base

	Items []*dto.Example `json:"items"`
}
