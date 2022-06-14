package pod

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/service/pod"
	"github.com/toolkits/pkg/ginx"
)

type PodController struct {
	Ctx        context.Context
	PodService pod.PodService
}

func NewPodController() *PodController {
	return &PodController{
		PodService: pod.NewPodService(),
	}
}

var podController PodController

func init() {
	podController = *NewPodController()
}

// 根据namespace下所有Pod名称
// @Tags 根据namespace下所有Pod名称
// @Summary: 根据namespace下所有Pod名称
// @Description: 根据namespace下所有Pod名称
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   cluster_id     query    string     true        "集群ID"
// @Param   namespace     query    string     true        "命名空间"
// @Success 200 {object} []string
// @Router /api/v1/pods [get]
func GetPods(ctx *gin.Context) {
	clusterID := ctx.Query("cluster_id")
	Namespace := ctx.Query("namespace")
	instance := entity.Instance{}
	instance.ClusterID = clusterID
	instance.Namespace = Namespace
	pods, err := podController.PodService.GetPods(instance)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(pods, nil)
}

// 获取Pod的日志
// @Tags 获取Pod的日志
// @Summary: 获取Pod的日志
// @Description: 获取Pod的日志
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   cluster_id     query    string     true        "集群ID"
// @Param   namespace     query    string     true        "命名空间"
// @Param   name     query    string     true        "pod名称"
// @Success 200 {object} string
// @Router /api/v1/pod/log [get]
func GetPodLogs(ctx *gin.Context) {
	clusterID := ctx.Query("cluster_id")
	Namespace := ctx.Query("namespace")
	Name := ctx.Query("pod_name")
	instance := entity.Instance{}
	instance.ClusterID = clusterID
	instance.Namespace = Namespace
	instance.Name = Name
	logs, err := podController.PodService.GetPogLog(instance)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(logs, nil)
}

// 获取Pod的状态
// @Tags 获取Pod的状态
// @Summary: 获取Pod的状态
// @Description: 获取Pod的状态
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   cluster_id     query    string     true        "集群ID"
// @Param   namespace     query    string     true        "命名空间"
// @Param   name     query    string     true        "pod名称"
// @Success 200 {object} entity.PodStatus
// @Router /api/v1/pod/status [get]
func GetPodStatus(ctx *gin.Context) {
	clusterID := ctx.Query("cluster_id")
	Namespace := ctx.Query("namespace")
	Name := ctx.Query("pod_name")
	instance := entity.Instance{}
	instance.ClusterID = clusterID
	instance.Namespace = Namespace
	instance.Name = Name
	status, err := podController.PodService.GetPodStatus(instance)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(status, nil)

}

// 根据namespace获取其所有Pod的状态
// @Tags 根据namespace获取其所有Pod的状态
// @Summary: 根据namespace获取其所有Pod的状态
// @Description: 根据namespace获取其所有Pod的状态
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   cluster_id     query    string     true        "集群ID"
// @Param   namespace     query    string     true        "命名空间"
// @Success 200 {object} []entity.PodStatus
// @Router /api/v1/pods/status [get]
func GetPodsStatus(ctx *gin.Context) {
	clusterID := ctx.Query("cluster_id")
	Namespace := ctx.Query("namespace")
	instance := entity.Instance{}
	instance.ClusterID = clusterID
	instance.Namespace = Namespace
	status, err := podController.PodService.GetPodsStatus(instance)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(status, nil)

}
