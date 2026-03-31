package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/resource/models"
	"gin-fast/plugins/resource/service"

	"github.com/gin-gonic/gin"
)

// ThitankFundController res_thinktank_fund控制器
type ThitankFundController struct {
	controllers.Common
	ThitankFundService *service.ThitankFundService
}

// NewThitankFundController 创建res_thinktank_fund控制器
func NewThitankFundController() *ThitankFundController {
	return &ThitankFundController{
		ThitankFundService: service.NewThitankFundService(),
	}
}

// Create 创建res_thinktank_fund
func (c *ThitankFundController) Create(ctx *gin.Context) {
	var req models.ThitankFundCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	thitankFund, err := c.ThitankFundService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建res_thinktank_fund失败", err)
	}

	c.Success(ctx, gin.H{
		"id": thitankFund.Id,
	})
}

// Update 更新res_thinktank_fund
func (c *ThitankFundController) Update(ctx *gin.Context) {
	var req models.ThitankFundUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ThitankFundService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新res_thinktank_fund失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除res_thinktank_fund
func (c *ThitankFundController) Delete(ctx *gin.Context) {
	var req models.ThitankFundDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ThitankFundService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除res_thinktank_fund失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取res_thinktank_fund信息
func (c *ThitankFundController) GetByID(ctx *gin.Context) {
	var req models.ThitankFundGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	thitankFund, err := c.ThitankFundService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "res_thinktank_fund不存在", err)
	}

	c.Success(ctx, thitankFund)
}

// List res_thinktank_fund列表（分页查询）
func (c *ThitankFundController) List(ctx *gin.Context) {
	var req models.ThitankFundListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	thitankFundList, total, err := c.ThitankFundService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取res_thinktank_fund列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  thitankFundList,
		"total": total,
	})
}