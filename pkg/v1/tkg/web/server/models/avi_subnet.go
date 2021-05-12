// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AviSubnet avi subnet
// swagger:model AviSubnet
type AviSubnet struct {

	// family
	Family string `json:"family,omitempty"`

	// subnet
	Subnet string `json:"subnet,omitempty"`
}

// Validate validates this avi subnet
func (m *AviSubnet) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AviSubnet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AviSubnet) UnmarshalBinary(b []byte) error {
	var res AviSubnet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
