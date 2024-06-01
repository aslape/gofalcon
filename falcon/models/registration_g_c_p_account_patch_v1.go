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

// RegistrationGCPAccountPatchV1 registration g c p account patch v1
//
// swagger:model registration.GCPAccountPatchV1
type RegistrationGCPAccountPatchV1 struct {

	// environment
	Environment string `json:"environment,omitempty"`

	// parent id
	// Required: true
	ParentID *string `json:"parent_id"`

	// service account
	// Required: true
	ServiceAccount *RegistrationGCPServiceAccountPatchV1 `json:"service_account"`
}

// Validate validates this registration g c p account patch v1
func (m *RegistrationGCPAccountPatchV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistrationGCPAccountPatchV1) validateParentID(formats strfmt.Registry) error {

	if err := validate.Required("parent_id", "body", m.ParentID); err != nil {
		return err
	}

	return nil
}

func (m *RegistrationGCPAccountPatchV1) validateServiceAccount(formats strfmt.Registry) error {

	if err := validate.Required("service_account", "body", m.ServiceAccount); err != nil {
		return err
	}

	if m.ServiceAccount != nil {
		if err := m.ServiceAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service_account")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this registration g c p account patch v1 based on the context it is used
func (m *RegistrationGCPAccountPatchV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServiceAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistrationGCPAccountPatchV1) contextValidateServiceAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.ServiceAccount != nil {

		if err := m.ServiceAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service_account")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegistrationGCPAccountPatchV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistrationGCPAccountPatchV1) UnmarshalBinary(b []byte) error {
	var res RegistrationGCPAccountPatchV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}