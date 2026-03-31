package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/resource/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册res_thinktank_fund插件路由
    var thitankFundControllers = controllers.NewThitankFundController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // res_thinktank_fund插件路由组
        thitankFund := engine.Group("/api/plugins/resource/thitankfund")
        thitankFund.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        thitankFund.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        thitankFund.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建res_thinktank_fund
            thitankFund.POST("/add", thitankFundControllers.Create)
            // 更新res_thinktank_fund
            thitankFund.PUT("/edit", thitankFundControllers.Update)
            // 删除res_thinktank_fund
            thitankFund.DELETE("/delete", thitankFundControllers.Delete)
            // res_thinktank_fund列表（分页查询）
            thitankFund.GET("/list", thitankFundControllers.List)
            // 根据ID获取res_thinktank_fund信息
            thitankFund.GET("/:id", thitankFundControllers.GetByID)
        }

        app.ZapLog.Info("res_thinktank_fund插件路由注册成功")
    })
}