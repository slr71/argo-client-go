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

// IoK8sAPICoreV1ResourceFieldSelector ResourceFieldSelector represents container resources (cpu, memory) and their output format
//
// swagger:model io.k8s.api.core.v1.ResourceFieldSelector
type IoK8sAPICoreV1ResourceFieldSelector struct {

	// Container name: required for volumes, optional for env vars
	ContainerName string `json:"containerName,omitempty"`

	// Specifies the output format of the exposed resources, defaults to "1"
	Divisor IoK8sApimachineryPkgAPIResourceQuantity `json:"divisor,omitempty"`

	// Required: resource to select
	// Required: true
	Resource *string `json:"resource"`
}

// Validate validates this io k8s api core v1 resource field selector
func (m *IoK8sAPICoreV1ResourceFieldSelector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDivisor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1ResourceFieldSelector) validateDivisor(formats strfmt.Registry) error {
	if swag.IsZero(m.Divisor) { // not required
		return nil
	}

	if err := m.Divisor.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("divisor")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1ResourceFieldSelector) validateResource(formats strfmt.Registry) error {

	if err := validate.Required("resource", "body", m.Resource); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this io k8s api core v1 resource field selector based on the context it is used
func (m *IoK8sAPICoreV1ResourceFieldSelector) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDivisor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1ResourceFieldSelector) contextValidateDivisor(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Divisor.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("divisor")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1ResourceFieldSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1ResourceFieldSelector) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1ResourceFieldSelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
