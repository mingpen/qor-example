package auth

import (
	"fmt"
	"net/http"

	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor-example/models/users"
	"github.com/qor/roles"
)

func init() {
	roles.Register("admin", func(req *http.Request, currentUser interface{}) bool {
		return currentUser != nil && currentUser.(*users.User).Role == "Admin"
	})
}

// AdminAuth ...
type AdminAuth struct {
}

// LoginURL ...
func (AdminAuth) LoginURL(c *admin.Context) string {
	return "/auth/login"
}

// LogoutURL ...
func (AdminAuth) LogoutURL(c *admin.Context) string {
	return "/auth/logout"
}

// GetCurrentUser ...
func (AdminAuth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	currentUser := Auth.GetCurrentUser(c.Request)
	if currentUser != nil {
			// http.Redirect(w, req, admin.Auth.LoginURL(context), http.StatusSeeOther)
		qorCurrentUser, ok := currentUser.(qor.CurrentUser)
		if !ok {
			fmt.Printf("User %#v haven't implement qor.CurrentUser interface\n", currentUser)
			return nil
		}
user ,ok := currentUser.(*users.User)
if !ok || user.Role != "Admin"{
	return nil
}
		return qorCurrentUser
	}
	return nil
}
