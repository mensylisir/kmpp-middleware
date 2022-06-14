package entity

import "github.com/mensylisir/kmpp-middleware/src/model"

type Instance struct {
	model.Instance
	UserId      string       `json:"user_id"`
	ServiceInfo ServiceInfo  `json:"service_info"`
	Secret      []SecretInfo `json:"secret"`
	Yaml        string       `json:"yaml"`
}

type InstancePage struct {
	Items []Instance `json:"items"`
	Total int        `json:"total"`
}

type Postgres struct {
	Name      string  `json:"name"`
	Namespace string  `json:"namespace"`
	ClusterId string  `json:"cluster_id"`
	Type      string  `json:"type"`
	Cluster   Cluster `json:"cluster"`
	UserId    string  `json:"user_id"`
}
