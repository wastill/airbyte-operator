package v1alpha1

import corev1 "k8s.io/api/core/v1"

type TemporalSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RoleConfig *RoleGroupTemporalSpec `json:"roleConfig,omitempty"`

	// +kubebuilder:validation:Optional
	RoleGroups map[string]*RoleGroupTemporalSpec `json:"roleGroups,omitempty"`
}

type RoleGroupTemporalSpec struct {
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

	// +kubebuilder:validation:Optional
	Minio *MinioSpec `json:"minio,omitempty"`

	// +kubebuilder:validation:Optional
	Postgres *PostgresSpec `json:"postgres,omitempty"`
}
