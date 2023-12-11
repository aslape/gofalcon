// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChangesPermissions changes permissions
//
// swagger:model changes.Permissions
type ChangesPermissions struct {

	// dacl
	Dacl *ChangesDACL `json:"dacl,omitempty"`

	// group
	Group *ChangesGroup `json:"group,omitempty"`

	// owner
	Owner *ChangesOwner `json:"owner,omitempty"`

	// Possible values: 0 - OWNER, 1 - GROUP, 2 - DACL, 3 - SACL
	SecurityInfo int32 `json:"security_info,omitempty"`
}

// Validate validates this changes permissions
func (m *ChangesPermissions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDacl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChangesPermissions) validateDacl(formats strfmt.Registry) error {
	if swag.IsZero(m.Dacl) { // not required
		return nil
	}

	if m.Dacl != nil {
		if err := m.Dacl.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dacl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dacl")
			}
			return err
		}
	}

	return nil
}

func (m *ChangesPermissions) validateGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *ChangesPermissions) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this changes permissions based on the context it is used
func (m *ChangesPermissions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDacl(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChangesPermissions) contextValidateDacl(ctx context.Context, formats strfmt.Registry) error {

	if m.Dacl != nil {

		if swag.IsZero(m.Dacl) { // not required
			return nil
		}

		if err := m.Dacl.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dacl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dacl")
			}
			return err
		}
	}

	return nil
}

func (m *ChangesPermissions) contextValidateGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.Group != nil {

		if swag.IsZero(m.Group) { // not required
			return nil
		}

		if err := m.Group.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *ChangesPermissions) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.Owner != nil {

		if swag.IsZero(m.Owner) { // not required
			return nil
		}

		if err := m.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChangesPermissions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangesPermissions) UnmarshalBinary(b []byte) error {
	var res ChangesPermissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}