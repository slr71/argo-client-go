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

// IoArgoprojWorkflowV1alpha1S3Artifact S3Artifact is the location of an S3 artifact
//
// swagger:model io.argoproj.workflow.v1alpha1.S3Artifact
type IoArgoprojWorkflowV1alpha1S3Artifact struct {

	// AccessKeySecret is the secret selector to the bucket's access key
	// Required: true
	AccessKeySecret *IoK8sAPICoreV1SecretKeySelector `json:"accessKeySecret"`

	// Bucket is the name of the bucket
	// Required: true
	Bucket *string `json:"bucket"`

	// Endpoint is the hostname of the bucket endpoint
	// Required: true
	Endpoint *string `json:"endpoint"`

	// Insecure will connect to the service with TLS
	Insecure bool `json:"insecure,omitempty"`

	// Key is the key in the bucket where the artifact resides
	// Required: true
	Key *string `json:"key"`

	// Region contains the optional bucket region
	Region string `json:"region,omitempty"`

	// RoleARN is the Amazon Resource Name (ARN) of the role to assume.
	RoleARN string `json:"roleARN,omitempty"`

	// SecretKeySecret is the secret selector to the bucket's secret key
	// Required: true
	SecretKeySecret *IoK8sAPICoreV1SecretKeySelector `json:"secretKeySecret"`

	// UseSDKCreds tells the driver to figure out credentials based on sdk defaults.
	UseSDKCreds bool `json:"useSDKCreds,omitempty"`
}

// Validate validates this io argoproj workflow v1alpha1 s3 artifact
func (m *IoArgoprojWorkflowV1alpha1S3Artifact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessKeySecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBucket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretKeySecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoArgoprojWorkflowV1alpha1S3Artifact) validateAccessKeySecret(formats strfmt.Registry) error {

	if err := validate.Required("accessKeySecret", "body", m.AccessKeySecret); err != nil {
		return err
	}

	if m.AccessKeySecret != nil {
		if err := m.AccessKeySecret.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessKeySecret")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1S3Artifact) validateBucket(formats strfmt.Registry) error {

	if err := validate.Required("bucket", "body", m.Bucket); err != nil {
		return err
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1S3Artifact) validateEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("endpoint", "body", m.Endpoint); err != nil {
		return err
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1S3Artifact) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1S3Artifact) validateSecretKeySecret(formats strfmt.Registry) error {

	if err := validate.Required("secretKeySecret", "body", m.SecretKeySecret); err != nil {
		return err
	}

	if m.SecretKeySecret != nil {
		if err := m.SecretKeySecret.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secretKeySecret")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io argoproj workflow v1alpha1 s3 artifact based on the context it is used
func (m *IoArgoprojWorkflowV1alpha1S3Artifact) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessKeySecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecretKeySecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoArgoprojWorkflowV1alpha1S3Artifact) contextValidateAccessKeySecret(ctx context.Context, formats strfmt.Registry) error {

	if m.AccessKeySecret != nil {
		if err := m.AccessKeySecret.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessKeySecret")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1S3Artifact) contextValidateSecretKeySecret(ctx context.Context, formats strfmt.Registry) error {

	if m.SecretKeySecret != nil {
		if err := m.SecretKeySecret.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secretKeySecret")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1S3Artifact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1S3Artifact) UnmarshalBinary(b []byte) error {
	var res IoArgoprojWorkflowV1alpha1S3Artifact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
