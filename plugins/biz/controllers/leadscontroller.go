package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/biz/models"
	"gin-fast/plugins/biz/service"

	"github.com/gin-gonic/gin"
)

// LeadsController biz_leads控制器
type LeadsController struct {
	controllers.Common
	LeadsService *service.LeadsService
}

// NewLeadsController 创建biz_leads控制器
func NewLeadsController() *LeadsController {
	return &LeadsController{
		LeadsService: service.NewLeadsService(),
	}
}

// Create 创建biz_leads
func (c *LeadsController) Create(ctx *gin.Context) {
	var req models.LeadsCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	leads, err := c.LeadsService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建biz_leads失败", err)
	}

	c.Success(ctx, gin.H{
		"id": leads.Id,
	})
}

// Update 更新biz_leads
func (c *LeadsController) Update(ctx *gin.Context) {
	var req models.LeadsUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.LeadsService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新biz_leads失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除biz_leads
func (c *LeadsController) Delete(ctx *gin.Context) {
	var req models.LeadsDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.LeadsService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除biz_leads失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取biz_leads信息
func (c *LeadsController) GetByID(ctx *gin.Context) {
	var req models.LeadsGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	leads, err := c.LeadsService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "biz_leads不存在", err)
	}

	c.Success(ctx, leads)
}

// List biz_leads列表（分页查询）
func (c *LeadsController) List(ctx *gin.Context) {
	var req models.LeadsListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	leadsList, total, err := c.LeadsService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取biz_leads列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  leadsList,
		"total": total,
	})
}