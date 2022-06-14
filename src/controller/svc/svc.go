package svc

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/service/svc"
	"github.com/toolkits/pkg/ginx"
)

type SvcController struct {
	Ctx        context.Context
	SvcService svc.SvcService
}

func NewSvcController() *SvcController {
	return &SvcController{
		SvcService: svc.NewSvcService(),
	}
}

var svcController SvcController

func init() {
	svcController = *NewSvcController()
}

// 更新ServiceType
// @Tags 更新ServiceType
// @Summary: 更新ServiceType
// @Description: 更新ServiceType
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   request body entity.Instance true "request"
// @Success 200
// @Router /api/v1/svc [patch]
func UpdateServiceType(ctx *gin.Context) {
	var instance entity.Instance
	err := ctx.ShouldBind(&instance)
	if err != nil {
		ginx.Dangerous(err)
	}
	err = svcController.SvcService.UpdateServiceType(instance)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(nil, nil)
}
