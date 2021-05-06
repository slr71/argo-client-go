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

// IoK8sAPICoreV1AzureDiskVolumeSource AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
//
// swagger:model io.k8s.api.core.v1.AzureDiskVolumeSource
type IoK8sAPICoreV1AzureDiskVolumeSource struct {

	// Host Caching mode: None, Read Only, Read Write.
	CachingMode string `json:"cachingMode,omitempty"`

	// The Name of the data disk in the blob storage
	// Required: true
	DiskName *string `json:"diskName"`

	// The URI the data disk in the blob storage
	// Required: true
	DiskURI *string `json:"diskURI"`

	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FsType string `json:"fsType,omitempty"`

	// Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared
	Kind string `json:"kind,omitempty"`

	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly bool `json:"readOnly,omitempty"`
}

// Validate validates this io k8s api core v1 azure disk volume source
func (m *IoK8sAPICoreV1AzureDiskVolumeSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiskName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiskURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1AzureDiskVolumeSource) validateDiskName(formats strfmt.Registry) error {

	if err := validate.Required("diskName", "body", m.DiskName); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1AzureDiskVolumeSource) validateDiskURI(formats strfmt.Registry) error {

	if err := validate.Required("diskURI", "body", m.DiskURI); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s api core v1 azure disk volume source based on context it is used
func (m *IoK8sAPICoreV1AzureDiskVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1AzureDiskVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1AzureDiskVolumeSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1AzureDiskVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
