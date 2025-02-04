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

// IoArgoprojWorkflowV1alpha1Artifact Artifact indicates an artifact to place at a specified path
//
// swagger:model io.argoproj.workflow.v1alpha1.Artifact
type IoArgoprojWorkflowV1alpha1Artifact struct {

	// Archive controls how the artifact will be saved to the artifact repository.
	Archive *IoArgoprojWorkflowV1alpha1ArchiveStrategy `json:"archive,omitempty"`

	// ArchiveLogs indicates if the container logs should be archived
	ArchiveLogs bool `json:"archiveLogs,omitempty"`

	// Artifactory contains artifactory artifact location details
	Artifactory *IoArgoprojWorkflowV1alpha1ArtifactoryArtifact `json:"artifactory,omitempty"`

	// From allows an artifact to reference an artifact from a previous step
	From string `json:"from,omitempty"`

	// GCS contains GCS artifact location details
	Gcs *IoArgoprojWorkflowV1alpha1GCSArtifact `json:"gcs,omitempty"`

	// Git contains git artifact location details
	Git *IoArgoprojWorkflowV1alpha1GitArtifact `json:"git,omitempty"`

	// GlobalName exports an output artifact to the global scope, making it available as '{{io.argoproj.workflow.v1alpha1.outputs.artifacts.XXXX}} and in workflow.status.outputs.artifacts
	GlobalName string `json:"globalName,omitempty"`

	// HDFS contains HDFS artifact location details
	Hdfs *IoArgoprojWorkflowV1alpha1HDFSArtifact `json:"hdfs,omitempty"`

	// HTTP contains HTTP artifact location details
	HTTP *IoArgoprojWorkflowV1alpha1HTTPArtifact `json:"http,omitempty"`

	// mode bits to use on this file, must be a value between 0 and 0777 set when loading input artifacts.
	Mode int32 `json:"mode,omitempty"`

	// name of the artifact. must be unique within a template's inputs/outputs.
	// Required: true
	Name *string `json:"name"`

	// Make Artifacts optional, if Artifacts doesn't generate or exist
	Optional bool `json:"optional,omitempty"`

	// OSS contains OSS artifact location details
	Oss *IoArgoprojWorkflowV1alpha1OSSArtifact `json:"oss,omitempty"`

	// Path is the container path to the artifact
	Path string `json:"path,omitempty"`

	// Raw contains raw artifact location details
	Raw *IoArgoprojWorkflowV1alpha1RawArtifact `json:"raw,omitempty"`

	// S3 contains S3 artifact location details
	S3 *IoArgoprojWorkflowV1alpha1S3Artifact `json:"s3,omitempty"`

	// SubPath allows an artifact to be sourced from a subpath within the specified source
	SubPath string `json:"subPath,omitempty"`
}

// Validate validates this io argoproj workflow v1alpha1 artifact
func (m *IoArgoprojWorkflowV1alpha1Artifact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArtifactory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHdfs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOss(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRaw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Artifact) validateArchive(formats strfmt.Registry) error {
	if swag.IsZero(m.Archive) { // not required
		return nil
	}

	if m.Archive != nil {
		if err := m.Archive.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("archive")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Artifact) validateArtifactory(formats strfmt.Registry) error {
	if swag.IsZero(m.Artifactory) { // not required
		return nil
	}

	if m.Artifactory != nil {
		if err := m.Artifactory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifactory")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Artifact) validateGcs(formats strfmt.Registry) error {
	if swag.IsZero(m.Gcs) { // not required
		return nil
	}

	if m.Gcs != nil {
		if err := m.Gcs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcs")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Artifact) validateGit(formats strfmt.Registry) error {
	if swag.IsZero(m.Git) { // not required
		return nil
	}

	if m.Git != nil {
		if err := m.Git.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("git")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Artifact) validateHdfs(formats strfmt.Registry) error {
	if swag.IsZero(m.Hdfs) { // not required
		return nil
	}

	if m.Hdfs != nil {
		if err := m.Hdfs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hdfs")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Artifact) validateHTTP(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTP) { // not required
		return nil
	}

	if m.HTTP != nil {
		if err := m.HTTP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Artifact) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Artifact) validateOss(formats strfmt.Registry) error {
	if swag.IsZero(m.Oss) { // not required
		return nil
	}

	if m.Oss != nil {
		if err := m.Oss.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oss")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Artifact) validateRaw(formats strfmt.Registry) error {
	if swag.IsZero(m.Raw) { // not required
		return nil
	}

	if m.Raw != nil {
		if err := m.Raw.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("raw")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Artifact) validateS3(formats strfmt.Registry) error {
	if swag.IsZero(m.S3) { // not required
		return nil
	}

	if m.S3 != nil {
		if err := m.S3.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io argoproj workflow v1alpha1 artifact based on the context it is used
func (m *IoArgoprojWorkflowV1alpha1Artifact) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArchive(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateArtifactory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGcs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHdfs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHTTP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOss(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRaw(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateS3(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Artifact) contextValidateArchive(ctx context.Context, formats strfmt.Registry) error {

	if m.Archive != nil {
		if err := m.Archive.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("archive")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Artifact) contextValidateArtifactory(ctx context.Context, formats strfmt.Registry) error {

	if m.Artifactory != nil {
		if err := m.Artifactory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifactory")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Artifact) contextValidateGcs(ctx context.Context, formats strfmt.Registry) error {

	if m.Gcs != nil {
		if err := m.Gcs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcs")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Artifact) contextValidateGit(ctx context.Context, formats strfmt.Registry) error {

	if m.Git != nil {
		if err := m.Git.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("git")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Artifact) contextValidateHdfs(ctx context.Context, formats strfmt.Registry) error {

	if m.Hdfs != nil {
		if err := m.Hdfs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hdfs")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Artifact) contextValidateHTTP(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTP != nil {
		if err := m.HTTP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Artifact) contextValidateOss(ctx context.Context, formats strfmt.Registry) error {

	if m.Oss != nil {
		if err := m.Oss.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oss")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Artifact) contextValidateRaw(ctx context.Context, formats strfmt.Registry) error {

	if m.Raw != nil {
		if err := m.Raw.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("raw")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1Artifact) contextValidateS3(ctx context.Context, formats strfmt.Registry) error {

	if m.S3 != nil {
		if err := m.S3.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1Artifact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1Artifact) UnmarshalBinary(b []byte) error {
	var res IoArgoprojWorkflowV1alpha1Artifact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
