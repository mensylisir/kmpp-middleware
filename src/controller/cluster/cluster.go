package cluster

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/service/cluster"
	"github.com/toolkits/pkg/ginx"
	"strconv"
)

type ClusterController struct {
	Ctx            context.Context
	ClusterService cluster.ClusterService
}

func NewClusterController() *ClusterController {
	return &ClusterController{
		ClusterService: cluster.NewClusterService(),
	}
}

var clusterController ClusterController

func init() {
	clusterController = *NewClusterController()
}

// 根据用户获取集群
// @Tags 根据用户获取集群
// @Summary: 根据用户获取集群
// @Description: 根据用户获取集群
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   page     query    string     true        "页码"
// @Param   size     query    string     true        "长度"
// @Param   user_id     query    string     true        "用户ID"
// @Success 200 {object} entity.ClusterPage
// @Router /api/v1/clusters [get]
func GetClusterByUserId(ctx *gin.Context) {
	page := ctx.Query("page")
	size := ctx.Query("size")
	userId := ctx.Query("user_id")
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		ginx.Dangerous(err)
	}
	pageSize, err := strconv.Atoi(size)
	if err != nil {
		ginx.Dangerous(err)
	}
	pageItem, err := clusterController.ClusterService.Page(pageNum, pageSize, userId)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(pageItem, nil)

}

// 根据集群ID获取集群
// @Tags 根据集群ID获取集群
// @Summary: 根据集群ID获取集群
// @Description: 根据集群ID获取集群
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   cluster_id     query    string     true        "集群ID"
// @Success 200 {object} entity.Cluster
// @Router /api/v1/clusters/:cluster_id [get]
func GetClusterByID(ctx *gin.Context) {
	clusterID := ctx.Param("cluster_id")
	clusterObj, err := clusterController.ClusterService.GetByID(clusterID)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(clusterObj, nil)
}

// 根据集群名获取集群
// @Tags 根据集群名获取集群
// @Summary: 根据集群名获取集群
// @Description: 根据集群名获取集群
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   cluster_name     query    string     true        "集群名"
// @Success 200 {object} entity.Cluster
// @Router /api/v1/cluster [get]
func GetClusterByName(ctx *gin.Context) {
	clusterName := ctx.Query("cluster_name")
	clusterObj, err := clusterController.ClusterService.Get(clusterName)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(clusterObj, nil)
}

// 根据集群名获取集群
// @Tags 根据集群名获取集群
// @Summary: 根据集群名获取集群
// @Description: 根据集群名获取集群
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   request body entity.Cluster true "request"
// @Success 200
// @Router /api/v1/cluster [post]
func ImportCluster(ctx *gin.Context) {
	var clusterDTO entity.Cluster
	if err := ctx.ShouldBind(&clusterDTO); err != nil {
		ginx.Dangerous(err)
	}
	err := clusterController.ClusterService.Import(clusterDTO)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(nil, nil)
}

// 根据用户ID和集群名称删除集群
// @Tags 根据用户ID和集群名称删除集群
// @Summary: 根据用户ID和集群名称删除集群
// @Description: 根据用户ID和集群名称删除集群
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   user_id     query    string     true        "用户ID"
// @Param   name     query    string     true        "集群名称"
// @Success 200
// @Router /api/v1/cluster [delete]
func DeleteCluster(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	name := ctx.Query("name")
	err := clusterController.ClusterService.Delete(userId, name)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(nil, nil)
}
