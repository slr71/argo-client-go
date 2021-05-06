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

// IoArgoprojWorkflowV1alpha1CronWorkflowStatus CronWorkflowStatus is the status of a CronWorkflow
//
// swagger:model io.argoproj.workflow.v1alpha1.CronWorkflowStatus
type IoArgoprojWorkflowV1alpha1CronWorkflowStatus struct {

	// Active is a list of active workflows stemming from this CronWorkflow
	Active []*IoK8sAPICoreV1ObjectReference `json:"active"`

	// Conditions is a list of conditions the CronWorkflow may have
	Conditions []*IoArgoprojWorkflowV1alpha1Condition `json:"conditions"`

	// LastScheduleTime is the last time the CronWorkflow was scheduled
	// Format: date-time
	LastScheduledTime IoK8sApimachineryPkgApisMetaV1Time `json:"lastScheduledTime,omitempty"`
}

// Validate validates this io argoproj workflow v1alpha1 cron workflow status
func (m *IoArgoprojWorkflowV1alpha1CronWorkflowStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastScheduledTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoArgoprojWorkflowV1alpha1CronWorkflowStatus) validateActive(formats strfmt.Registry) error {
	if swag.IsZero(m.Active) { // not required
		return nil
	}

	for i := 0; i < len(m.Active); i++ {
		if swag.IsZero(m.Active[i]) { // not required
			continue
		}

		if m.Active[i] != nil {
			if err := m.Active[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("active" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1CronWorkflowStatus) validateConditions(formats strfmt.Registry) error {
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

func (m *IoArgoprojWorkflowV1alpha1CronWorkflowStatus) validateLastScheduledTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastScheduledTime) { // not required
		return nil
	}

	if err := m.LastScheduledTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastScheduledTime")
		}
		return err
	}

	return nil
}

// ContextValidate validate this io argoproj workflow v1alpha1 cron workflow status based on the context it is used
func (m *IoArgoprojWorkflowV1alpha1CronWorkflowStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActive(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastScheduledTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoArgoprojWorkflowV1alpha1CronWorkflowStatus) contextValidateActive(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Active); i++ {

		if m.Active[i] != nil {
			if err := m.Active[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("active" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1CronWorkflowStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IoArgoprojWorkflowV1alpha1CronWorkflowStatus) contextValidateLastScheduledTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LastScheduledTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastScheduledTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1CronWorkflowStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1CronWorkflowStatus) UnmarshalBinary(b []byte) error {
	var res IoArgoprojWorkflowV1alpha1CronWorkflowStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
