// Code generated by go-swagger; DO NOT EDIT.

package docker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu-private/core/pkg/v1/tkg/web/server/models"
)

// GenerateTKGConfigForDockerOKCode is the HTTP code returned for type GenerateTKGConfigForDockerOK
const GenerateTKGConfigForDockerOKCode int = 200

/*GenerateTKGConfigForDockerOK Generated TKG configuration successfully

swagger:response generateTKGConfigForDockerOK
*/
type GenerateTKGConfigForDockerOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGenerateTKGConfigForDockerOK creates GenerateTKGConfigForDockerOK with default headers values
func NewGenerateTKGConfigForDockerOK() *GenerateTKGConfigForDockerOK {

	return &GenerateTKGConfigForDockerOK{}
}

// WithPayload adds the payload to the generate t k g config for docker o k response
func (o *GenerateTKGConfigForDockerOK) WithPayload(payload string) *GenerateTKGConfigForDockerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the generate t k g config for docker o k response
func (o *GenerateTKGConfigForDockerOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GenerateTKGConfigForDockerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GenerateTKGConfigForDockerBadRequestCode is the HTTP code returned for type GenerateTKGConfigForDockerBadRequest
const GenerateTKGConfigForDockerBadRequestCode int = 400

/*GenerateTKGConfigForDockerBadRequest Bad request

swagger:response generateTKGConfigForDockerBadRequest
*/
type GenerateTKGConfigForDockerBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGenerateTKGConfigForDockerBadRequest creates GenerateTKGConfigForDockerBadRequest with default headers values
func NewGenerateTKGConfigForDockerBadRequest() *GenerateTKGConfigForDockerBadRequest {

	return &GenerateTKGConfigForDockerBadRequest{}
}

// WithPayload adds the payload to the generate t k g config for docker bad request response
func (o *GenerateTKGConfigForDockerBadRequest) WithPayload(payload *models.Error) *GenerateTKGConfigForDockerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the generate t k g config for docker bad request response
func (o *GenerateTKGConfigForDockerBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GenerateTKGConfigForDockerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GenerateTKGConfigForDockerInternalServerErrorCode is the HTTP code returned for type GenerateTKGConfigForDockerInternalServerError
const GenerateTKGConfigForDockerInternalServerErrorCode int = 500

/*GenerateTKGConfigForDockerInternalServerError Internal server error

swagger:response generateTKGConfigForDockerInternalServerError
*/
type GenerateTKGConfigForDockerInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGenerateTKGConfigForDockerInternalServerError creates GenerateTKGConfigForDockerInternalServerError with default headers values
func NewGenerateTKGConfigForDockerInternalServerError() *GenerateTKGConfigForDockerInternalServerError {

	return &GenerateTKGConfigForDockerInternalServerError{}
}

// WithPayload adds the payload to the generate t k g config for docker internal server error response
func (o *GenerateTKGConfigForDockerInternalServerError) WithPayload(payload *models.Error) *GenerateTKGConfigForDockerInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the generate t k g config for docker internal server error response
func (o *GenerateTKGConfigForDockerInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GenerateTKGConfigForDockerInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
