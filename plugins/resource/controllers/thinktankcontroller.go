package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/resource/models"
	"gin-fast/plugins/resource/service"

	"github.com/gin-gonic/gin"
)

// ThinktankController res_thinktank控制器
type ThinktankController struct {
	controllers.Common
	ThinktankService *service.ThinktankService
}

// NewThinktankController 创建res_thinktank控制器
func NewThinktankController() *ThinktankController {
	return &ThinktankController{
		ThinktankService: service.NewThinktankService(),
	}
}

// Create 创建res_thinktank
func (c *ThinktankController) Create(ctx *gin.Context) {
	var req models.ThinktankCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	thinktank, err := c.ThinktankService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建res_thinktank失败", err)
	}

	c.Success(ctx, gin.H{
		"id": thinktank.Id,
	})
}

// Update 更新res_thinktank
func (c *ThinktankController) Update(ctx *gin.Context) {
	var req models.ThinktankUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ThinktankService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新res_thinktank失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除res_thinktank
func (c *ThinktankController) Delete(ctx *gin.Context) {
	var req models.ThinktankDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ThinktankService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除res_thinktank失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取res_thinktank信息
func (c *ThinktankController) GetByID(ctx *gin.Context) {
	var req models.ThinktankGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	thinktank, err := c.ThinktankService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "res_thinktank不存在", err)
	}

	c.Success(ctx, thinktank)
}

// List res_thinktank列表（分页查询）
func (c *ThinktankController) List(ctx *gin.Context) {
	var req models.ThinktankListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	thinktankList, total, err := c.ThinktankService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取res_thinktank列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  thinktankList,
		"total": total,
	})
}