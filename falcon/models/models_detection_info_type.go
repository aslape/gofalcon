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

// ModelsDetectionInfoType models detection info type
//
// swagger:model models.DetectionInfoType
type ModelsDetectionInfoType struct {

	// detection
	// Required: true
	Detection *ModelsDetection `json:"Detection"`
}

// Validate validates this models detection info type
func (m *ModelsDetectionInfoType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsDetectionInfoType) validateDetection(formats strfmt.Registry) error {

	if err := validate.Required("Detection", "body", m.Detection); err != nil {
		return err
	}

	if m.Detection != nil {
		if err := m.Detection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Detection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Detection")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this models detection info type based on the context it is used
func (m *ModelsDetectionInfoType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDetection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsDetectionInfoType) contextValidateDetection(ctx context.Context, formats strfmt.Registry) error {

	if m.Detection != nil {

		if err := m.Detection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Detection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Detection")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsDetectionInfoType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsDetectionInfoType) UnmarshalBinary(b []byte) error {
	var res ModelsDetectionInfoType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
