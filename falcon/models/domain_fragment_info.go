// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DomainFragmentInfo domain fragment info
//
// swagger:model domain.FragmentInfo
type DomainFragmentInfo struct {

	// Offset of the content field from the start of data, in characters
	ContentOffset int64 `json:"content_offset,omitempty"`

	// Total number of fragments for this group
	Count int64 `json:"count,omitempty"`

	//  List of fields that have been split, such as: content, iocs, translated_content, ...
	FragmentedFields []string `json:"fragmented_fields"`

	// HEX string, similar to stream_id, ties all fragments together
	GroupID string `json:"group_id,omitempty"`

	// Zero-based index of fragment in the group
	Index int64 `json:"index,omitempty"`

	// Offset of the translated_content field from the start of data, in characters
	TranslatedContentOffset int64 `json:"translated_content_offset,omitempty"`

	// List of fields that have been truncated or deleted
	TruncatedFields []string `json:"truncated_fields"`
}

// Validate validates this domain fragment info
func (m *DomainFragmentInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this domain fragment info based on context it is used
func (m *DomainFragmentInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainFragmentInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainFragmentInfo) UnmarshalBinary(b []byte) error {
	var res DomainFragmentInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
