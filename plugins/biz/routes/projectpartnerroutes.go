package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/biz/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册biz_project_partner插件路由
    var projectPartnerControllers = controllers.NewProjectPartnerController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // biz_project_partner插件路由组
        projectPartner := engine.Group("/api/plugins/biz/projectpartner")
        projectPartner.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        projectPartner.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        projectPartner.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建biz_project_partner
            projectPartner.POST("/add", projectPartnerControllers.Create)
            // 更新biz_project_partner
            projectPartner.PUT("/edit", projectPartnerControllers.Update)
            // 删除biz_project_partner
            projectPartner.DELETE("/delete", projectPartnerControllers.Delete)
            // biz_project_partner列表（分页查询）
            projectPartner.GET("/list", projectPartnerControllers.List)
            // 根据ID获取biz_project_partner信息
            projectPartner.GET("/:id", projectPartnerControllers.GetByID)
        }

        app.ZapLog.Info("biz_project_partner插件路由注册成功")
    })
}