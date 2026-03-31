package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/resource/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册res_thinktank_member插件路由
    var thitankMemberControllers = controllers.NewThitankMemberController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // res_thinktank_member插件路由组
        thitankMember := engine.Group("/api/plugins/resource/thitankmember")
        thitankMember.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        thitankMember.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        thitankMember.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建res_thinktank_member
            thitankMember.POST("/add", thitankMemberControllers.Create)
            // 更新res_thinktank_member
            thitankMember.PUT("/edit", thitankMemberControllers.Update)
            // 删除res_thinktank_member
            thitankMember.DELETE("/delete", thitankMemberControllers.Delete)
            // res_thinktank_member列表（分页查询）
            thitankMember.GET("/list", thitankMemberControllers.List)
            // 根据ID获取res_thinktank_member信息
            thitankMember.GET("/:id", thitankMemberControllers.GetByID)
        }

        app.ZapLog.Info("res_thinktank_member插件路由注册成功")
    })
}