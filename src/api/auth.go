package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"user-service/src/common/security"
	"user-service/src/dao"
	"user-service/src/models"
)

type AuthAPI struct {
	BaseController
}

// Prepare ...
func (au *AuthAPI) Prepare() {
}

// Post ...
func (au *AuthAPI) Register() {
	user := models.User{}
	if err := au.DecodeJSONReq(&user); err != nil {
		au.SendBadRequestError(err)
		return
	}

	userExist, err := dao.UserExists(&user)
	if err != nil {
		fmt.Printf("Error occurred in Register: %v", err)
		au.SendInternalServerError(errors.New("internal error"))
		return
	}
	if userExist {
		fmt.Println("username has already been used!")
		au.SendConflictError(errors.New("username has already been used"))
		return
	}

	userID, err := dao.Register(&user)
	if err != nil {
		fmt.Printf("Error occurred in Register: %v", err)
		au.SendInternalServerError(errors.New("internal error"))
		return
	}

	au.Redirect(http.StatusCreated, strconv.FormatInt(userID, 10))
}

func (au *AuthAPI) Login() {
	username := au.GetString("username")
	password := au.GetString("password")

	user, err := dao.LoginByDb(models.AuthModel{
		Username: username,
		Password: password,
	})
	if err != nil {
		fmt.Printf("Error occurred in UserLogin: %v", err)
		au.SendUnAuthorizedError(errors.New(""))
		return
	}

	if user == nil {
		au.SendUnAuthorizedError(errors.New(""))
		return
	}

	td, err := security.CreateToken(user.UserID)
	if err != nil {
		au.SendStatusServiceUnavailableError(err)
		return
	}

	saveErr := security.CreateAuth(user.UserID, td)
	if saveErr != nil {
		au.SendStatusServiceUnavailableError(err)
		return
	}

	tokens := map[string]string{
		"access_token":  td.AccessToken,
		"refresh_token": td.RefreshToken,
	}

	au.Data["json"] = tokens
	au.ServeJSON()
}

func (au *AuthAPI) RefreshToken() {
	userId, err := security.RefreshToken(au.Ctx.Request)
	if err != nil {
		au.SendUnAuthorizedError(err)
		return
	}

	td, err := security.CreateToken(userId)
	if err != nil {
		au.SendStatusServiceUnavailableError(err)
		return
	}

	saveErr := security.CreateAuth(userId, td)
	if saveErr != nil {
		au.SendStatusServiceUnavailableError(err)
		return
	}

	tokens := map[string]string{
		"access_token":  td.AccessToken,
		"refresh_token": td.RefreshToken,
	}

	au.Data["json"] = tokens
	au.ServeJSON()
}

func (au *AuthAPI) Logout() {
	tokenAuth, err := security.ExtractTokenMetadata(au.Ctx.Request)
	if err != nil {
		au.SendUnAuthorizedError(errors.New("unauthorized"))
		return
	}

	deleted, delErr := security.DeleteAuth(tokenAuth.AccessUUID)
	if delErr != nil || deleted == 0 { //if any goes wrong
		au.SendUnAuthorizedError(errors.New("unauthorized"))
		return
	}

	au.Data["json"] = "Successfully logged out"
	au.ServeJSON()
}
