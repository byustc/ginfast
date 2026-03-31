package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/resource/models"
	"gin-fast/plugins/resource/service"

	"github.com/gin-gonic/gin"
)

// ThitankPlatController res_thinktank_plat控制器
type ThitankPlatController struct {
	controllers.Common
	ThitankPlatService *service.ThitankPlatService
}

// NewThitankPlatController 创建res_thinktank_plat控制器
func NewThitankPlatController() *ThitankPlatController {
	return &ThitankPlatController{
		ThitankPlatService: service.NewThitankPlatService(),
	}
}

// Create 创建res_thinktank_plat
func (c *ThitankPlatController) Create(ctx *gin.Context) {
	var req models.ThitankPlatCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	thitankPlat, err := c.ThitankPlatService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建res_thinktank_plat失败", err)
	}

	c.Success(ctx, gin.H{
		"id": thitankPlat.Id,
	})
}

// Update 更新res_thinktank_plat
func (c *ThitankPlatController) Update(ctx *gin.Context) {
	var req models.ThitankPlatUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ThitankPlatService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新res_thinktank_plat失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除res_thinktank_plat
func (c *ThitankPlatController) Delete(ctx *gin.Context) {
	var req models.ThitankPlatDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ThitankPlatService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除res_thinktank_plat失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取res_thinktank_plat信息
func (c *ThitankPlatController) GetByID(ctx *gin.Context) {
	var req models.ThitankPlatGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	thitankPlat, err := c.ThitankPlatService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "res_thinktank_plat不存在", err)
	}

	c.Success(ctx, thitankPlat)
}

// List res_thinktank_plat列表（分页查询）
func (c *ThitankPlatController) List(ctx *gin.Context) {
	var req models.ThitankPlatListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	thitankPlatList, total, err := c.ThitankPlatService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取res_thinktank_plat列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  thitankPlatList,
		"total": total,
	})
}