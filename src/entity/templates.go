package entity

import "github.com/mensylisir/kmpp-middleware/src/model"

type Templates struct {
	model.Templates
}

type OperateTemplates struct {
	Operation string            `json:"operation"`
	Items     []model.Templates `json:"items"`
}
