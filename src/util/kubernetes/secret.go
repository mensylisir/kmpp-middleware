package kubernetes

import (
	"context"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetSecretInfo(instance *entity.Instance) ([]entity.SecretInfo, error) {
	client, err := NewKubernetesClient(&Config{
		ApiServer: instance.Cluster.ApiServer,
		Token:     instance.Cluster.Token,
	})
	if err != nil {
		return nil, err
	}
	sts, err := client.CoreV1().Secrets(instance.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var infos []entity.SecretInfo
	for _, st := range sts.Items {
		if st.Type != corev1.SecretTypeOpaque {
			continue
		}
		info := entity.SecretInfo{}
		info.Name = st.Name
		data := make(map[string]string)
		for key, value := range st.Data {
			data[key] = string(value)
		}
		for key, value := range st.StringData {
			data[key] = value
		}
		info.Data = data
		infos = append(infos, info)
	}
	return infos, nil
}
