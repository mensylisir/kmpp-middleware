package job

import (
	"github.com/mensylisir/kmpp-middleware/src/db"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/model"
	"github.com/mensylisir/kmpp-middleware/src/service/cluster"
	"sync"
)

type WatchClusterInfo struct {
	ClusterService cluster.ClusterService
}

func NewWatchClusterInfo() *WatchClusterInfo {
	return &WatchClusterInfo{
		ClusterService: cluster.NewClusterService(),
	}
}

func (w *WatchClusterInfo) Run() {
	var clusters []model.Cluster
	var wg sync.WaitGroup
	sem := make(chan struct{}, 2)
	db.DB.Find(&clusters)
	for _, clusterObj := range clusters {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			_, err := w.ClusterService.Sync(name)
			if err != nil {
				logger.Log.Errorf("gather cluster info error: %s", err.Error())
			}
		}(clusterObj.Name)
	}
	wg.Wait()
}
