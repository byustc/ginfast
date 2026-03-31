package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/resource/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册res_enterprise_member插件路由
    var entMemberControllers = controllers.NewEntMemberController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // res_enterprise_member插件路由组
        entMember := engine.Group("/api/plugins/resource/entmember")
        entMember.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        entMember.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        entMember.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建res_enterprise_member
            entMember.POST("/add", entMemberControllers.Create)
            // 更新res_enterprise_member
            entMember.PUT("/edit", entMemberControllers.Update)
            // 删除res_enterprise_member
            entMember.DELETE("/delete", entMemberControllers.Delete)
            // res_enterprise_member列表（分页查询）
            entMember.GET("/list", entMemberControllers.List)
            // 根据ID获取res_enterprise_member信息
            entMember.GET("/:id", entMemberControllers.GetByID)
        }

        app.ZapLog.Info("res_enterprise_member插件路由注册成功")
    })
}