package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/resource/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册res_invest_fund插件路由
    var invFundControllers = controllers.NewInvFundController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // res_invest_fund插件路由组
        invFund := engine.Group("/api/plugins/resource/invfund")
        invFund.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        invFund.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        invFund.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建res_invest_fund
            invFund.POST("/add", invFundControllers.Create)
            // 更新res_invest_fund
            invFund.PUT("/edit", invFundControllers.Update)
            // 删除res_invest_fund
            invFund.DELETE("/delete", invFundControllers.Delete)
            // res_invest_fund列表（分页查询）
            invFund.GET("/list", invFundControllers.List)
            // 根据ID获取res_invest_fund信息
            invFund.GET("/:id", invFundControllers.GetByID)
        }

        app.ZapLog.Info("res_invest_fund插件路由注册成功")
    })
}