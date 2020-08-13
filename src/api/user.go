package api

import "errors"

// UserAPI handles requests for user management
type UserAPI struct {
	BaseController
}

// Prepare ...
func (ua *UserAPI) Prepare() {
	ua.BaseController.Prepare()
}

func (ua *UserAPI) Get() {
	if !ua.SecurityCtx.IsAuthenticated() {
		ua.SendUnAuthorizedError(errors.New("UnAuthorized"))
		return
	}

	ua.Data["json"] = ua.SecurityCtx.User()
	ua.ServeJSON()
}
