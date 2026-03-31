package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/resource/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册res_invest_member插件路由
    var invMemberControllers = controllers.NewInvMemberController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // res_invest_member插件路由组
        invMember := engine.Group("/api/plugins/resource/invmember")
        invMember.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        invMember.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        invMember.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建res_invest_member
            invMember.POST("/add", invMemberControllers.Create)
            // 更新res_invest_member
            invMember.PUT("/edit", invMemberControllers.Update)
            // 删除res_invest_member
            invMember.DELETE("/delete", invMemberControllers.Delete)
            // res_invest_member列表（分页查询）
            invMember.GET("/list", invMemberControllers.List)
            // 根据ID获取res_invest_member信息
            invMember.GET("/:id", invMemberControllers.GetByID)
        }

        app.ZapLog.Info("res_invest_member插件路由注册成功")
    })
}