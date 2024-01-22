/*
Copyright 2023 zncdata-labs.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"github.com/zncdata-labs/operator-go/pkg/status"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	apimeta "k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

var (
	PullPolicy                 = "IfNotPresent"
	ServerRepo                 = "airbyte/server"
	ServerTag                  = "0.50.30"
	ServerPort                 = 8001
	WorkerRepo                 = "airbyte/worker"
	WorkerTag                  = "0.50.30"
	AirbyteApiServerRepo       = "airbyte/airbyte-api-server"
	AirbyteApiServerTag        = "0.50.30"
	AirbyteApiServerPort       = 8006
	WebAppRepo                 = "airbyte/webapp"
	WebAppTag                  = "0.50.30"
	WebAppPort                 = 80
	PodSweeperRepo             = "bitnami/kubectl"
	PodSweeperTag              = "latest"
	ConnectorBuilderServerRepo = "airbyte/connector-builder-server"
	ConnectorBuilderServerTag  = "0.50.30"
	ConnectorBuilderServerPort = 80
	TemporalRepo               = "temporalio/auto-setup"
	TemporalTag                = "1.20.1"
	TemporalPort               = 7233
	KeycloakRepo               = "airbyte/keycloak"
	KeycloakTag                = "0.50.30"
	KeycloakPort               = 8180
	CronRepo                   = "airbyte/cron"
	CronTag                    = "0.50.30"
)

// AirbyteSpec defines the desired state of Airbyte
type AirbyteSpec struct {
	// +kubebuilder:validation:Required
	SecurityContext *corev1.PodSecurityContext `json:"securityContext,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterConfig *ClusterConfigSpec `json:"clusterConfig,omitempty"`

	// +kubebuilder:validation:Required
	Server *ServerSpec `json:"server,omitempty"`

	// +kubebuilder:validation:Required
	Worker *WorkerSpec `json:"worker,omitempty"`

	// +kubebuilder:validation:Required
	AirbyteApiServer *AirbyteApiServerSpec `json:"airbyteApiServer,omitempty"`

	// +kubebuilder:validation:Required
	WebApp *WebAppSpec `json:"webApp,omitempty"`

	// +kubebuilder:validation:Required
	PodSweeper *PodSweeperSpec `json:"podSweeper,omitempty"`

	// +kubebuilder:validation:Required
	ConnectorBuilderServer *ConnectorBuilderServerSpec `json:"connectorBuilderServer,omitempty"`

	// +kubebuilder:validation:Required
	AirbyteBootloader *AirbyteBootloaderSpec `json:"airbyteBootloader,omitempty"`

	// +kubebuilder:validation:Required
	Temporal *TemporalSpec `json:"temporal,omitempty"`

	// +kubebuilder:validation:Required
	Keycloak *KeycloakSpec `json:"keycloak,omitempty"`

	// +kubebuilder:validation:Required
	Cron *CronSpec `json:"cron,omitempty"`
}

type MinioSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default="minio"
	RootUser string `json:"rootUser,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="minio123"
	RootPassword string `json:"rootPassword,omitempty"`
}

type DatabaseClusterConfigSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	SecretName string `json:"secretName,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	SecretValue string `json:"secretValue,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="example.com"
	Host string `json:"host,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=5432
	Port int32 `json:"port,omitempty"`
}

type JobsClusterConfigSpec struct {
	// +kubebuilder:validation:Optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	// +kubebuilder:validation:Optional
	Kube *JobsKubeClusterConfigSpec `json:"kube,omitempty"`
}

type JobsKubeClusterConfigSpec struct {
	// +kubebuilder:validation:Optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// +kubebuilder:validation:Optional
	Tolerations *corev1.Toleration `json:"tolerations,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	MainContainerImagePullSecret string `json:"mainContainerImagePullSecret,omitempty"`

	// +kubebuilder:validation:Optional
	Images *JobsKubeImagesClusterConfigSpec `json:"images,omitempty"`
}

type JobsKubeImagesClusterConfigSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	Busybox string `json:"busybox,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	Socat string `json:"socat,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	Curl string `json:"curl,omitempty"`
}

type LogsClusterConfigSpec struct {
	// +kubebuilder:validation:Optional
	AccessKey *LogsAccessKeyClusterConfigSpec `json:"accessKey,omitempty"`

	// +kubebuilder:validation:Optional
	SecretKey *LogsSecretKeyClusterConfigSpec `json:"secretKey,omitempty"`

	// +kubebuilder:validation:Optional
	Minio *LogMinioClusterConfigSpec `json:"minio,omitempty"`

	// +kubebuilder:validation:Optional
	ExternalMinio *LogExternalMinioClusterConfigSpec `json:"externalMinio,omitempty"`

	// +kubebuilder:validation:Optional
	S3 *LogS3ClusterConfigSpec `json:"s3,omitempty"`

	// +kubebuilder:validation:Optional
	Gcs *LogGcsClusterConfigSpec `json:"gcs,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="MINIO"
	StorageType string `json:"storageType,omitempty"`
}

type LogsAccessKeyClusterConfigSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	Password string `json:"password,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	ExistingSecret string `json:"existingSecret,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	ExistingSecretKey string `json:"existingSecretKey,omitempty"`
}

type LogsSecretKeyClusterConfigSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	Password string `json:"password,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	ExistingSecret string `json:"existingSecret,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	ExistingSecretKey string `json:"existingSecretKey,omitempty"`
}

type MetricsClusterConfigSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	MetricClient string `json:"metricClient,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	OtelCollectorEndpoint string `json:"otelCollectorEndpoint,omitempty"`
}

type LogMinioClusterConfigSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`
}

type LogExternalMinioClusterConfigSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	Enabled bool `json:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	Endpoint string `json:"endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="localhost"
	Host string `json:"host,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=9000
	Port int32 `json:"port,omitempty"`
}

type LogS3ClusterConfigSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:default=false
	Enabled bool `json:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="airbyte-dev-logs"
	Bucket string `json:"bucket,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	BucketRegion string `json:"bucketRegion,omitempty"`
}

type LogGcsClusterConfigSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	Bucket string `json:"bucket,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	Credentials string `json:"credentials,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	CredentialsJson string `json:"credentialsJson,omitempty"`
}

type PostgresSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default="postgresql"
	Host string `json:"host,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default="5432"
	Port string `json:"port,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default="airbyte"
	UserName string `json:"username,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default="airbyte"
	Password string `json:"password,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default="db-airbyte"
	DataBase string `json:"database,omitempty"`
}

type ServiceSpec struct {
	// +kubebuilder:validation:Optional
	Annotations map[string]string `json:"annotations,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:enum=ClusterIP;NodePort;LoadBalancer;ExternalName
	// +kubebuilder:default=ClusterIP
	Type corev1.ServiceType `json:"type,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default=8001
	Port int32 `json:"port,omitempty"`
}

type ProbeSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=5
	InitialDelaySeconds int32 `json:"initialDelaySeconds,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=30
	PeriodSeconds int32 `json:"periodSeconds,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=1
	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=3
	FailureThreshold int32 `json:"failureThreshold,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=1
	SuccessThreshold int32 `json:"successThreshold,omitempty"`
}

type PodSweeperTimeToDeletePodsSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	Running string `json:"running,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=120
	Succeeded int32 `json:"succeeded,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=1440
	Unsuccessful int32 `json:"unsuccessful,omitempty"`
}

type DebugSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	Enabled bool `json:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=5005
	RemoteDebugPort int32 `json:"RemoteDebugPort,omitempty"`
}

type WebAppIngressSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:=true
	Enabled bool `json:"enabled,omitempty"`
	// +kubebuilder:validation:Optional
	TLS *networkingv1.IngressTLS `json:"tls,omitempty"`
	// +kubebuilder:validation:Optional
	Annotations map[string]string `json:"annotations,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:="webapp.example.com"
	Host string `json:"host,omitempty"`
}

// SetStatusCondition updates the status condition using the provided arguments.
// If the condition already exists, it updates the condition; otherwise, it appends the condition.
// If the condition status has changed, it updates the condition's LastTransitionTime.
func (airbyte *Airbyte) SetStatusCondition(condition metav1.Condition) {
	// if the condition already exists, update it
	existingCondition := apimeta.FindStatusCondition(airbyte.Status.Conditions, condition.Type)
	if existingCondition == nil {
		condition.ObservedGeneration = airbyte.GetGeneration()
		condition.LastTransitionTime = metav1.Now()
		airbyte.Status.Conditions = append(airbyte.Status.Conditions, condition)
	} else if existingCondition.Status != condition.Status || existingCondition.Reason != condition.Reason || existingCondition.Message != condition.Message {
		existingCondition.Status = condition.Status
		existingCondition.Reason = condition.Reason
		existingCondition.Message = condition.Message
		existingCondition.ObservedGeneration = airbyte.GetGeneration()
		existingCondition.LastTransitionTime = metav1.Now()
	}
}

// InitStatusConditions initializes the status conditions to the provided conditions.
func (airbyte *Airbyte) InitStatusConditions() {
	airbyte.Status.InitStatus(airbyte)
	airbyte.Status.InitStatusConditions()
}

// AirbyteStatus defines the observed state of Airbyte
type AirbyteStatus struct {
	// +kubebuilder:validation:Optional
	Conditions []metav1.Condition `json:"condition,omitempty"`
	// +kubebuilder:validation:Optional
	URLs []StatusURL `json:"urls,omitempty"`
}

type StatusURL struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

//+kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="Image",type="string",JSONPath=".spec.image.repository",description="The Docker Image of Airbyte"
// +kubebuilder:printcolumn:name="Tag",type="string",JSONPath=".spec.image.tag",description="The Docker Tag of Airbyte"
//+kubebuilder:subresource:status

// Airbyte is the Schema for the airbytes API
type Airbyte struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AirbyteSpec   `json:"spec,omitempty"`
	Status status.Status `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AirbyteList contains a list of Airbyte
type AirbyteList struct {
	metav1.TypeMeta `json:",inline,omitempty"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Airbyte `json:"items,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Airbyte{}, &AirbyteList{})
}
