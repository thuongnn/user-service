package api

import (
	"user-service/src/common/api"
	"user-service/src/common/security"
)

// BaseController ...
type BaseController struct {
	api.BaseAPI
	// SecurityCtx is the security context used to authN &authZ
	SecurityCtx security.Context
}

// Prepare inits security context
// context
func (b *BaseController) Prepare() {
	b.SecurityCtx = security.FromRequest(b.Ctx.Request)
}