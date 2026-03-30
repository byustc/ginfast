package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/resource/models"
	"gin-fast/plugins/resource/service"

	"github.com/gin-gonic/gin"
)

// DistrictController res_district控制器
type DistrictController struct {
	controllers.Common
	DistrictService *service.DistrictService
}

// NewDistrictController 创建res_district控制器
func NewDistrictController() *DistrictController {
	return &DistrictController{
		DistrictService: service.NewDistrictService(),
	}
}

// Create 创建res_district
func (c *DistrictController) Create(ctx *gin.Context) {
	var req models.DistrictCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	district, err := c.DistrictService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建res_district失败", err)
	}

	c.Success(ctx, gin.H{
		"id": district.Id,
	})
}

// Update 更新res_district
func (c *DistrictController) Update(ctx *gin.Context) {
	var req models.DistrictUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.DistrictService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新res_district失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除res_district
func (c *DistrictController) Delete(ctx *gin.Context) {
	var req models.DistrictDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.DistrictService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除res_district失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取res_district信息
func (c *DistrictController) GetByID(ctx *gin.Context) {
	var req models.DistrictGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	district, err := c.DistrictService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "res_district不存在", err)
	}

	c.Success(ctx, district)
}

// List res_district列表（分页查询）
func (c *DistrictController) List(ctx *gin.Context) {
	var req models.DistrictListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	districtList, total, err := c.DistrictService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取res_district列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  districtList,
		"total": total,
	})
}