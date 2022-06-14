package entity

type SecretInfo struct {
	Name string            `json:"name"`
	Data map[string]string `json:"data"`
}
