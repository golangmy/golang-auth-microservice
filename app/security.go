// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Authentication API": Application Security
//
// Command:
// $ goagen
// --design=github.com/rymccue/golang-auth-microservice/design
// --out=$(GOPATH)/src/github.com/rymccue/golang-auth-microservice
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

type (
	// Private type used to store auth handler info in request context
	authMiddlewareKey string
)

// UseJWTMiddleware mounts the jwt auth middleware onto the service.
func UseJWTMiddleware(service *goa.Service, middleware goa.Middleware) {
	service.Context = context.WithValue(service.Context, authMiddlewareKey("jwt"), middleware)
}

// NewJWTSecurity creates a jwt security definition.
func NewJWTSecurity() *goa.JWTSecurity {
	def := goa.JWTSecurity{
		In:       goa.LocHeader,
		Name:     "Authorization",
		TokenURL: "",
	}
	def.Description = "Use JWT to authenticate"
	return &def
}

// handleSecurity creates a handler that runs the auth middleware for the security scheme.
func handleSecurity(schemeName string, h goa.Handler, scopes ...string) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		scheme := ctx.Value(authMiddlewareKey(schemeName))
		am, ok := scheme.(goa.Middleware)
		if !ok {
			return goa.NoAuthMiddleware(schemeName)
		}
		ctx = goa.WithRequiredScopes(ctx, scopes)
		return am(h)(ctx, rw, req)
	}
}