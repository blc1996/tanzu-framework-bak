// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

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

	models "github.com/vmware-tanzu-private/core/pkg/v1/tkg/web/server/models"
)

// NewSetAzureEndpointParams creates a new SetAzureEndpointParams object
// with the default values initialized.
func NewSetAzureEndpointParams() *SetAzureEndpointParams {
	var ()
	return &SetAzureEndpointParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetAzureEndpointParamsWithTimeout creates a new SetAzureEndpointParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetAzureEndpointParamsWithTimeout(timeout time.Duration) *SetAzureEndpointParams {
	var ()
	return &SetAzureEndpointParams{

		timeout: timeout,
	}
}

// NewSetAzureEndpointParamsWithContext creates a new SetAzureEndpointParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetAzureEndpointParamsWithContext(ctx context.Context) *SetAzureEndpointParams {
	var ()
	return &SetAzureEndpointParams{

		Context: ctx,
	}
}

// NewSetAzureEndpointParamsWithHTTPClient creates a new SetAzureEndpointParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetAzureEndpointParamsWithHTTPClient(client *http.Client) *SetAzureEndpointParams {
	var ()
	return &SetAzureEndpointParams{
		HTTPClient: client,
	}
}

/*SetAzureEndpointParams contains all the parameters to send to the API endpoint
for the set azure endpoint operation typically these are written to a http.Request
*/
type SetAzureEndpointParams struct {

	/*AccountParams
	  Azure account parameters

	*/
	AccountParams *models.AzureAccountParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set azure endpoint params
func (o *SetAzureEndpointParams) WithTimeout(timeout time.Duration) *SetAzureEndpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set azure endpoint params
func (o *SetAzureEndpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set azure endpoint params
func (o *SetAzureEndpointParams) WithContext(ctx context.Context) *SetAzureEndpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set azure endpoint params
func (o *SetAzureEndpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set azure endpoint params
func (o *SetAzureEndpointParams) WithHTTPClient(client *http.Client) *SetAzureEndpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set azure endpoint params
func (o *SetAzureEndpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountParams adds the accountParams to the set azure endpoint params
func (o *SetAzureEndpointParams) WithAccountParams(accountParams *models.AzureAccountParams) *SetAzureEndpointParams {
	o.SetAccountParams(accountParams)
	return o
}

// SetAccountParams adds the accountParams to the set azure endpoint params
func (o *SetAzureEndpointParams) SetAccountParams(accountParams *models.AzureAccountParams) {
	o.AccountParams = accountParams
}

// WriteToRequest writes these params to a swagger request
func (o *SetAzureEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountParams != nil {
		if err := r.SetBodyParam(o.AccountParams); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
