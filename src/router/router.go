package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mensylisir/kmpp-middleware/docs"
	"github.com/mensylisir/kmpp-middleware/src/controller/auth"
	"github.com/mensylisir/kmpp-middleware/src/controller/cluster"
	"github.com/mensylisir/kmpp-middleware/src/controller/instance"
	"github.com/mensylisir/kmpp-middleware/src/controller/pod"
	"github.com/mensylisir/kmpp-middleware/src/controller/postgres"
	"github.com/mensylisir/kmpp-middleware/src/controller/svc"
	"github.com/mensylisir/kmpp-middleware/src/controller/templates"
	"github.com/mensylisir/kmpp-middleware/src/controller/user"
	"github.com/mensylisir/kmpp-middleware/src/util/aop"
	"github.com/spf13/viper"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"os"
)

func New(version string) *gin.Engine {

	PrintAccessLog := viper.GetBool("bind.print_access_log")
	RunMode := viper.GetString("app.run_mode")
	gin.SetMode(RunMode)

	loggerMid := aop.Logrus()
	recoveryMid := aop.Recovery()
	r := gin.New()
	r.Use(recoveryMid)
	if PrintAccessLog {
		r.Use(loggerMid)
	}
	configSwagger(r)
	configRoute(r, version)
	return r
}

func configSwagger(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func configRoute(r *gin.Engine, version string) {
	httpRouter := r.Group("/api/v1")
	configHttpRouter(httpRouter, version)
	//
	//websocket := r.Group("api/ws/v1")
	//configWebsocketRouter(websocket)
}

func configHttpRouter(rg *gin.RouterGroup, version string) {
	rg.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	rg.GET("/pid", func(c *gin.Context) {
		c.String(200, fmt.Sprintf("%d", os.Getpid()))
	})
	rg.GET("/addr", func(c *gin.Context) {
		c.String(200, c.Request.RemoteAddr)
	})
	rg.GET("/version", func(c *gin.Context) {
		c.String(200, version)
	})

	//rg.POST("/auth/login", auth.Login)
	//rg.POST("/auth/register", auth.Register)
	//rg.POST("/auth/refresh", aop.Auth(), auth.Refresh)
	//rg.POST("/auth/logout", aop.Auth(), auth.Logout)
	//
	//rg.GET("/template", aop.Auth(), templates.GetByName)
	//rg.GET("/template/:template_id", aop.Auth(), templates.GetById)
	//rg.GET("/templates", aop.Auth(), templates.Page)
	//rg.POST("/template", aop.Auth(), templates.Save)
	//rg.PATCH("/template", aop.Auth(), templates.Update)
	//rg.DELETE("/template", aop.Auth(), templates.Delete)
	//
	//rg.GET("/user", aop.Auth(), user.Get)
	//rg.GET("/user/:user_id", aop.Auth(), user.GetById)
	//rg.GET("/users", aop.Auth(), user.Page)
	//rg.POST("/user", aop.Auth(), user.Create)
	//rg.PATCH("/user", aop.Auth(), user.Update)
	//rg.DELETE("/user", aop.Auth(), user.Delete)
	//rg.PATCH("/user/password", aop.Auth(), user.ChangePassword)
	//
	//rg.GET("/clusters", aop.Auth(), cluster.GetClusterByUserId)
	//rg.GET("/clusters/:cluster_id", aop.Auth(), cluster.GetClusterByID)
	//rg.GET("/cluster", aop.Auth(), cluster.GetClusterByName)
	//rg.POST("/cluster", aop.Auth(), cluster.ImportCluster)
	//rg.DELETE("cluster", aop.Auth(), cluster.DeleteCluster)
	//
	//rg.GET("/instances", aop.Auth(), instance.GetInstanceByUserId)
	//rg.GET("/instances/:instance_id", aop.Auth(), instance.GetInstanceByID)
	//rg.GET("/instance", aop.Auth(), instance.GetInstanceByName)
	//
	//rg.GET("/pods", aop.Auth(), pod.GetPods)
	//rg.GET("/pod/logs", aop.Auth(), pod.GetPodLogs)
	//rg.GET("/pod/status", aop.Auth(), pod.GetPodStatus)
	//rg.GET("/pods/status", aop.Auth(), pod.GetPodsStatus)
	//
	//rg.GET("/postgres", aop.Auth(), postgres.Get)
	//rg.GET("/postgres/:postgres_id", aop.Auth(), postgres.GetById)
	//rg.GET("/postgres/:postgres_id/status", aop.Auth(), postgres.GetStatus)
	//rg.GET("/postgressqls", aop.Auth(), postgres.Page)
	//rg.POST("/postgres", aop.Auth(), postgres.Save)
	//rg.POST("/postgres/template", aop.Auth(), postgres.Create)
	//rg.PATCH("/postgres", aop.Auth(), postgres.Update)
	//rg.DELETE("/postgres", aop.Auth(), postgres.Delete)
	//rg.POST("/postgres/:postgres_id", aop.Auth(), postgres.Edit)

	//rg.PATCH("/svc", aop.Auth(), svc.UpdateServiceType)

	rg.POST("/auth/login", auth.Login)
	rg.POST("/auth/register", auth.Register)
	rg.POST("/auth/refresh", auth.Refresh)
	rg.POST("/auth/logout", auth.Logout)

	rg.GET("/template", templates.GetByName)
	rg.GET("/template/:template_id", templates.GetById)
	rg.GET("/templates", templates.Page)
	rg.POST("/template", templates.Save)
	rg.PATCH("/template", templates.Update)
	rg.DELETE("/template", templates.Delete)

	rg.GET("/user", user.Get)
	rg.GET("/user/:user_id", user.GetById)
	rg.GET("/users", user.Page)
	rg.POST("/user", user.Create)
	rg.PATCH("/user", user.Update)
	rg.DELETE("/user", user.Delete)
	rg.PATCH("/user/password", user.ChangePassword)

	rg.GET("/clusters", cluster.GetClusterByUserId)
	rg.GET("/clusters/:cluster_id", cluster.GetClusterByID)
	rg.GET("/cluster", cluster.GetClusterByName)
	rg.POST("/cluster", cluster.ImportCluster)
	rg.DELETE("cluster", cluster.DeleteCluster)

	rg.GET("/instances", instance.GetInstanceByUserId)
	rg.GET("/instances/:instance_id", instance.GetInstanceByID)
	rg.GET("/instance", instance.GetInstanceByName)

	rg.GET("/pods", pod.GetPods)
	rg.GET("/pod/logs", pod.GetPodLogs)
	rg.GET("/pod/status", pod.GetPodStatus)
	rg.GET("/pods/status", pod.GetPodsStatus)

	rg.GET("/postgres", postgres.Get)
	rg.GET("/postgres/:postgres_id", postgres.GetById)
	rg.GET("/postgres/:postgres_id/status", postgres.GetStatus)
	rg.GET("/postgressqls", postgres.Page)
	rg.POST("/postgres", postgres.Save)
	rg.POST("/postgres/template", postgres.Create)
	rg.PATCH("/postgres", postgres.Update)
	rg.DELETE("/postgres", postgres.Delete)
	rg.POST("/postgres/:postgres_id", postgres.Edit)

	rg.PATCH("/svc", svc.UpdateServiceType)
}
