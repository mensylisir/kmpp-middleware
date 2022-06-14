package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/service/user"
	"github.com/toolkits/pkg/ginx"
	"strconv"
)

type UserController struct {
	Ctx         context.Context
	UserService user.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: user.NewUserService(),
	}
}

var userController UserController

func init() {
	userController = *NewUserController()
}

// 根据用户名获取用户对象
// @Tags 根据用户名获取用户对象
// @Summary: 根据用户名获取用户对象
// @Description: 根据用户名获取用户对象
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   name     query    string     true        "用户名"
// @Success 200 {object} entity.User
// @Router /api/v1/user [get]
func Get(ctx *gin.Context) {
	name := ctx.Query("name")
	user1, err := userController.UserService.Get(name)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(user1, nil)
}

// 根据用户ID获取用户对象
// @Tags 根据用户ID获取用户对象
// @Summary: 根据用户ID获取用户对象
// @Description: 根据用户ID获取用户对象
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   id     query    string     true        "用户id"
// @Success 200 {object} entity.User
// @Router /api/v1/user/:user_id [get]
func GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	user1, err := userController.UserService.GetUserById(id)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(user1, nil)
}

func List(ctx *gin.Context) {
	users, err := userController.UserService.List()
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(users, nil)
}

// 根据用户ID获取用户对象
// @Tags 根据用户ID获取用户对象
// @Summary: 根据用户ID获取用户对象
// @Description: 根据用户ID获取用户对象
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   page     query    string     true        "页码"
// @Param   size     query    string     true        "长度"
// @Success 200 {object} entity.Page
// @Router /api/v1/users [get]
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
	pageItem, err := userController.UserService.Page(pageNum, pageSize)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(pageItem, nil)

}

// 创建用户
// @Tags 创建用户
// @Summary: 创建用户
// @Description: 创建用户
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   request body entity.User true "request"
// @Success 200 {object} entity.User
// @Router /api/v1/user [post]
func Create(ctx *gin.Context) {
	var usr entity.User
	if err := ctx.ShouldBind(&usr); err != nil {
		ginx.Dangerous(err)
	}
	usr1, err := userController.UserService.Create(usr)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(usr1, nil)
}

// 更新用户
// @Tags 更新用户
// @Summary: 更新用户
// @Description: 更新用户
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   request body entity.User true "request"
// @Success 200 {object} entity.User
// @Router /api/v1/user [patch]
func Update(ctx *gin.Context) {
	var usr entity.User
	if err := ctx.ShouldBind(&usr); err != nil {
		ginx.Dangerous(err)
	}
	usr1, err := userController.UserService.Update(usr)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(usr1, nil)
}

// 删除用户
// @Tags 删除用户
// @Summary: 删除用户
// @Description: 删除用户
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   name     query    string     true        "用户名"
// @Success 200
// @Router /api/v1/user [delete]
func Delete(ctx *gin.Context) {
	name := ctx.Query("name")
	err := userController.UserService.Delete(name)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(nil, nil)
}

func Batch(ctx *gin.Context) {
	var users entity.OperateUser
	if err := ctx.ShouldBind(&users); err != nil {
		ginx.Dangerous(err)
	}
	err := userController.UserService.Batch(users)
	if err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(nil, nil)
}

// 修改用户密码
// @Tags 修改用户密码
// @Summary: 修改用户密码
// @Description: 修改用户密码
// @Accept json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param   request body entity.UserChangePassword true "request"
// @Success 200
// @Router /api/v1/user/password [delete]
func ChangePassword(ctx *gin.Context) {
	var usr entity.UserChangePassword
	if err := ctx.ShouldBind(&usr); err != nil {
		ginx.Dangerous(err)
	}
	if err := userController.UserService.ChangePassword(usr); err != nil {
		ginx.Dangerous(err)
	}
	ginx.NewRender(ctx).Data(nil, nil)
}
