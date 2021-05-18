// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu-private/core/pkg/v1/tkg/web/server/models"
)

// GenerateTKGConfigForAzureOKCode is the HTTP code returned for type GenerateTKGConfigForAzureOK
const GenerateTKGConfigForAzureOKCode int = 200

/*GenerateTKGConfigForAzureOK Generated TKG configuration successfully

swagger:response generateTKGConfigForAzureOK
*/
type GenerateTKGConfigForAzureOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGenerateTKGConfigForAzureOK creates GenerateTKGConfigForAzureOK with default headers values
func NewGenerateTKGConfigForAzureOK() *GenerateTKGConfigForAzureOK {

	return &GenerateTKGConfigForAzureOK{}
}

// WithPayload adds the payload to the generate t k g config for azure o k response
func (o *GenerateTKGConfigForAzureOK) WithPayload(payload string) *GenerateTKGConfigForAzureOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the generate t k g config for azure o k response
func (o *GenerateTKGConfigForAzureOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GenerateTKGConfigForAzureOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GenerateTKGConfigForAzureBadRequestCode is the HTTP code returned for type GenerateTKGConfigForAzureBadRequest
const GenerateTKGConfigForAzureBadRequestCode int = 400

/*GenerateTKGConfigForAzureBadRequest Bad request

swagger:response generateTKGConfigForAzureBadRequest
*/
type GenerateTKGConfigForAzureBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGenerateTKGConfigForAzureBadRequest creates GenerateTKGConfigForAzureBadRequest with default headers values
func NewGenerateTKGConfigForAzureBadRequest() *GenerateTKGConfigForAzureBadRequest {

	return &GenerateTKGConfigForAzureBadRequest{}
}

// WithPayload adds the payload to the generate t k g config for azure bad request response
func (o *GenerateTKGConfigForAzureBadRequest) WithPayload(payload *models.Error) *GenerateTKGConfigForAzureBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the generate t k g config for azure bad request response
func (o *GenerateTKGConfigForAzureBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GenerateTKGConfigForAzureBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GenerateTKGConfigForAzureUnauthorizedCode is the HTTP code returned for type GenerateTKGConfigForAzureUnauthorized
const GenerateTKGConfigForAzureUnauthorizedCode int = 401

/*GenerateTKGConfigForAzureUnauthorized Incorrect credentials

swagger:response generateTKGConfigForAzureUnauthorized
*/
type GenerateTKGConfigForAzureUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGenerateTKGConfigForAzureUnauthorized creates GenerateTKGConfigForAzureUnauthorized with default headers values
func NewGenerateTKGConfigForAzureUnauthorized() *GenerateTKGConfigForAzureUnauthorized {

	return &GenerateTKGConfigForAzureUnauthorized{}
}

// WithPayload adds the payload to the generate t k g config for azure unauthorized response
func (o *GenerateTKGConfigForAzureUnauthorized) WithPayload(payload *models.Error) *GenerateTKGConfigForAzureUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the generate t k g config for azure unauthorized response
func (o *GenerateTKGConfigForAzureUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GenerateTKGConfigForAzureUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GenerateTKGConfigForAzureInternalServerErrorCode is the HTTP code returned for type GenerateTKGConfigForAzureInternalServerError
const GenerateTKGConfigForAzureInternalServerErrorCode int = 500

/*GenerateTKGConfigForAzureInternalServerError Internal server error

swagger:response generateTKGConfigForAzureInternalServerError
*/
type GenerateTKGConfigForAzureInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGenerateTKGConfigForAzureInternalServerError creates GenerateTKGConfigForAzureInternalServerError with default headers values
func NewGenerateTKGConfigForAzureInternalServerError() *GenerateTKGConfigForAzureInternalServerError {

	return &GenerateTKGConfigForAzureInternalServerError{}
}

// WithPayload adds the payload to the generate t k g config for azure internal server error response
func (o *GenerateTKGConfigForAzureInternalServerError) WithPayload(payload *models.Error) *GenerateTKGConfigForAzureInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the generate t k g config for azure internal server error response
func (o *GenerateTKGConfigForAzureInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GenerateTKGConfigForAzureInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
