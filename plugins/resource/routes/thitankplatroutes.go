package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/resource/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册res_thinktank_plat插件路由
    var thitankPlatControllers = controllers.NewThitankPlatController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // res_thinktank_plat插件路由组
        thitankPlat := engine.Group("/api/plugins/resource/thitankplat")
        thitankPlat.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        thitankPlat.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        thitankPlat.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建res_thinktank_plat
            thitankPlat.POST("/add", thitankPlatControllers.Create)
            // 更新res_thinktank_plat
            thitankPlat.PUT("/edit", thitankPlatControllers.Update)
            // 删除res_thinktank_plat
            thitankPlat.DELETE("/delete", thitankPlatControllers.Delete)
            // res_thinktank_plat列表（分页查询）
            thitankPlat.GET("/list", thitankPlatControllers.List)
            // 根据ID获取res_thinktank_plat信息
            thitankPlat.GET("/:id", thitankPlatControllers.GetByID)
        }

        app.ZapLog.Info("res_thinktank_plat插件路由注册成功")
    })
}