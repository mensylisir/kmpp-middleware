package entity

import "github.com/mensylisir/kmpp-middleware/src/model"

type Cluster struct {
	model.Cluster
	UserId string `json:"user_id"`
}

type ClusterPage struct {
	Items []Cluster `json:"items"`
	Total int       `json:"total"`
}
