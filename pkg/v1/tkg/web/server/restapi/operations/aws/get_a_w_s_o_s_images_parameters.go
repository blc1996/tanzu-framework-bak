// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAWSOSImagesParams creates a new GetAWSOSImagesParams object
// no default values defined in spec.
func NewGetAWSOSImagesParams() GetAWSOSImagesParams {

	return GetAWSOSImagesParams{}
}

// GetAWSOSImagesParams contains all the bound params for the get a w s o s images operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAWSOSImages
type GetAWSOSImagesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*AWS region, e.g. us-west-2
	  Required: true
	  In: query
	*/
	Region string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAWSOSImagesParams() beforehand.
func (o *GetAWSOSImagesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qRegion, qhkRegion, _ := qs.GetOK("region")
	if err := o.bindRegion(qRegion, qhkRegion, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindRegion binds and validates parameter Region from query.
func (o *GetAWSOSImagesParams) bindRegion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("region", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("region", "query", raw); err != nil {
		return err
	}

	o.Region = raw

	return nil
}
