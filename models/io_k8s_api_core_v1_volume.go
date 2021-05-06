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

// IoK8sAPICoreV1Volume Volume represents a named volume in a pod that may be accessed by any container in the pod.
//
// swagger:model io.k8s.api.core.v1.Volume
type IoK8sAPICoreV1Volume struct {

	// AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	AwsElasticBlockStore *IoK8sAPICoreV1AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty"`

	// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	AzureDisk *IoK8sAPICoreV1AzureDiskVolumeSource `json:"azureDisk,omitempty"`

	// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
	AzureFile *IoK8sAPICoreV1AzureFileVolumeSource `json:"azureFile,omitempty"`

	// CephFS represents a Ceph FS mount on the host that shares a pod's lifetime
	Cephfs *IoK8sAPICoreV1CephFSVolumeSource `json:"cephfs,omitempty"`

	// Cinder represents a cinder volume attached and mounted on kubelets host machine More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
	Cinder *IoK8sAPICoreV1CinderVolumeSource `json:"cinder,omitempty"`

	// ConfigMap represents a configMap that should populate this volume
	ConfigMap *IoK8sAPICoreV1ConfigMapVolumeSource `json:"configMap,omitempty"`

	// CSI (Container Storage Interface) represents storage that is handled by an external CSI driver (Alpha feature).
	Csi *IoK8sAPICoreV1CSIVolumeSource `json:"csi,omitempty"`

	// DownwardAPI represents downward API about the pod that should populate this volume
	DownwardAPI *IoK8sAPICoreV1DownwardAPIVolumeSource `json:"downwardAPI,omitempty"`

	// EmptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	EmptyDir *IoK8sAPICoreV1EmptyDirVolumeSource `json:"emptyDir,omitempty"`

	// FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
	Fc *IoK8sAPICoreV1FCVolumeSource `json:"fc,omitempty"`

	// FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
	FlexVolume *IoK8sAPICoreV1FlexVolumeSource `json:"flexVolume,omitempty"`

	// Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running
	Flocker *IoK8sAPICoreV1FlockerVolumeSource `json:"flocker,omitempty"`

	// GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	GcePersistentDisk *IoK8sAPICoreV1GCEPersistentDiskVolumeSource `json:"gcePersistentDisk,omitempty"`

	// GitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.
	GitRepo *IoK8sAPICoreV1GitRepoVolumeSource `json:"gitRepo,omitempty"`

	// Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md
	Glusterfs *IoK8sAPICoreV1GlusterfsVolumeSource `json:"glusterfs,omitempty"`

	// HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	HostPath *IoK8sAPICoreV1HostPathVolumeSource `json:"hostPath,omitempty"`

	// ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://releases.k8s.io/HEAD/examples/volumes/iscsi/README.md
	Iscsi *IoK8sAPICoreV1ISCSIVolumeSource `json:"iscsi,omitempty"`

	// Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// Required: true
	Name *string `json:"name"`

	// NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Nfs *IoK8sAPICoreV1NFSVolumeSource `json:"nfs,omitempty"`

	// PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	PersistentVolumeClaim *IoK8sAPICoreV1PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty"`

	// PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine
	PhotonPersistentDisk *IoK8sAPICoreV1PhotonPersistentDiskVolumeSource `json:"photonPersistentDisk,omitempty"`

	// PortworxVolume represents a portworx volume attached and mounted on kubelets host machine
	PortworxVolume *IoK8sAPICoreV1PortworxVolumeSource `json:"portworxVolume,omitempty"`

	// Items for all in one resources secrets, configmaps, and downward API
	Projected *IoK8sAPICoreV1ProjectedVolumeSource `json:"projected,omitempty"`

	// Quobyte represents a Quobyte mount on the host that shares a pod's lifetime
	Quobyte *IoK8sAPICoreV1QuobyteVolumeSource `json:"quobyte,omitempty"`

	// RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md
	Rbd *IoK8sAPICoreV1RBDVolumeSource `json:"rbd,omitempty"`

	// ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
	ScaleIO *IoK8sAPICoreV1ScaleIOVolumeSource `json:"scaleIO,omitempty"`

	// Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
	Secret *IoK8sAPICoreV1SecretVolumeSource `json:"secret,omitempty"`

	// StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.
	Storageos *IoK8sAPICoreV1StorageOSVolumeSource `json:"storageos,omitempty"`

	// VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine
	VsphereVolume *IoK8sAPICoreV1VsphereVirtualDiskVolumeSource `json:"vsphereVolume,omitempty"`
}

