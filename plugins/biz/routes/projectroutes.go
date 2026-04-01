package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/biz/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册biz_project插件路由
    var projectControllers = controllers.NewProjectController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // biz_project插件路由组
        project := engine.Group("/api/plugins/biz/project")
        project.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        project.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        project.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建biz_project
            project.POST("/add", projectControllers.Create)
            // 更新biz_project
            project.PUT("/edit", projectControllers.Update)
            // 删除biz_project
            project.DELETE("/delete", projectControllers.Delete)
            // biz_project列表（分页查询）
            project.GET("/list", projectControllers.List)
            // 根据ID获取biz_project信息
            project.GET("/:id", projectControllers.GetByID)
        }

        app.ZapLog.Info("biz_project插件路由注册成功")
    })
}