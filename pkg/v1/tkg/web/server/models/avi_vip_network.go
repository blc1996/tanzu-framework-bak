// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AviVipNetwork avi vip network
// swagger:model AviVipNetwork
type AviVipNetwork struct {

	// cloud
	Cloud string `json:"cloud,omitempty"`

	// configed subnets
	ConfigedSubnets []*AviSubnet `json:"configedSubnets"`

	// name
	Name string `json:"name,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this avi vip network
func (m *AviVipNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigedSubnets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AviVipNetwork) validateConfigedSubnets(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigedSubnets) { // not required
		return nil
	}

	for i := 0; i < len(m.ConfigedSubnets); i++ {
		if swag.IsZero(m.ConfigedSubnets[i]) { // not required
			continue
		}

		if m.ConfigedSubnets[i] != nil {
			if err := m.ConfigedSubnets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configedSubnets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AviVipNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AviVipNetwork) UnmarshalBinary(b []byte) error {
	var res AviVipNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
