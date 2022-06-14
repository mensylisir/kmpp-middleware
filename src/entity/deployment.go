package entity

type DeploymentStatus struct {
	Replicas            int32                 `json:"replicas"`
	ReadyReplicas       int32                 `json:"ready_replicas"`
	AvailableReplicas   int32                 `json:"available_replicas"`
	UnavailableReplicas int32                 `json:"unavailable_replicas"`
	Conditions          []DeploymentCondition `json:"conditions"`
}

type DeploymentCondition struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
}
