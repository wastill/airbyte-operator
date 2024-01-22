package v1alpha1

import (
	"github.com/zncdata-labs/operator-go/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"reflect"
)

type ClusterConfigSpec struct {

	// +kubebuilder:validation:Required
	SecurityContext *corev1.PodSecurityContext `json:"securityContext,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="airbyte-admin"
	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="oss"
	DeploymentMode string `json:"deploymentMode,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={}
	EnvVars map[string]string `json:"envVars,omitempty"`

	// +kubebuilder:validation:Optional
	Database *DatabaseClusterConfigSpec `json:"database,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="community"
	Edition string `json:"edition,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="MINIO"
	StateStorageType string `json:"stateStorageType,omitempty"`

	// +kubebuilder:validation:Optional
	Logs *LogsClusterConfigSpec `json:"logs,omitempty"`

	// +kubebuilder:validation:Optional
	Metrics *MetricsClusterConfigSpec `json:"metrics,omitempty"`

	// +kubebuilder:validation:Optional
	Jobs *JobsClusterConfigSpec `json:"jobs,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

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

	// +kubebuilder:validation:Optional
	Postgres *PostgresSpec `json:"postgres,omitempty"`

	// +kubebuilder:validation:Optional
	Minio *MinioSpec `json:"minio,omitempty"`

	// +kubebuilder:validation:Optional
	ExtraEnv *corev1.EnvVar `json:"extraEnv,omitempty"`
}

func (c *ClusterConfigSpec) GetConfigValue(key string) (any, error) {
	rv := reflect.ValueOf(c)
	if v := rv.FieldByName(key); v.IsValid() {
		return v.Interface(), nil
	}
	return nil, errors.New("field not found.")
}
