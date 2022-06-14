package model

import (
	"github.com/mensylisir/kmpp-middleware/src/model/common"
	uuid "github.com/satori/go.uuid"
)

type Cluster struct {
	common.BaseModel
	ID        string `json:"id"`
	Name      string `json:"name" gorm:"not null;unique"`
	ApiServer string `json:"api_server" gorm:"type:varchar(64)"`
	Version   string `json:"version" gorm:"type:varchar(64)"`
	Token     string `json:"token"  gorm:"type:varchar(64)"`
	Type      string `json:"type"  gorm:"type:varchar(64)"`
	Status    string `json:"Status"  gorm:"type:varchar(64)"`
}

func (c *Cluster) BeforeCreate() error {
	c.ID = uuid.NewV4().String()
	return nil
}
