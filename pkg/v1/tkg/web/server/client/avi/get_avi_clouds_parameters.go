// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package avi

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

// NewGetAviCloudsParams creates a new GetAviCloudsParams object
// with the default values initialized.
func NewGetAviCloudsParams() *GetAviCloudsParams {
	return &GetAviCloudsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAviCloudsParamsWithTimeout creates a new GetAviCloudsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAviCloudsParamsWithTimeout(timeout time.Duration) *GetAviCloudsParams {
	return &GetAviCloudsParams{

		timeout: timeout,
	}
}

// NewGetAviCloudsParamsWithContext creates a new GetAviCloudsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAviCloudsParamsWithContext(ctx context.Context) *GetAviCloudsParams {
	return &GetAviCloudsParams{

		Context: ctx,
	}
}

// NewGetAviCloudsParamsWithHTTPClient creates a new GetAviCloudsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAviCloudsParamsWithHTTPClient(client *http.Client) *GetAviCloudsParams {
	return &GetAviCloudsParams{
		HTTPClient: client,
	}
}

/*GetAviCloudsParams contains all the parameters to send to the API endpoint
for the get avi clouds operation typically these are written to a http.Request
*/
type GetAviCloudsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get avi clouds params
func (o *GetAviCloudsParams) WithTimeout(timeout time.Duration) *GetAviCloudsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get avi clouds params
func (o *GetAviCloudsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get avi clouds params
func (o *GetAviCloudsParams) WithContext(ctx context.Context) *GetAviCloudsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get avi clouds params
func (o *GetAviCloudsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get avi clouds params
func (o *GetAviCloudsParams) WithHTTPClient(client *http.Client) *GetAviCloudsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get avi clouds params
func (o *GetAviCloudsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAviCloudsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
