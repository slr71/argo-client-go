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
)

// GrpcGatewayRuntimeStreamError grpc gateway runtime stream error
//
// swagger:model grpc.gateway.runtime.StreamError
type GrpcGatewayRuntimeStreamError struct {

	// details
	Details []*GoogleProtobufAny `json:"details"`

	// grpc code
	GrpcCode int32 `json:"grpc_code,omitempty"`

	// http code
	HTTPCode int32 `json:"http_code,omitempty"`

	// http status
	HTTPStatus string `json:"http_status,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this grpc gateway runtime stream error
func (m *GrpcGatewayRuntimeStreamError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GrpcGatewayRuntimeStreamError) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.Details) { // not required
		return nil
	}

	for i := 0; i < len(m.Details); i++ {
		if swag.IsZero(m.Details[i]) { // not required
			continue
		}

		if m.Details[i] != nil {
			if err := m.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this grpc gateway runtime stream error based on the context it is used
func (m *GrpcGatewayRuntimeStreamError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GrpcGatewayRuntimeStreamError) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Details); i++ {

		if m.Details[i] != nil {
			if err := m.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GrpcGatewayRuntimeStreamError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GrpcGatewayRuntimeStreamError) UnmarshalBinary(b []byte) error {
	var res GrpcGatewayRuntimeStreamError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
