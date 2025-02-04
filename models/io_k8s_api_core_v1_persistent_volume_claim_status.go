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
)

// IoK8sAPICoreV1PersistentVolumeClaimStatus PersistentVolumeClaimStatus is the current status of a persistent volume claim.
//
// swagger:model io.k8s.api.core.v1.PersistentVolumeClaimStatus
type IoK8sAPICoreV1PersistentVolumeClaimStatus struct {

	// AccessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	AccessModes []string `json:"accessModes"`

	// Represents the actual resources of the underlying volume.
	Capacity map[string]IoK8sApimachineryPkgAPIResourceQuantity `json:"capacity,omitempty"`

	// Current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'ResizeStarted'.
	Conditions []*IoK8sAPICoreV1PersistentVolumeClaimCondition `json:"conditions"`

	// Phase represents the current phase of PersistentVolumeClaim.
	Phase string `json:"phase,omitempty"`
}

// Validate validates this io k8s api core v1 persistent volume claim status
func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) validateCapacity(formats strfmt.Registry) error {
	if swag.IsZero(m.Capacity) { // not required
		return nil
	}

	for k := range m.Capacity {

		if val, ok := m.Capacity[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this io k8s api core v1 persistent volume claim status based on the context it is used
func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapacity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) contextValidateCapacity(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Capacity {

		if val, ok := m.Capacity[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Conditions); i++ {

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1PersistentVolumeClaimStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
