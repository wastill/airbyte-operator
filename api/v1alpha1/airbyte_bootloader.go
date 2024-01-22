package v1alpha1

import (
	"github.com/zncdata-labs/operator-go/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"reflect"
)

const (
	AirbyteBootloaderRepo     = "airbyte/airbyte-bootloader"
	AirbyteBootloaderTag      = "0.50.30"
	AirbyteBootloaderPort     = 80
	AirbyteBootloaderReplicas = 1
)

type AirbyteBootloaderSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RoleConfig *RoleAirbyteBootloaderSpec `json:"roleConfig,omitempty"`

	// +kubebuilder:validation:Optional
	RoleGroups map[string]*RoleAirbyteBootloaderSpec `json:"roleGroups,omitempty"`
}

func (bootloader *AirbyteBootloaderSpec) InitRoleGroups() {
	for _, role := range bootloader.RoleGroups {
		role.InitRoleGroup(bootloader.RoleConfig)
	}
}

func (bootloaderRole *RoleAirbyteBootloaderSpec) InitRoleGroup(roleConfig *RoleAirbyteBootloaderSpec) {
	outConfig := roleConfig
	inConfig := bootloaderRole
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
	if inConfig.RunDatabaseMigrationsOnStartup == nil {
		inConfig.RunDatabaseMigrationsOnStartup = outConfig.RunDatabaseMigrationsOnStartup
	}

}

func (bootloader *AirbyteBootloaderSpec) InitRoleConfig(clusterConfig *ClusterConfigSpec) {
	outConfig := clusterConfig
	inConfig := bootloader.RoleConfig
	if inConfig.Replicas == 0 {
		inConfig.Replicas = AirbyteBootloaderReplicas
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
	if inConfig.RunDatabaseMigrationsOnStartup == nil {
		inConfig.RunDatabaseMigrationsOnStartup = outConfig.RunDatabaseMigrationsOnStartup
	}

}

type RoleAirbyteBootloaderSpec struct {
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

	// +kubebuilder:validation:Optional
	RunDatabaseMigrationsOnStartup *bool `json:"runDatabaseMigrationsOnStartup,omitempty"`
}

func (r *RoleAirbyteBootloaderSpec) GetConfigValue(key string) (any, error) {
	rv := reflect.ValueOf(r)
	if v := rv.FieldByName(key); v.IsValid() {
		return v.Interface(), nil
	}
	return nil, errors.New("field not found.")
}
