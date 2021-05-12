// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetVSphereNodeTypesHandlerFunc turns a function with the right signature into a get v sphere node types handler
type GetVSphereNodeTypesHandlerFunc func(GetVSphereNodeTypesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetVSphereNodeTypesHandlerFunc) Handle(params GetVSphereNodeTypesParams) middleware.Responder {
	return fn(params)
}

// GetVSphereNodeTypesHandler interface for that can handle valid get v sphere node types params
type GetVSphereNodeTypesHandler interface {
	Handle(GetVSphereNodeTypesParams) middleware.Responder
}

// NewGetVSphereNodeTypes creates a new http.Handler for the get v sphere node types operation
func NewGetVSphereNodeTypes(ctx *middleware.Context, handler GetVSphereNodeTypesHandler) *GetVSphereNodeTypes {
	return &GetVSphereNodeTypes{Context: ctx, Handler: handler}
}

/*GetVSphereNodeTypes swagger:route GET /api/providers/vsphere/nodetypes vsphere getVSphereNodeTypes

Retrieve vSphere supported kubernetes control plane node types

*/
type GetVSphereNodeTypes struct {
	Context *middleware.Context
	Handler GetVSphereNodeTypesHandler
}

func (o *GetVSphereNodeTypes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetVSphereNodeTypesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
