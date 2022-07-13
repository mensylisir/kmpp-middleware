package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/mensylisir/kmpp-middleware/src/db"
	"github.com/mensylisir/kmpp-middleware/src/model"
)

type UserInstanceRepository interface {
	Get(userID string) ([]model.UserInstance, error)
	Save(item *model.UserInstance) error
	Delete(UserID string) error
	DeleteByInstanceId(instanceId string) error
	Update(ID string, values map[string]interface{}) error
}

type userInstanceRepository struct {
}

func NewUserInstanceRepository() UserInstanceRepository {
	return &userInstanceRepository{}
}

func (u userInstanceRepository) Get(userID string) ([]model.UserInstance, error) {
	var userInstances []model.UserInstance
	if err := db.DB.Where("user_id = ?", userID).Find(&userInstances).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, nil
		}
		return userInstances, err
	}
	return userInstances, nil
}

func (u userInstanceRepository) Save(item *model.UserInstance) error {
	if db.DB.NewRecord(item) {
		return db.DB.Create(&item).Error
	} else {
		return db.DB.Save(&item).Error
	}
}

func (u userInstanceRepository) Delete(UserID string) error {
	userInstance := model.UserInstance{
		UserID: UserID,
	}
	return db.DB.Delete(&userInstance).Error
}

func (u userInstanceRepository) DeleteByInstanceId(instanceId string) error {
	userInstance := model.UserInstance{
		InstanceID: instanceId,
	}
	return db.DB.Delete(&userInstance).Error
}

func (u userInstanceRepository) Update(ID string, values map[string]interface{}) error {
	err := db.DB.Model(model.UserInstance{}).Where("id = ?", ID).Updates(values)
	if err != nil {
		return err.Error
	}
	return nil
}
