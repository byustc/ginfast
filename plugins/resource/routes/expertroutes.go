package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/resource/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册res_expert插件路由
    var expertControllers = controllers.NewExpertController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // res_expert插件路由组
        expert := engine.Group("/api/plugins/resource/expert")
        expert.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        expert.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        expert.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建res_expert
            expert.POST("/add", expertControllers.Create)
            // 更新res_expert
            expert.PUT("/edit", expertControllers.Update)
            // 删除res_expert
            expert.DELETE("/delete", expertControllers.Delete)
            // res_expert列表（分页查询）
            expert.GET("/list", expertControllers.List)
            // 根据ID获取res_expert信息
            expert.GET("/:id", expertControllers.GetByID)
        }

        app.ZapLog.Info("res_expert插件路由注册成功")
    })
}