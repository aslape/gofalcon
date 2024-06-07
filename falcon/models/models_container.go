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
	"github.com/go-openapi/validate"
)

// ModelsContainer models container
//
// swagger:model models.Container
type ModelsContainer struct {

	// agents
	Agents []ModelsContainerAgents `json:"agents"`

	// allow privilege escalation
	AllowPrivilegeEscalation bool `json:"allow_privilege_escalation,omitempty"`

	// cid
	// Required: true
	Cid *string `json:"cid"`

	// cloud
	Cloud string `json:"cloud,omitempty"`

	// cloud account id
	CloudAccountID string `json:"cloud_account_id,omitempty"`

	// cloud region
	CloudRegion string `json:"cloud_region,omitempty"`

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// cluster name
	ClusterName string `json:"cluster_name,omitempty"`

	// config labels
	ConfigLabels string `json:"config_labels,omitempty"`

	// config user
	ConfigUser string `json:"config_user,omitempty"`

	// container id
	ContainerID string `json:"container_id,omitempty"`

	// container name
	ContainerName string `json:"container_name,omitempty"`

	// container image id
	ContainerImageID string `json:"container_image_id,omitempty"`

	// created at
	CreatedAt int64 `json:"created_at,omitempty"`

	// cve ids
	CveIds []string `json:"cve_ids"`

	// first seen
	FirstSeen string `json:"first_seen,omitempty"`

	// host config devices
	HostConfigDevices string `json:"host_config_devices,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// image application package count
	ImageApplicationPackageCount int32 `json:"image_application_package_count,omitempty"`

	// image assessed at
	ImageAssessedAt int64 `json:"image_assessed_at,omitempty"`

	// image detection count
	ImageDetectionCount int32 `json:"image_detection_count,omitempty"`

	// image detection id list
	ImageDetectionIDList []string `json:"image_detection_id_list"`

	// image detection name list
	ImageDetectionNameList []string `json:"image_detection_name_list"`

	// image detection severity by type
	ImageDetectionSeverityByType map[string]int64 `json:"image_detection_severity_by_type,omitempty"`

	// image digest
	ImageDigest string `json:"image_digest,omitempty"`

	// image has been assessed
	ImageHasBeenAssessed bool `json:"image_has_been_assessed,omitempty"`

	// image highest severity vulnerability
	ImageHighestSeverityVulnerability string `json:"image_highest_severity_vulnerability,omitempty"`

	// image id
	ImageID string `json:"image_id,omitempty"`

	// image package count
	ImagePackageCount int32 `json:"image_package_count,omitempty"`

	// image registry
	ImageRegistry string `json:"image_registry,omitempty"`

	// image repository
	ImageRepository string `json:"image_repository,omitempty"`

	// image tag
	ImageTag string `json:"image_tag,omitempty"`

	// image vulnerability count
	ImageVulnerabilityCount int32 `json:"image_vulnerability_count,omitempty"`

	// image vulnerability severity by type
	ImageVulnerabilitySeverityByType map[string]int64 `json:"image_vulnerability_severity_by_type,omitempty"`

	// insecure mount source
	InsecureMountSource string `json:"insecure_mount_source,omitempty"`

	// insecure mount type
	InsecureMountType string `json:"insecure_mount_type,omitempty"`

	// insecure propagation mode
	InsecurePropagationMode bool `json:"insecure_propagation_mode,omitempty"`

	// interactive mode
	InteractiveMode bool `json:"interactive_mode,omitempty"`

	// ipv4
	IPV4 string `json:"ipv4,omitempty"`

	// ipv6
	IPV6 string `json:"ipv6,omitempty"`

	// kpa coverage
	KpaCoverage bool `json:"kpa_coverage,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// labels list
	LabelsList []string `json:"labels_list"`

	// last seen
	LastSeen string `json:"last_seen,omitempty"`

	// linux sensor aid
	LinuxSensorAid string `json:"linux_sensor_aid,omitempty"`

	// linux sensor config build
	LinuxSensorConfigBuild string `json:"linux_sensor_config_build,omitempty"`

	// linux sensor coverage
	LinuxSensorCoverage bool `json:"linux_sensor_coverage,omitempty"`

	// lumos sensor aid
	LumosSensorAid string `json:"lumos_sensor_aid,omitempty"`

	// lumos sensor config build
	LumosSensorConfigBuild string `json:"lumos_sensor_config_build,omitempty"`

	// lumos sensor coverage
	LumosSensorCoverage bool `json:"lumos_sensor_coverage,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// node id
	NodeID string `json:"node_uid,omitempty"`

	// node name
	NodeName string `json:"node_name,omitempty"`

	// pod id
	PodID string `json:"pod_id,omitempty"`

	// pod name
	PodName string `json:"pod_name,omitempty"`

	// port list
	PortList []ModelsContainerPortList `json:"port_list"`

	// privileged
	Privileged bool `json:"privileged,omitempty"`

	// root write access
	RootWriteAccess bool `json:"root_write_access,omitempty"`

	// rpd
	Rpd []int64 `json:"rpd"`

	// run as root group
	RunAsRootGroup bool `json:"run_as_root_group,omitempty"`

	// run as root user
	RunAsRootUser bool `json:"run_as_root_user,omitempty"`

	// running status
	RunningStatus bool `json:"running_status,omitempty"`

	// snapshot coverage
	SnapshotCoverage bool `json:"snapshot_coverage,omitempty"`

	// unidentified
	Unidentified bool `json:"unidentified,omitempty"`

	// volume mounts
	VolumeMounts string `json:"volume_mounts,omitempty"`
}

// Validate validates this models container
func (m *ModelsContainer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsContainer) validateAgents(formats strfmt.Registry) error {
	if swag.IsZero(m.Agents) { // not required
		return nil
	}

	for i := 0; i < len(m.Agents); i++ {

		if m.Agents[i] != nil {
			if err := m.Agents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("agents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsContainer) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *ModelsContainer) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ModelsContainer) validatePortList(formats strfmt.Registry) error {
	if swag.IsZero(m.PortList) { // not required
		return nil
	}

	for i := 0; i < len(m.PortList); i++ {

		if m.PortList[i] != nil {
			if err := m.PortList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("port_list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("port_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this models container based on the context it is used
func (m *ModelsContainer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAgents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePortList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsContainer) contextValidateAgents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Agents); i++ {

		if swag.IsZero(m.Agents[i]) { // not required
			return nil
		}

		if err := m.Agents[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agents" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agents" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ModelsContainer) contextValidatePortList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PortList); i++ {

		if swag.IsZero(m.PortList[i]) { // not required
			return nil
		}

		if err := m.PortList[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("port_list" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("port_list" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsContainer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsContainer) UnmarshalBinary(b []byte) error {
	var res ModelsContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
