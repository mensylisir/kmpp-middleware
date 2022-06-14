package repository

import (
	"github.com/mensylisir/kmpp-middleware/src/constant"
	"github.com/mensylisir/kmpp-middleware/src/db"
	"github.com/mensylisir/kmpp-middleware/src/model"
)

type TemplatesRepository interface {
	Save(item *model.Templates) error
	Get(name string) (model.Templates, error)
	GetById(id string) (model.Templates, error)
	Page(num, size int) (int, []model.Templates, error)
	List() ([]model.Templates, error)
	Delete(name string) error
	Batch(operation string, items []model.Templates) error
	Update(ID string, values map[string]interface{}) error
}

type templatesRepository struct {
}

func NewTemplatesRepository() TemplatesRepository {
	return &templatesRepository{}
}

func (t templatesRepository) Save(item *model.Templates) error {
	if db.DB.NewRecord(item) {
		return db.DB.Create(&item).Error
	} else {
		return db.DB.Save(&item).Error
	}
}

func (t templatesRepository) Get(name string) (model.Templates, error) {
	var template model.Templates
	if err := db.DB.Where("name = ?", name).First(&template).Error; err != nil {
		return template, err
	}
	return template, nil
}

func (t templatesRepository) GetById(id string) (model.Templates, error) {
	var template model.Templates
	if err := db.DB.Where("id = ?", id).First(&template).Error; err != nil {
		return template, err
	}
	return template, nil
}

func (t templatesRepository) Page(num, size int) (int, []model.Templates, error) {
	var total int
	var templates []model.Templates
	err := db.DB.Model(&model.Templates{}).Count(&total).Order("name").Offset((num - 1) * size).Limit(size).Find(&templates).Error
	return total, templates, err
}

func (t templatesRepository) List() ([]model.Templates, error) {
	var templates []model.Templates
	err := db.DB.Order("name").Find(&templates).Error
	return templates, err
}

func (t templatesRepository) Delete(name string) error {
	template, err := t.Get(name)
	if err != nil {
		return err
	}
	return db.DB.Delete(&template).Error
}

func (t templatesRepository) Batch(operation string, items []model.Templates) error {
	switch operation {
	case constant.BatchOperationDelete:
		tx := db.DB.Begin()
		for _, item := range items {
			err := tx.Delete(&item).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()
	default:
		return constant.NotSupportedBatchOperation
	}
	return nil
}

func (t templatesRepository) Update(ID string, values map[string]interface{}) error {
	err := db.DB.Model(model.Templates{}).Where("id = ?", ID).Updates(values)
	if err != nil {
		return err.Error
	}
	return nil
}
