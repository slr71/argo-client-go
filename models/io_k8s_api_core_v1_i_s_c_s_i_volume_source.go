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

// IoK8sAPICoreV1ISCSIVolumeSource Represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.
//
// swagger:model io.k8s.api.core.v1.ISCSIVolumeSource
type IoK8sAPICoreV1ISCSIVolumeSource struct {

	// whether support iSCSI Discovery CHAP authentication
	ChapAuthDiscovery bool `json:"chapAuthDiscovery,omitempty"`

	// whether support iSCSI Session CHAP authentication
	ChapAuthSession bool `json:"chapAuthSession,omitempty"`

	// Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
	FsType string `json:"fsType,omitempty"`

	// Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.
	InitiatorName string `json:"initiatorName,omitempty"`

	// Target iSCSI Qualified Name.
	// Required: true
	Iqn *string `json:"iqn"`

	// iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).
	IscsiInterface string `json:"iscsiInterface,omitempty"`

	// iSCSI Target Lun number.
	// Required: true
	Lun *int32 `json:"lun"`

	// iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	Portals []string `json:"portals"`

	// ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.
	ReadOnly bool `json:"readOnly,omitempty"`

	// CHAP Secret for iSCSI target and initiator authentication
	SecretRef *IoK8sAPICoreV1LocalObjectReference `json:"secretRef,omitempty"`

	// iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	// Required: true
	TargetPortal *string `json:"targetPortal"`
}

// Validate validates this io k8s api core v1 i s c s i volume source
func (m *IoK8sAPICoreV1ISCSIVolumeSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIqn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLun(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetPortal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1ISCSIVolumeSource) validateIqn(formats strfmt.Registry) error {

	if err := validate.Required("iqn", "body", m.Iqn); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1ISCSIVolumeSource) validateLun(formats strfmt.Registry) error {

	if err := validate.Required("lun", "body", m.Lun); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1ISCSIVolumeSource) validateSecretRef(formats strfmt.Registry) error {
	if swag.IsZero(m.SecretRef) { // not required
		return nil
	}

	if m.SecretRef != nil {
		if err := m.SecretRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secretRef")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1ISCSIVolumeSource) validateTargetPortal(formats strfmt.Registry) error {

	if err := validate.Required("targetPortal", "body", m.TargetPortal); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this io k8s api core v1 i s c s i volume source based on the context it is used
func (m *IoK8sAPICoreV1ISCSIVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSecretRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1ISCSIVolumeSource) contextValidateSecretRef(ctx context.Context, formats strfmt.Registry) error {

	if m.SecretRef != nil {
		if err := m.SecretRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secretRef")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1ISCSIVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1ISCSIVolumeSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1ISCSIVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
