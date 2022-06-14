package cluster

import (
	"encoding/json"
	"fmt"
	"github.com/mensylisir/kmpp-middleware/src/constant"
	"github.com/mensylisir/kmpp-middleware/src/db"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/model"
	"github.com/mensylisir/kmpp-middleware/src/repository"
	"github.com/mensylisir/kmpp-middleware/src/util/kubernetes"
	"github.com/sirupsen/logrus"
	"strings"
)

type ClusterService interface {
	Get(name string) (entity.Cluster, error)
	GetByID(ID string) (entity.Cluster, error)
	Page(num, size int, userId string) (*entity.ClusterPage, error)
	Import(clusterImport entity.Cluster) error
	Sync(name string) (entity.Cluster, error)
	Delete(userId, name string) error
}

type clusterService struct {
	clusterRepo     repository.ClusterRepository
	userClusterRepo repository.UserClusterRepository
	userRepo        repository.UserRepository
}

func NewClusterService() ClusterService {
	return &clusterService{
		clusterRepo:     repository.NewClusterRepository(),
		userClusterRepo: repository.NewUserClusterRepository(),
		userRepo:        repository.NewUserRepository(),
	}
}

func (c clusterService) Get(name string) (entity.Cluster, error) {
	var clusterDTO entity.Cluster
	mo, err := c.clusterRepo.Get(name)
	if err != nil {
		return clusterDTO, err
	}
	clusterDTO.Cluster = mo
	return clusterDTO, nil
}

func (c clusterService) GetByID(ID string) (entity.Cluster, error) {
	var clusterDTO entity.Cluster
	mo, err := c.clusterRepo.GetByID(ID)
	if err != nil {
		return clusterDTO, err
	}
	clusterDTO.Cluster = mo
	return clusterDTO, nil
}

func (c clusterService) Page(num, size int, userID string) (*entity.ClusterPage, error) {
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

func (c clusterService) Import(clusterImport entity.Cluster) error {
	loginfo, _ := json.Marshal(clusterImport)
	logger.Log.WithFields(logrus.Fields{"cluster_import_info": string(loginfo)}).Debugf("start to import the cluster %s", clusterImport.Name)
	if strings.HasSuffix(clusterImport.ApiServer, "/") {
		clusterImport.ApiServer = strings.Replace(clusterImport.ApiServer, "/", "", -1)
	}
	clusterImport.ApiServer = strings.Replace(clusterImport.ApiServer, "http://", "", -1)
	clusterImport.ApiServer = strings.Replace(clusterImport.ApiServer, "https://", "", -1)
	tx := db.DB.Begin()
	cluster := model.Cluster{
		Name:      clusterImport.Name,
		ApiServer: clusterImport.ApiServer,
		Token:     clusterImport.Token,
		Type:      clusterImport.Type,
	}
	var entityCluster entity.Cluster
	entityCluster.Cluster = cluster
	if err := kubernetes.GatherClusterInfo(&entityCluster); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(&entityCluster.Cluster).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("can not import cluster %s", err.Error())
	}
	user, err := c.userRepo.GetByID(clusterImport.UserId)
	if err != nil {
		return err
	}
	if user.Role != 0 {
		userCluster := model.UserCluster{
			UserID:    user.ID,
			ClusterID: entityCluster.Cluster.ID,
		}
		if err := tx.Create(&userCluster).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("can not import cluster%s", err.Error())
		}
	}
	tx.Commit()
	return nil
}

func (c clusterService) Sync(name string) (entity.Cluster, error) {
	cluster, err := c.clusterRepo.Get(name)
	if err != nil {
		logger.Log.Errorf("instance of %s not found error: %s", name, err.Error())
		return entity.Cluster{Cluster: cluster}, err
	}
	var entityCluster entity.Cluster
	entityCluster.Cluster = cluster
	tx := db.DB.Begin()
	if err := kubernetes.GatherClusterInfo(&entityCluster); err != nil {
		tx.Rollback()
		return entityCluster, err
	}
	if err := tx.Model(model.Cluster{}).Where("id = ?", cluster.ID).Updates(map[string]string{
		"status": entityCluster.Status,
	}); err != nil {
		tx.Rollback()
		return entityCluster, err.Error
	}
	tx.Commit()
	return entityCluster, nil
}

func (c clusterService) Delete(userId, name string) error {
	user, err := c.userRepo.GetByID(userId)
	if err != nil {
		return err
	}
	if user.Role == 0 {
		return c.clusterRepo.Delete(name)
	}
	cluster, err := c.clusterRepo.Get(name)
	if err != nil {
		return err
	}
	err = c.clusterRepo.Delete(name)
	if err != nil {
		return err
	}
	return c.userClusterRepo.Delete(userId, cluster.ID)
}
func (c clusterService) getPageForUser(num, size int, userID string) (*entity.ClusterPage, error) {
	var (
		page       entity.ClusterPage
		clusters   []model.Cluster
		clusterIds []string
	)
	clusters, err := c.clusterRepo.GetByType(constant.CenterCluster)
	if err != nil {
		return nil, err
	}

	for _, cluster := range clusters {
		clusterIds = append(clusterIds, cluster.ID)
	}

	userCluster, err := c.userClusterRepo.Get(userID)
	if err != nil {
		return nil, err
	}

	for _, cluster := range userCluster {
		clusterIds = append(clusterIds, cluster.ClusterID)
	}

	if err := db.DB.Model(&model.Cluster{}).
		Where("id in (?)", clusterIds).
		Count(&page.Total).
		Offset((num - 1) * size).
		Limit(size).
		Order("created_at ASC").
		Find(&clusters).Error; err != nil {
		return nil, err
	}

	for _, mo := range clusters {
		clusterDTO := entity.Cluster{
			Cluster: mo,
		}
		page.Items = append(page.Items, clusterDTO)
	}
	return &page, nil
}

func (c clusterService) getPageForAdmin(num, size int) (*entity.ClusterPage, error) {
	var (
		page     entity.ClusterPage
		clusters []model.Cluster
	)

	if err := db.DB.Model(&model.Cluster{}).
		Count(&page.Total).
		Offset((num - 1) * size).
		Limit(size).
		Order("created_at ASC").
		Find(&clusters).Error; err != nil {
		return nil, err
	}

	for _, mo := range clusters {
		clusterDTO := entity.Cluster{
			Cluster: mo,
		}
		page.Items = append(page.Items, clusterDTO)
	}
	return &page, nil
}
