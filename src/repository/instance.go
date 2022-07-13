package repository

import (
	"github.com/mensylisir/kmpp-middleware/src/db"
	"github.com/mensylisir/kmpp-middleware/src/model"
)

type InstanceRepository interface {
	Get(name string) (model.Instance, error)
	GetByID(ID string) (model.Instance, error)
	Page(num, size int) (int, []model.Instance, error)
	List() ([]model.Instance, error)
	Save(instance *model.Instance) error
	Delete(name string) error
	GetByType(name string) ([]model.Instance, error)
	Update(ID string, values map[string]interface{}) error
}

func NewInstanceRepository() InstanceRepository {
	return &instanceRepository{}
}

type instanceRepository struct {
}

func (c instanceRepository) Get(name string) (model.Instance, error) {
	var instance model.Instance
	if err := db.DB.
		Where("name = ?", name).
		Preload("Cluster").
		Preload("Template").
		Find(&instance).Error; err != nil {
		return instance, err
	}
	return instance, nil
}

func (c instanceRepository) GetByMap(values map[string]interface{}) (model.Instance, error) {
	var instance model.Instance
	if err := db.DB.
		Where("? = ?", values).
		Preload("Cluster").
		Preload("Template").
		Find(&instance).Error; err != nil {
		return instance, err
	}
	return instance, nil
}

func (c instanceRepository) GetByID(ID string) (model.Instance, error) {
	var instance model.Instance
	if err := db.DB.
		Where("id = ?", ID).
		Preload("Cluster").
		Preload("Template").
		Find(&instance).Error; err != nil {
		return instance, err
	}
	return instance, nil
}

func (c instanceRepository) GetByType(instanceType string) ([]model.Instance, error) {
	var instances []model.Instance
	if err := db.DB.
		Where("type = ?", instanceType).
		Find(&instances).Error; err != nil {
		return instances, err
	}
	return instances, nil
}

func (c instanceRepository) List() ([]model.Instance, error) {
	var instances []model.Instance
	if err := db.DB.
		Find(&instances).Error; err != nil {
		return instances, err
	}
	return instances, nil
}

func (c instanceRepository) Page(num, size int) (int, []model.Instance, error) {
	var (
		total     int
		instances []model.Instance
	)

	if err := db.DB.Model(&model.Instance{}).
		Count(&total).
		Offset((num - 1) * size).
		Limit(size).
		Find(&instances).Error; err != nil {
		return total, instances, err
	}

	return total, instances, nil
}

func (c instanceRepository) Save(instance *model.Instance) error {
	if db.DB.NewRecord(instance) {
		if err := db.DB.Create(instance).Error; err != nil {
			return err
		}
	} else {
		if err := db.DB.Save(instance).Error; err != nil {
			return err
		}
	}
	return nil
}

func (c instanceRepository) Delete(name string) error {
	err := db.DB.Where("name = ?", name).Delete(&model.Instance{}).Error
	return err
}

func (c instanceRepository) Update(ID string, values map[string]interface{}) error {
	err := db.DB.Model(model.Instance{}).Where("id = ?", ID).Updates(values).Error
	if err != nil {
		return err
	}
	return nil
}
