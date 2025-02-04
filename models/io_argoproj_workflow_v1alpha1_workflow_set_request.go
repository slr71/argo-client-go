// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoArgoprojWorkflowV1alpha1WorkflowSetRequest io argoproj workflow v1alpha1 workflow set request
//
// swagger:model io.argoproj.workflow.v1alpha1.WorkflowSetRequest
type IoArgoprojWorkflowV1alpha1WorkflowSetRequest struct {

	// message
	Message string `json:"message,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// node field selector
	NodeFieldSelector string `json:"nodeFieldSelector,omitempty"`

	// output parameters
	OutputParameters string `json:"outputParameters,omitempty"`

	// phase
	Phase string `json:"phase,omitempty"`
}

// Validate validates this io argoproj workflow v1alpha1 workflow set request
func (m *IoArgoprojWorkflowV1alpha1WorkflowSetRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io argoproj workflow v1alpha1 workflow set request based on context it is used
func (m *IoArgoprojWorkflowV1alpha1WorkflowSetRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1WorkflowSetRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1WorkflowSetRequest) UnmarshalBinary(b []byte) error {
	var res IoArgoprojWorkflowV1alpha1WorkflowSetRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
