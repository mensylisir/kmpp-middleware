package job

import (
	"github.com/mensylisir/kmpp-middleware/src/constant"
	"github.com/mensylisir/kmpp-middleware/src/db"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/model"
	"github.com/mensylisir/kmpp-middleware/src/service/postgresql"
	"sync"
)

type WatchPostgresqlInfo struct {
	PostgresService postgresql.PostgresService
}

func NewWatchPostgresqlInfo() *WatchPostgresqlInfo {
	return &WatchPostgresqlInfo{
		PostgresService: postgresql.NewPostgresService(),
	}
}

func (w *WatchPostgresqlInfo) Run() {
	var instances []model.Instance
	var wg sync.WaitGroup
	sem := make(chan struct{}, 2)
	db.DB.Where("type = ?", constant.POSTGRESQL).Find(&instances)
	for _, instance := range instances {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			_, err := w.PostgresService.Sync(name)
			if err != nil {
				logger.Log.Errorf("gather postgres info error: %s", err.Error())
			}
		}(instance.Name)
	}
	wg.Wait()
}
