package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mensylisir/kmpp-middleware/src/constant"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/service/pod"
	"github.com/mensylisir/kmpp-middleware/src/service/postgresql"
	"github.com/mensylisir/kmpp-middleware/src/util/aop"
	"github.com/toolkits/pkg/ginx"
	"regexp"
	"strconv"
	"strings"
)

type PostgresController struct {
	Ctx             context.Context
	PostgresService postgresql.PostgresService
	PodService      pod.PodService
}

func NewPostgresController() *PostgresController {
	return &PostgresController{
		PostgresService: postgresql.NewPostgresService(),
		PodService:      pod.NewPodService(),
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
// @Router /api/ws/v1/pod/log [get]
func GetPodLogs(ctx *gin.Context) {
	var postgresLog entity.PostgresLog
	ws, err := aop.UpGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ws.WriteMessage(websocket.TextMessage, []byte(err.Error()))
	}
	err = ws.ReadJSON(&postgresLog)
	if err != nil {
		ws.WriteMessage(websocket.TextMessage, []byte(err.Error()))
	}
	log := make(chan string)
	operatorInstance := postgresLog.Instance
	operatorInstance.Name = postgresLog.OperatorName
	operatorInstance.Namespace = postgresLog.OperatorNamespace
	postgresController.PodService.GetPogLog(operatorInstance, log)
	//logMap := make(map[string]string)
	logFormat := entity.LogFormat{}
	for {
		data := <-log
		cn := fmt.Sprintf("%s/%s", postgresLog.Namespace, postgresLog.Name)
		if ok := strings.Contains(data, cn); ok {
			timeRegex := regexp.MustCompile("time=\"(.*)\"\\s*?level")
			logFormat.Time = timeRegex.FindStringSubmatch(data)[1]
			levelRegex := regexp.MustCompile("level=(.*)\\s+?msg")
			logFormat.Level = levelRegex.FindStringSubmatch(data)[1]
			msgRegex := regexp.MustCompile("msg=\"(.*)\"")
			logFormat.Msg = msgRegex.FindStringSubmatch(data)[1]
			mJson, err := json.Marshal(logFormat)
			if err != nil {
				ws.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			}
			ws.WriteMessage(websocket.TextMessage, mJson)
		}
		//ws.WriteMessage(websocket.TextMessage, []byte(data))
	}
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
	instanceId := ctx.Param("postgres_id")
	instanceObj, err := postgresController.PostgresService.GetStatusById(instanceId)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(instanceObj, nil)
}

func GetPostgresOperatorName(ctx *gin.Context) {
	clusterId := ctx.Query("cluster_id")
	podName, err := postgresController.PostgresService.GetPostgresOperatorPodName(clusterId)
	if err != nil {
		ginx.Dangerous(err)
	}
	operatorMap := make(map[string]string)
	operatorMap["namespace"] = constant.PostgresOperatorNamespace
	operatorMap["pod_name"] = podName
	ginx.NewRender(ctx).Data(operatorMap, nil)
}
