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

// IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec WorkflowTemplateSpec is a spec of WorkflowTemplate.
//
// swagger:model io.argoproj.workflow.v1alpha1.WorkflowTemplateSpec
type IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec struct {

	// Optional duration in seconds relative to the workflow start time which the workflow is allowed to run before the controller terminates the io.argoproj.workflow.v1alpha1. A value of zero is used to terminate a Running workflow
	ActiveDeadlineSeconds int64 `json:"activeDeadlineSeconds,omitempty"`

	// Affinity sets the scheduling constraints for all pods in the io.argoproj.workflow.v1alpha1. Can be overridden by an affinity specified in the template
	Affinity *IoK8sAPICoreV1Affinity `json:"affinity,omitempty"`

	// Arguments contain the parameters and artifacts sent to the workflow entrypoint Parameters are referencable globally using the 'workflow' variable prefix. e.g. {{io.argoproj.workflow.v1alpha1.parameters.myparam}}
	Arguments *IoArgoprojWorkflowV1alpha1Arguments `json:"arguments,omitempty"`

	// ArtifactRepositoryRef specifies the configMap name and key containing the artifact repository config.
	ArtifactRepositoryRef *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRef `json:"artifactRepositoryRef,omitempty"`

	// AutomountServiceAccountToken indicates whether a service account token should be automatically mounted in pods. ServiceAccountName of ExecutorConfig must be specified if this value is false.
	AutomountServiceAccountToken bool `json:"automountServiceAccountToken,omitempty"`

	// PodDNSConfig defines the DNS parameters of a pod in addition to those generated from DNSPolicy.
	DNSConfig *IoK8sAPICoreV1PodDNSConfig `json:"dnsConfig,omitempty"`

	// Set DNS policy for the pod. Defaults to "ClusterFirst". Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.
	DNSPolicy string `json:"dnsPolicy,omitempty"`

	// Entrypoint is a template reference to the starting point of the io.argoproj.workflow.v1alpha1.
	Entrypoint string `json:"entrypoint,omitempty"`

	// Executor holds configurations of executor containers of the io.argoproj.workflow.v1alpha1.
	Executor *IoArgoprojWorkflowV1alpha1ExecutorConfig `json:"executor,omitempty"`

	// host aliases
	HostAliases []*IoK8sAPICoreV1HostAlias `json:"hostAliases"`

	// Host networking requested for this workflow pod. Default to false.
	HostNetwork bool `json:"hostNetwork,omitempty"`

	// ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
	ImagePullSecrets []*IoK8sAPICoreV1LocalObjectReference `json:"imagePullSecrets"`

	// Metrics are a list of metrics emitted from this Workflow
	Metrics *IoArgoprojWorkflowV1alpha1Metrics `json:"metrics,omitempty"`

	// NodeSelector is a selector which will result in all pods of the workflow to be scheduled on the selected node(s). This is able to be overridden by a nodeSelector specified in the template.
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// OnExit is a template reference which is invoked at the end of the workflow, irrespective of the success, failure, or error of the primary io.argoproj.workflow.v1alpha1.
	OnExit string `json:"onExit,omitempty"`

	// Parallelism limits the max total parallel pods that can execute at the same time in a workflow
	Parallelism int64 `json:"parallelism,omitempty"`

	// PodDisruptionBudget holds the number of concurrent disruptions that you allow for Workflow's Pods. Controller will automatically add the selector with workflow name, if selector is empty. Optional: Defaults to empty.
	PodDisruptionBudget *IoK8sAPIPolicyV1beta1PodDisruptionBudgetSpec `json:"podDisruptionBudget,omitempty"`

	// PodGC describes the strategy to use when to deleting completed pods
	PodGC *IoArgoprojWorkflowV1alpha1PodGC `json:"podGC,omitempty"`

	// Priority to apply to workflow pods.
	PodPriority int32 `json:"podPriority,omitempty"`

	// PriorityClassName to apply to workflow pods.
	PodPriorityClassName string `json:"podPriorityClassName,omitempty"`

	// PodSpecPatch holds strategic merge patch to apply against the pod spec. Allows parameterization of container fields which are not strings (e.g. resource limits).
	PodSpecPatch string `json:"podSpecPatch,omitempty"`

	// Priority is used if controller is configured to process limited number of workflows in parallel. Workflows with higher priority are processed first.
	Priority int32 `json:"priority,omitempty"`

	// Set scheduler name for all pods. Will be overridden if container/script template's scheduler name is set. Default scheduler will be used if neither specified.
	SchedulerName string `json:"schedulerName,omitempty"`

	// SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty.  See type description for default values of each field.
	SecurityContext *IoK8sAPICoreV1PodSecurityContext `json:"securityContext,omitempty"`

	// ServiceAccountName is the name of the ServiceAccount to run all pods of the workflow as.
	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// Shutdown will shutdown the workflow according to its ShutdownStrategy
	Shutdown string `json:"shutdown,omitempty"`

	// Suspend will suspend the workflow and prevent execution of any future steps in the workflow
	Suspend bool `json:"suspend,omitempty"`

	// Synchronization holds synchronization lock configuration for this Workflow
	Synchronization *IoArgoprojWorkflowV1alpha1Synchronization `json:"synchronization,omitempty"`

	// Templates is a list of workflow templates used in a workflow
	Templates []*IoArgoprojWorkflowV1alpha1Template `json:"templates"`

	// Tolerations to apply to workflow pods.
	Tolerations []*IoK8sAPICoreV1Toleration `json:"tolerations"`

	// TTLSecondsAfterFinished limits the lifetime of a Workflow that has finished execution (Succeeded, Failed, Error). If this field is set, once the Workflow finishes, it will be deleted after ttlSecondsAfterFinished expires. If this field is unset, ttlSecondsAfterFinished will not expire. If this field is set to zero, ttlSecondsAfterFinished expires immediately after the Workflow finishes. DEPRECATED: Use TTLStrategy.SecondsAfterCompletion instead.
	TTLSecondsAfterFinished int32 `json:"ttlSecondsAfterFinished,omitempty"`

	// TTLStrategy limits the lifetime of a Workflow that has finished execution depending on if it Succeeded or Failed. If this struct is set, once the Workflow finishes, it will be deleted after the time to live expires. If this field is unset, the controller config map will hold the default values.
	TTLStrategy *IoArgoprojWorkflowV1alpha1TTLStrategy `json:"ttlStrategy,omitempty"`

	// VolumeClaimTemplates is a list of claims that containers are allowed to reference. The Workflow controller will create the claims at the beginning of the workflow and delete the claims upon completion of the workflow
	VolumeClaimTemplates []*IoK8sAPICoreV1PersistentVolumeClaim `json:"volumeClaimTemplates"`

	// Volumes is a list of volumes that can be mounted by containers in a io.argoproj.workflow.v1alpha1.
	Volumes []*IoK8sAPICoreV1Volume `json:"volumes"`

	// WorkflowMetadata contains some metadata of the workflow to be refer
	WorkflowMetadata *IoK8sApimachineryPkgApisMetaV1ObjectMeta `json:"workflowMetadata,omitempty"`

	// WorkflowTemplateRef holds a reference to a WorkflowTemplate for execution
	WorkflowTemplateRef *IoArgoprojWorkflowV1alpha1WorkflowTemplateRef `json:"workflowTemplateRef,omitempty"`
}

