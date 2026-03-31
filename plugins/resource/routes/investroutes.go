package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/resource/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册res_invest插件路由
    var investControllers = controllers.NewInvestController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // res_invest插件路由组
        invest := engine.Group("/api/plugins/resource/invest")
        invest.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        invest.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        invest.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建res_invest
            invest.POST("/add", investControllers.Create)
            // 更新res_invest
            invest.PUT("/edit", investControllers.Update)
            // 删除res_invest
            invest.DELETE("/delete", investControllers.Delete)
            // res_invest列表（分页查询）
            invest.GET("/list", investControllers.List)
            // 根据ID获取res_invest信息
            invest.GET("/:id", investControllers.GetByID)
        }

        app.ZapLog.Info("res_invest插件路由注册成功")
    })
}