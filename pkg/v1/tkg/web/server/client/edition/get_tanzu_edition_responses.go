// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware-tanzu-private/core/pkg/v1/tkg/web/server/models"
)

// GetTanzuEditionReader is a Reader for the GetTanzuEdition structure.
type GetTanzuEditionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTanzuEditionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTanzuEditionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTanzuEditionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTanzuEditionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTanzuEditionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTanzuEditionOK creates a GetTanzuEditionOK with default headers values
func NewGetTanzuEditionOK() *GetTanzuEditionOK {
	return &GetTanzuEditionOK{}
}

/*GetTanzuEditionOK handles this case with default header values.

Successful retrieval of tanzu edition
*/
type GetTanzuEditionOK struct {
	Payload string
}

func (o *GetTanzuEditionOK) Error() string {
	return fmt.Sprintf("[GET /api/edition][%d] getTanzuEditionOK  %+v", 200, o.Payload)
}

func (o *GetTanzuEditionOK) GetPayload() string {
	return o.Payload
}

func (o *GetTanzuEditionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTanzuEditionBadRequest creates a GetTanzuEditionBadRequest with default headers values
func NewGetTanzuEditionBadRequest() *GetTanzuEditionBadRequest {
	return &GetTanzuEditionBadRequest{}
}

/*GetTanzuEditionBadRequest handles this case with default header values.

Bad Request
*/
type GetTanzuEditionBadRequest struct {
	Payload *models.Error
}

func (o *GetTanzuEditionBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/edition][%d] getTanzuEditionBadRequest  %+v", 400, o.Payload)
}

func (o *GetTanzuEditionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTanzuEditionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTanzuEditionUnauthorized creates a GetTanzuEditionUnauthorized with default headers values
func NewGetTanzuEditionUnauthorized() *GetTanzuEditionUnauthorized {
	return &GetTanzuEditionUnauthorized{}
}

/*GetTanzuEditionUnauthorized handles this case with default header values.

Incorrect credentials
*/
type GetTanzuEditionUnauthorized struct {
	Payload *models.Error
}

func (o *GetTanzuEditionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/edition][%d] getTanzuEditionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTanzuEditionUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTanzuEditionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTanzuEditionInternalServerError creates a GetTanzuEditionInternalServerError with default headers values
func NewGetTanzuEditionInternalServerError() *GetTanzuEditionInternalServerError {
	return &GetTanzuEditionInternalServerError{}
}

/*GetTanzuEditionInternalServerError handles this case with default header values.

Internal server error
*/
type GetTanzuEditionInternalServerError struct {
	Payload *models.Error
}

func (o *GetTanzuEditionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/edition][%d] getTanzuEditionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTanzuEditionInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTanzuEditionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
