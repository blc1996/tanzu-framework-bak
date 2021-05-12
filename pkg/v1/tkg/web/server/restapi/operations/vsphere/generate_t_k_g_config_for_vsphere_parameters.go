// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "github.com/vmware-tanzu-private/core/pkg/v1/tkg/web/server/models"
)

// NewGenerateTKGConfigForVsphereParams creates a new GenerateTKGConfigForVsphereParams object
// no default values defined in spec.
func NewGenerateTKGConfigForVsphereParams() GenerateTKGConfigForVsphereParams {

	return GenerateTKGConfigForVsphereParams{}
}

// GenerateTKGConfigForVsphereParams contains all the bound params for the generate t k g config for vsphere operation
// typically these are obtained from a http.Request
//
// swagger:parameters generateTKGConfigForVsphere
type GenerateTKGConfigForVsphereParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*params to generate tkg configuration for vsphere
	  Required: true
	  In: body
	*/
	Params *models.VsphereRegionalClusterParams
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGenerateTKGConfigForVsphereParams() beforehand.
func (o *GenerateTKGConfigForVsphereParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.VsphereRegionalClusterParams
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("params", "body"))
			} else {
				res = append(res, errors.NewParseError("params", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Params = &body
			}
		}
	} else {
		res = append(res, errors.Required("params", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
