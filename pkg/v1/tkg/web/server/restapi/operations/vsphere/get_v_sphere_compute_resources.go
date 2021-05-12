// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetVSphereComputeResourcesHandlerFunc turns a function with the right signature into a get v sphere compute resources handler
type GetVSphereComputeResourcesHandlerFunc func(GetVSphereComputeResourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetVSphereComputeResourcesHandlerFunc) Handle(params GetVSphereComputeResourcesParams) middleware.Responder {
	return fn(params)
}

// GetVSphereComputeResourcesHandler interface for that can handle valid get v sphere compute resources params
type GetVSphereComputeResourcesHandler interface {
	Handle(GetVSphereComputeResourcesParams) middleware.Responder
}

// NewGetVSphereComputeResources creates a new http.Handler for the get v sphere compute resources operation
func NewGetVSphereComputeResources(ctx *middleware.Context, handler GetVSphereComputeResourcesHandler) *GetVSphereComputeResources {
	return &GetVSphereComputeResources{Context: ctx, Handler: handler}
}

/*GetVSphereComputeResources swagger:route GET /api/providers/vsphere/compute vsphere getVSphereComputeResources

Retrieve vSphere compute resources

*/
type GetVSphereComputeResources struct {
	Context *middleware.Context
	Handler GetVSphereComputeResourcesHandler
}

func (o *GetVSphereComputeResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetVSphereComputeResourcesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
