package api

type HomeAPI struct {
	BaseController
}

// Prepare ...
func (h *HomeAPI) Prepare() {
}

// default home page
func (h *HomeAPI) Get() {
	h.Data["json"] = "Welcome to User API !"
	h.ServeJSON()
}
