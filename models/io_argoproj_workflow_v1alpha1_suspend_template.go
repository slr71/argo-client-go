// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoArgoprojWorkflowV1alpha1SuspendTemplate SuspendTemplate is a template subtype to suspend a workflow at a predetermined point in time
//
// swagger:model io.argoproj.workflow.v1alpha1.SuspendTemplate
type IoArgoprojWorkflowV1alpha1SuspendTemplate struct {

	// Duration is the seconds to wait before automatically resuming a template
	Duration string `json:"duration,omitempty"`
}

// Validate validates this io argoproj workflow v1alpha1 suspend template
func (m *IoArgoprojWorkflowV1alpha1SuspendTemplate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io argoproj workflow v1alpha1 suspend template based on context it is used
func (m *IoArgoprojWorkflowV1alpha1SuspendTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1SuspendTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1SuspendTemplate) UnmarshalBinary(b []byte) error {
	var res IoArgoprojWorkflowV1alpha1SuspendTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
