// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu-private/core/pkg/v1/tkg/web/server/models"
)

// SetAWSEndpointCreatedCode is the HTTP code returned for type SetAWSEndpointCreated
const SetAWSEndpointCreatedCode int = 201

/*SetAWSEndpointCreated Connection successful

swagger:response setAWSEndpointCreated
*/
type SetAWSEndpointCreated struct {
}

// NewSetAWSEndpointCreated creates SetAWSEndpointCreated with default headers values
func NewSetAWSEndpointCreated() *SetAWSEndpointCreated {

	return &SetAWSEndpointCreated{}
}

// WriteResponse to the client
func (o *SetAWSEndpointCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// SetAWSEndpointBadRequestCode is the HTTP code returned for type SetAWSEndpointBadRequest
const SetAWSEndpointBadRequestCode int = 400

/*SetAWSEndpointBadRequest Bad request

swagger:response setAWSEndpointBadRequest
*/
type SetAWSEndpointBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSetAWSEndpointBadRequest creates SetAWSEndpointBadRequest with default headers values
func NewSetAWSEndpointBadRequest() *SetAWSEndpointBadRequest {

	return &SetAWSEndpointBadRequest{}
}

// WithPayload adds the payload to the set a w s endpoint bad request response
func (o *SetAWSEndpointBadRequest) WithPayload(payload *models.Error) *SetAWSEndpointBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set a w s endpoint bad request response
func (o *SetAWSEndpointBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetAWSEndpointBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SetAWSEndpointUnauthorizedCode is the HTTP code returned for type SetAWSEndpointUnauthorized
const SetAWSEndpointUnauthorizedCode int = 401

/*SetAWSEndpointUnauthorized Incorrect credentials

swagger:response setAWSEndpointUnauthorized
*/
type SetAWSEndpointUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSetAWSEndpointUnauthorized creates SetAWSEndpointUnauthorized with default headers values
func NewSetAWSEndpointUnauthorized() *SetAWSEndpointUnauthorized {

	return &SetAWSEndpointUnauthorized{}
}

// WithPayload adds the payload to the set a w s endpoint unauthorized response
func (o *SetAWSEndpointUnauthorized) WithPayload(payload *models.Error) *SetAWSEndpointUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set a w s endpoint unauthorized response
func (o *SetAWSEndpointUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetAWSEndpointUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SetAWSEndpointInternalServerErrorCode is the HTTP code returned for type SetAWSEndpointInternalServerError
const SetAWSEndpointInternalServerErrorCode int = 500

/*SetAWSEndpointInternalServerError Internal server error

swagger:response setAWSEndpointInternalServerError
*/
type SetAWSEndpointInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSetAWSEndpointInternalServerError creates SetAWSEndpointInternalServerError with default headers values
func NewSetAWSEndpointInternalServerError() *SetAWSEndpointInternalServerError {

	return &SetAWSEndpointInternalServerError{}
}

// WithPayload adds the payload to the set a w s endpoint internal server error response
func (o *SetAWSEndpointInternalServerError) WithPayload(payload *models.Error) *SetAWSEndpointInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set a w s endpoint internal server error response
func (o *SetAWSEndpointInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetAWSEndpointInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
