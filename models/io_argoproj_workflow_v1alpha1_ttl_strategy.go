// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoArgoprojWorkflowV1alpha1TTLStrategy TTLStrategy is the strategy for the time to live depending on if the workflow succeeded or failed
//
// swagger:model io.argoproj.workflow.v1alpha1.TTLStrategy
type IoArgoprojWorkflowV1alpha1TTLStrategy struct {

	// SecondsAfterCompletion is the number of seconds to live after completion
	SecondsAfterCompletion int32 `json:"secondsAfterCompletion,omitempty"`

	// SecondsAfterFailure is the number of seconds to live after failure
	SecondsAfterFailure int32 `json:"secondsAfterFailure,omitempty"`

	// SecondsAfterSuccess is the number of seconds to live after success
	SecondsAfterSuccess int32 `json:"secondsAfterSuccess,omitempty"`
}

// Validate validates this io argoproj workflow v1alpha1 TTL strategy
func (m *IoArgoprojWorkflowV1alpha1TTLStrategy) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io argoproj workflow v1alpha1 TTL strategy based on context it is used
func (m *IoArgoprojWorkflowV1alpha1TTLStrategy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1TTLStrategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1TTLStrategy) UnmarshalBinary(b []byte) error {
	var res IoArgoprojWorkflowV1alpha1TTLStrategy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
