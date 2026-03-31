package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/resource/models"
	"gin-fast/plugins/resource/service"

	"github.com/gin-gonic/gin"
)

// InvMemberController res_invest_member控制器
type InvMemberController struct {
	controllers.Common
	InvMemberService *service.InvMemberService
}

// NewInvMemberController 创建res_invest_member控制器
func NewInvMemberController() *InvMemberController {
	return &InvMemberController{
		InvMemberService: service.NewInvMemberService(),
	}
}

// Create 创建res_invest_member
func (c *InvMemberController) Create(ctx *gin.Context) {
	var req models.InvMemberCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	invMember, err := c.InvMemberService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建res_invest_member失败", err)
	}

	c.Success(ctx, gin.H{
		"id": invMember.Id,
	})
}

// Update 更新res_invest_member
func (c *InvMemberController) Update(ctx *gin.Context) {
	var req models.InvMemberUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.InvMemberService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新res_invest_member失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除res_invest_member
func (c *InvMemberController) Delete(ctx *gin.Context) {
	var req models.InvMemberDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.InvMemberService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除res_invest_member失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取res_invest_member信息
func (c *InvMemberController) GetByID(ctx *gin.Context) {
	var req models.InvMemberGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	invMember, err := c.InvMemberService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "res_invest_member不存在", err)
	}

	c.Success(ctx, invMember)
}

// List res_invest_member列表（分页查询）
func (c *InvMemberController) List(ctx *gin.Context) {
	var req models.InvMemberListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	invMemberList, total, err := c.InvMemberService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取res_invest_member列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  invMemberList,
		"total": total,
	})
}