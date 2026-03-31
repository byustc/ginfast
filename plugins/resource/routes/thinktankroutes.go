package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/resource/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册res_thinktank插件路由
    var thinktankControllers = controllers.NewThinktankController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // res_thinktank插件路由组
        thinktank := engine.Group("/api/plugins/resource/thinktank")
        thinktank.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        thinktank.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        thinktank.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建res_thinktank
            thinktank.POST("/add", thinktankControllers.Create)
            // 更新res_thinktank
            thinktank.PUT("/edit", thinktankControllers.Update)
            // 删除res_thinktank
            thinktank.DELETE("/delete", thinktankControllers.Delete)
            // res_thinktank列表（分页查询）
            thinktank.GET("/list", thinktankControllers.List)
            // 根据ID获取res_thinktank信息
            thinktank.GET("/:id", thinktankControllers.GetByID)
        }

        app.ZapLog.Info("res_thinktank插件路由注册成功")
    })
}