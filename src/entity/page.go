package entity

type Page struct {
	Items interface{} `json:"items"`
	Total int         `json:"total"`
}
