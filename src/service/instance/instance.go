package instance

import (
	"github.com/mensylisir/kmpp-middleware/src/db"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/model"
	"github.com/mensylisir/kmpp-middleware/src/repository"
)

type InstanceService interface {
	Get(name string) (entity.Instance, error)
	GetById(ID string) (entity.Instance, error)
	Page(num, size int, userId string) (*entity.InstancePage, error)
}

type instanceService struct {
	instanceRepo     repository.InstanceRepository
	userInstanceRepo repository.UserInstanceRepository
	userRepo         repository.UserRepository
}

func NewInstanceService() InstanceService {
	return &instanceService{
		instanceRepo:     repository.NewInstanceRepository(),
		userInstanceRepo: repository.NewUserInstanceRepository(),
		userRepo:         repository.NewUserRepository(),
	}
}

func (c instanceService) Get(name string) (entity.Instance, error) {
	var instanceDTO entity.Instance
	mo, err := c.instanceRepo.Get(name)
	if err != nil {
		return instanceDTO, err
	}
	instanceDTO.Instance = mo
	return instanceDTO, nil
}

func (c instanceService) GetById(ID string) (entity.Instance, error) {
	var instanceDTO entity.Instance
	mo, err := c.instanceRepo.GetByID(ID)
	if err != nil {
		return instanceDTO, err
	}
	instanceDTO.Instance = mo
	return instanceDTO, nil
}

func (c instanceService) Page(num, size int, userID string) (*entity.InstancePage, error) {
	user, err := c.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	if user.Role == 0 {
		return c.getPageForAdmin(num, size)
	} else {
		return c.getPageForUser(num, size, userID)
	}
}

func (c instanceService) getPageForUser(num, size int, userID string) (*entity.InstancePage, error) {
	var (
		page        entity.InstancePage
		instances   []model.Instance
		instanceIds []string
	)
	userInstance, err := c.userInstanceRepo.Get(userID)
	if err != nil {
		return nil, err
	}

	for _, instance := range userInstance {
		instanceIds = append(instanceIds, instance.InstanceID)
	}

	if err := db.DB.Model(&model.Instance{}).
		Where("id in (?)", instanceIds).
		Preload("Cluster").
		Preload("Template").
		Count(&page.Total).
		Offset((num - 1) * size).
		Limit(size).
		Order("created_at ASC").
		Find(&instances).Error; err != nil {
		return nil, err
	}

	for _, mo := range instances {
		instanceDTO := entity.Instance{
			Instance: mo,
		}
		page.Items = append(page.Items, instanceDTO)
	}
	return &page, nil
}

func (c instanceService) getPageForAdmin(num, size int) (*entity.InstancePage, error) {
	var (
		page      entity.InstancePage
		instances []model.Instance
	)

	if err := db.DB.Model(&model.Instance{}).
		Preload("Cluster").
		Preload("Template").
		Count(&page.Total).
		Offset((num - 1) * size).
		Limit(size).
		Order("created_at ASC").
		Find(&instances).Error; err != nil {
		return nil, err
	}

	for _, mo := range instances {
		instanceDTO := entity.Instance{
			Instance: mo,
		}
		page.Items = append(page.Items, instanceDTO)
	}
	return &page, nil
}
