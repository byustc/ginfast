package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/resource/models"
	"gin-fast/plugins/resource/service"

	"github.com/gin-gonic/gin"
)

// IndustryChainController res_industry_chain控制器
type IndustryChainController struct {
	controllers.Common
	IndustryChainService *service.IndustryChainService
}

// NewIndustryChainController 创建res_industry_chain控制器
func NewIndustryChainController() *IndustryChainController {
	return &IndustryChainController{
		IndustryChainService: service.NewIndustryChainService(),
	}
}

// Create 创建res_industry_chain
func (c *IndustryChainController) Create(ctx *gin.Context) {
	var req models.IndustryChainCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	industryChain, err := c.IndustryChainService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建res_industry_chain失败", err)
	}

	c.Success(ctx, gin.H{
		"id": industryChain.Id,
	})
}

// Update 更新res_industry_chain
func (c *IndustryChainController) Update(ctx *gin.Context) {
	var req models.IndustryChainUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.IndustryChainService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新res_industry_chain失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除res_industry_chain
func (c *IndustryChainController) Delete(ctx *gin.Context) {
	var req models.IndustryChainDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.IndustryChainService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除res_industry_chain失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取res_industry_chain信息
func (c *IndustryChainController) GetByID(ctx *gin.Context) {
	var req models.IndustryChainGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	industryChain, err := c.IndustryChainService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "res_industry_chain不存在", err)
	}

	c.Success(ctx, industryChain)
}

// List res_industry_chain列表（分页查询）
func (c *IndustryChainController) List(ctx *gin.Context) {
	var req models.IndustryChainListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	industryChainList, total, err := c.IndustryChainService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取res_industry_chain列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  industryChainList,
		"total": total,
	})
}