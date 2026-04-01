package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/biz/models"
	"gin-fast/plugins/biz/service"

	"github.com/gin-gonic/gin"
)

// ProjectPartnerController biz_project_partner控制器
type ProjectPartnerController struct {
	controllers.Common
	ProjectPartnerService *service.ProjectPartnerService
}

// NewProjectPartnerController 创建biz_project_partner控制器
func NewProjectPartnerController() *ProjectPartnerController {
	return &ProjectPartnerController{
		ProjectPartnerService: service.NewProjectPartnerService(),
	}
}

// Create 创建biz_project_partner
func (c *ProjectPartnerController) Create(ctx *gin.Context) {
	var req models.ProjectPartnerCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	projectPartner, err := c.ProjectPartnerService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建biz_project_partner失败", err)
	}

	c.Success(ctx, gin.H{
		"id": projectPartner.Id,
	})
}

// Update 更新biz_project_partner
func (c *ProjectPartnerController) Update(ctx *gin.Context) {
	var req models.ProjectPartnerUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ProjectPartnerService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新biz_project_partner失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除biz_project_partner
func (c *ProjectPartnerController) Delete(ctx *gin.Context) {
	var req models.ProjectPartnerDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ProjectPartnerService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除biz_project_partner失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取biz_project_partner信息
func (c *ProjectPartnerController) GetByID(ctx *gin.Context) {
	var req models.ProjectPartnerGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	projectPartner, err := c.ProjectPartnerService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "biz_project_partner不存在", err)
	}

	c.Success(ctx, projectPartner)
}

// List biz_project_partner列表（分页查询）
func (c *ProjectPartnerController) List(ctx *gin.Context) {
	var req models.ProjectPartnerListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	projectPartnerList, total, err := c.ProjectPartnerService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取biz_project_partner列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  projectPartnerList,
		"total": total,
	})
}