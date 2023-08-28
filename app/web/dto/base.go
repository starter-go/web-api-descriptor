package dto

import "github.com/starter-go/security/rbac"

// Base 基本的 DTO (for web-api-descriptor)
type Base struct {
	rbac.BaseDTO
}
