// Code generated by go-swagger; DO NOT EDIT.

package provider

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu-private/core/pkg/v1/tkg/web/server/models"
)

// GetProviderOKCode is the HTTP code returned for type GetProviderOK
const GetProviderOKCode int = 200

/*GetProviderOK Successful operation

swagger:response getProviderOK
*/
type GetProviderOK struct {

	/*
	  In: Body
	*/
	Payload *models.ProviderInfo `json:"body,omitempty"`
}

// NewGetProviderOK creates GetProviderOK with default headers values
func NewGetProviderOK() *GetProviderOK {

	return &GetProviderOK{}
}

// WithPayload adds the payload to the get provider o k response
func (o *GetProviderOK) WithPayload(payload *models.ProviderInfo) *GetProviderOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get provider o k response
func (o *GetProviderOK) SetPayload(payload *models.ProviderInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProviderOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProviderBadRequestCode is the HTTP code returned for type GetProviderBadRequest
const GetProviderBadRequestCode int = 400

/*GetProviderBadRequest Bad request

swagger:response getProviderBadRequest
*/
type GetProviderBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetProviderBadRequest creates GetProviderBadRequest with default headers values
func NewGetProviderBadRequest() *GetProviderBadRequest {

	return &GetProviderBadRequest{}
}

// WithPayload adds the payload to the get provider bad request response
func (o *GetProviderBadRequest) WithPayload(payload *models.Error) *GetProviderBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get provider bad request response
func (o *GetProviderBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProviderBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProviderInternalServerErrorCode is the HTTP code returned for type GetProviderInternalServerError
const GetProviderInternalServerErrorCode int = 500

/*GetProviderInternalServerError Internal server error

swagger:response getProviderInternalServerError
*/
type GetProviderInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetProviderInternalServerError creates GetProviderInternalServerError with default headers values
func NewGetProviderInternalServerError() *GetProviderInternalServerError {

	return &GetProviderInternalServerError{}
}

// WithPayload adds the payload to the get provider internal server error response
func (o *GetProviderInternalServerError) WithPayload(payload *models.Error) *GetProviderInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get provider internal server error response
func (o *GetProviderInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProviderInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
