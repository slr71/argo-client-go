// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoArgoprojWorkflowV1alpha1Condition io argoproj workflow v1alpha1 condition
//
// swagger:model io.argoproj.workflow.v1alpha1.Condition
type IoArgoprojWorkflowV1alpha1Condition struct {

	// Message is the condition message
	Message string `json:"message,omitempty"`

	// Status is the status of the condition
	Status string `json:"status,omitempty"`

	// Type is the type of condition
	Type string `json:"type,omitempty"`
}

// Validate validates this io argoproj workflow v1alpha1 condition
func (m *IoArgoprojWorkflowV1alpha1Condition) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io argoproj workflow v1alpha1 condition based on context it is used
func (m *IoArgoprojWorkflowV1alpha1Condition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1Condition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1Condition) UnmarshalBinary(b []byte) error {
	var res IoArgoprojWorkflowV1alpha1Condition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