// Validate validates this io k8s api core v1 volume
func (m *IoK8sAPICoreV1Volume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsElasticBlockStore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCephfs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCinder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigMap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCsi(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownwardAPI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmptyDir(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlexVolume(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlocker(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcePersistentDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitRepo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlusterfs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIscsi(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNfs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersistentVolumeClaim(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhotonPersistentDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortworxVolume(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjected(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuobyte(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRbd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScaleIO(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVsphereVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1Volume) validateAwsElasticBlockStore(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsElasticBlockStore) { // not required
		return nil
	}

	if m.AwsElasticBlockStore != nil {
		if err := m.AwsElasticBlockStore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsElasticBlockStore")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateAzureDisk(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureDisk) { // not required
		return nil
	}

	if m.AzureDisk != nil {
		if err := m.AzureDisk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureDisk")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateAzureFile(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureFile) { // not required
		return nil
	}

	if m.AzureFile != nil {
		if err := m.AzureFile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureFile")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateCephfs(formats strfmt.Registry) error {
	if swag.IsZero(m.Cephfs) { // not required
		return nil
	}

	if m.Cephfs != nil {
		if err := m.Cephfs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cephfs")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateCinder(formats strfmt.Registry) error {
	if swag.IsZero(m.Cinder) { // not required
		return nil
	}

	if m.Cinder != nil {
		if err := m.Cinder.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cinder")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateConfigMap(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigMap) { // not required
		return nil
	}

	if m.ConfigMap != nil {
		if err := m.ConfigMap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configMap")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateCsi(formats strfmt.Registry) error {
	if swag.IsZero(m.Csi) { // not required
		return nil
	}

	if m.Csi != nil {
		if err := m.Csi.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("csi")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateDownwardAPI(formats strfmt.Registry) error {
	if swag.IsZero(m.DownwardAPI) { // not required
		return nil
	}

	if m.DownwardAPI != nil {
		if err := m.DownwardAPI.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downwardAPI")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateEmptyDir(formats strfmt.Registry) error {
	if swag.IsZero(m.EmptyDir) { // not required
		return nil
	}

	if m.EmptyDir != nil {
		if err := m.EmptyDir.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emptyDir")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateFc(formats strfmt.Registry) error {
	if swag.IsZero(m.Fc) { // not required
		return nil
	}

	if m.Fc != nil {
		if err := m.Fc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fc")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateFlexVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.FlexVolume) { // not required
		return nil
	}

	if m.FlexVolume != nil {
		if err := m.FlexVolume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flexVolume")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateFlocker(formats strfmt.Registry) error {
	if swag.IsZero(m.Flocker) { // not required
		return nil
	}

	if m.Flocker != nil {
		if err := m.Flocker.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flocker")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateGcePersistentDisk(formats strfmt.Registry) error {
	if swag.IsZero(m.GcePersistentDisk) { // not required
		return nil
	}

	if m.GcePersistentDisk != nil {
		if err := m.GcePersistentDisk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcePersistentDisk")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateGitRepo(formats strfmt.Registry) error {
	if swag.IsZero(m.GitRepo) { // not required
		return nil
	}

	if m.GitRepo != nil {
		if err := m.GitRepo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gitRepo")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateGlusterfs(formats strfmt.Registry) error {
	if swag.IsZero(m.Glusterfs) { // not required
		return nil
	}

	if m.Glusterfs != nil {
		if err := m.Glusterfs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("glusterfs")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateHostPath(formats strfmt.Registry) error {
	if swag.IsZero(m.HostPath) { // not required
		return nil
	}

	if m.HostPath != nil {
		if err := m.HostPath.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostPath")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateIscsi(formats strfmt.Registry) error {
	if swag.IsZero(m.Iscsi) { // not required
		return nil
	}

	if m.Iscsi != nil {
		if err := m.Iscsi.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iscsi")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateNfs(formats strfmt.Registry) error {
	if swag.IsZero(m.Nfs) { // not required
		return nil
	}

	if m.Nfs != nil {
		if err := m.Nfs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfs")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validatePersistentVolumeClaim(formats strfmt.Registry) error {
	if swag.IsZero(m.PersistentVolumeClaim) { // not required
		return nil
	}

	if m.PersistentVolumeClaim != nil {
		if err := m.PersistentVolumeClaim.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("persistentVolumeClaim")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validatePhotonPersistentDisk(formats strfmt.Registry) error {
	if swag.IsZero(m.PhotonPersistentDisk) { // not required
		return nil
	}

	if m.PhotonPersistentDisk != nil {
		if err := m.PhotonPersistentDisk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("photonPersistentDisk")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validatePortworxVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.PortworxVolume) { // not required
		return nil
	}

	if m.PortworxVolume != nil {
		if err := m.PortworxVolume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("portworxVolume")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateProjected(formats strfmt.Registry) error {
	if swag.IsZero(m.Projected) { // not required
		return nil
	}

	if m.Projected != nil {
		if err := m.Projected.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("projected")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateQuobyte(formats strfmt.Registry) error {
	if swag.IsZero(m.Quobyte) { // not required
		return nil
	}

	if m.Quobyte != nil {
		if err := m.Quobyte.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quobyte")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateRbd(formats strfmt.Registry) error {
	if swag.IsZero(m.Rbd) { // not required
		return nil
	}

	if m.Rbd != nil {
		if err := m.Rbd.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rbd")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateScaleIO(formats strfmt.Registry) error {
	if swag.IsZero(m.ScaleIO) { // not required
		return nil
	}

	if m.ScaleIO != nil {
		if err := m.ScaleIO.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scaleIO")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.Secret) { // not required
		return nil
	}

	if m.Secret != nil {
		if err := m.Secret.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secret")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateStorageos(formats strfmt.Registry) error {
	if swag.IsZero(m.Storageos) { // not required
		return nil
	}

	if m.Storageos != nil {
		if err := m.Storageos.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageos")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) validateVsphereVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.VsphereVolume) { // not required
		return nil
	}

	if m.VsphereVolume != nil {
		if err := m.VsphereVolume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsphereVolume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io k8s api core v1 volume based on the context it is used
func (m *IoK8sAPICoreV1Volume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsElasticBlockStore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureDisk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureFile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCephfs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCinder(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfigMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCsi(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDownwardAPI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmptyDir(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlexVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlocker(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGcePersistentDisk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGitRepo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGlusterfs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHostPath(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIscsi(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNfs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePersistentVolumeClaim(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhotonPersistentDisk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePortworxVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProjected(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuobyte(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRbd(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScaleIO(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageos(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVsphereVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateAwsElasticBlockStore(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsElasticBlockStore != nil {
		if err := m.AwsElasticBlockStore.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsElasticBlockStore")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateAzureDisk(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureDisk != nil {
		if err := m.AzureDisk.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureDisk")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateAzureFile(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureFile != nil {
		if err := m.AzureFile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureFile")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateCephfs(ctx context.Context, formats strfmt.Registry) error {

	if m.Cephfs != nil {
		if err := m.Cephfs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cephfs")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateCinder(ctx context.Context, formats strfmt.Registry) error {

	if m.Cinder != nil {
		if err := m.Cinder.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cinder")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateConfigMap(ctx context.Context, formats strfmt.Registry) error {

	if m.ConfigMap != nil {
		if err := m.ConfigMap.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configMap")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateCsi(ctx context.Context, formats strfmt.Registry) error {

	if m.Csi != nil {
		if err := m.Csi.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("csi")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateDownwardAPI(ctx context.Context, formats strfmt.Registry) error {

	if m.DownwardAPI != nil {
		if err := m.DownwardAPI.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downwardAPI")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateEmptyDir(ctx context.Context, formats strfmt.Registry) error {

	if m.EmptyDir != nil {
		if err := m.EmptyDir.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emptyDir")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateFc(ctx context.Context, formats strfmt.Registry) error {

	if m.Fc != nil {
		if err := m.Fc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fc")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateFlexVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.FlexVolume != nil {
		if err := m.FlexVolume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flexVolume")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateFlocker(ctx context.Context, formats strfmt.Registry) error {

	if m.Flocker != nil {
		if err := m.Flocker.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flocker")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateGcePersistentDisk(ctx context.Context, formats strfmt.Registry) error {

	if m.GcePersistentDisk != nil {
		if err := m.GcePersistentDisk.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcePersistentDisk")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateGitRepo(ctx context.Context, formats strfmt.Registry) error {

	if m.GitRepo != nil {
		if err := m.GitRepo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gitRepo")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateGlusterfs(ctx context.Context, formats strfmt.Registry) error {

	if m.Glusterfs != nil {
		if err := m.Glusterfs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("glusterfs")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateHostPath(ctx context.Context, formats strfmt.Registry) error {

	if m.HostPath != nil {
		if err := m.HostPath.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostPath")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateIscsi(ctx context.Context, formats strfmt.Registry) error {

	if m.Iscsi != nil {
		if err := m.Iscsi.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iscsi")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateNfs(ctx context.Context, formats strfmt.Registry) error {

	if m.Nfs != nil {
		if err := m.Nfs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfs")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidatePersistentVolumeClaim(ctx context.Context, formats strfmt.Registry) error {

	if m.PersistentVolumeClaim != nil {
		if err := m.PersistentVolumeClaim.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("persistentVolumeClaim")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidatePhotonPersistentDisk(ctx context.Context, formats strfmt.Registry) error {

	if m.PhotonPersistentDisk != nil {
		if err := m.PhotonPersistentDisk.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("photonPersistentDisk")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidatePortworxVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.PortworxVolume != nil {
		if err := m.PortworxVolume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("portworxVolume")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateProjected(ctx context.Context, formats strfmt.Registry) error {

	if m.Projected != nil {
		if err := m.Projected.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("projected")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateQuobyte(ctx context.Context, formats strfmt.Registry) error {

	if m.Quobyte != nil {
		if err := m.Quobyte.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quobyte")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateRbd(ctx context.Context, formats strfmt.Registry) error {

	if m.Rbd != nil {
		if err := m.Rbd.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rbd")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateScaleIO(ctx context.Context, formats strfmt.Registry) error {

	if m.ScaleIO != nil {
		if err := m.ScaleIO.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scaleIO")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateSecret(ctx context.Context, formats strfmt.Registry) error {

	if m.Secret != nil {
		if err := m.Secret.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secret")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateStorageos(ctx context.Context, formats strfmt.Registry) error {

	if m.Storageos != nil {
		if err := m.Storageos.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageos")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Volume) contextValidateVsphereVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.VsphereVolume != nil {
		if err := m.VsphereVolume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsphereVolume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1Volume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1Volume) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1Volume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
