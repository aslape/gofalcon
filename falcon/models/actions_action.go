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

// ActionsAction actions action
//
// swagger:model actions.Action
type ActionsAction struct {

	// expected change count
	// Required: true
	ExpectedChangeCount *int64 `json:"expected_change_count"`

	// id
	// Required: true
	ID *string `json:"id"`

	// Possible values: SUPPRESS, PURGE, UNSUPPRESS.
	// Required: true
	OperationType *string `json:"operation_type"`

	// previous change count
	// Required: true
	PreviousChangeCount *int64 `json:"previous_change_count"`

	// reason
	Reason string `json:"reason,omitempty"`

	// Possible values: RUNNING, DONE, FAILED
	// Required: true
	Status *string `json:"status"`

	// total change count
	// Required: true
	TotalChangeCount *int64 `json:"total_change_count"`

	// updated by
	// Required: true
	UpdatedBy *string `json:"updated_by"`

	// updated date
	// Required: true
	UpdatedDate *string `json:"updated_date"`
}

// Validate validates this actions action
func (m *ActionsAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpectedChangeCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreviousChangeCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalChangeCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActionsAction) validateExpectedChangeCount(formats strfmt.Registry) error {

	if err := validate.Required("expected_change_count", "body", m.ExpectedChangeCount); err != nil {
		return err
	}

	return nil
}

func (m *ActionsAction) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ActionsAction) validateOperationType(formats strfmt.Registry) error {

	if err := validate.Required("operation_type", "body", m.OperationType); err != nil {
		return err
	}

	return nil
}

func (m *ActionsAction) validatePreviousChangeCount(formats strfmt.Registry) error {

	if err := validate.Required("previous_change_count", "body", m.PreviousChangeCount); err != nil {
		return err
	}

	return nil
}

func (m *ActionsAction) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ActionsAction) validateTotalChangeCount(formats strfmt.Registry) error {

	if err := validate.Required("total_change_count", "body", m.TotalChangeCount); err != nil {
		return err
	}

	return nil
}

func (m *ActionsAction) validateUpdatedBy(formats strfmt.Registry) error {

	if err := validate.Required("updated_by", "body", m.UpdatedBy); err != nil {
		return err
	}

	return nil
}

func (m *ActionsAction) validateUpdatedDate(formats strfmt.Registry) error {

	if err := validate.Required("updated_date", "body", m.UpdatedDate); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this actions action based on context it is used
func (m *ActionsAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ActionsAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionsAction) UnmarshalBinary(b []byte) error {
	var res ActionsAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
