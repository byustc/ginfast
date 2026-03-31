package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/resource/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册res_policy插件路由
    var policyControllers = controllers.NewPolicyController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // res_policy插件路由组
        policy := engine.Group("/api/plugins/resource/policy")
        policy.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        policy.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        policy.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建res_policy
            policy.POST("/add", policyControllers.Create)
            // 更新res_policy
            policy.PUT("/edit", policyControllers.Update)
            // 删除res_policy
            policy.DELETE("/delete", policyControllers.Delete)
            // res_policy列表（分页查询）
            policy.GET("/list", policyControllers.List)
            // 根据ID获取res_policy信息
            policy.GET("/:id", policyControllers.GetByID)
        }

        app.ZapLog.Info("res_policy插件路由注册成功")
    })
}