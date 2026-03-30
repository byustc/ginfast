package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/resource/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册res_carrier插件路由
    var carrierControllers = controllers.NewCarrierController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // res_carrier插件路由组
        carrier := engine.Group("/api/plugins/resource/carrier")
        carrier.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        carrier.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        carrier.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建res_carrier
            carrier.POST("/add", carrierControllers.Create)
            // 更新res_carrier
            carrier.PUT("/edit", carrierControllers.Update)
            // 删除res_carrier
            carrier.DELETE("/delete", carrierControllers.Delete)
            // res_carrier列表（分页查询）
            carrier.GET("/list", carrierControllers.List)
            // 根据ID获取res_carrier信息
            carrier.GET("/:id", carrierControllers.GetByID)
        }

        app.ZapLog.Info("res_carrier插件路由注册成功")
    })
}