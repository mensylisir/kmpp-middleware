package model

import (
	"github.com/mensylisir/kmpp-middleware/src/model/common"
	uuid "github.com/satori/go.uuid"
)

type Instance struct {
	common.BaseModel
	ID            string    `json:"id" gorm:"not null;unique"`
	Name          string    `json:"name" gorm:"not null;unique"`
	Type          string    `json:"type" gorm:"type:varchar(64)"`
	ClusterID     string    `json:"cluster_id" gorm:"type:varchar(64)"`
	TemplateID    string    `json:"template_id" gorm:"type:varchar(64)"`
	Namespace     string    `json:"namespace" gorm:"type:varchar(64)"`
	Count         int32     `json:"count"  gorm:"type:int(64)"`
	RequestCpu    string    `json:"request_cpu" gorm:"type:varchar(64)"`
	RequestMemory string    `json:"request_memory" gorm:"type:varchar(64)"`
	LimitCpu      string    `json:"limit_cpu" gorm:"type:varchar(64)"`
	LimitMemory   string    `json:"limit_memory" gorm:"type:varchar(64)"`
	Volume        string    `json:"volume" gorm:"type:varchar(64)"`
	Status        string    `json:"status" gorm:"type:varchar(64)"`
	Cluster       Cluster   `gorm:"save_associations:false" json:"cluster"`
	Template      Templates `gorm:"save_associations:false" json:"template"`
}

func (c *Instance) BeforeCreate() error {
	c.ID = uuid.NewV4().String()
	return nil
}
