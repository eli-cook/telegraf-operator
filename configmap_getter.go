package main

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ConfigMapGetter interface {
	Get(namespace string, name string) (*corev1.ConfigMap, error)
}

// a struct that contains a kubeclient and a method to get a configmap from a namespace
type configMapGetterImpl struct {
	kubeClient kubernetes.Interface
}

func NewConfigMapGetter(kubeClient kubernetes.Interface) ConfigMapGetter {
	return &configMapGetterImpl{
		kubeClient: kubeClient,
	}
}

// get a configmap from a namespace
func (c *configMapGetterImpl) Get(namespace string, name string) (*corev1.ConfigMap, error) {
	return c.kubeClient.CoreV1().ConfigMaps(namespace).Get(context.Background(), name, metav1.GetOptions{})
}
