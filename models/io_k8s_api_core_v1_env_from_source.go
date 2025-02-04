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

// IoK8sAPICoreV1EnvFromSource EnvFromSource represents the source of a set of ConfigMaps
//
// swagger:model io.k8s.api.core.v1.EnvFromSource
type IoK8sAPICoreV1EnvFromSource struct {

	// The ConfigMap to select from
	ConfigMapRef *IoK8sAPICoreV1ConfigMapEnvSource `json:"configMapRef,omitempty"`

	// An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
	Prefix string `json:"prefix,omitempty"`

	// The Secret to select from
	SecretRef *IoK8sAPICoreV1SecretEnvSource `json:"secretRef,omitempty"`
}

// Validate validates this io k8s api core v1 env from source
func (m *IoK8sAPICoreV1EnvFromSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigMapRef(formats); err != nil {
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

func (m *IoK8sAPICoreV1EnvFromSource) validateConfigMapRef(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigMapRef) { // not required
		return nil
	}

	if m.ConfigMapRef != nil {
		if err := m.ConfigMapRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configMapRef")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1EnvFromSource) validateSecretRef(formats strfmt.Registry) error {
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

// ContextValidate validate this io k8s api core v1 env from source based on the context it is used
func (m *IoK8sAPICoreV1EnvFromSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfigMapRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecretRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1EnvFromSource) contextValidateConfigMapRef(ctx context.Context, formats strfmt.Registry) error {

	if m.ConfigMapRef != nil {
		if err := m.ConfigMapRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configMapRef")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1EnvFromSource) contextValidateSecretRef(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IoK8sAPICoreV1EnvFromSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1EnvFromSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1EnvFromSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
