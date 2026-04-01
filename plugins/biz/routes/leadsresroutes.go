package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/biz/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册biz_leads_res插件路由
    var leadsResControllers = controllers.NewLeadsResController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // biz_leads_res插件路由组
        leadsRes := engine.Group("/api/plugins/biz/leadsres")
        leadsRes.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        leadsRes.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        leadsRes.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建biz_leads_res
            leadsRes.POST("/add", leadsResControllers.Create)
            // 更新biz_leads_res
            leadsRes.PUT("/edit", leadsResControllers.Update)
            // 删除biz_leads_res
            leadsRes.DELETE("/delete", leadsResControllers.Delete)
            // biz_leads_res列表（分页查询）
            leadsRes.GET("/list", leadsResControllers.List)
            // 根据ID获取biz_leads_res信息
            leadsRes.GET("/:id", leadsResControllers.GetByID)
        }

        app.ZapLog.Info("biz_leads_res插件路由注册成功")
    })
}