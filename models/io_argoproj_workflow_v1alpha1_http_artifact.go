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

// IoArgoprojWorkflowV1alpha1HTTPArtifact HTTPArtifact allows an file served on HTTP to be placed as an input artifact in a container
//
// swagger:model io.argoproj.workflow.v1alpha1.HTTPArtifact
type IoArgoprojWorkflowV1alpha1HTTPArtifact struct {

	// URL of the artifact
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this io argoproj workflow v1alpha1 HTTP artifact
func (m *IoArgoprojWorkflowV1alpha1HTTPArtifact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoArgoprojWorkflowV1alpha1HTTPArtifact) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io argoproj workflow v1alpha1 HTTP artifact based on context it is used
func (m *IoArgoprojWorkflowV1alpha1HTTPArtifact) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1HTTPArtifact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1HTTPArtifact) UnmarshalBinary(b []byte) error {
	var res IoArgoprojWorkflowV1alpha1HTTPArtifact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
