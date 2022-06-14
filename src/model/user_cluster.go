package model

type UserCluster struct {
	UserID    string `json:"user_id" gorm:"type:varchar(64)"`
	ClusterID string `json:"cluster_id" gorm:"type:varchar(64)"`
}
