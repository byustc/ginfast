package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/biz/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册biz_project_phase_plan插件路由
    var projectPhasePlanControllers = controllers.NewProjectPhasePlanController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // biz_project_phase_plan插件路由组
        projectPhasePlan := engine.Group("/api/plugins/biz/projectphaseplan")
        projectPhasePlan.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        projectPhasePlan.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        projectPhasePlan.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建biz_project_phase_plan
            projectPhasePlan.POST("/add", projectPhasePlanControllers.Create)
            // 更新biz_project_phase_plan
            projectPhasePlan.PUT("/edit", projectPhasePlanControllers.Update)
            // 删除biz_project_phase_plan
            projectPhasePlan.DELETE("/delete", projectPhasePlanControllers.Delete)
            // biz_project_phase_plan列表（分页查询）
            projectPhasePlan.GET("/list", projectPhasePlanControllers.List)
            // 根据ID获取biz_project_phase_plan信息
            projectPhasePlan.GET("/:id", projectPhasePlanControllers.GetByID)
        }

        app.ZapLog.Info("biz_project_phase_plan插件路由注册成功")
    })
}