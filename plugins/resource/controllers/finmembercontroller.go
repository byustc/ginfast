package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/resource/models"
	"gin-fast/plugins/resource/service"

	"github.com/gin-gonic/gin"
)

// FinMemberController res_finance_member控制器
type FinMemberController struct {
	controllers.Common
	FinMemberService *service.FinMemberService
}

// NewFinMemberController 创建res_finance_member控制器
func NewFinMemberController() *FinMemberController {
	return &FinMemberController{
		FinMemberService: service.NewFinMemberService(),
	}
}

// Create 创建res_finance_member
func (c *FinMemberController) Create(ctx *gin.Context) {
	var req models.FinMemberCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	finMember, err := c.FinMemberService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建res_finance_member失败", err)
	}

	c.Success(ctx, gin.H{
		"id": finMember.Id,
	})
}

// Update 更新res_finance_member
func (c *FinMemberController) Update(ctx *gin.Context) {
	var req models.FinMemberUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.FinMemberService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新res_finance_member失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除res_finance_member
func (c *FinMemberController) Delete(ctx *gin.Context) {
	var req models.FinMemberDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.FinMemberService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除res_finance_member失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取res_finance_member信息
func (c *FinMemberController) GetByID(ctx *gin.Context) {
	var req models.FinMemberGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	finMember, err := c.FinMemberService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "res_finance_member不存在", err)
	}

	c.Success(ctx, finMember)
}

// List res_finance_member列表（分页查询）
func (c *FinMemberController) List(ctx *gin.Context) {
	var req models.FinMemberListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	finMemberList, total, err := c.FinMemberService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取res_finance_member列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  finMemberList,
		"total": total,
	})
}