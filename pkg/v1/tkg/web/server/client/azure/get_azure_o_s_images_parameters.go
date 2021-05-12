// Code generated by go-swagger; DO NOT EDIT.

package azure

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

// NewGetAzureOSImagesParams creates a new GetAzureOSImagesParams object
// with the default values initialized.
func NewGetAzureOSImagesParams() *GetAzureOSImagesParams {

	return &GetAzureOSImagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAzureOSImagesParamsWithTimeout creates a new GetAzureOSImagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAzureOSImagesParamsWithTimeout(timeout time.Duration) *GetAzureOSImagesParams {

	return &GetAzureOSImagesParams{

		timeout: timeout,
	}
}

// NewGetAzureOSImagesParamsWithContext creates a new GetAzureOSImagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAzureOSImagesParamsWithContext(ctx context.Context) *GetAzureOSImagesParams {

	return &GetAzureOSImagesParams{

		Context: ctx,
	}
}

// NewGetAzureOSImagesParamsWithHTTPClient creates a new GetAzureOSImagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAzureOSImagesParamsWithHTTPClient(client *http.Client) *GetAzureOSImagesParams {

	return &GetAzureOSImagesParams{
		HTTPClient: client,
	}
}

/*GetAzureOSImagesParams contains all the parameters to send to the API endpoint
for the get azure o s images operation typically these are written to a http.Request
*/
type GetAzureOSImagesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get azure o s images params
func (o *GetAzureOSImagesParams) WithTimeout(timeout time.Duration) *GetAzureOSImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get azure o s images params
func (o *GetAzureOSImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get azure o s images params
func (o *GetAzureOSImagesParams) WithContext(ctx context.Context) *GetAzureOSImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get azure o s images params
func (o *GetAzureOSImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get azure o s images params
func (o *GetAzureOSImagesParams) WithHTTPClient(client *http.Client) *GetAzureOSImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get azure o s images params
func (o *GetAzureOSImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAzureOSImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
