package kubernetes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mensylisir/kmpp-middleware/src/db"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"strings"
)

func GetServiceInfo(instance *entity.Instance) (*entity.ServiceInfo, error) {
	client, err := NewKubernetesClient(&Config{
		ApiServer: instance.Cluster.ApiServer,
		Token:     instance.Cluster.Token,
	})
	if err != nil {
		return nil, err
	}
	svc, err := client.CoreV1().Services(instance.Namespace).Get(context.TODO(), instance.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	var serviceInfo entity.ServiceInfo
	serviceInfo.ServiceType = string(svc.Spec.Type)
	var addresses []entity.ServiceAddr
	switch serviceInfo.ServiceType {
	case "ClusterIP":
		for _, pt := range svc.Spec.Ports {
			address := entity.ServiceAddr{}
			address.Host = svc.Spec.ClusterIP
			address.Port = pt.Port
			addresses = append(addresses, address)
		}
	case "NodePort":
		cluster, err := getCluster(instance)
		if err != nil {
			return nil, err
		}
		for _, pt := range svc.Spec.Ports {
			address := entity.ServiceAddr{}
			address.Host = strings.Split(cluster.ApiServer, ":")[0]
			address.Port = pt.NodePort
			addresses = append(addresses, address)
		}
	}
	serviceInfo.Addresses = addresses
	return &serviceInfo, nil
}

func EditServiceType(instance *entity.Instance) error {
	client, err := NewKubernetesClient(&Config{
		ApiServer: instance.Cluster.ApiServer,
		Token:     instance.Cluster.Token,
	})
	if err != nil {
		return err
	}

	patches := fmt.Sprintf("{\"spec\": {\"type\": \"%s\"}}", instance.ServiceType)
	payloadBYtes, err := json.Marshal(patches)
	if err != nil {
		return err
	}
	_, err = client.CoreV1().Services(instance.Namespace).Patch(context.TODO(), instance.Name, types.StrategicMergePatchType, payloadBYtes, metav1.PatchOptions{})
	return err
}

func getCluster(instance *entity.Instance) (*model.Cluster, error) {
	var cluster model.Cluster
	if err := db.DB.Where("id = ?", instance.ClusterID).First(&cluster).Error; err != nil {
		return nil, err
	}
	return &cluster, nil
}
