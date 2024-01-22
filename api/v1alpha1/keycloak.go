package v1alpha1

import corev1 "k8s.io/api/core/v1"

type KeycloakSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:default=false
	Enabled bool `json:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RoleConfig *RoleConfigKeycloakSpec `json:"roleConfig,omitempty"`

	// +kubebuilder:validation:Optional
	RoleGroups map[string]*RoleGroupKeycloakSpec `json:"roleGroups,omitempty"`
}

type RoleConfigKeycloakSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=INFO
	LogLevel string `json:"logLevel,omitempty"`
}

type RoleGroupKeycloakSpec struct {

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
	Service *ServiceSpec `json:"service,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=INFO
	LogLevel string `json:"logLevel,omitempty"`
}
