package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/resource/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册res_enterprise插件路由
    var enterpriseControllers = controllers.NewEnterpriseController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // res_enterprise插件路由组
        enterprise := engine.Group("/api/plugins/resource/enterprise")
        enterprise.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        enterprise.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        enterprise.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建res_enterprise
            enterprise.POST("/add", enterpriseControllers.Create)
            // 更新res_enterprise
            enterprise.PUT("/edit", enterpriseControllers.Update)
            // 删除res_enterprise
            enterprise.DELETE("/delete", enterpriseControllers.Delete)
            // res_enterprise列表（分页查询）
            enterprise.GET("/list", enterpriseControllers.List)
            // 根据ID获取res_enterprise信息
            enterprise.GET("/:id", enterpriseControllers.GetByID)
        }

        app.ZapLog.Info("res_enterprise插件路由注册成功")
    })
}