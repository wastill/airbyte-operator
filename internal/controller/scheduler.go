package controller

import (
	stackv1alpha1 "github.com/zncdata-labs/airbyte-operator/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func ServerScheduler(instance *stackv1alpha1.Airbyte, dep *appsv1.Deployment) {

	scheduler := instance.Spec.Server

	if scheduler.NodeSelector != nil {
		dep.Spec.Template.Spec.NodeSelector = scheduler.NodeSelector
	}

	if scheduler.Tolerations != nil {
		toleration := *scheduler.Tolerations
		dep.Spec.Template.Spec.Tolerations = []corev1.Toleration{
			*toleration.DeepCopy(),
		}
	}
	if scheduler.Affinity != nil {
		dep.Spec.Template.Spec.Affinity = scheduler.Affinity.DeepCopy()
	}
}

func WorkerScheduler(instance *stackv1alpha1.Airbyte, dep *appsv1.Deployment) {
	scheduler := instance.Spec.Worker
	if scheduler.NodeSelector != nil {
		dep.Spec.Template.Spec.NodeSelector = scheduler.NodeSelector
	}
	if scheduler.Tolerations != nil {
		dep.Spec.Template.Spec.Tolerations = []corev1.Toleration{
			*scheduler.Tolerations.DeepCopy(),
		}
	}
	if scheduler.Affinity != nil {
		dep.Spec.Template.Spec.Affinity = scheduler.Affinity.DeepCopy()
	}
}

func AirbyteApiServerScheduler(instance *stackv1alpha1.Airbyte, dep *appsv1.Deployment) {

	scheduler := instance.Spec.AirbyteApiServer
	if scheduler.NodeSelector != nil {
		dep.Spec.Template.Spec.NodeSelector = scheduler.NodeSelector
	}

	if scheduler.Tolerations != nil {
		dep.Spec.Template.Spec.Tolerations = []corev1.Toleration{
			*scheduler.Tolerations.DeepCopy(),
		}
	}

	if scheduler.Affinity != nil {
		dep.Spec.Template.Spec.Affinity = scheduler.Affinity.DeepCopy()
	}
}

func ConnectorBuilderServerScheduler(instance *stackv1alpha1.Airbyte, dep *appsv1.Deployment) {

	scheduler := instance.Spec.ConnectorBuilderServer
	if scheduler.NodeSelector != nil {
		dep.Spec.Template.Spec.NodeSelector = scheduler.NodeSelector
	}

	if scheduler.Tolerations != nil {
		dep.Spec.Template.Spec.Tolerations = []corev1.Toleration{
			*scheduler.Tolerations.DeepCopy(),
		}
	}

	if scheduler.Affinity != nil {
		dep.Spec.Template.Spec.Affinity = scheduler.Affinity.DeepCopy()
	}

}

func CronScheduler(instance *stackv1alpha1.Airbyte, dep *appsv1.Deployment) {

	scheduler := instance.Spec.Cron
	if scheduler.NodeSelector != nil {
		dep.Spec.Template.Spec.NodeSelector = scheduler.NodeSelector
	}

	if scheduler.Tolerations != nil {
		dep.Spec.Template.Spec.Tolerations = []corev1.Toleration{
			*scheduler.Tolerations.DeepCopy(),
		}
	}

	if scheduler.Affinity != nil {
		dep.Spec.Template.Spec.Affinity = scheduler.Affinity.DeepCopy()
	}

}
