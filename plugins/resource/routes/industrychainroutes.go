package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/resource/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册res_industry_chain插件路由
    var industryChainControllers = controllers.NewIndustryChainController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // res_industry_chain插件路由组
        industryChain := engine.Group("/api/plugins/resource/industrychain")
        industryChain.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        industryChain.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        industryChain.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建res_industry_chain
            industryChain.POST("/add", industryChainControllers.Create)
            // 更新res_industry_chain
            industryChain.PUT("/edit", industryChainControllers.Update)
            // 删除res_industry_chain
            industryChain.DELETE("/delete", industryChainControllers.Delete)
            // res_industry_chain列表（分页查询）
            industryChain.GET("/list", industryChainControllers.List)
            // 根据ID获取res_industry_chain信息
            industryChain.GET("/:id", industryChainControllers.GetByID)
        }

        app.ZapLog.Info("res_industry_chain插件路由注册成功")
    })
}