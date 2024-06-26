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

// DomainCPSRatingHistoryEntry domain c p s rating history entry
//
// swagger:model domain.CPSRatingHistoryEntry
type DomainCPSRatingHistoryEntry struct {

	// date recorded
	// Required: true
	// Format: date-time
	DateRecorded *strfmt.DateTime `json:"DateRecorded"`

	// rating
	// Required: true
	Rating *string `json:"Rating"`
}

// Validate validates this domain c p s rating history entry
func (m *DomainCPSRatingHistoryEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateRecorded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRating(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainCPSRatingHistoryEntry) validateDateRecorded(formats strfmt.Registry) error {

	if err := validate.Required("DateRecorded", "body", m.DateRecorded); err != nil {
		return err
	}

	if err := validate.FormatOf("DateRecorded", "body", "date-time", m.DateRecorded.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainCPSRatingHistoryEntry) validateRating(formats strfmt.Registry) error {

	if err := validate.Required("Rating", "body", m.Rating); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain c p s rating history entry based on context it is used
func (m *DomainCPSRatingHistoryEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainCPSRatingHistoryEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainCPSRatingHistoryEntry) UnmarshalBinary(b []byte) error {
	var res DomainCPSRatingHistoryEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
