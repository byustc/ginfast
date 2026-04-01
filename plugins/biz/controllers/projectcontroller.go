package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/biz/models"
	"gin-fast/plugins/biz/service"

	"github.com/gin-gonic/gin"
)

// ProjectController biz_project控制器
type ProjectController struct {
	controllers.Common
	ProjectService *service.ProjectService
}

// NewProjectController 创建biz_project控制器
func NewProjectController() *ProjectController {
	return &ProjectController{
		ProjectService: service.NewProjectService(),
	}
}

// Create 创建biz_project
func (c *ProjectController) Create(ctx *gin.Context) {
	var req models.ProjectCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	project, err := c.ProjectService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建biz_project失败", err)
	}

	c.Success(ctx, gin.H{
		"id": project.Id,
	})
}

// Update 更新biz_project
func (c *ProjectController) Update(ctx *gin.Context) {
	var req models.ProjectUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ProjectService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新biz_project失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除biz_project
func (c *ProjectController) Delete(ctx *gin.Context) {
	var req models.ProjectDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ProjectService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除biz_project失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取biz_project信息
func (c *ProjectController) GetByID(ctx *gin.Context) {
	var req models.ProjectGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	project, err := c.ProjectService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "biz_project不存在", err)
	}

	c.Success(ctx, project)
}

// List biz_project列表（分页查询）
func (c *ProjectController) List(ctx *gin.Context) {
	var req models.ProjectListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	projectList, total, err := c.ProjectService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取biz_project列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  projectList,
		"total": total,
	})
}