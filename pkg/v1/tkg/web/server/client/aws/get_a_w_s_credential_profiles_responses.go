// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware-tanzu-private/core/pkg/v1/tkg/web/server/models"
)

// GetAWSCredentialProfilesReader is a Reader for the GetAWSCredentialProfiles structure.
type GetAWSCredentialProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAWSCredentialProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAWSCredentialProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAWSCredentialProfilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAWSCredentialProfilesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAWSCredentialProfilesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAWSCredentialProfilesOK creates a GetAWSCredentialProfilesOK with default headers values
func NewGetAWSCredentialProfilesOK() *GetAWSCredentialProfilesOK {
	return &GetAWSCredentialProfilesOK{}
}

/*GetAWSCredentialProfilesOK handles this case with default header values.

Successful retrieval of AWS credentials profiles
*/
type GetAWSCredentialProfilesOK struct {
	Payload []string
}

func (o *GetAWSCredentialProfilesOK) Error() string {
	return fmt.Sprintf("[GET /api/providers/aws/profiles][%d] getAWSCredentialProfilesOK  %+v", 200, o.Payload)
}

func (o *GetAWSCredentialProfilesOK) GetPayload() []string {
	return o.Payload
}

func (o *GetAWSCredentialProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAWSCredentialProfilesBadRequest creates a GetAWSCredentialProfilesBadRequest with default headers values
func NewGetAWSCredentialProfilesBadRequest() *GetAWSCredentialProfilesBadRequest {
	return &GetAWSCredentialProfilesBadRequest{}
}

/*GetAWSCredentialProfilesBadRequest handles this case with default header values.

Bad request
*/
type GetAWSCredentialProfilesBadRequest struct {
	Payload *models.Error
}

func (o *GetAWSCredentialProfilesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/providers/aws/profiles][%d] getAWSCredentialProfilesBadRequest  %+v", 400, o.Payload)
}

func (o *GetAWSCredentialProfilesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAWSCredentialProfilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAWSCredentialProfilesUnauthorized creates a GetAWSCredentialProfilesUnauthorized with default headers values
func NewGetAWSCredentialProfilesUnauthorized() *GetAWSCredentialProfilesUnauthorized {
	return &GetAWSCredentialProfilesUnauthorized{}
}

/*GetAWSCredentialProfilesUnauthorized handles this case with default header values.

Incorrect credentials
*/
type GetAWSCredentialProfilesUnauthorized struct {
	Payload *models.Error
}

func (o *GetAWSCredentialProfilesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/providers/aws/profiles][%d] getAWSCredentialProfilesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAWSCredentialProfilesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAWSCredentialProfilesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAWSCredentialProfilesInternalServerError creates a GetAWSCredentialProfilesInternalServerError with default headers values
func NewGetAWSCredentialProfilesInternalServerError() *GetAWSCredentialProfilesInternalServerError {
	return &GetAWSCredentialProfilesInternalServerError{}
}

/*GetAWSCredentialProfilesInternalServerError handles this case with default header values.

Internal server error
*/
type GetAWSCredentialProfilesInternalServerError struct {
	Payload *models.Error
}

func (o *GetAWSCredentialProfilesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/providers/aws/profiles][%d] getAWSCredentialProfilesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAWSCredentialProfilesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAWSCredentialProfilesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
