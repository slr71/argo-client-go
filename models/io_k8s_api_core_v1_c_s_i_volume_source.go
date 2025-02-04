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

// IoK8sAPICoreV1CSIVolumeSource Represents a source location of a volume to mount, managed by an external CSI driver
//
// swagger:model io.k8s.api.core.v1.CSIVolumeSource
type IoK8sAPICoreV1CSIVolumeSource struct {

	// Driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.
	// Required: true
	Driver *string `json:"driver"`

	// Filesystem type to mount. Ex. "ext4", "xfs", "ntfs". If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.
	FsType string `json:"fsType,omitempty"`

	// NodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.
	NodePublishSecretRef *IoK8sAPICoreV1LocalObjectReference `json:"nodePublishSecretRef,omitempty"`

	// Specifies a read-only configuration for the volume. Defaults to false (read/write).
	ReadOnly bool `json:"readOnly,omitempty"`

	// VolumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.
	VolumeAttributes map[string]string `json:"volumeAttributes,omitempty"`
}

// Validate validates this io k8s api core v1 c s i volume source
func (m *IoK8sAPICoreV1CSIVolumeSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDriver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodePublishSecretRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1CSIVolumeSource) validateDriver(formats strfmt.Registry) error {

	if err := validate.Required("driver", "body", m.Driver); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1CSIVolumeSource) validateNodePublishSecretRef(formats strfmt.Registry) error {
	if swag.IsZero(m.NodePublishSecretRef) { // not required
		return nil
	}

	if m.NodePublishSecretRef != nil {
		if err := m.NodePublishSecretRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodePublishSecretRef")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io k8s api core v1 c s i volume source based on the context it is used
func (m *IoK8sAPICoreV1CSIVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNodePublishSecretRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1CSIVolumeSource) contextValidateNodePublishSecretRef(ctx context.Context, formats strfmt.Registry) error {

	if m.NodePublishSecretRef != nil {
		if err := m.NodePublishSecretRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodePublishSecretRef")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1CSIVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1CSIVolumeSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1CSIVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
