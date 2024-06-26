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

// ModelsMockExecutionCreateRequestV1 models mock execution create request v1
//
// swagger:model models.MockExecutionCreateRequestV1
type ModelsMockExecutionCreateRequestV1 struct {

	// definition to be executed with provided mock results and on-demand trigger data
	Definition *ModelsDefinitionCreateRequestV2 `json:"definition,omitempty"`

	// Mock activity data and trigger data for non-on-demand executions, keyed by node ID, may include trigger and/or activity nodes
	// Required: true
	Mocks *string `json:"mocks"`

	// Trigger data for on-demand executions
	OnDemandTrigger string `json:"on_demand_trigger,omitempty"`
}

// Validate validates this models mock execution create request v1
func (m *ModelsMockExecutionCreateRequestV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMocks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsMockExecutionCreateRequestV1) validateDefinition(formats strfmt.Registry) error {
	if swag.IsZero(m.Definition) { // not required
		return nil
	}

	if m.Definition != nil {
		if err := m.Definition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("definition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("definition")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsMockExecutionCreateRequestV1) validateMocks(formats strfmt.Registry) error {

	if err := validate.Required("mocks", "body", m.Mocks); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this models mock execution create request v1 based on the context it is used
func (m *ModelsMockExecutionCreateRequestV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDefinition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsMockExecutionCreateRequestV1) contextValidateDefinition(ctx context.Context, formats strfmt.Registry) error {

	if m.Definition != nil {

		if swag.IsZero(m.Definition) { // not required
			return nil
		}

		if err := m.Definition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("definition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("definition")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsMockExecutionCreateRequestV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsMockExecutionCreateRequestV1) UnmarshalBinary(b []byte) error {
	var res ModelsMockExecutionCreateRequestV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
