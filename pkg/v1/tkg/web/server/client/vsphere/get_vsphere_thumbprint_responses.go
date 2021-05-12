// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware-tanzu-private/core/pkg/v1/tkg/web/server/models"
)

// GetVsphereThumbprintReader is a Reader for the GetVsphereThumbprint structure.
type GetVsphereThumbprintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVsphereThumbprintReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVsphereThumbprintOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVsphereThumbprintBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetVsphereThumbprintUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVsphereThumbprintInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVsphereThumbprintOK creates a GetVsphereThumbprintOK with default headers values
func NewGetVsphereThumbprintOK() *GetVsphereThumbprintOK {
	return &GetVsphereThumbprintOK{}
}

/*GetVsphereThumbprintOK handles this case with default header values.

Successful retrieval of vSphere thumbprint
*/
type GetVsphereThumbprintOK struct {
	Payload *models.VSphereThumbprint
}

func (o *GetVsphereThumbprintOK) Error() string {
	return fmt.Sprintf("[GET /api/providers/vsphere/thumbprint][%d] getVsphereThumbprintOK  %+v", 200, o.Payload)
}

func (o *GetVsphereThumbprintOK) GetPayload() *models.VSphereThumbprint {
	return o.Payload
}

func (o *GetVsphereThumbprintOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VSphereThumbprint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVsphereThumbprintBadRequest creates a GetVsphereThumbprintBadRequest with default headers values
func NewGetVsphereThumbprintBadRequest() *GetVsphereThumbprintBadRequest {
	return &GetVsphereThumbprintBadRequest{}
}

/*GetVsphereThumbprintBadRequest handles this case with default header values.

Bad request
*/
type GetVsphereThumbprintBadRequest struct {
	Payload *models.Error
}

func (o *GetVsphereThumbprintBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/providers/vsphere/thumbprint][%d] getVsphereThumbprintBadRequest  %+v", 400, o.Payload)
}

func (o *GetVsphereThumbprintBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVsphereThumbprintBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVsphereThumbprintUnauthorized creates a GetVsphereThumbprintUnauthorized with default headers values
func NewGetVsphereThumbprintUnauthorized() *GetVsphereThumbprintUnauthorized {
	return &GetVsphereThumbprintUnauthorized{}
}

/*GetVsphereThumbprintUnauthorized handles this case with default header values.

Incorrect credentials
*/
type GetVsphereThumbprintUnauthorized struct {
	Payload *models.Error
}

func (o *GetVsphereThumbprintUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/providers/vsphere/thumbprint][%d] getVsphereThumbprintUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVsphereThumbprintUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVsphereThumbprintUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVsphereThumbprintInternalServerError creates a GetVsphereThumbprintInternalServerError with default headers values
func NewGetVsphereThumbprintInternalServerError() *GetVsphereThumbprintInternalServerError {
	return &GetVsphereThumbprintInternalServerError{}
}

/*GetVsphereThumbprintInternalServerError handles this case with default header values.

Internal server error
*/
type GetVsphereThumbprintInternalServerError struct {
	Payload *models.Error
}

func (o *GetVsphereThumbprintInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/providers/vsphere/thumbprint][%d] getVsphereThumbprintInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVsphereThumbprintInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVsphereThumbprintInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
