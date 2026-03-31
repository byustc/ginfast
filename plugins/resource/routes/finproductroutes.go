package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/resource/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册res_finance_product插件路由
    var finProductControllers = controllers.NewFinProductController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // res_finance_product插件路由组
        finProduct := engine.Group("/api/plugins/resource/finproduct")
        finProduct.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        finProduct.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        finProduct.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建res_finance_product
            finProduct.POST("/add", finProductControllers.Create)
            // 更新res_finance_product
            finProduct.PUT("/edit", finProductControllers.Update)
            // 删除res_finance_product
            finProduct.DELETE("/delete", finProductControllers.Delete)
            // res_finance_product列表（分页查询）
            finProduct.GET("/list", finProductControllers.List)
            // 根据ID获取res_finance_product信息
            finProduct.GET("/:id", finProductControllers.GetByID)
        }

        app.ZapLog.Info("res_finance_product插件路由注册成功")
    })
}