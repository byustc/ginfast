package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/resource/models"
	"gin-fast/plugins/resource/service"

	"github.com/gin-gonic/gin"
)

// PolicyController res_policy控制器
type PolicyController struct {
	controllers.Common
	PolicyService *service.PolicyService
}

// NewPolicyController 创建res_policy控制器
func NewPolicyController() *PolicyController {
	return &PolicyController{
		PolicyService: service.NewPolicyService(),
	}
}

// Create 创建res_policy
func (c *PolicyController) Create(ctx *gin.Context) {
	var req models.PolicyCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	policy, err := c.PolicyService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建res_policy失败", err)
	}

	c.Success(ctx, gin.H{
		"id": policy.Id,
	})
}

// Update 更新res_policy
func (c *PolicyController) Update(ctx *gin.Context) {
	var req models.PolicyUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.PolicyService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新res_policy失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除res_policy
func (c *PolicyController) Delete(ctx *gin.Context) {
	var req models.PolicyDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.PolicyService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除res_policy失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取res_policy信息
func (c *PolicyController) GetByID(ctx *gin.Context) {
	var req models.PolicyGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	policy, err := c.PolicyService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "res_policy不存在", err)
	}

	c.Success(ctx, policy)
}

// List res_policy列表（分页查询）
func (c *PolicyController) List(ctx *gin.Context) {
	var req models.PolicyListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	policyList, total, err := c.PolicyService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取res_policy列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  policyList,
		"total": total,
	})
}