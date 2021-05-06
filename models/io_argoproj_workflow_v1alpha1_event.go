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

// IoArgoprojWorkflowV1alpha1Event io argoproj workflow v1alpha1 event
//
// swagger:model io.argoproj.workflow.v1alpha1.Event
type IoArgoprojWorkflowV1alpha1Event struct {

	// Selector (https://github.com/antonmedv/expr) that we must must match the io.argoproj.workflow.v1alpha1. E.g. `payload.message == "test"`
	// Required: true
	Selector *string `json:"selector"`
}

// Validate validates this io argoproj workflow v1alpha1 event
func (m *IoArgoprojWorkflowV1alpha1Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Event) validateSelector(formats strfmt.Registry) error {

	if err := validate.Required("selector", "body", m.Selector); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io argoproj workflow v1alpha1 event based on context it is used
func (m *IoArgoprojWorkflowV1alpha1Event) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1Event) UnmarshalBinary(b []byte) error {
	var res IoArgoprojWorkflowV1alpha1Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
