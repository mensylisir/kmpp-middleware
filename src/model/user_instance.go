package model

type UserInstance struct {
	UserID     string `json:"user_id" gorm:"type:varchar(64)"`
	InstanceID string `json:"instance_id" gorm:"type:varchar(64)"`
}
