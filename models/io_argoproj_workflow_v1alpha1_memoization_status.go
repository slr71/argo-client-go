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

// IoArgoprojWorkflowV1alpha1MemoizationStatus io argoproj workflow v1alpha1 memoization status
//
// swagger:model io.argoproj.workflow.v1alpha1.MemoizationStatus
type IoArgoprojWorkflowV1alpha1MemoizationStatus struct {

	// cache name
	// Required: true
	CacheName *string `json:"cacheName"`

	// hit
	// Required: true
	Hit *bool `json:"hit"`

	// key
	// Required: true
	Key *string `json:"key"`
}

// Validate validates this io argoproj workflow v1alpha1 memoization status
func (m *IoArgoprojWorkflowV1alpha1MemoizationStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCacheName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoArgoprojWorkflowV1alpha1MemoizationStatus) validateCacheName(formats strfmt.Registry) error {

	if err := validate.Required("cacheName", "body", m.CacheName); err != nil {
		return err
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1MemoizationStatus) validateHit(formats strfmt.Registry) error {

	if err := validate.Required("hit", "body", m.Hit); err != nil {
		return err
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1MemoizationStatus) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io argoproj workflow v1alpha1 memoization status based on context it is used
func (m *IoArgoprojWorkflowV1alpha1MemoizationStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1MemoizationStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1MemoizationStatus) UnmarshalBinary(b []byte) error {
	var res IoArgoprojWorkflowV1alpha1MemoizationStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
