package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/resource/models"
	"gin-fast/plugins/resource/service"

	"github.com/gin-gonic/gin"
)

// ExpertController res_expert控制器
type ExpertController struct {
	controllers.Common
	ExpertService *service.ExpertService
}

// NewExpertController 创建res_expert控制器
func NewExpertController() *ExpertController {
	return &ExpertController{
		ExpertService: service.NewExpertService(),
	}
}

// Create 创建res_expert
func (c *ExpertController) Create(ctx *gin.Context) {
	var req models.ExpertCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	expert, err := c.ExpertService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建res_expert失败", err)
	}

	c.Success(ctx, gin.H{
		"id": expert.Id,
	})
}

// Update 更新res_expert
func (c *ExpertController) Update(ctx *gin.Context) {
	var req models.ExpertUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ExpertService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新res_expert失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除res_expert
func (c *ExpertController) Delete(ctx *gin.Context) {
	var req models.ExpertDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ExpertService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除res_expert失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取res_expert信息
func (c *ExpertController) GetByID(ctx *gin.Context) {
	var req models.ExpertGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	expert, err := c.ExpertService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "res_expert不存在", err)
	}

	c.Success(ctx, expert)
}

// List res_expert列表（分页查询）
func (c *ExpertController) List(ctx *gin.Context) {
	var req models.ExpertListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	expertList, total, err := c.ExpertService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取res_expert列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  expertList,
		"total": total,
	})
}