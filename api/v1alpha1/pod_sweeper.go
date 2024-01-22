package v1alpha1

import corev1 "k8s.io/api/core/v1"

type PodSweeperSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RoleConfig *RoleGroupPodSweeperSpec `json:"roleConfig,omitempty"`

	// +kubebuilder:validation:Optional
	RoleGroups map[string]*RoleGroupPodSweeperSpec `json:"roleGroups,omitempty"`
}

type RoleGroupPodSweeperSpec struct {

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
	SecurityContext *corev1.PodSecurityContext `json:"securityContext,omitempty"`

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
	LivenessProbe *ProbeSpec `json:"livenessProbe,omitempty"`

	// +kubebuilder:validation:Optional
	ReadinessProbe *ProbeSpec `json:"readinessProbe,omitempty"`

	// +kubebuilder:validation:Optional
	ExtraVolumes []corev1.Volume `json:"extraVolumes,omitempty"`

	// +kubebuilder:validation:Optional
	ExtraVolumeMounts []corev1.VolumeMount `json:"extraVolumeMounts,omitempty"`

	// +kubebuilder:validation:Optional
	TimeToDeletePods *PodSweeperTimeToDeletePodsSpec `json:"timeToDeletePods,omitempty"`
}
