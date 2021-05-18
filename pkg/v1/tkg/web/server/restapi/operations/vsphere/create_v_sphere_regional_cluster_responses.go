// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu-private/core/pkg/v1/tkg/web/server/models"
)

// CreateVSphereRegionalClusterOKCode is the HTTP code returned for type CreateVSphereRegionalClusterOK
const CreateVSphereRegionalClusterOKCode int = 200

/*CreateVSphereRegionalClusterOK Creating regional cluster started successfully

swagger:response createVSphereRegionalClusterOK
*/
type CreateVSphereRegionalClusterOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCreateVSphereRegionalClusterOK creates CreateVSphereRegionalClusterOK with default headers values
func NewCreateVSphereRegionalClusterOK() *CreateVSphereRegionalClusterOK {

	return &CreateVSphereRegionalClusterOK{}
}

// WithPayload adds the payload to the create v sphere regional cluster o k response
func (o *CreateVSphereRegionalClusterOK) WithPayload(payload string) *CreateVSphereRegionalClusterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create v sphere regional cluster o k response
func (o *CreateVSphereRegionalClusterOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateVSphereRegionalClusterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// CreateVSphereRegionalClusterBadRequestCode is the HTTP code returned for type CreateVSphereRegionalClusterBadRequest
const CreateVSphereRegionalClusterBadRequestCode int = 400

/*CreateVSphereRegionalClusterBadRequest Bad request

swagger:response createVSphereRegionalClusterBadRequest
*/
type CreateVSphereRegionalClusterBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateVSphereRegionalClusterBadRequest creates CreateVSphereRegionalClusterBadRequest with default headers values
func NewCreateVSphereRegionalClusterBadRequest() *CreateVSphereRegionalClusterBadRequest {

	return &CreateVSphereRegionalClusterBadRequest{}
}

// WithPayload adds the payload to the create v sphere regional cluster bad request response
func (o *CreateVSphereRegionalClusterBadRequest) WithPayload(payload *models.Error) *CreateVSphereRegionalClusterBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create v sphere regional cluster bad request response
func (o *CreateVSphereRegionalClusterBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateVSphereRegionalClusterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateVSphereRegionalClusterUnauthorizedCode is the HTTP code returned for type CreateVSphereRegionalClusterUnauthorized
const CreateVSphereRegionalClusterUnauthorizedCode int = 401

/*CreateVSphereRegionalClusterUnauthorized Incorrect credentials

swagger:response createVSphereRegionalClusterUnauthorized
*/
type CreateVSphereRegionalClusterUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateVSphereRegionalClusterUnauthorized creates CreateVSphereRegionalClusterUnauthorized with default headers values
func NewCreateVSphereRegionalClusterUnauthorized() *CreateVSphereRegionalClusterUnauthorized {

	return &CreateVSphereRegionalClusterUnauthorized{}
}

// WithPayload adds the payload to the create v sphere regional cluster unauthorized response
func (o *CreateVSphereRegionalClusterUnauthorized) WithPayload(payload *models.Error) *CreateVSphereRegionalClusterUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create v sphere regional cluster unauthorized response
func (o *CreateVSphereRegionalClusterUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateVSphereRegionalClusterUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateVSphereRegionalClusterInternalServerErrorCode is the HTTP code returned for type CreateVSphereRegionalClusterInternalServerError
const CreateVSphereRegionalClusterInternalServerErrorCode int = 500

/*CreateVSphereRegionalClusterInternalServerError Internal server error

swagger:response createVSphereRegionalClusterInternalServerError
*/
type CreateVSphereRegionalClusterInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateVSphereRegionalClusterInternalServerError creates CreateVSphereRegionalClusterInternalServerError with default headers values
func NewCreateVSphereRegionalClusterInternalServerError() *CreateVSphereRegionalClusterInternalServerError {

	return &CreateVSphereRegionalClusterInternalServerError{}
}

// WithPayload adds the payload to the create v sphere regional cluster internal server error response
func (o *CreateVSphereRegionalClusterInternalServerError) WithPayload(payload *models.Error) *CreateVSphereRegionalClusterInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create v sphere regional cluster internal server error response
func (o *CreateVSphereRegionalClusterInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateVSphereRegionalClusterInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
