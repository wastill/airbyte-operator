package v1alpha1

import corev1 "k8s.io/api/core/v1"

type CronSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RoleConfig *RoleGroupCronSpec `json:"roleConfig,omitempty"`

	// +kubebuilder:validation:Optional
	RoleGroups map[string]*RoleGroupCronSpec `json:"roleGroups,omitempty"`
}

type RoleGroupCronSpec struct {

	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=1
	Replicas int32 `json:"replicas,omitempty"`

	// +kubebuilder:validation:Optional
	Image *ImageSpec `json:"image,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// +kubebuilder:validation:Optional
	MatchLabels map[string]string `json:"matchLabels,omitempty"`

	// +kubebuilder:validation:Optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// +kubebuilder:validation:Optional
	Tolerations *corev1.Toleration `json:"tolerations,omitempty"`

	// +kubebuilder:validation:Optional
	Affinity *corev1.Affinity `json:"affinity,omitempty"`

	// +kubebuilder:validation:Optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={}
	Secret map[string]string `json:"secret,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=INFO
	LogLevel string `json:"logLevel,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={}
	EnvVars map[string]string `json:"envVars,omitempty"`

	// +kubebuilder:validation:Optional
	ExtraEnv corev1.EnvVar `json:"extraEnv,omitempty"`
}

type ImageSpec struct {
	// +kubebuilder:validation:Optional
	Repository string `json:"repository,omitempty"`

	// +kubebuilder:validation:Optional
	Tag string `json:"tag,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=IfNotPresent
	PullPolicy corev1.PullPolicy `json:"pullPolicy,omitempty"`
}

func (airbyte *Airbyte) GetNameWithSuffix(name string) string {
	return airbyte.GetName() + "" + name
}
