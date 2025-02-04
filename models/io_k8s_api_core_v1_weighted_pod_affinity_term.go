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

// IoK8sAPICoreV1WeightedPodAffinityTerm The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)
//
// swagger:model io.k8s.api.core.v1.WeightedPodAffinityTerm
type IoK8sAPICoreV1WeightedPodAffinityTerm struct {

	// Required. A pod affinity term, associated with the corresponding weight.
	// Required: true
	PodAffinityTerm *IoK8sAPICoreV1PodAffinityTerm `json:"podAffinityTerm"`

	// weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
	// Required: true
	Weight *int32 `json:"weight"`
}

// Validate validates this io k8s api core v1 weighted pod affinity term
func (m *IoK8sAPICoreV1WeightedPodAffinityTerm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePodAffinityTerm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1WeightedPodAffinityTerm) validatePodAffinityTerm(formats strfmt.Registry) error {

	if err := validate.Required("podAffinityTerm", "body", m.PodAffinityTerm); err != nil {
		return err
	}

	if m.PodAffinityTerm != nil {
		if err := m.PodAffinityTerm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podAffinityTerm")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1WeightedPodAffinityTerm) validateWeight(formats strfmt.Registry) error {

	if err := validate.Required("weight", "body", m.Weight); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this io k8s api core v1 weighted pod affinity term based on the context it is used
func (m *IoK8sAPICoreV1WeightedPodAffinityTerm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePodAffinityTerm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1WeightedPodAffinityTerm) contextValidatePodAffinityTerm(ctx context.Context, formats strfmt.Registry) error {

	if m.PodAffinityTerm != nil {
		if err := m.PodAffinityTerm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podAffinityTerm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1WeightedPodAffinityTerm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1WeightedPodAffinityTerm) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1WeightedPodAffinityTerm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
