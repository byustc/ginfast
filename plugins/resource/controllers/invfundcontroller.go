package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/resource/models"
	"gin-fast/plugins/resource/service"

	"github.com/gin-gonic/gin"
)

// InvFundController res_invest_fund控制器
type InvFundController struct {
	controllers.Common
	InvFundService *service.InvFundService
}

// NewInvFundController 创建res_invest_fund控制器
func NewInvFundController() *InvFundController {
	return &InvFundController{
		InvFundService: service.NewInvFundService(),
	}
}

// Create 创建res_invest_fund
func (c *InvFundController) Create(ctx *gin.Context) {
	var req models.InvFundCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	invFund, err := c.InvFundService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建res_invest_fund失败", err)
	}

	c.Success(ctx, gin.H{
		"id": invFund.Id,
	})
}

// Update 更新res_invest_fund
func (c *InvFundController) Update(ctx *gin.Context) {
	var req models.InvFundUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.InvFundService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新res_invest_fund失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除res_invest_fund
func (c *InvFundController) Delete(ctx *gin.Context) {
	var req models.InvFundDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.InvFundService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除res_invest_fund失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取res_invest_fund信息
func (c *InvFundController) GetByID(ctx *gin.Context) {
	var req models.InvFundGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	invFund, err := c.InvFundService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "res_invest_fund不存在", err)
	}

	c.Success(ctx, invFund)
}

// List res_invest_fund列表（分页查询）
func (c *InvFundController) List(ctx *gin.Context) {
	var req models.InvFundListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	invFundList, total, err := c.InvFundService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取res_invest_fund列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  invFundList,
		"total": total,
	})
}