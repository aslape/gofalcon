// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChangesPermissionsMac changes permissions mac
//
// swagger:model changes.PermissionsMac
type ChangesPermissionsMac struct {

	// acl
	ACL bool `json:"acl,omitempty"`

	// basic
	Basic []*ChangesBasic `json:"basic"`

	// common attributes
	CommonAttributes []string `json:"common_attributes"`

	// group
	Group *ChangesGroup `json:"group,omitempty"`

	// owner
	Owner *ChangesOwner `json:"owner,omitempty"`

	// security info
	SecurityInfo int32 `json:"security_info,omitempty"`
}

// Validate validates this changes permissions mac
func (m *ChangesPermissionsMac) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBasic(formats); err != nil {
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

func (m *ChangesPermissionsMac) validateBasic(formats strfmt.Registry) error {
	if swag.IsZero(m.Basic) { // not required
		return nil
	}

	for i := 0; i < len(m.Basic); i++ {
		if swag.IsZero(m.Basic[i]) { // not required
			continue
		}

		if m.Basic[i] != nil {
			if err := m.Basic[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("basic" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("basic" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ChangesPermissionsMac) validateGroup(formats strfmt.Registry) error {
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

func (m *ChangesPermissionsMac) validateOwner(formats strfmt.Registry) error {
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

// ContextValidate validate this changes permissions mac based on the context it is used
func (m *ChangesPermissionsMac) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBasic(ctx, formats); err != nil {
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

func (m *ChangesPermissionsMac) contextValidateBasic(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Basic); i++ {

		if m.Basic[i] != nil {

			if swag.IsZero(m.Basic[i]) { // not required
				return nil
			}

			if err := m.Basic[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("basic" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("basic" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ChangesPermissionsMac) contextValidateGroup(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ChangesPermissionsMac) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ChangesPermissionsMac) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangesPermissionsMac) UnmarshalBinary(b []byte) error {
	var res ChangesPermissionsMac
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
