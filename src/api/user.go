package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"user-service/src/models"
)

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

	var b []models.Book
	url := "http://127.0.0.1:8081"
	fmt.Println("URL: ", url)

	client := &http.Client{}
	resp, err := client.Get(url + "/all")
	if err != nil {
		ua.SendStatusServiceUnavailableError(err)
		return
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&b); err != nil {
		ua.SendStatusServiceUnavailableError(err)
		return
	}

	ua.Data["json"] = b
	ua.ServeJSON()
}
