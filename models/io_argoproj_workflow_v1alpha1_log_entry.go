// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoArgoprojWorkflowV1alpha1LogEntry io argoproj workflow v1alpha1 log entry
//
// swagger:model io.argoproj.workflow.v1alpha1.LogEntry
type IoArgoprojWorkflowV1alpha1LogEntry struct {

	// content
	Content string `json:"content,omitempty"`

	// pod name
	PodName string `json:"podName,omitempty"`
}

// Validate validates this io argoproj workflow v1alpha1 log entry
func (m *IoArgoprojWorkflowV1alpha1LogEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io argoproj workflow v1alpha1 log entry based on context it is used
func (m *IoArgoprojWorkflowV1alpha1LogEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1LogEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1LogEntry) UnmarshalBinary(b []byte) error {
	var res IoArgoprojWorkflowV1alpha1LogEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
