package svc

import (
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/service/cluster"
	"github.com/mensylisir/kmpp-middleware/src/util/kubernetes"
)

type SvcService interface {
	UpdateServiceType(instance entity.Instance) error
}

type svcService struct {
	clusterService cluster.ClusterService
}

func NewSvcService() SvcService {
	return &svcService{
		clusterService: cluster.NewClusterService(),
	}
}

func (c svcService) UpdateServiceType(instance entity.Instance) error {
	clusterObj, err := c.clusterService.GetByID(instance.ClusterID)
	if err != nil {
		return err
	}
	instance.Cluster = clusterObj.Cluster
	return kubernetes.EditServiceType(&instance)
}
