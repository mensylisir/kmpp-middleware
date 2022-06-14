package templates

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/service/templates"
	"github.com/toolkits/pkg/ginx"
	"strconv"
)

type TemplatesController struct {
	Ctx              context.Context
	TemplatesService templates.TemplatesService
}

func NewTemplatesController() *TemplatesController {
	return &TemplatesController{
		TemplatesService: templates.NewTemplatesService(),
	}
}

var templatesController TemplatesController

func init() {
	templatesController = *NewTemplatesController()
}

// 根据中间件类型获取模板
// @Tags 使用中间件类型获取模板
// @Summary: 使用中间件类型获取模板
// @Description: 使用中间件类型获取模板
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   name     query    string     true        "中间件类型"
// @Success 200 {object} entity.Templates
// @Router /api/v1/template [get]
func GetByName(ctx *gin.Context) {
	name := ctx.Query("name")
	template, err := templatesController.TemplatesService.Get(name)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(template, nil)
}

// 根据中间件ID获取模板
// @Tags 使用中间件ID获取模板
// @Summary: 使用中间件ID获取模板
// @Description: 使用中间件ID获取模板
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   id     query    string     true        "中间件ID"
// @Success 200 {object} entity.Templates
// @Router /api/v1/template/:template_id [get]
func GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	template, err := templatesController.TemplatesService.GetById(id)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(template, nil)
}

// 获取所有模板
// @Tags 获取所有模板
// @Summary: 使用中间件ID获取模板
// @Description: 使用中间件ID获取模板
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   page     query    string     true        "页码"
// @Param   size     query    string     true        "长度"
// @Success 200 {object} entity.Page
// @Router /api/v1/templates [get]
func Page(ctx *gin.Context) {
	page := ctx.Query("page")
	size := ctx.Query("size")
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		ginx.Dangerous(err)
	}
	pageSize, err := strconv.Atoi(size)
	if err != nil {
		ginx.Dangerous(err)
	}
	templatePage, err := templatesController.TemplatesService.Page(pageNum, pageSize)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(templatePage, nil)
}

func List(ctx *gin.Context) {
	templates1, err := templatesController.TemplatesService.List()
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(templates1, nil)
}

// 创建模板
// @Tags 创建模板
// @Summary: 创建模板
// @Description: 创建模板
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   request body entity.Templates true "request"
// @Success 200 {object} entity.Templates
// @Router /api/v1/template [post]
func Save(ctx *gin.Context) {
	var template entity.Templates
	if err := ctx.ShouldBind(&template); err != nil {
		ginx.Dangerous(err)
	}
	err := templatesController.TemplatesService.Save(template)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(template, nil)
}

// 更新模板
// @Tags 更新模板
// @Summary: 更新模板
// @Description: 更新模板
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   request body entity.Templates true "request"
// @Success 200 {object} entity.Templates
// @Router /api/v1/template [patch]
func Update(ctx *gin.Context) {
	var template entity.Templates
	if err := ctx.ShouldBind(&template); err != nil {
		ginx.Dangerous(err)
	}
	err := templatesController.TemplatesService.Update(template)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(template, nil)
}

// 删除模板
// @Tags 删除模板
// @Summary: 删除模板
// @Description: 删除模板
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   name     query    string     true        "模板类型"
// @Success 200
// @Router /api/v1/template [delete]
func Delete(ctx *gin.Context) {
	name := ctx.Query("name")
	err := templatesController.TemplatesService.Delete(name)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(nil, nil)
}

func Batch(ctx *gin.Context) {
	var templates entity.OperateTemplates
	if err := ctx.ShouldBind(&templates); err != nil {
		ginx.Dangerous(err)
	}
	err := templatesController.TemplatesService.Batch(templates)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(nil, nil)
}
