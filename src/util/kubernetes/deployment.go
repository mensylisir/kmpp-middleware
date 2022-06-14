package kubernetes

import (
	"context"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetDeploymentStatus(instance *entity.Instance) (*entity.DeploymentStatus, error) {
	client, err := NewKubernetesClient(&Config{
		ApiServer: instance.Cluster.ApiServer,
		Token:     instance.Cluster.Token,
	})
	if err != nil {
		return nil, err
	}
	deploy, err := client.AppsV1().Deployments(instance.Namespace).Get(context.TODO(), instance.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	var deployStatus entity.DeploymentStatus
	deployStatus.Replicas = deploy.Status.Replicas
	deployStatus.ReadyReplicas = deploy.Status.ReadyReplicas
	deployStatus.AvailableReplicas = deploy.Status.AvailableReplicas
	deployStatus.UnavailableReplicas = deploy.Status.UpdatedReplicas
	var conditions []entity.DeploymentCondition
	for _, condit := range deploy.Status.Conditions {
		condition := entity.DeploymentCondition{}
		condition.Status = string(condit.Status)
		condition.Message = condit.Message
		condition.Reason = condit.Reason
		conditions = append(conditions, condition)
	}
	deployStatus.Conditions = conditions
	return &deployStatus, nil
}