// Validate validates this io argoproj workflow v1alpha1 workflow template spec
func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAffinity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArguments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArtifactRepositoryRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecutor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostAliases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImagePullSecrets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodDisruptionBudget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodGC(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSynchronization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTolerations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTTLStrategy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeClaimTemplates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowTemplateRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) validateAffinity(formats strfmt.Registry) error {
	if swag.IsZero(m.Affinity) { // not required
		return nil
	}

	if m.Affinity != nil {
		if err := m.Affinity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("affinity")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) validateArguments(formats strfmt.Registry) error {
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

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) validateArtifactRepositoryRef(formats strfmt.Registry) error {
	if swag.IsZero(m.ArtifactRepositoryRef) { // not required
		return nil
	}

	if m.ArtifactRepositoryRef != nil {
		if err := m.ArtifactRepositoryRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifactRepositoryRef")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) validateDNSConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSConfig) { // not required
		return nil
	}

	if m.DNSConfig != nil {
		if err := m.DNSConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dnsConfig")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) validateExecutor(formats strfmt.Registry) error {
	if swag.IsZero(m.Executor) { // not required
		return nil
	}

	if m.Executor != nil {
		if err := m.Executor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("executor")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) validateHostAliases(formats strfmt.Registry) error {
	if swag.IsZero(m.HostAliases) { // not required
		return nil
	}

	for i := 0; i < len(m.HostAliases); i++ {
		if swag.IsZero(m.HostAliases[i]) { // not required
			continue
		}

		if m.HostAliases[i] != nil {
			if err := m.HostAliases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hostAliases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) validateImagePullSecrets(formats strfmt.Registry) error {
	if swag.IsZero(m.ImagePullSecrets) { // not required
		return nil
	}

	for i := 0; i < len(m.ImagePullSecrets); i++ {
		if swag.IsZero(m.ImagePullSecrets[i]) { // not required
			continue
		}

		if m.ImagePullSecrets[i] != nil {
			if err := m.ImagePullSecrets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("imagePullSecrets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) validateMetrics(formats strfmt.Registry) error {
	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	if m.Metrics != nil {
		if err := m.Metrics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) validatePodDisruptionBudget(formats strfmt.Registry) error {
	if swag.IsZero(m.PodDisruptionBudget) { // not required
		return nil
	}

	if m.PodDisruptionBudget != nil {
		if err := m.PodDisruptionBudget.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podDisruptionBudget")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) validatePodGC(formats strfmt.Registry) error {
	if swag.IsZero(m.PodGC) { // not required
		return nil
	}

	if m.PodGC != nil {
		if err := m.PodGC.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podGC")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) validateSecurityContext(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityContext) { // not required
		return nil
	}

	if m.SecurityContext != nil {
		if err := m.SecurityContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityContext")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) validateSynchronization(formats strfmt.Registry) error {
	if swag.IsZero(m.Synchronization) { // not required
		return nil
	}

	if m.Synchronization != nil {
		if err := m.Synchronization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("synchronization")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) validateTemplates(formats strfmt.Registry) error {
	if swag.IsZero(m.Templates) { // not required
		return nil
	}

	for i := 0; i < len(m.Templates); i++ {
		if swag.IsZero(m.Templates[i]) { // not required
			continue
		}

		if m.Templates[i] != nil {
			if err := m.Templates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("templates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) validateTolerations(formats strfmt.Registry) error {
	if swag.IsZero(m.Tolerations) { // not required
		return nil
	}

	for i := 0; i < len(m.Tolerations); i++ {
		if swag.IsZero(m.Tolerations[i]) { // not required
			continue
		}

		if m.Tolerations[i] != nil {
			if err := m.Tolerations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tolerations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) validateTTLStrategy(formats strfmt.Registry) error {
	if swag.IsZero(m.TTLStrategy) { // not required
		return nil
	}

	if m.TTLStrategy != nil {
		if err := m.TTLStrategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ttlStrategy")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) validateVolumeClaimTemplates(formats strfmt.Registry) error {
	if swag.IsZero(m.VolumeClaimTemplates) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumeClaimTemplates); i++ {
		if swag.IsZero(m.VolumeClaimTemplates[i]) { // not required
			continue
		}

		if m.VolumeClaimTemplates[i] != nil {
			if err := m.VolumeClaimTemplates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeClaimTemplates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) validateVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.Volumes) { // not required
		return nil
	}

	for i := 0; i < len(m.Volumes); i++ {
		if swag.IsZero(m.Volumes[i]) { // not required
			continue
		}

		if m.Volumes[i] != nil {
			if err := m.Volumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) validateWorkflowMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkflowMetadata) { // not required
		return nil
	}

	if m.WorkflowMetadata != nil {
		if err := m.WorkflowMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workflowMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) validateWorkflowTemplateRef(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkflowTemplateRef) { // not required
		return nil
	}

	if m.WorkflowTemplateRef != nil {
		if err := m.WorkflowTemplateRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workflowTemplateRef")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io argoproj workflow v1alpha1 workflow template spec based on the context it is used
func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAffinity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateArguments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateArtifactRepositoryRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDNSConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExecutor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHostAliases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImagePullSecrets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePodDisruptionBudget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePodGC(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSynchronization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTemplates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTolerations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTTLStrategy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumeClaimTemplates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkflowMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkflowTemplateRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) contextValidateAffinity(ctx context.Context, formats strfmt.Registry) error {

	if m.Affinity != nil {
		if err := m.Affinity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("affinity")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) contextValidateArguments(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) contextValidateArtifactRepositoryRef(ctx context.Context, formats strfmt.Registry) error {

	if m.ArtifactRepositoryRef != nil {
		if err := m.ArtifactRepositoryRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifactRepositoryRef")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) contextValidateDNSConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.DNSConfig != nil {
		if err := m.DNSConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dnsConfig")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) contextValidateExecutor(ctx context.Context, formats strfmt.Registry) error {

	if m.Executor != nil {
		if err := m.Executor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("executor")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) contextValidateHostAliases(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.HostAliases); i++ {

		if m.HostAliases[i] != nil {
			if err := m.HostAliases[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hostAliases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) contextValidateImagePullSecrets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ImagePullSecrets); i++ {

		if m.ImagePullSecrets[i] != nil {
			if err := m.ImagePullSecrets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("imagePullSecrets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) contextValidateMetrics(ctx context.Context, formats strfmt.Registry) error {

	if m.Metrics != nil {
		if err := m.Metrics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) contextValidatePodDisruptionBudget(ctx context.Context, formats strfmt.Registry) error {

	if m.PodDisruptionBudget != nil {
		if err := m.PodDisruptionBudget.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podDisruptionBudget")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) contextValidatePodGC(ctx context.Context, formats strfmt.Registry) error {

	if m.PodGC != nil {
		if err := m.PodGC.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podGC")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) contextValidateSecurityContext(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityContext != nil {
		if err := m.SecurityContext.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityContext")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) contextValidateSynchronization(ctx context.Context, formats strfmt.Registry) error {

	if m.Synchronization != nil {
		if err := m.Synchronization.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("synchronization")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) contextValidateTemplates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Templates); i++ {

		if m.Templates[i] != nil {
			if err := m.Templates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("templates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) contextValidateTolerations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tolerations); i++ {

		if m.Tolerations[i] != nil {
			if err := m.Tolerations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tolerations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) contextValidateTTLStrategy(ctx context.Context, formats strfmt.Registry) error {

	if m.TTLStrategy != nil {
		if err := m.TTLStrategy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ttlStrategy")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) contextValidateVolumeClaimTemplates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VolumeClaimTemplates); i++ {

		if m.VolumeClaimTemplates[i] != nil {
			if err := m.VolumeClaimTemplates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeClaimTemplates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) contextValidateVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Volumes); i++ {

		if m.Volumes[i] != nil {
			if err := m.Volumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) contextValidateWorkflowMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.WorkflowMetadata != nil {
		if err := m.WorkflowMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workflowMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) contextValidateWorkflowTemplateRef(ctx context.Context, formats strfmt.Registry) error {

	if m.WorkflowTemplateRef != nil {
		if err := m.WorkflowTemplateRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workflowTemplateRef")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) UnmarshalBinary(b []byte) error {
	var res IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
