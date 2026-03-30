package routes

import (
	"gin-fast/app/global/app"
	"gin-fast/app/middleware"
	"gin-fast/plugins/resource/controllers"
	"gin-fast/app/utils/ginhelper"
	"github.com/gin-gonic/gin"
)

func init() {
    // RegisterRoutes 注册res_district插件路由
    var districtControllers = controllers.NewDistrictController()
	ginhelper.RegisterPluginRoutes(func (engine *gin.Engine) {
        // res_district插件路由组
        district := engine.Group("/api/plugins/resource/district")
        district.Use(middleware.JWTAuthMiddleware())     // 认证中间件
        district.Use(middleware.DemoAccountMiddleware()) // 添加演示账号中间件
        district.Use(middleware.CasbinMiddleware())      // 权限中间件
        {
            // 创建res_district
            district.POST("/add", districtControllers.Create)
            // 更新res_district
            district.PUT("/edit", districtControllers.Update)
            // 删除res_district
            district.DELETE("/delete", districtControllers.Delete)
            // res_district列表（分页查询）
            district.GET("/list", districtControllers.List)
            // 根据ID获取res_district信息
            district.GET("/:id", districtControllers.GetByID)
        }

        app.ZapLog.Info("res_district插件路由注册成功")
    })
}