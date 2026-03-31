package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/resource/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册res_finance_member插件路由
    var finMemberControllers = controllers.NewFinMemberController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // res_finance_member插件路由组
        finMember := engine.Group("/api/plugins/resource/finmember")
        finMember.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        finMember.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        finMember.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建res_finance_member
            finMember.POST("/add", finMemberControllers.Create)
            // 更新res_finance_member
            finMember.PUT("/edit", finMemberControllers.Update)
            // 删除res_finance_member
            finMember.DELETE("/delete", finMemberControllers.Delete)
            // res_finance_member列表（分页查询）
            finMember.GET("/list", finMemberControllers.List)
            // 根据ID获取res_finance_member信息
            finMember.GET("/:id", finMemberControllers.GetByID)
        }

        app.ZapLog.Info("res_finance_member插件路由注册成功")
    })
}