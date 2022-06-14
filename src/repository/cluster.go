package repository

import (
	"github.com/mensylisir/kmpp-middleware/src/db"
	"github.com/mensylisir/kmpp-middleware/src/model"
)

type ClusterRepository interface {
	Get(name string) (model.Cluster, error)
	List() ([]model.Cluster, error)
	Save(cluster *model.Cluster) error
	Delete(name string) error
	Page(num, size int) (int, []model.Cluster, error)
	GetByID(ID string) (model.Cluster, error)
	GetByType(name string) ([]model.Cluster, error)
	Update(ID string, values map[string]interface{}) error
}

func NewClusterRepository() ClusterRepository {
	return &clusterRepository{}
}

type clusterRepository struct {
}

func (c clusterRepository) Get(name string) (model.Cluster, error) {
	var cluster model.Cluster
	if err := db.DB.
		Where("name = ?", name).
		Find(&cluster).Error; err != nil {
		return cluster, err
	}
	return cluster, nil
}

func (c clusterRepository) GetByID(ID string) (model.Cluster, error) {
	var cluster model.Cluster
	if err := db.DB.
		Where("ID = ?", ID).
		Find(&cluster).Error; err != nil {
		return cluster, err
	}
	return cluster, nil
}

func (c clusterRepository) GetByType(clusterType string) ([]model.Cluster, error) {
	var clusters []model.Cluster
	if err := db.DB.
		Where("type = ?", clusterType).
		Find(&clusters).Error; err != nil {
		return clusters, err
	}
	return clusters, nil
}

func (c clusterRepository) List() ([]model.Cluster, error) {
	var clusters []model.Cluster
	if err := db.DB.
		Find(&clusters).Error; err != nil {
		return clusters, err
	}
	return clusters, nil
}

func (c clusterRepository) Page(num, size int) (int, []model.Cluster, error) {
	var (
		total    int
		clusters []model.Cluster
	)

	if err := db.DB.Model(&model.Cluster{}).
		Count(&total).
		Offset((num - 1) * size).
		Limit(size).
		Find(&clusters).Error; err != nil {
		return total, clusters, err
	}

	return total, clusters, nil
}

func (c clusterRepository) Save(cluster *model.Cluster) error {
	if db.DB.NewRecord(cluster) {
		if err := db.DB.Create(cluster).Error; err != nil {
			return err
		}
	} else {
		if err := db.DB.Save(cluster).Error; err != nil {
			return err
		}
	}
	return nil
}

func (c clusterRepository) Delete(name string) error {
	err := db.DB.Where("name = ?", name).Delete(&model.Cluster{}).Error
	return err
}

func (c clusterRepository) Update(ID string, values map[string]interface{}) error {
	err := db.DB.Model(model.Cluster{}).Where("id = ?", ID).Updates(values)
	if err != nil {
		return err.Error
	}
	return nil
}
