// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPICoreV1ConfigMapEnvSource ConfigMapEnvSource selects a ConfigMap to populate the environment variables with.
//
// The contents of the target ConfigMap's Data field will represent the key-value pairs as environment variables.
//
// swagger:model io.k8s.api.core.v1.ConfigMapEnvSource
type IoK8sAPICoreV1ConfigMapEnvSource struct {

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Specify whether the ConfigMap must be defined
	Optional bool `json:"optional,omitempty"`
}

// Validate validates this io k8s api core v1 config map env source
func (m *IoK8sAPICoreV1ConfigMapEnvSource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io k8s api core v1 config map env source based on context it is used
func (m *IoK8sAPICoreV1ConfigMapEnvSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1ConfigMapEnvSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1ConfigMapEnvSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1ConfigMapEnvSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
