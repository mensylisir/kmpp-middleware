package kubernetes

import (
	"context"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetStatefulSetStatus(instance *entity.Instance) (*entity.StatefulsetStatus, error) {
	client, err := NewKubernetesClient(&Config{
		ApiServer: instance.Cluster.ApiServer,
		Token:     instance.Cluster.Token,
	})
	if err != nil {
		return nil, err
	}
	sts, err := client.AppsV1().StatefulSets(instance.Namespace).Get(context.TODO(), instance.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	var stsStatus entity.StatefulsetStatus
	stsStatus.AvailableReplicas = sts.Status.AvailableReplicas
	stsStatus.CurrentReplicas = sts.Status.CurrentReplicas
	stsStatus.ReadyReplicas = sts.Status.ReadyReplicas
	stsStatus.Replicas = sts.Status.Replicas
	var conditions []entity.StatefulsetCondition
	for _, cdt := range sts.Status.Conditions {
		condition := entity.StatefulsetCondition{}
		condition.Status = string(cdt.Status)
		condition.Message = cdt.Message
		condition.Reason = cdt.Reason
		conditions = append(conditions, condition)
	}
	return &stsStatus, nil
}
