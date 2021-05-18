// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAWSEndpointHandlerFunc turns a function with the right signature into a get a w s endpoint handler
type GetAWSEndpointHandlerFunc func(GetAWSEndpointParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAWSEndpointHandlerFunc) Handle(params GetAWSEndpointParams) middleware.Responder {
	return fn(params)
}

// GetAWSEndpointHandler interface for that can handle valid get a w s endpoint params
type GetAWSEndpointHandler interface {
	Handle(GetAWSEndpointParams) middleware.Responder
}

// NewGetAWSEndpoint creates a new http.Handler for the get a w s endpoint operation
func NewGetAWSEndpoint(ctx *middleware.Context, handler GetAWSEndpointHandler) *GetAWSEndpoint {
	return &GetAWSEndpoint{Context: ctx, Handler: handler}
}

/*GetAWSEndpoint swagger:route GET /api/providers/aws aws getAWSEndpoint

Retrieve aws account params from environment variables

*/
type GetAWSEndpoint struct {
	Context *middleware.Context
	Handler GetAWSEndpointHandler
}

func (o *GetAWSEndpoint) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAWSEndpointParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
