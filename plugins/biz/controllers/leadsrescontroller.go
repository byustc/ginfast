package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/biz/models"
	"gin-fast/plugins/biz/service"

	"github.com/gin-gonic/gin"
)

// LeadsResController biz_leads_res控制器
type LeadsResController struct {
	controllers.Common
	LeadsResService *service.LeadsResService
}

// NewLeadsResController 创建biz_leads_res控制器
func NewLeadsResController() *LeadsResController {
	return &LeadsResController{
		LeadsResService: service.NewLeadsResService(),
	}
}

// Create 创建biz_leads_res
func (c *LeadsResController) Create(ctx *gin.Context) {
	var req models.LeadsResCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	leadsRes, err := c.LeadsResService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建biz_leads_res失败", err)
	}

	c.Success(ctx, gin.H{
		"id": leadsRes.Id,
	})
}

// Update 更新biz_leads_res
func (c *LeadsResController) Update(ctx *gin.Context) {
	var req models.LeadsResUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.LeadsResService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新biz_leads_res失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除biz_leads_res
func (c *LeadsResController) Delete(ctx *gin.Context) {
	var req models.LeadsResDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.LeadsResService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除biz_leads_res失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取biz_leads_res信息
func (c *LeadsResController) GetByID(ctx *gin.Context) {
	var req models.LeadsResGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	leadsRes, err := c.LeadsResService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "biz_leads_res不存在", err)
	}

	c.Success(ctx, leadsRes)
}

// List biz_leads_res列表（分页查询）
func (c *LeadsResController) List(ctx *gin.Context) {
	var req models.LeadsResListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	leadsResList, total, err := c.LeadsResService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取biz_leads_res列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  leadsResList,
		"total": total,
	})
}