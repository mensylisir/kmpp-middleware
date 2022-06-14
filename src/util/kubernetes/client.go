package kubernetes

import (
	"fmt"
	mclient "github.com/minio/operator/pkg/client/clientset/versioned"
	"github.com/pkg/errors"
	pclient "github.com/zalando/postgres-operator/pkg/generated/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Config struct {
	ApiServer string
	Token     string
}

func NewKubernetesClient(c *Config) (*kubernetes.Clientset, error) {
	kubeConf := &rest.Config{
		Host:        c.ApiServer,
		BearerToken: c.Token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
	client, err := kubernetes.NewForConfig(kubeConf)
	if err != nil {
		return client, errors.Wrap(err, fmt.Sprintf("new kubernetes client with config failed: %v", err))
	}
	return client, nil
}

func NewPostgresqlClient(c *Config) (*pclient.Clientset, error) {
	kubeConf := &rest.Config{
		Host:        c.ApiServer,
		BearerToken: c.Token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
	client, err := pclient.NewForConfig(kubeConf)
	if err != nil {
		return client, errors.Wrap(err, fmt.Sprintf("new postgresql client with config failed: %v", err))
	}
	return client, nil
}

func NewMinioClient(c *Config) (*mclient.Clientset, error) {
	kubeConf := &rest.Config{
		Host:        c.ApiServer,
		BearerToken: c.Token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
	client, err := mclient.NewForConfig(kubeConf)
	if err != nil {
		return client, errors.Wrap(err, fmt.Sprintf("new postgresql client with config failed: %v", err))
	}
	return client, nil
}
