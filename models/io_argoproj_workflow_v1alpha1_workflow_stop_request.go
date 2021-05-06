// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoArgoprojWorkflowV1alpha1WorkflowStopRequest io argoproj workflow v1alpha1 workflow stop request
//
// swagger:model io.argoproj.workflow.v1alpha1.WorkflowStopRequest
type IoArgoprojWorkflowV1alpha1WorkflowStopRequest struct {

	// message
	Message string `json:"message,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// node field selector
	NodeFieldSelector string `json:"nodeFieldSelector,omitempty"`
}

// Validate validates this io argoproj workflow v1alpha1 workflow stop request
func (m *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io argoproj workflow v1alpha1 workflow stop request based on context it is used
func (m *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) UnmarshalBinary(b []byte) error {
	var res IoArgoprojWorkflowV1alpha1WorkflowStopRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
