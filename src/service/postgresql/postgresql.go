package postgresql

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/mensylisir/kmpp-middleware/src/constant"
	"github.com/mensylisir/kmpp-middleware/src/db"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/model"
	"github.com/mensylisir/kmpp-middleware/src/repository"
	"github.com/mensylisir/kmpp-middleware/src/service/cluster"
	"github.com/mensylisir/kmpp-middleware/src/service/templates"
	"github.com/mensylisir/kmpp-middleware/src/util/kubernetes"
	"github.com/sirupsen/logrus"
	v1 "github.com/zalando/postgres-operator/pkg/apis/acid.zalan.do/v1"
)

type PostgresService interface {
	Get(name string) (entity.Instance, error)
	GetById(ID string) (entity.Instance, error)
	GetStatusById(ID string) (string, error)
	Page(num, size int, userId string) (*entity.InstancePage, error)
	Save(instance entity.Instance) error
	Create(postgres entity.Postgres) (string, error)
	Sync(name string) (entity.Instance, error)
	Delete(name string) error
	Update(instance entity.Instance) error
	Edit(instance entity.Instance) error
}

type postgresService struct {
	instanceRepo     repository.InstanceRepository
	userInstanceRepo repository.UserInstanceRepository
	userRepo         repository.UserRepository
	clusterService   cluster.ClusterService
	templateService  templates.TemplatesService
}

func NewPostgresService() PostgresService {
	return &postgresService{
		instanceRepo:     repository.NewInstanceRepository(),
		userInstanceRepo: repository.NewUserInstanceRepository(),
		userRepo:         repository.NewUserRepository(),
		clusterService:   cluster.NewClusterService(),
		templateService:  templates.NewTemplatesService(),
	}
}

func (c postgresService) Get(name string) (entity.Instance, error) {
	var instanceDTO entity.Instance
	mo, err := c.instanceRepo.Get(name)
	if err != nil {
		return instanceDTO, err
	}
	instanceDTO.Instance = mo

	clusterModel, err := c.clusterService.GetByID(mo.ClusterID)
	if err != nil {
		return instanceDTO, err
	}
	instanceDTO.Instance.Cluster = clusterModel.Cluster

	template, err := c.templateService.Get(constant.POSTGRESQL)
	if err != nil {
		return instanceDTO, err
	}
	instanceDTO.Instance.Template = template.Templates

	if err := kubernetes.GatherPostgresInfo(&instanceDTO); err != nil {
		return instanceDTO, err
	}
	if err := kubernetes.GatherPostgresStatus(&instanceDTO); err != nil {
		return instanceDTO, err
	}
	return instanceDTO, nil
}

func (c postgresService) GetById(ID string) (entity.Instance, error) {
	var instanceDTO entity.Instance
	mo, err := c.instanceRepo.GetByID(ID)
	if err != nil {
		return instanceDTO, err
	}
	instanceDTO.Instance = mo
	clusterModel, err := c.clusterService.GetByID(mo.ClusterID)
	if err != nil {
		return instanceDTO, err
	}
	instanceDTO.Instance.Cluster = clusterModel.Cluster

	template, err := c.templateService.Get(constant.POSTGRESQL)
	if err != nil {
		return instanceDTO, err
	}
	instanceDTO.Instance.Template = template.Templates

	if err := kubernetes.GatherPostgresInfo(&instanceDTO); err != nil {
		return instanceDTO, err
	}
	if err := kubernetes.GatherPostgresStatus(&instanceDTO); err != nil {
		return instanceDTO, err
	}
	return instanceDTO, nil
}

func (c postgresService) GetStatusById(ID string) (string, error) {
	mo, err := c.instanceRepo.GetByID(ID)
	if err != nil {
		return "", err
	}
	return mo.Status, nil
}

func (c postgresService) Page(num, size int, userID string) (*entity.InstancePage, error) {
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

func (c postgresService) Save(instance entity.Instance) error {
	loginfo, _ := json.Marshal(instance)
	logger.Log.WithFields(logrus.Fields{"instance_info": string(loginfo)}).Debugf("start to add the instance %s", instance.Name)
	tx := db.DB.Begin()
	modelInstance := model.Instance{
		Name:          instance.Name,
		Type:          constant.POSTGRESQL,
		ClusterID:     instance.ClusterID,
		Namespace:     instance.Namespace,
		Count:         instance.Count,
		RequestCpu:    instance.RequestCpu,
		RequestMemory: instance.RequestMemory,
		LimitCpu:      instance.LimitCpu,
		LimitMemory:   instance.LimitMemory,
		Volume:        instance.Volume,
	}

	clusterObj, err := c.clusterService.GetByID(instance.ClusterID)
	if err != nil {
		return err
	}
	modelInstance.Cluster = clusterObj.Cluster

	template1, err := c.templateService.Get(constant.POSTGRESQL)
	if err != nil {
		return err
	}
	modelInstance.TemplateID = template1.Templates.ID
	modelInstance.Template = template1.Templates

	var inst entity.Instance
	inst.Instance = modelInstance

	res, err := kubernetes.CreatePostgres(&inst)
	if err != nil {
		return err
	}

	inst.Instance.Status = res.Status.PostgresClusterStatus
	if err := tx.Create(&inst.Instance).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("can not create postgres %s", err.Error())
	}

	user, err := c.userRepo.GetByID(instance.UserId)
	if err != nil {
		return err
	}
	userInstance := model.UserInstance{
		UserID:     user.ID,
		InstanceID: inst.Instance.ID,
	}
	if err := tx.Create(&userInstance).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("can not create postgres%s", err.Error())
	}
	tx.Commit()
	return nil
}

