package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"net/http"
	"strconv"
	"user-service/src/common/utils"
)

// BaseAPI wraps common methods for controllers to host API
type BaseAPI struct {
	beego.Controller
}

// GetStringFromPath gets the param from path and returns it as string
func (b *BaseAPI) GetStringFromPath(key string) string {
	return b.Ctx.Input.Param(key)
}

// GetInt64FromPath gets the param from path and returns it as int64
func (b *BaseAPI) GetInt64FromPath(key string) (int64, error) {
	value := b.Ctx.Input.Param(key)
	return strconv.ParseInt(value, 10, 64)
}

// ParamExistsInPath returns true when param exists in the path
func (b *BaseAPI) ParamExistsInPath(key string) bool {
	return b.GetStringFromPath(key) != ""
}

// DecodeJSONReq decodes a json request
func (b *BaseAPI) DecodeJSONReq(v interface{}) error {
	err := json.Unmarshal(b.Ctx.Input.CopyBody(1<<32), v)
	if err != nil {
		fmt.Errorf("Error while decoding the json request, error: %v, %v",
			err, string(b.Ctx.Input.CopyBody(1 << 32)[:]))
		return errors.New("Invalid json request")
	}
	return nil
}

// Validate validates v if it implements interface validation.ValidFormer
func (b *BaseAPI) Validate(v interface{}) (bool, error) {
	validator := validation.Validation{}
	isValid, err := validator.Valid(v)
	if err != nil {
		fmt.Errorf("failed to validate: %v", err)
		return false, err
	}

	if !isValid {
		message := ""
		for _, e := range validator.Errors {
			message += fmt.Sprintf("%s %s \n", e.Field, e.Message)
		}
		return false, errors.New(message)
	}
	return true, nil
}

// DecodeJSONReqAndValidate does both decoding and validation
func (b *BaseAPI) DecodeJSONReqAndValidate(v interface{}) (bool, error) {
	if err := b.DecodeJSONReq(v); err != nil {
		return false, err
	}
	return b.Validate(v)
}

// Redirect does redirection to resource URI with http header status code.
func (b *BaseAPI) Redirect(statusCode int, resouceID string) {
	requestURI := b.Ctx.Request.RequestURI
	resourceURI := requestURI + "/" + resouceID

	b.Ctx.Redirect(statusCode, resourceURI)
}

// RenderError provides shortcut to render http error
func (b *BaseAPI) RenderError(code int, text string) {
	errPayload := &utils.Error{
		Code:    code,
		Message: text,
	}

	w := b.Ctx.ResponseWriter
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Println(w, errPayload.String())
}

// SendUnAuthorizedError sends unauthorized error to the client.
func (b *BaseAPI) SendUnAuthorizedError(err error) {
	b.RenderError(http.StatusUnauthorized, err.Error())
}

// SendConflictError sends conflict error to the client.
func (b *BaseAPI) SendConflictError(err error) {
	b.RenderError(http.StatusConflict, err.Error())
}

// SendNotFoundError sends not found error to the client.
func (b *BaseAPI) SendNotFoundError(err error) {
	b.RenderError(http.StatusNotFound, err.Error())
}

// SendBadRequestError sends bad request error to the client.
func (b *BaseAPI) SendBadRequestError(err error) {
	b.RenderError(http.StatusBadRequest, err.Error())
}

// SendInternalServerError sends internal server error to the client.
// Note the detail info of err will not include in the response body.
// When you send an internal server error  to the client, you expect user to check the log
// to find out the root cause.
func (b *BaseAPI) SendInternalServerError(err error) {
	b.RenderError(http.StatusInternalServerError, err.Error())
}

// SendForbiddenError sends forbidden error to the client.
func (b *BaseAPI) SendForbiddenError(err error) {
	b.RenderError(http.StatusForbidden, err.Error())
}

// SendPreconditionFailedError sends conflict error to the client.
func (b *BaseAPI) SendPreconditionFailedError(err error) {
	b.RenderError(http.StatusPreconditionFailed, err.Error())
}

// SendStatusServiceUnavailableError sends service unavailable error to the client.
func (b *BaseAPI) SendStatusServiceUnavailableError(err error) {
	b.RenderError(http.StatusServiceUnavailable, err.Error())
}
