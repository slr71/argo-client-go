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

// IoK8sAPICoreV1KeyToPath Maps a string key to a path within a volume.
//
// swagger:model io.k8s.api.core.v1.KeyToPath
type IoK8sAPICoreV1KeyToPath struct {

	// The key to project.
	// Required: true
	Key *string `json:"key"`

	// Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	Mode int32 `json:"mode,omitempty"`

	// The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.
	// Required: true
	Path *string `json:"path"`
}

// Validate validates this io k8s api core v1 key to path
func (m *IoK8sAPICoreV1KeyToPath) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1KeyToPath) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1KeyToPath) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s api core v1 key to path based on context it is used
func (m *IoK8sAPICoreV1KeyToPath) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1KeyToPath) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1KeyToPath) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1KeyToPath
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
