// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoArgoprojWorkflowV1alpha1TemplateRef TemplateRef is a reference of template resource.
//
// swagger:model io.argoproj.workflow.v1alpha1.TemplateRef
type IoArgoprojWorkflowV1alpha1TemplateRef struct {

	// ClusterScope indicates the referred template is cluster scoped (i.e. a ClusterWorkflowTemplate).
	ClusterScope bool `json:"clusterScope,omitempty"`

	// Name is the resource name of the template.
	Name string `json:"name,omitempty"`

	// RuntimeResolution skips validation at creation time. By enabling this option, you can create the referred workflow template before the actual runtime.
	RuntimeResolution bool `json:"runtimeResolution,omitempty"`

	// Template is the name of referred template in the resource.
	Template string `json:"template,omitempty"`
}

// Validate validates this io argoproj workflow v1alpha1 template ref
func (m *IoArgoprojWorkflowV1alpha1TemplateRef) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io argoproj workflow v1alpha1 template ref based on context it is used
func (m *IoArgoprojWorkflowV1alpha1TemplateRef) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1TemplateRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1TemplateRef) UnmarshalBinary(b []byte) error {
	var res IoArgoprojWorkflowV1alpha1TemplateRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
