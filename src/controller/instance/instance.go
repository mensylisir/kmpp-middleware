package instance

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mensylisir/kmpp-middleware/src/service/instance"
	"github.com/toolkits/pkg/ginx"
	"strconv"
)

type InstanceController struct {
	Ctx             context.Context
	InstanceService instance.InstanceService
}

func NewInstanceController() *InstanceController {
	return &InstanceController{
		InstanceService: instance.NewInstanceService(),
	}
}

var instanceController InstanceController

func init() {
	instanceController = *NewInstanceController()
}

// 根据用户获取实例
// @Tags 根据用户获取实例
// @Summary: 根据用户获取实例
// @Description: 根据用户获取实例
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   page     query    string     true        "页码"
// @Param   size     query    string     true        "长度"
// @Param   user_id     query    string     true        "用户ID"
// @Success 200 {object} entity.InstancePage
// @Router /api/v1/instances [get]
func GetInstanceByUserId(ctx *gin.Context) {
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
	pageItem, err := instanceController.InstanceService.Page(pageNum, pageSize, userId)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(pageItem, nil)

}

// 根据实例ID获取实例
// @Tags 根据实例ID获取实例
// @Summary: 根据实例ID获取实例
// @Description: 根据实例ID获取实例
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   id     query    string     true        "实例ID"
// @Success 200 {object} entity.Instance
// @Router /api/v1/instances/:instance_id [get]
func GetInstanceByID(ctx *gin.Context) {
	instanceId := ctx.Param("id")
	instanceObj, err := instanceController.InstanceService.GetById(instanceId)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(instanceObj, nil)
}

// 根据实例名获取实例
// @Tags 根据实例名获取实例
// @Summary: 根据实例名获取实例
// @Description: 根据实例名获取实例
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   name     query    string     true        "实例name"
// @Success 200 {object} entity.Instance
// @Router /api/v1/instance/ [get]
func GetInstanceByName(ctx *gin.Context) {
	instanceName := ctx.Query("name")
	instanceObj, err := instanceController.InstanceService.Get(instanceName)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(instanceObj, nil)
}
