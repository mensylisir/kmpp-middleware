package cron

import (
	"fmt"
	"github.com/mensylisir/kmpp-middleware/src/cron/job"
	"github.com/robfig/cron/v3"
)

var Cron *cron.Cron

const phaseName = "cron"

type InitCronPhase struct {
	Enable bool
}

func (c *InitCronPhase) Init() error {
	Cron = cron.New()
	if c.Enable {
		_, err := Cron.AddJob("@every 30s", job.NewWatchPostgresqlInfo())
		if err != nil {
			return fmt.Errorf("can not add corn job: %s", err.Error())
		}
		_, err = Cron.AddJob("@every 30s", job.NewWatchClusterInfo())
		if err != nil {
			return fmt.Errorf("can not add corn job: %s", err.Error())
		}
		Cron.Start()
	}
	return nil
}

func (c *InitCronPhase) PhaseName() string {
	return phaseName
}
