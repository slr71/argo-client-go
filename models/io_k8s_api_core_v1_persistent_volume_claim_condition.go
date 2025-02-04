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

// IoK8sAPICoreV1PersistentVolumeClaimCondition PersistentVolumeClaimCondition contails details about state of pvc
//
// swagger:model io.k8s.api.core.v1.PersistentVolumeClaimCondition
type IoK8sAPICoreV1PersistentVolumeClaimCondition struct {

	// Last time we probed the condition.
	// Format: date-time
	LastProbeTime IoK8sApimachineryPkgApisMetaV1Time `json:"lastProbeTime,omitempty"`

	// Last time the condition transitioned from one status to another.
	// Format: date-time
	LastTransitionTime IoK8sApimachineryPkgApisMetaV1Time `json:"lastTransitionTime,omitempty"`

	// Human-readable message indicating details about last transition.
	Message string `json:"message,omitempty"`

	// Unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If it reports "ResizeStarted" that means the underlying persistent volume is being resized.
	Reason string `json:"reason,omitempty"`

	// status
	// Required: true
	Status *string `json:"status"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this io k8s api core v1 persistent volume claim condition
func (m *IoK8sAPICoreV1PersistentVolumeClaimCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastProbeTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastTransitionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimCondition) validateLastProbeTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastProbeTime) { // not required
		return nil
	}

	if err := m.LastProbeTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastProbeTime")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimCondition) validateLastTransitionTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastTransitionTime) { // not required
		return nil
	}

	if err := m.LastTransitionTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastTransitionTime")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimCondition) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimCondition) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this io k8s api core v1 persistent volume claim condition based on the context it is used
func (m *IoK8sAPICoreV1PersistentVolumeClaimCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastProbeTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastTransitionTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimCondition) contextValidateLastProbeTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LastProbeTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastProbeTime")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimCondition) contextValidateLastTransitionTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LastTransitionTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastTransitionTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1PersistentVolumeClaimCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1PersistentVolumeClaimCondition) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1PersistentVolumeClaimCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
