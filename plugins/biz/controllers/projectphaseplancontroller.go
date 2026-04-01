package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/biz/models"
	"gin-fast/plugins/biz/service"

	"github.com/gin-gonic/gin"
)

// ProjectPhasePlanController biz_project_phase_plan控制器
type ProjectPhasePlanController struct {
	controllers.Common
	ProjectPhasePlanService *service.ProjectPhasePlanService
}

// NewProjectPhasePlanController 创建biz_project_phase_plan控制器
func NewProjectPhasePlanController() *ProjectPhasePlanController {
	return &ProjectPhasePlanController{
		ProjectPhasePlanService: service.NewProjectPhasePlanService(),
	}
}

// Create 创建biz_project_phase_plan
func (c *ProjectPhasePlanController) Create(ctx *gin.Context) {
	var req models.ProjectPhasePlanCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	projectPhasePlan, err := c.ProjectPhasePlanService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建biz_project_phase_plan失败", err)
	}

	c.Success(ctx, gin.H{
		"id": projectPhasePlan.Id,
	})
}

// Update 更新biz_project_phase_plan
func (c *ProjectPhasePlanController) Update(ctx *gin.Context) {
	var req models.ProjectPhasePlanUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ProjectPhasePlanService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新biz_project_phase_plan失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除biz_project_phase_plan
func (c *ProjectPhasePlanController) Delete(ctx *gin.Context) {
	var req models.ProjectPhasePlanDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ProjectPhasePlanService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除biz_project_phase_plan失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取biz_project_phase_plan信息
func (c *ProjectPhasePlanController) GetByID(ctx *gin.Context) {
	var req models.ProjectPhasePlanGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	projectPhasePlan, err := c.ProjectPhasePlanService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "biz_project_phase_plan不存在", err)
	}

	c.Success(ctx, projectPhasePlan)
}

// List biz_project_phase_plan列表（分页查询）
func (c *ProjectPhasePlanController) List(ctx *gin.Context) {
	var req models.ProjectPhasePlanListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	projectPhasePlanList, total, err := c.ProjectPhasePlanService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取biz_project_phase_plan列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  projectPhasePlanList,
		"total": total,
	})
}