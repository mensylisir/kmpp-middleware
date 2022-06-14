package entity

type PodStatus struct {
	Name       string         `json:"name"`
	Phase      string         `json:"phase"`
	Conditions []PodCondition `json:"conditions"`
}

type PodCondition struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
}
