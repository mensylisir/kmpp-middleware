package entity

type StatefulsetStatus struct {
	AvailableReplicas int32                  `json:"available_replicas"`
	CurrentReplicas   int32                  `json:"current_replicas"`
	ReadyReplicas     int32                  `json:"ready_replicas"`
	Replicas          int32                  `json:"replicas"`
	Conditions        []StatefulsetCondition `json:"conditions"`
}

type StatefulsetCondition struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
}
