package v1alpha1

import corev1 "k8s.io/api/core/v1"

type AirbyteApiServerSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RoleConfig *RoleGroupAirbyteApiServerSpec `json:"roleConfig,omitempty"`

	// +kubebuilder:validation:Optional
	RoleGroups map[string]*RoleGroupAirbyteApiServerSpec `json:"roleGroups,omitempty"`
}

func (apiServer *AirbyteApiServerSpec) InitRoleGroups() {
	for _, role := range apiServer.RoleGroups {
		role.InitRoleGroup(apiServer.RoleConfig)
	}
}

func (apiServerRole *RoleGroupAirbyteApiServerSpec) InitRoleGroup(roleConfig *RoleGroupAirbyteApiServerSpec) {
	outConfig := roleConfig
	inConfig := apiServerRole
	if inConfig.Replicas == 0 {
		inConfig.Replicas = AirbyteBootloaderReplicas
	}
	if inConfig.Image == nil {
		inConfig.Image = outConfig.Image
	}
	if inConfig.Annotations == nil {
		inConfig.Annotations = outConfig.Annotations
	}
	if inConfig.MatchLabels == nil {
		inConfig.MatchLabels = outConfig.MatchLabels
	}
	if inConfig.NodeSelector == nil {
		inConfig.NodeSelector = outConfig.NodeSelector
	}
	if inConfig.Tolerations == nil {
		inConfig.Tolerations = outConfig.Tolerations.DeepCopy()
	}
	if inConfig.Affinity == nil {
		inConfig.Affinity = outConfig.Affinity.DeepCopy()
	}
	if inConfig.Resources == nil {
		inConfig.Resources = outConfig.Resources.DeepCopy()
	}
	if inConfig.Service == nil {
		inConfig.Service = outConfig.Service.DeepCopy()
	}
	if inConfig.LogLevel == "" {
		inConfig.LogLevel = outConfig.LogLevel
	}
	if inConfig.Labels == nil {
		inConfig.Labels = outConfig.Labels
	}
	if inConfig.Secret == nil {
		inConfig.Secret = outConfig.Secret
	}
	if inConfig.SecurityContext == nil {
		inConfig.SecurityContext = outConfig.SecurityContext.DeepCopy()
	}
	if inConfig.EnvVars == nil {
		inConfig.EnvVars = outConfig.EnvVars
	}
	if inConfig.ExtraEnv == nil {
		inConfig.ExtraEnv = outConfig.ExtraEnv.DeepCopy()
	}
	if inConfig.Debug == nil {
		inConfig.Debug = outConfig.Debug.DeepCopy()
	}
}

func (apiServer *AirbyteApiServerSpec) InitRoleConfig(clusterConfig *ClusterConfigSpec) {
	outConfig := clusterConfig
	inConfig := apiServer.RoleConfig
	if inConfig.Replicas == 0 {
		inConfig.Replicas = AirbyteBootloaderReplicas
	}
	if inConfig.Annotations == nil {
		inConfig.Annotations = outConfig.Annotations
	}
	if inConfig.Labels == nil {
		inConfig.Labels = outConfig.Labels
	}
	if inConfig.MatchLabels == nil {
		inConfig.MatchLabels = outConfig.MatchLabels
	}
	if inConfig.NodeSelector == nil {
		inConfig.NodeSelector = outConfig.NodeSelector
	}
	if inConfig.Tolerations == nil {
		inConfig.Tolerations = outConfig.Tolerations.DeepCopy()
	}
	if inConfig.Affinity == nil {
		inConfig.Affinity = outConfig.Affinity.DeepCopy()
	}
	if inConfig.Resources == nil {
		inConfig.Resources = outConfig.Resources.DeepCopy()
	}
	if inConfig.Service == nil {
		inConfig.Service = outConfig.Service.DeepCopy()
	}
	if inConfig.LogLevel == "" {
		inConfig.LogLevel = outConfig.LogLevel
	}
	if inConfig.EnvVars == nil {
		inConfig.EnvVars = outConfig.EnvVars
	}
	if inConfig.ExtraEnv == nil {
		inConfig.ExtraEnv = outConfig.ExtraEnv.DeepCopy()
	}

}

type RoleGroupAirbyteApiServerSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=1
	Replicas int32 `json:"replicas,omitempty"`

	// +kubebuilder:validation:Optional
	Image *ImageSpec `json:"image,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

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
	Service *ServiceSpec `json:"service,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={}
	Secret map[string]string `json:"secret,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=INFO
	LogLevel string `json:"logLevel,omitempty"`

	// +kubebuilder:validation:Optional
	Debug *DebugSpec `json:"debug,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={}
	EnvVars map[string]string `json:"envVars,omitempty"`

	// +kubebuilder:validation:Optional
	ExtraEnv *corev1.EnvVar `json:"extraEnv,omitempty"`
}
