package kubernetes

import (
	"context"
	"github.com/mensylisir/kmpp-middleware/src/constant"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sync"
)

type GatherClusterInfoFunc func(cluster *entity.Cluster, client *kubernetes.Clientset, wg *sync.WaitGroup)

var funcList = []GatherClusterInfoFunc{
	GetServerVersion,
	GetKubernetesStatus,
}

func GetServerVersion(cluster *entity.Cluster, client *kubernetes.Clientset, wg *sync.WaitGroup) {
	defer wg.Done()
	v, err := client.ServerVersion()
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
	cluster.Version = v.GitVersion
}

func GetKubernetesStatus(cluster *entity.Cluster, client *kubernetes.Clientset, wg *sync.WaitGroup) {
	defer wg.Done()
	nodes, err := client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
	for _, node := range nodes.Items {
		if node.Status.Phase != corev1.NodeRunning {
			cluster.Status = constant.ClusterInnormal
		}
	}
	cluster.Status = constant.ClusterNormal
}

func GatherClusterInfo(cluster *entity.Cluster) error {
	c, err := NewKubernetesClient(&Config{
		ApiServer: cluster.ApiServer,
		Token:     cluster.Token,
	})
	if err != nil {
		return err
	}
	_, err = c.ServerVersion()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	for _, f := range funcList {
		wg.Add(1)
		go f(cluster, c, &wg)
	}
	wg.Wait()
	return nil
}
