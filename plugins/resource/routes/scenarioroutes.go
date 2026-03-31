package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/resource/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册res_scenario插件路由
    var scenarioControllers = controllers.NewScenarioController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // res_scenario插件路由组
        scenario := engine.Group("/api/plugins/resource/scenario")
        scenario.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        scenario.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        scenario.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建res_scenario
            scenario.POST("/add", scenarioControllers.Create)
            // 更新res_scenario
            scenario.PUT("/edit", scenarioControllers.Update)
            // 删除res_scenario
            scenario.DELETE("/delete", scenarioControllers.Delete)
            // res_scenario列表（分页查询）
            scenario.GET("/list", scenarioControllers.List)
            // 根据ID获取res_scenario信息
            scenario.GET("/:id", scenarioControllers.GetByID)
        }

        app.ZapLog.Info("res_scenario插件路由注册成功")
    })
}