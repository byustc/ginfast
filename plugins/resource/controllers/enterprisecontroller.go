package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/resource/models"
	"gin-fast/plugins/resource/service"

	"github.com/gin-gonic/gin"
)

// EnterpriseController res_enterprise控制器
type EnterpriseController struct {
	controllers.Common
	EnterpriseService *service.EnterpriseService
}

// NewEnterpriseController 创建res_enterprise控制器
func NewEnterpriseController() *EnterpriseController {
	return &EnterpriseController{
		EnterpriseService: service.NewEnterpriseService(),
	}
}

// Create 创建res_enterprise
func (c *EnterpriseController) Create(ctx *gin.Context) {
	var req models.EnterpriseCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	enterprise, err := c.EnterpriseService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建res_enterprise失败", err)
	}

	c.Success(ctx, gin.H{
		"id": enterprise.Id,
	})
}

// Update 更新res_enterprise
func (c *EnterpriseController) Update(ctx *gin.Context) {
	var req models.EnterpriseUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.EnterpriseService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新res_enterprise失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除res_enterprise
func (c *EnterpriseController) Delete(ctx *gin.Context) {
	var req models.EnterpriseDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.EnterpriseService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除res_enterprise失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取res_enterprise信息
func (c *EnterpriseController) GetByID(ctx *gin.Context) {
	var req models.EnterpriseGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	enterprise, err := c.EnterpriseService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "res_enterprise不存在", err)
	}

	c.Success(ctx, enterprise)
}

// List res_enterprise列表（分页查询）
func (c *EnterpriseController) List(ctx *gin.Context) {
	var req models.EnterpriseListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	enterpriseList, total, err := c.EnterpriseService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取res_enterprise列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  enterpriseList,
		"total": total,
	})
}