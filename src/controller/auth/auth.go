package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/service/auth"
	"github.com/toolkits/pkg/ginx"
)

type AuthController struct {
	Ctx         context.Context
	AuthService auth.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		AuthService: auth.NewAuthService(),
	}
}

var authController AuthController

func init() {
	authController = *NewAuthController()
}

// 登录
// @Tags 登录
// @Summary: 用户登录
// @Description: 用户登录
// @Accept json
// @Param request body entity.LoginProfile true "request"
// @Success 200 {object} entity.Profile
// @Router /api/v1/auth/login/ [post]
func Login(ctx *gin.Context) {
	var aul entity.LoginProfile
	if err := ctx.ShouldBind(&aul); err != nil {
		ginx.Dangerous(err)
	}
	profile, err := authController.AuthService.Login(aul)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(profile, nil)
}

// 刷新token
// @Tags 刷新token
// @Summary: 刷新token
// @Description: 刷新token
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Success 200 {object} entity.Profile
// @Router /api/v1/auth/refresh/ [post]
func Refresh(ctx *gin.Context) {
	profile, err := authController.AuthService.Refresh(ctx.Request)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(profile, nil)
}

// 登出
// @Tags 登出
// @Summary: 登出
// @Description: 登出
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Router /api/v1/auth/logout/ [post]
func Logout(ctx *gin.Context) {
	err := authController.AuthService.Logout(ctx.Request)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(nil, nil)

}

// 注册
// @Tags 注册
// @Summary: 注册
// @Description: 注册
// @Accept json
// @Param request body entity.User true "request"
// @Success 200 {object} entity.Profile
// @Router /api/v1/auth/register/ [post]
func Register(ctx *gin.Context) {
	var usr entity.User
	if err := ctx.ShouldBind(&usr); err != nil {
		ginx.Dangerous(err)
	}
	profile, err := authController.AuthService.Register(usr)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(profile, nil)

}
