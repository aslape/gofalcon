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

// ModelsAPIVulnByPublicationDate models API vuln by publication date
//
// swagger:model models.APIVulnByPublicationDate
type ModelsAPIVulnByPublicationDate struct {

	// containers impacted
	// Required: true
	ContainersImpacted *int64 `json:"containers_impacted"`

	// cve id
	// Required: true
	CveID *string `json:"cve_id"`

	// cvss score
	// Required: true
	CvssScore *float32 `json:"cvss_score"`

	// images impacted
	// Required: true
	ImagesImpacted *int64 `json:"images_impacted"`

	// published date
	// Required: true
	PublishedDate *string `json:"published_date"`

	// severity
	// Required: true
	Severity *string `json:"severity"`
}

// Validate validates this models API vuln by publication date
func (m *ModelsAPIVulnByPublicationDate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainersImpacted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCveID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCvssScore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImagesImpacted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsAPIVulnByPublicationDate) validateContainersImpacted(formats strfmt.Registry) error {

	if err := validate.Required("containers_impacted", "body", m.ContainersImpacted); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAPIVulnByPublicationDate) validateCveID(formats strfmt.Registry) error {

	if err := validate.Required("cve_id", "body", m.CveID); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAPIVulnByPublicationDate) validateCvssScore(formats strfmt.Registry) error {

	if err := validate.Required("cvss_score", "body", m.CvssScore); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAPIVulnByPublicationDate) validateImagesImpacted(formats strfmt.Registry) error {

	if err := validate.Required("images_impacted", "body", m.ImagesImpacted); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAPIVulnByPublicationDate) validatePublishedDate(formats strfmt.Registry) error {

	if err := validate.Required("published_date", "body", m.PublishedDate); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAPIVulnByPublicationDate) validateSeverity(formats strfmt.Registry) error {

	if err := validate.Required("severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this models API vuln by publication date based on context it is used
func (m *ModelsAPIVulnByPublicationDate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsAPIVulnByPublicationDate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsAPIVulnByPublicationDate) UnmarshalBinary(b []byte) error {
	var res ModelsAPIVulnByPublicationDate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
