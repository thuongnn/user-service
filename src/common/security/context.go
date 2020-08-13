package security

import (
	"net/http"
	"user-service/src/dao"
	"user-service/src/models"
)

type Context struct {
	user *models.User
}

// IsAuthenticated returns true if the user has been authenticated
func (c *Context) IsAuthenticated() bool {
	return c.user != nil
}

// GetUsername returns the username of the authenticated user
// It returns null if the user has not been authenticated
func (c *Context) GetUsername() string {
	if !c.IsAuthenticated() {
		return ""
	}
	return c.user.Username
}

// User get the current user
func (c *Context) User() *models.User {
	c.user.Password = ""
	return c.user
}

// FromContext returns security context from the context
func FromRequest(req *http.Request) Context {
	tokenAuth, err := ExtractTokenMetadata(req)
	if err != nil {
		return Context{nil}
	}

	userId, err := FetchAuth(tokenAuth)
	if err != nil {
		return Context{nil}
	}

	user, err := dao.GetUser(models.User{UserID: userId})
	if err != nil {
		return Context{nil}
	}

	return Context{user}
}
