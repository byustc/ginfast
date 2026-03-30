package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/resource/models"
	"gin-fast/plugins/resource/service"

	"github.com/gin-gonic/gin"
)

// CarrierController res_carrier控制器
type CarrierController struct {
	controllers.Common
	CarrierService *service.CarrierService
}

// NewCarrierController 创建res_carrier控制器
func NewCarrierController() *CarrierController {
	return &CarrierController{
		CarrierService: service.NewCarrierService(),
	}
}

// Create 创建res_carrier
func (c *CarrierController) Create(ctx *gin.Context) {
	var req models.CarrierCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	carrier, err := c.CarrierService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建res_carrier失败", err)
	}

	c.Success(ctx, gin.H{
		"id": carrier.Id,
	})
}

// Update 更新res_carrier
func (c *CarrierController) Update(ctx *gin.Context) {
	var req models.CarrierUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.CarrierService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新res_carrier失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除res_carrier
func (c *CarrierController) Delete(ctx *gin.Context) {
	var req models.CarrierDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.CarrierService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除res_carrier失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取res_carrier信息
func (c *CarrierController) GetByID(ctx *gin.Context) {
	var req models.CarrierGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	carrier, err := c.CarrierService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "res_carrier不存在", err)
	}

	c.Success(ctx, carrier)
}

// List res_carrier列表（分页查询）
func (c *CarrierController) List(ctx *gin.Context) {
	var req models.CarrierListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	carrierList, total, err := c.CarrierService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取res_carrier列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  carrierList,
		"total": total,
	})
}