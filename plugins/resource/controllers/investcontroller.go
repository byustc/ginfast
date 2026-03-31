package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/resource/models"
	"gin-fast/plugins/resource/service"

	"github.com/gin-gonic/gin"
)

// InvestController res_invest控制器
type InvestController struct {
	controllers.Common
	InvestService *service.InvestService
}

// NewInvestController 创建res_invest控制器
func NewInvestController() *InvestController {
	return &InvestController{
		InvestService: service.NewInvestService(),
	}
}

// Create 创建res_invest
func (c *InvestController) Create(ctx *gin.Context) {
	var req models.InvestCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	invest, err := c.InvestService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建res_invest失败", err)
	}

	c.Success(ctx, gin.H{
		"id": invest.Id,
	})
}

// Update 更新res_invest
func (c *InvestController) Update(ctx *gin.Context) {
	var req models.InvestUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.InvestService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新res_invest失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除res_invest
func (c *InvestController) Delete(ctx *gin.Context) {
	var req models.InvestDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.InvestService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除res_invest失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取res_invest信息
func (c *InvestController) GetByID(ctx *gin.Context) {
	var req models.InvestGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	invest, err := c.InvestService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "res_invest不存在", err)
	}

	c.Success(ctx, invest)
}

// List res_invest列表（分页查询）
func (c *InvestController) List(ctx *gin.Context) {
	var req models.InvestListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	investList, total, err := c.InvestService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取res_invest列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  investList,
		"total": total,
	})
}