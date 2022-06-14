package entity

type ServiceInfo struct {
	ServiceType string        `json:"service_type"`
	Addresses   []ServiceAddr `json:"addresses"`
}

type ServiceAddr struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}
