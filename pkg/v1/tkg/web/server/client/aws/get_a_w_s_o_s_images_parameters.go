// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAWSOSImagesParams creates a new GetAWSOSImagesParams object
// with the default values initialized.
func NewGetAWSOSImagesParams() *GetAWSOSImagesParams {
	var ()
	return &GetAWSOSImagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAWSOSImagesParamsWithTimeout creates a new GetAWSOSImagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAWSOSImagesParamsWithTimeout(timeout time.Duration) *GetAWSOSImagesParams {
	var ()
	return &GetAWSOSImagesParams{

		timeout: timeout,
	}
}

// NewGetAWSOSImagesParamsWithContext creates a new GetAWSOSImagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAWSOSImagesParamsWithContext(ctx context.Context) *GetAWSOSImagesParams {
	var ()
	return &GetAWSOSImagesParams{

		Context: ctx,
	}
}

// NewGetAWSOSImagesParamsWithHTTPClient creates a new GetAWSOSImagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAWSOSImagesParamsWithHTTPClient(client *http.Client) *GetAWSOSImagesParams {
	var ()
	return &GetAWSOSImagesParams{
		HTTPClient: client,
	}
}

/*GetAWSOSImagesParams contains all the parameters to send to the API endpoint
for the get a w s o s images operation typically these are written to a http.Request
*/
type GetAWSOSImagesParams struct {

	/*Region
	  AWS region, e.g. us-west-2

	*/
	Region string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get a w s o s images params
func (o *GetAWSOSImagesParams) WithTimeout(timeout time.Duration) *GetAWSOSImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get a w s o s images params
func (o *GetAWSOSImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get a w s o s images params
func (o *GetAWSOSImagesParams) WithContext(ctx context.Context) *GetAWSOSImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get a w s o s images params
func (o *GetAWSOSImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get a w s o s images params
func (o *GetAWSOSImagesParams) WithHTTPClient(client *http.Client) *GetAWSOSImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get a w s o s images params
func (o *GetAWSOSImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegion adds the region to the get a w s o s images params
func (o *GetAWSOSImagesParams) WithRegion(region string) *GetAWSOSImagesParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get a w s o s images params
func (o *GetAWSOSImagesParams) SetRegion(region string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *GetAWSOSImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param region
	qrRegion := o.Region
	qRegion := qrRegion
	if qRegion != "" {
		if err := r.SetQueryParam("region", qRegion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
