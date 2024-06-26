// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APITokenCreateRequestV1 api token create request v1
//
// swagger:model api.TokenCreateRequestV1
type APITokenCreateRequestV1 struct {

	// The token's expiration time (RFC-3339). Null, if the token never expires.
	// Format: date-time
	ExpiresTimestamp strfmt.DateTime `json:"expires_timestamp,omitempty"`

	// The token label.
	Label string `json:"label,omitempty"`

	// The token type.
	Type string `json:"type,omitempty"`
}

// Validate validates this api token create request v1
func (m *APITokenCreateRequestV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiresTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APITokenCreateRequestV1) validateExpiresTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiresTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("expires_timestamp", "body", "date-time", m.ExpiresTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this api token create request v1 based on context it is used
func (m *APITokenCreateRequestV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APITokenCreateRequestV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APITokenCreateRequestV1) UnmarshalBinary(b []byte) error {
	var res APITokenCreateRequestV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
