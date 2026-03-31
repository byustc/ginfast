package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/resource/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册res_finance插件路由
    var financeControllers = controllers.NewFinanceController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // res_finance插件路由组
        finance := engine.Group("/api/plugins/resource/finance")
        finance.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        finance.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        finance.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建res_finance
            finance.POST("/add", financeControllers.Create)
            // 更新res_finance
            finance.PUT("/edit", financeControllers.Update)
            // 删除res_finance
            finance.DELETE("/delete", financeControllers.Delete)
            // res_finance列表（分页查询）
            finance.GET("/list", financeControllers.List)
            // 根据ID获取res_finance信息
            finance.GET("/:id", financeControllers.GetByID)
        }

        app.ZapLog.Info("res_finance插件路由注册成功")
    })
}