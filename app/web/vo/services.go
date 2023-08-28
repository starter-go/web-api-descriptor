package vo

import "github.com/bitwormhole/web-api-descriptor/app/web/dto"

// Services ...
type Services struct {
	Base

	Services []*dto.Service `json:"services"`
}
