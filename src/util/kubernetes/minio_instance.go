package kubernetes

import (
	"context"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	v1 "github.com/minio/operator/pkg/apis/minio.min.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateMinio(instance *entity.Instance) (*v1.Tenant, error) {
	client, err := NewMinioClient(&Config{
		ApiServer: instance.Cluster.ApiServer,
		Token:     instance.Cluster.Token,
	})
	if err != nil {
		return nil, err
	}
	minioInstance := &v1.Tenant{
		TypeMeta: metav1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
	}
	tenant, err := client.MinioV1().Tenants(instance.Namespace).Create(context.TODO(), minioInstance, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return tenant, nil
}
