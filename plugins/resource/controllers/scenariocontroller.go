package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/resource/models"
	"gin-fast/plugins/resource/service"

	"github.com/gin-gonic/gin"
)

// ScenarioController res_scenario控制器
type ScenarioController struct {
	controllers.Common
	ScenarioService *service.ScenarioService
}

// NewScenarioController 创建res_scenario控制器
func NewScenarioController() *ScenarioController {
	return &ScenarioController{
		ScenarioService: service.NewScenarioService(),
	}
}

// Create 创建res_scenario
func (c *ScenarioController) Create(ctx *gin.Context) {
	var req models.ScenarioCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	scenario, err := c.ScenarioService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建res_scenario失败", err)
	}

	c.Success(ctx, gin.H{
		"id": scenario.Id,
	})
}

// Update 更新res_scenario
func (c *ScenarioController) Update(ctx *gin.Context) {
	var req models.ScenarioUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ScenarioService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新res_scenario失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除res_scenario
func (c *ScenarioController) Delete(ctx *gin.Context) {
	var req models.ScenarioDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.ScenarioService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除res_scenario失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取res_scenario信息
func (c *ScenarioController) GetByID(ctx *gin.Context) {
	var req models.ScenarioGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	scenario, err := c.ScenarioService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "res_scenario不存在", err)
	}

	c.Success(ctx, scenario)
}

// List res_scenario列表（分页查询）
func (c *ScenarioController) List(ctx *gin.Context) {
	var req models.ScenarioListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	scenarioList, total, err := c.ScenarioService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取res_scenario列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  scenarioList,
		"total": total,
	})
}