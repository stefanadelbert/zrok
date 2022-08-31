// Code generated by go-swagger; DO NOT EDIT.

package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateAccountHandlerFunc turns a function with the right signature into a create account handler
type CreateAccountHandlerFunc func(CreateAccountParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAccountHandlerFunc) Handle(params CreateAccountParams) middleware.Responder {
	return fn(params)
}

// CreateAccountHandler interface for that can handle valid create account params
type CreateAccountHandler interface {
	Handle(CreateAccountParams) middleware.Responder
}

// NewCreateAccount creates a new http.Handler for the create account operation
func NewCreateAccount(ctx *middleware.Context, handler CreateAccountHandler) *CreateAccount {
	return &CreateAccount{Context: ctx, Handler: handler}
}

/* CreateAccount swagger:route POST /account identity createAccount

CreateAccount create account API

*/
type CreateAccount struct {
	Context *middleware.Context
	Handler CreateAccountHandler
}

func (o *CreateAccount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateAccountParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // backend params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
