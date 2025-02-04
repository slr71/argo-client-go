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

// IoK8sAPICoreV1AzureFileVolumeSource AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
//
// swagger:model io.k8s.api.core.v1.AzureFileVolumeSource
type IoK8sAPICoreV1AzureFileVolumeSource struct {

	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly bool `json:"readOnly,omitempty"`

	// the name of secret that contains Azure Storage Account Name and Key
	// Required: true
	SecretName *string `json:"secretName"`

	// Share Name
	// Required: true
	ShareName *string `json:"shareName"`
}

// Validate validates this io k8s api core v1 azure file volume source
func (m *IoK8sAPICoreV1AzureFileVolumeSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecretName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShareName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1AzureFileVolumeSource) validateSecretName(formats strfmt.Registry) error {

	if err := validate.Required("secretName", "body", m.SecretName); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1AzureFileVolumeSource) validateShareName(formats strfmt.Registry) error {

	if err := validate.Required("shareName", "body", m.ShareName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s api core v1 azure file volume source based on context it is used
func (m *IoK8sAPICoreV1AzureFileVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1AzureFileVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1AzureFileVolumeSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1AzureFileVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
