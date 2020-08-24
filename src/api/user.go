package api

import (
	"encoding/json"
	"errors"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
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
	url, err := lookupServiceWithConsul("book-service")
	fmt.Println("URL: ", url)
	if err != nil {
		ua.SendStatusServiceUnavailableError(err)
		return
	}
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

func lookupServiceWithConsul(serviceName string) (string, error) {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		return "", err
	}
	services, err := consul.Agent().Services()
	if err != nil {
		return "", err
	}
	srvc := services[serviceName]
	address := srvc.Address
	port := srvc.Port
	return fmt.Sprintf("http://%s:%v", address, port), nil
}
