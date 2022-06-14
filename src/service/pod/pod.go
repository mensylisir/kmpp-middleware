package pod

import (
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/service/cluster"
	"github.com/mensylisir/kmpp-middleware/src/util/kubernetes"
)

type PodService interface {
	GetPods(instance entity.Instance) ([]string, error)
	GetPogLog(instance entity.Instance) (string, error)
	GetPodStatus(instance entity.Instance) (*entity.PodStatus, error)
	GetPodsStatus(instance entity.Instance) ([]entity.PodStatus, error)
}

type podService struct {
	clusterService cluster.ClusterService
}

func NewPodService() PodService {
	return &podService{
		clusterService: cluster.NewClusterService(),
	}
}

func (c podService) GetPods(instance entity.Instance) ([]string, error) {
	clusterObj, err := c.clusterService.GetByID(instance.ClusterID)
	if err != nil {
		return nil, err
	}
	instance.Cluster = clusterObj.Cluster
	return kubernetes.GetPods(&instance)
}

func (c podService) GetPogLog(instance entity.Instance) (string, error) {
	clusterObj, err := c.clusterService.GetByID(instance.ClusterID)
	if err != nil {
		return "", err
	}
	instance.Cluster = clusterObj.Cluster
	return kubernetes.GetLogs(&instance)
}

func (c podService) GetPodStatus(instance entity.Instance) (*entity.PodStatus, error) {
	clusterObj, err := c.clusterService.GetByID(instance.ClusterID)
	if err != nil {
		return nil, err
	}
	instance.Cluster = clusterObj.Cluster
	return kubernetes.GetPodStatus(&instance)
}

func (c podService) GetPodsStatus(instance entity.Instance) ([]entity.PodStatus, error) {
	clusterObj, err := c.clusterService.GetByID(instance.ClusterID)
	if err != nil {
		return nil, err
	}
	instance.Cluster = clusterObj.Cluster
	return kubernetes.GetPodsStatus(&instance)
}