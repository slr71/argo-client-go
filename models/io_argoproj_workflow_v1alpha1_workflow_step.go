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

// IoArgoprojWorkflowV1alpha1WorkflowStep WorkflowStep is a reference to a template to execute in a series of step
//
// swagger:model io.argoproj.workflow.v1alpha1.WorkflowStep
type IoArgoprojWorkflowV1alpha1WorkflowStep struct {

	// Arguments hold arguments to the template
	Arguments *IoArgoprojWorkflowV1alpha1Arguments `json:"arguments,omitempty"`

	// ContinueOn makes argo to proceed with the following step even if this step fails. Errors and Failed states can be specified
	ContinueOn *IoArgoprojWorkflowV1alpha1ContinueOn `json:"continueOn,omitempty"`

	// Name of the step
	Name string `json:"name,omitempty"`

	// OnExit is a template reference which is invoked at the end of the template, irrespective of the success, failure, or error of the primary template.
	OnExit string `json:"onExit,omitempty"`

	// Template is the name of the template to execute as the step
	Template string `json:"template,omitempty"`

	// TemplateRef is the reference to the template resource to execute as the step.
	TemplateRef *IoArgoprojWorkflowV1alpha1TemplateRef `json:"templateRef,omitempty"`

	// When is an expression in which the step should conditionally execute
	When string `json:"when,omitempty"`

	// WithItems expands a step into multiple parallel steps from the items in the list
	WithItems []IoArgoprojWorkflowV1alpha1Item `json:"withItems"`

	// WithParam expands a step into multiple parallel steps from the value in the parameter, which is expected to be a JSON list.
	WithParam string `json:"withParam,omitempty"`

	// WithSequence expands a step into a numeric sequence
	WithSequence *IoArgoprojWorkflowV1alpha1Sequence `json:"withSequence,omitempty"`
}

// Validate validates this io argoproj workflow v1alpha1 workflow step
func (m *IoArgoprojWorkflowV1alpha1WorkflowStep) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArguments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContinueOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplateRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWithSequence(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowStep) validateArguments(formats strfmt.Registry) error {
	if swag.IsZero(m.Arguments) { // not required
		return nil
	}

	if m.Arguments != nil {
		if err := m.Arguments.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("arguments")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowStep) validateContinueOn(formats strfmt.Registry) error {
	if swag.IsZero(m.ContinueOn) { // not required
		return nil
	}

	if m.ContinueOn != nil {
		if err := m.ContinueOn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("continueOn")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowStep) validateTemplateRef(formats strfmt.Registry) error {
	if swag.IsZero(m.TemplateRef) { // not required
		return nil
	}

	if m.TemplateRef != nil {
		if err := m.TemplateRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("templateRef")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowStep) validateWithSequence(formats strfmt.Registry) error {
	if swag.IsZero(m.WithSequence) { // not required
		return nil
	}

	if m.WithSequence != nil {
		if err := m.WithSequence.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("withSequence")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io argoproj workflow v1alpha1 workflow step based on the context it is used
func (m *IoArgoprojWorkflowV1alpha1WorkflowStep) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArguments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContinueOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTemplateRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWithSequence(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowStep) contextValidateArguments(ctx context.Context, formats strfmt.Registry) error {

	if m.Arguments != nil {
		if err := m.Arguments.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("arguments")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowStep) contextValidateContinueOn(ctx context.Context, formats strfmt.Registry) error {

	if m.ContinueOn != nil {
		if err := m.ContinueOn.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("continueOn")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowStep) contextValidateTemplateRef(ctx context.Context, formats strfmt.Registry) error {

	if m.TemplateRef != nil {
		if err := m.TemplateRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("templateRef")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowStep) contextValidateWithSequence(ctx context.Context, formats strfmt.Registry) error {

	if m.WithSequence != nil {
		if err := m.WithSequence.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("withSequence")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1WorkflowStep) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1WorkflowStep) UnmarshalBinary(b []byte) error {
	var res IoArgoprojWorkflowV1alpha1WorkflowStep
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