func (c postgresService) Create(postgres entity.Postgres) (string, error) {
	template, err := c.templateService.Get(constant.POSTGRESQL)
	if err != nil {
		return "", err
	}
	clusterObj, err := c.clusterService.GetByID(postgres.ClusterId)
	if err != nil {
		return "", err
	}
	postgres.Cluster = clusterObj
	var postgresql *v1.Postgresql
	if postgres.Type == constant.BASIC {
		err = yaml.Unmarshal([]byte(template.BaseTemplate), &postgresql)
		if err != nil {
			return "", err
		}
	} else if postgres.Type == constant.ADVANCE {
		err = yaml.Unmarshal([]byte(template.AdvanceTemplate), &postgresql)
		if err != nil {
			return "", err
		}
	}
	pInst, err := kubernetes.CreatePostgresFromTemplate(&postgres, postgresql)
	if err != nil {
		return "", err
	}
	tx := db.DB.Begin()
	modelInstance := model.Instance{
		Name:          postgres.Name,
		Type:          constant.POSTGRESQL,
		ClusterID:     postgres.ClusterId,
		Namespace:     postgres.Namespace,
		Count:         pInst.Spec.NumberOfInstances,
		RequestCpu:    pInst.Spec.Resources.ResourceRequests.CPU,
		RequestMemory: pInst.Spec.Resources.ResourceRequests.Memory,
		LimitCpu:      pInst.Spec.Resources.ResourceLimits.CPU,
		LimitMemory:   pInst.Spec.Resources.ResourceLimits.Memory,
		Volume:        pInst.Spec.Volume.Size,
		Status:        pInst.Status.PostgresClusterStatus,
	}
	modelInstance.TemplateID = template.Templates.ID
	modelInstance.Template = template.Templates

	if err := tx.Create(&modelInstance).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("can not create postgres %s", err.Error())
	}

	user, err := c.userRepo.GetByID(postgres.UserId)
	if err != nil {
		return "", err
	}
	userInstance := model.UserInstance{
		UserID:     user.ID,
		InstanceID: modelInstance.ID,
	}
	if err := tx.Create(&userInstance).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("can not create postgres%s", err.Error())
	}
	tx.Commit()
	return modelInstance.ID, nil
}

func (c postgresService) Sync(name string) (entity.Instance, error) {
	instance, err := c.instanceRepo.Get(name)
	if err != nil {
		logger.Log.Errorf("instance of %s not found error: %s", name, err.Error())
		return entity.Instance{Instance: instance}, err
	}
	var inst entity.Instance
	inst.Instance = instance

	clusterModel, err := c.clusterService.GetByID(instance.ClusterID)
	if err != nil {
		logger.Log.Errorf("instance of %s not found error: %s", name, err.Error())
		return entity.Instance{Instance: instance}, err
	}
	inst.Instance.Cluster = clusterModel.Cluster
	tx := db.DB.Begin()

	if err := kubernetes.GatherPostgresStatus(&inst); err != nil {
		tx.Rollback()
		return inst, err
	}
	err = c.instanceRepo.Update(instance.ID, map[string]interface{}{
		"status": inst.Status,
	})
	if err != nil {
		return inst, err
	}
	return inst, nil
}

func (c postgresService) Delete(name string) error {
	return c.instanceRepo.Delete(name)
}

func (c postgresService) Update(instance entity.Instance) error {
	clusterObj, err := c.clusterService.GetByID(instance.ClusterID)
	if err != nil {
		return err
	}
	instance.Cluster = clusterObj.Cluster
	_, err = kubernetes.UpdatePostgres(&instance)
	if err != nil {
		return err
	}
	jsonInstance, err := json.Marshal(instance)
	if err != nil {
		return err
	}
	mapInstance := make(map[string]interface{})
	err = json.Unmarshal(jsonInstance, &mapInstance)
	if err != nil {
		return err
	}
	return c.instanceRepo.Update(instance.ID, mapInstance)
}

func (c postgresService) Edit(instance entity.Instance) error {
	clusterObj, err := c.clusterService.GetByID(instance.ClusterID)
	if err != nil {
		return err
	}
	instance.Cluster = clusterObj.Cluster
	_, err = kubernetes.EditPostgres(&instance)
	if err != nil {
		return err
	}
	jsonInstance, err := json.Marshal(instance)
	if err != nil {
		return err
	}
	mapInstance := make(map[string]interface{})
	err = json.Unmarshal(jsonInstance, &mapInstance)
	if err != nil {
		return err
	}
	return c.instanceRepo.Update(instance.ID, mapInstance)
}

func (c postgresService) getPageForUser(num, size int, userID string) (*entity.InstancePage, error) {
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
		Where("id in (?) and type = ?", instanceIds, constant.POSTGRESQL).
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

func (c postgresService) getPageForAdmin(num, size int) (*entity.InstancePage, error) {
	var (
		page      entity.InstancePage
		instances []model.Instance
	)
	if err := db.DB.Model(&model.Instance{}).
		Where("type = ?", constant.POSTGRESQL).
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
