package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/biz/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册biz_project_res插件路由
    var projectResControllers = controllers.NewProjectResController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // biz_project_res插件路由组
        projectRes := engine.Group("/api/plugins/biz/projectres")
        projectRes.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        projectRes.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        projectRes.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建biz_project_res
            projectRes.POST("/add", projectResControllers.Create)
            // 更新biz_project_res
            projectRes.PUT("/edit", projectResControllers.Update)
            // 删除biz_project_res
            projectRes.DELETE("/delete", projectResControllers.Delete)
            // biz_project_res列表（分页查询）
            projectRes.GET("/list", projectResControllers.List)
            // 根据ID获取biz_project_res信息
            projectRes.GET("/:id", projectResControllers.GetByID)
        }

        app.ZapLog.Info("biz_project_res插件路由注册成功")
    })
}