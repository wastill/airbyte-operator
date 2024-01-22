package v1alpha1

import corev1 "k8s.io/api/core/v1"

type WorkerSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RoleConfig *RoleGroupWorkerSpec `json:"roleConfig,omitempty"`

	// +kubebuilder:validation:Optional
	RoleGroups map[string]*RoleGroupWorkerSpec `json:"roleGroups,omitempty"`
}

type RoleGroupWorkerSpec struct {

	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=1
	Replicas int32 `json:"replicas,omitempty"`

	// +kubebuilder:validation:Optional
	Image *ImageSpec `json:"image,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityContext *corev1.PodSecurityContext `json:"securityContext,omitempty"`

	// +kubebuilder:validation:Required
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

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
	// +kubebuilder:default=INFO
	LogLevel string `json:"logLevel,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	Debug bool `json:"debug,omitempty"`

	// +kubebuilder:validation:Optional
	ContainerOrchestrator *ContainerOrchestratorSpec `json:"containerOrchestrator,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	ActivityMaxAttempt string `json:"activityMaxAttempt,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	ActivityInitialDelayBetweenAttemptsSeconds string `json:"activityInitialDelayBetweenAttemptsSeconds,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	ActivityMaxDelayBetweenAttemptsSeconds string `json:"activityMaxDelayBetweenAttemptsSeconds,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="5"
	MaxNotifyWorkers string `json:"maxNotifyWorkers,omitempty"`
}

type ContainerOrchestratorSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	Image string `json:"image,omitempty"`
}
