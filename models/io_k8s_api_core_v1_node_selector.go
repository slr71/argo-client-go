// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IoK8sAPICoreV1NodeSelector A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
//
// swagger:model io.k8s.api.core.v1.NodeSelector
type IoK8sAPICoreV1NodeSelector struct {

	// Required. A list of node selector terms. The terms are ORed.
	// Required: true
	NodeSelectorTerms []*IoK8sAPICoreV1NodeSelectorTerm `json:"nodeSelectorTerms"`
}

// Validate validates this io k8s api core v1 node selector
func (m *IoK8sAPICoreV1NodeSelector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeSelectorTerms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1NodeSelector) validateNodeSelectorTerms(formats strfmt.Registry) error {

	if err := validate.Required("nodeSelectorTerms", "body", m.NodeSelectorTerms); err != nil {
		return err
	}

	for i := 0; i < len(m.NodeSelectorTerms); i++ {
		if swag.IsZero(m.NodeSelectorTerms[i]) { // not required
			continue
		}

		if m.NodeSelectorTerms[i] != nil {
			if err := m.NodeSelectorTerms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeSelectorTerms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this io k8s api core v1 node selector based on the context it is used
func (m *IoK8sAPICoreV1NodeSelector) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNodeSelectorTerms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1NodeSelector) contextValidateNodeSelectorTerms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NodeSelectorTerms); i++ {

		if m.NodeSelectorTerms[i] != nil {
			if err := m.NodeSelectorTerms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeSelectorTerms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1NodeSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1NodeSelector) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1NodeSelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
