package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/resource/models"
	"gin-fast/plugins/resource/service"

	"github.com/gin-gonic/gin"
)

// EntMemberController res_enterprise_member控制器
type EntMemberController struct {
	controllers.Common
	EntMemberService *service.EntMemberService
}

// NewEntMemberController 创建res_enterprise_member控制器
func NewEntMemberController() *EntMemberController {
	return &EntMemberController{
		EntMemberService: service.NewEntMemberService(),
	}
}

// Create 创建res_enterprise_member
func (c *EntMemberController) Create(ctx *gin.Context) {
	var req models.EntMemberCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	entMember, err := c.EntMemberService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建res_enterprise_member失败", err)
	}

	c.Success(ctx, gin.H{
		"id": entMember.Id,
	})
}

// Update 更新res_enterprise_member
func (c *EntMemberController) Update(ctx *gin.Context) {
	var req models.EntMemberUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.EntMemberService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新res_enterprise_member失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除res_enterprise_member
func (c *EntMemberController) Delete(ctx *gin.Context) {
	var req models.EntMemberDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.EntMemberService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除res_enterprise_member失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取res_enterprise_member信息
func (c *EntMemberController) GetByID(ctx *gin.Context) {
	var req models.EntMemberGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	entMember, err := c.EntMemberService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "res_enterprise_member不存在", err)
	}

	c.Success(ctx, entMember)
}

// List res_enterprise_member列表（分页查询）
func (c *EntMemberController) List(ctx *gin.Context) {
	var req models.EntMemberListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	entMemberList, total, err := c.EntMemberService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取res_enterprise_member列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  entMemberList,
		"total": total,
	})
}