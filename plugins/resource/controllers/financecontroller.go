package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/resource/models"
	"gin-fast/plugins/resource/service"

	"github.com/gin-gonic/gin"
)

// FinanceController res_finance控制器
type FinanceController struct {
	controllers.Common
	FinanceService *service.FinanceService
}

// NewFinanceController 创建res_finance控制器
func NewFinanceController() *FinanceController {
	return &FinanceController{
		FinanceService: service.NewFinanceService(),
	}
}

// Create 创建res_finance
func (c *FinanceController) Create(ctx *gin.Context) {
	var req models.FinanceCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	finance, err := c.FinanceService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建res_finance失败", err)
	}

	c.Success(ctx, gin.H{
		"id": finance.Id,
	})
}

// Update 更新res_finance
func (c *FinanceController) Update(ctx *gin.Context) {
	var req models.FinanceUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.FinanceService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新res_finance失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除res_finance
func (c *FinanceController) Delete(ctx *gin.Context) {
	var req models.FinanceDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.FinanceService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除res_finance失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取res_finance信息
func (c *FinanceController) GetByID(ctx *gin.Context) {
	var req models.FinanceGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	finance, err := c.FinanceService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "res_finance不存在", err)
	}

	c.Success(ctx, finance)
}

// List res_finance列表（分页查询）
func (c *FinanceController) List(ctx *gin.Context) {
	var req models.FinanceListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	financeList, total, err := c.FinanceService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取res_finance列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  financeList,
		"total": total,
	})
}