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

// IoK8sAPICoreV1FlexVolumeSource FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
//
// swagger:model io.k8s.api.core.v1.FlexVolumeSource
type IoK8sAPICoreV1FlexVolumeSource struct {

	// Driver is the name of the driver to use for this volume.
	// Required: true
	Driver *string `json:"driver"`

	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
	FsType string `json:"fsType,omitempty"`

	// Optional: Extra command options if any.
	Options map[string]string `json:"options,omitempty"`

	// Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly bool `json:"readOnly,omitempty"`

	// Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.
	SecretRef *IoK8sAPICoreV1LocalObjectReference `json:"secretRef,omitempty"`
}

// Validate validates this io k8s api core v1 flex volume source
func (m *IoK8sAPICoreV1FlexVolumeSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDriver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1FlexVolumeSource) validateDriver(formats strfmt.Registry) error {

	if err := validate.Required("driver", "body", m.Driver); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1FlexVolumeSource) validateSecretRef(formats strfmt.Registry) error {
	if swag.IsZero(m.SecretRef) { // not required
		return nil
	}

	if m.SecretRef != nil {
		if err := m.SecretRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secretRef")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io k8s api core v1 flex volume source based on the context it is used
func (m *IoK8sAPICoreV1FlexVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSecretRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1FlexVolumeSource) contextValidateSecretRef(ctx context.Context, formats strfmt.Registry) error {

	if m.SecretRef != nil {
		if err := m.SecretRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secretRef")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1FlexVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1FlexVolumeSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1FlexVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
