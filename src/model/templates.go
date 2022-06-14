package model

import "github.com/mensylisir/kmpp-middleware/src/model/common"

type Templates struct {
	common.BaseModel
	ID              string `json:"id"`
	Name            string `json:"name" gorm:"not null;unique"`
	Icon            []byte `json:"icon"`
	BaseTemplate    string `json:"base_template"`
	AdvanceTemplate string `json:"advance_template"`
}
