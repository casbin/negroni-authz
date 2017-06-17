package authz

// This plugin is based on Casbin: an authorization library that supports ACL, RBAC, ABAC
// View source at:
// https://github.com/casbin/casbin

import (
	"net/http"

	"github.com/casbin/casbin"
	"github.com/urfave/negroni"
)

// Authz is a middleware that controls the access to the HTTP service, it is based
// on Casbin, which supports access control models like ACL, RBAC, ABAC.
// The plugin determines whether to allow a request based on (user, path, method).
// user: the authenticated user name.
// path: the URL for the requested resource.
// method: one of HTTP methods like GET, POST, PUT, DELETE.
//
// This middleware should be inserted fairly early in the middleware stack to
// protect subsequent layers. All the denied requests will not go further.
//
// It's notable that this middleware should be behind the authentication (e.g.,
// HTTP basic authentication, OAuth), so this plugin can get the logged-in user name
// to perform the authorization.
func Authorizer(e *casbin.Enforcer) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if !CheckPermission(e, r) {
			http.Error(w, http.StatusText(403), 403)
		} else {
			next(w, r)
		}
	}
}

// GetUserName gets the user name from the request.
// Currently, only HTTP basic authentication is supported
func GetUserName(r *http.Request) string {
	username, _, _ := r.BasicAuth()
	return username
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func CheckPermission(e *casbin.Enforcer, r *http.Request) bool {
	user := GetUserName(r)
	method := r.Method
	path := r.URL.Path
	return e.Enforce(user, path, method)
}
