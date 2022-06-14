package postgres

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/service/postgresql"
	"github.com/toolkits/pkg/ginx"
	"strconv"
)

type PostgresController struct {
	Ctx             context.Context
	PostgresService postgresql.PostgresService
}

func NewPostgresController() *PostgresController {
	return &PostgresController{
		PostgresService: postgresql.NewPostgresService(),
	}
}

var postgresController PostgresController

func init() {
	postgresController = *NewPostgresController()
}

// 根据postgresql名称获取postgres对象
// @Tags 根据postgresql名称获取postgres对象
// @Summary: 根据postgresql名称获取postgres对象
// @Description: 根据postgresql名称获取postgres对象
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   instance_name     query    string     true        "postgresql实例名称"
// @Success 200 {object} entity.Instance
// @Router /api/v1/postgres [get]
func Get(ctx *gin.Context) {
	instanceName := ctx.Query("instance_name")
	instanceObj, err := postgresController.PostgresService.Get(instanceName)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(instanceObj, nil)
}

// 根据postgresql ID获取postgres对象
// @Tags 根据postgresql ID获取postgres对象
// @Summary: 根据postgresql ID获取postgres对象
// @Description: 根据postgresql ID获取postgres对象
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   postgres_id     query    string     true        "postgresql实例ID"
// @Success 200 {object} entity.Instance
// @Router /api/v1/postgres/:postgres_id [get]
func GetById(ctx *gin.Context) {
	instanceId := ctx.Param("postgres_id")
	instanceObj, err := postgresController.PostgresService.GetById(instanceId)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(instanceObj, nil)
}

// 根据postgresql ID获取postgres状态
// @Tags 根据postgresql ID获取postgres状态
// @Summary: 根据postgresql ID获取postgres状态
// @Description: 根据postgresql ID获取postgres状态
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   postgres_id     query    string     true        "postgresql实例ID"
// @Success 200
// @Router /api/v1/postgres/:postgres_id/status [get]
func GetStatus(ctx *gin.Context) {
	instanceId := ctx.Query("instance_id")
	instanceObj, err := postgresController.PostgresService.GetStatusById(instanceId)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(instanceObj, nil)
}

// 根据用户ID获取postgres实例列表
// @Tags 根据用户ID获取postgres实例列表
// @Summary: 根据用户ID获取postgres实例列表
// @Description: 根据用户ID获取postgres实例列表
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   page     query    string     true        "页码"
// @Param   size     query    string     true        "长度"
// @Param   user_id     query    string     true        "用户ID"
// @Success 200 {object} entity.InstancePage
// @Router /api/v1/postgressqls [get]
func Page(ctx *gin.Context) {
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
	pageItem, err := postgresController.PostgresService.Page(pageNum, pageSize, userId)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(pageItem, nil)

}

// 创建postgres实例
// @Tags 创建postgres实例
// @Summary: 创建postgres实例
// @Description: 创建postgres实例
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   request body entity.Instance true "request"
// @Success 200
// @Router /api/v1/postgres [post]
func Save(ctx *gin.Context) {
	var instanceDTO entity.Instance
	if err := ctx.ShouldBind(&instanceDTO); err != nil {
		ginx.Dangerous(err)
	}
	err := postgresController.PostgresService.Save(instanceDTO)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(nil, nil)
}

// 从模板创建postgres实例
// @Tags 从模板创建postgres实例
// @Summary: 从模板创建postgres实例
// @Description: 从模板创建postgres实例
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   request body entity.Postgres true "request"
// @Success 200
// @Router /api/v1/postgres/template [post]
func Create(ctx *gin.Context) {
	var postgres entity.Postgres
	if err := ctx.ShouldBind(&postgres); err != nil {
		ginx.Dangerous(err)
	}
	instanceId, err := postgresController.PostgresService.Create(postgres)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(instanceId, nil)
}

// 删除实例
// @Tags 删除实例
// @Summary: 删除实例
// @Description: 删除实例
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   name     query    string     true        "实例名"
// @Success 200
// @Router /api/v1/postgres [delete]
func Delete(ctx *gin.Context) {
	name := ctx.Query("name")
	err := postgresController.PostgresService.Delete(name)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(nil, nil)
}

// 更新实例
// @Tags 更新实例
// @Summary: 更新实例
// @Description: 更新实例
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   request body entity.Instance true "request"
// @Success 200
// @Router /api/v1/postgres [patch]
func Update(ctx *gin.Context) {
	var instance entity.Instance
	if err := ctx.ShouldBind(&instance); err != nil {
		ginx.Dangerous(err)
	}
	err := postgresController.PostgresService.Update(instance)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(nil, nil)
}

// 编辑实例
// @Tags 编辑实例
// @Summary: 编辑实例
// @Description: 编辑实例
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   request body entity.Instance true "request"
// @Success 200
// @Router /api/v1/postgres/:postgres_id [post]
func Edit(ctx *gin.Context) {
	var instance entity.Instance
	if err := ctx.ShouldBind(&instance); err != nil {
		ginx.Dangerous(err)
	}
	err := postgresController.PostgresService.Edit(instance)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(nil, nil)
}
