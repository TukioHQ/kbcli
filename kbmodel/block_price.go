// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// BlockPrice block price
// swagger:model BlockPrice
type BlockPrice struct {

	// max
	Max float64 `json:"max,omitempty"`

	// price
	Price float64 `json:"price,omitempty"`

	// size
	Size float64 `json:"size,omitempty"`

	// unit name
	UnitName string `json:"unitName,omitempty"`
}

// Validate validates this block price
func (m *BlockPrice) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *BlockPrice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockPrice) UnmarshalBinary(b []byte) error {
	var res BlockPrice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}