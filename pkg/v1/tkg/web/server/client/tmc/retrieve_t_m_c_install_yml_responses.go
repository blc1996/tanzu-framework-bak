// Code generated by go-swagger; DO NOT EDIT.

package tmc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware-tanzu-private/core/pkg/v1/tkg/web/server/models"
)

// RetrieveTMCInstallYmlReader is a Reader for the RetrieveTMCInstallYml structure.
type RetrieveTMCInstallYmlReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveTMCInstallYmlReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveTMCInstallYmlOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRetrieveTMCInstallYmlBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRetrieveTMCInstallYmlInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewRetrieveTMCInstallYmlBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRetrieveTMCInstallYmlOK creates a RetrieveTMCInstallYmlOK with default headers values
func NewRetrieveTMCInstallYmlOK() *RetrieveTMCInstallYmlOK {
	return &RetrieveTMCInstallYmlOK{}
}

/*RetrieveTMCInstallYmlOK handles this case with default header values.

Successfully retrieved TMC install yml.
*/
type RetrieveTMCInstallYmlOK struct {
	Payload string
}

func (o *RetrieveTMCInstallYmlOK) Error() string {
	return fmt.Sprintf("[GET /api/integration/tmc][%d] retrieveTMCInstallYmlOK  %+v", 200, o.Payload)
}

func (o *RetrieveTMCInstallYmlOK) GetPayload() string {
	return o.Payload
}

func (o *RetrieveTMCInstallYmlOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveTMCInstallYmlBadRequest creates a RetrieveTMCInstallYmlBadRequest with default headers values
func NewRetrieveTMCInstallYmlBadRequest() *RetrieveTMCInstallYmlBadRequest {
	return &RetrieveTMCInstallYmlBadRequest{}
}

/*RetrieveTMCInstallYmlBadRequest handles this case with default header values.

Bad request
*/
type RetrieveTMCInstallYmlBadRequest struct {
	Payload *models.Error
}

func (o *RetrieveTMCInstallYmlBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/integration/tmc][%d] retrieveTMCInstallYmlBadRequest  %+v", 400, o.Payload)
}

func (o *RetrieveTMCInstallYmlBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RetrieveTMCInstallYmlBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveTMCInstallYmlInternalServerError creates a RetrieveTMCInstallYmlInternalServerError with default headers values
func NewRetrieveTMCInstallYmlInternalServerError() *RetrieveTMCInstallYmlInternalServerError {
	return &RetrieveTMCInstallYmlInternalServerError{}
}

/*RetrieveTMCInstallYmlInternalServerError handles this case with default header values.

Internal server error
*/
type RetrieveTMCInstallYmlInternalServerError struct {
	Payload *models.Error
}

func (o *RetrieveTMCInstallYmlInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/integration/tmc][%d] retrieveTMCInstallYmlInternalServerError  %+v", 500, o.Payload)
}

func (o *RetrieveTMCInstallYmlInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *RetrieveTMCInstallYmlInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveTMCInstallYmlBadGateway creates a RetrieveTMCInstallYmlBadGateway with default headers values
func NewRetrieveTMCInstallYmlBadGateway() *RetrieveTMCInstallYmlBadGateway {
	return &RetrieveTMCInstallYmlBadGateway{}
}

/*RetrieveTMCInstallYmlBadGateway handles this case with default header values.

Bad Gateway
*/
type RetrieveTMCInstallYmlBadGateway struct {
	Payload *models.Error
}

func (o *RetrieveTMCInstallYmlBadGateway) Error() string {
	return fmt.Sprintf("[GET /api/integration/tmc][%d] retrieveTMCInstallYmlBadGateway  %+v", 502, o.Payload)
}

func (o *RetrieveTMCInstallYmlBadGateway) GetPayload() *models.Error {
	return o.Payload
}

func (o *RetrieveTMCInstallYmlBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
