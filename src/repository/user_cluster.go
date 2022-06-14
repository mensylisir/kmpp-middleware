package repository

import (
	"github.com/mensylisir/kmpp-middleware/src/db"
	"github.com/mensylisir/kmpp-middleware/src/model"
)

type UserClusterRepository interface {
	Get(userID string) ([]model.UserCluster, error)
	Save(item *model.UserCluster) error
	Delete(userId, clusterId string) error
	Update(ID string, values map[string]interface{}) error
}

type userClusterRepository struct {
}

func NewUserClusterRepository() UserClusterRepository {
	return &userClusterRepository{}
}

func (u userClusterRepository) Get(userID string) ([]model.UserCluster, error) {
	var userClusters []model.UserCluster
	if err := db.DB.Where("user_id = ?", userID).First(&userClusters).Error; err != nil {
		return userClusters, err
	}
	return userClusters, nil
}

func (u userClusterRepository) Save(item *model.UserCluster) error {
	if db.DB.NewRecord(item) {
		return db.DB.Create(&item).Error
	} else {
		return db.DB.Save(&item).Error
	}
}

func (u userClusterRepository) Delete(userId, clusterId string) error {
	userCluster := model.UserCluster{
		UserID:    userId,
		ClusterID: clusterId,
	}
	return db.DB.Delete(&userCluster).Error
}

func (u userClusterRepository) Update(ID string, values map[string]interface{}) error {
	err := db.DB.Model(model.UserCluster{}).Where("id = ?", ID).Updates(values)
	if err != nil {
		return err.Error
	}
	return nil
}
