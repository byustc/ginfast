package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/resource/models"
	"gin-fast/plugins/resource/service"

	"github.com/gin-gonic/gin"
)

// ThitankMemberController res_thinktank_member控制器
type ThitankMemberController struct {
	controllers.Common
	ThitankMemberService *service.ThitankMemberService
}

// NewThitankMemberController 创建res_thinktank_member控制器
func NewThitankMemberController() *ThitankMemberController {
	return &ThitankMemberController{
		ThitankMemberService: service.NewThitankMemberService(),
	}
}

// Create 创建res_thinktank_member
func (c *ThitankMemberController) Create(ctx *gin.Context) {
	var req models.ThitankMemberCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	thitankMember, err := c.ThitankMemberService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建res_thinktank_member失败", err)
	}

	c.Success(ctx, gin.H{
		"id": thitankMember.Id,
	})
}

// Update 更新res_thinktank_member
func (c *ThitankMemberController) Update(ctx *gin.Context) {
	var req models.ThitankMemberUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ThitankMemberService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新res_thinktank_member失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除res_thinktank_member
func (c *ThitankMemberController) Delete(ctx *gin.Context) {
	var req models.ThitankMemberDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ThitankMemberService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除res_thinktank_member失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取res_thinktank_member信息
func (c *ThitankMemberController) GetByID(ctx *gin.Context) {
	var req models.ThitankMemberGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	thitankMember, err := c.ThitankMemberService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "res_thinktank_member不存在", err)
	}

	c.Success(ctx, thitankMember)
}

// List res_thinktank_member列表（分页查询）
func (c *ThitankMemberController) List(ctx *gin.Context) {
	var req models.ThitankMemberListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	thitankMemberList, total, err := c.ThitankMemberService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取res_thinktank_member列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  thitankMemberList,
		"total": total,
	})
}