package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/biz/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册biz_leads插件路由
    var leadsControllers = controllers.NewLeadsController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // biz_leads插件路由组
        leads := engine.Group("/api/plugins/biz/leads")
        leads.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        leads.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        leads.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建biz_leads
            leads.POST("/add", leadsControllers.Create)
            // 更新biz_leads
            leads.PUT("/edit", leadsControllers.Update)
            // 删除biz_leads
            leads.DELETE("/delete", leadsControllers.Delete)
            // biz_leads列表（分页查询）
            leads.GET("/list", leadsControllers.List)
            // 根据ID获取biz_leads信息
            leads.GET("/:id", leadsControllers.GetByID)
        }

        app.ZapLog.Info("biz_leads插件路由注册成功")
    })
}