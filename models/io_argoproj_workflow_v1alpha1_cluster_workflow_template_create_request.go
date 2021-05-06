// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest io argoproj workflow v1alpha1 cluster workflow template create request
//
// swagger:model io.argoproj.workflow.v1alpha1.ClusterWorkflowTemplateCreateRequest
type IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest struct {

	// create options
	CreateOptions *IoK8sApimachineryPkgApisMetaV1CreateOptions `json:"createOptions,omitempty"`

	// template
	Template *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate `json:"template,omitempty"`
}

// Validate validates this io argoproj workflow v1alpha1 cluster workflow template create request
func (m *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) validateCreateOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateOptions) { // not required
		return nil
	}

	if m.CreateOptions != nil {
		if err := m.CreateOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createOptions")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) validateTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.Template) { // not required
		return nil
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io argoproj workflow v1alpha1 cluster workflow template create request based on the context it is used
func (m *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreateOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) contextValidateCreateOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.CreateOptions != nil {
		if err := m.CreateOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createOptions")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) contextValidateTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.Template != nil {
		if err := m.Template.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) UnmarshalBinary(b []byte) error {
	var res IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
