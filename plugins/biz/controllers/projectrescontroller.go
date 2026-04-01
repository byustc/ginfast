package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/biz/models"
	"gin-fast/plugins/biz/service"

	"github.com/gin-gonic/gin"
)

// ProjectResController biz_project_res控制器
type ProjectResController struct {
	controllers.Common
	ProjectResService *service.ProjectResService
}

// NewProjectResController 创建biz_project_res控制器
func NewProjectResController() *ProjectResController {
	return &ProjectResController{
		ProjectResService: service.NewProjectResService(),
	}
}

// Create 创建biz_project_res
func (c *ProjectResController) Create(ctx *gin.Context) {
	var req models.ProjectResCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	projectRes, err := c.ProjectResService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建biz_project_res失败", err)
	}

	c.Success(ctx, gin.H{
		"id": projectRes.Id,
	})
}

// Update 更新biz_project_res
func (c *ProjectResController) Update(ctx *gin.Context) {
	var req models.ProjectResUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ProjectResService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新biz_project_res失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除biz_project_res
func (c *ProjectResController) Delete(ctx *gin.Context) {
	var req models.ProjectResDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ProjectResService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除biz_project_res失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取biz_project_res信息
func (c *ProjectResController) GetByID(ctx *gin.Context) {
	var req models.ProjectResGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	projectRes, err := c.ProjectResService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "biz_project_res不存在", err)
	}

	c.Success(ctx, projectRes)
}

// List biz_project_res列表（分页查询）
func (c *ProjectResController) List(ctx *gin.Context) {
	var req models.ProjectResListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	projectResList, total, err := c.ProjectResService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取biz_project_res列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  projectResList,
		"total": total,
	})
}