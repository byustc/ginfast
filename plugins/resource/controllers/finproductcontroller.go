package controllers

import (
	"gin-fast/app/controllers"
	"gin-fast/plugins/resource/models"
	"gin-fast/plugins/resource/service"

	"github.com/gin-gonic/gin"
)

// FinProductController res_finance_product控制器
type FinProductController struct {
	controllers.Common
	FinProductService *service.FinProductService
}

// NewFinProductController 创建res_finance_product控制器
func NewFinProductController() *FinProductController {
	return &FinProductController{
		FinProductService: service.NewFinProductService(),
	}
}

// Create 创建res_finance_product
func (c *FinProductController) Create(ctx *gin.Context) {
	var req models.FinProductCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}
	
	finProduct, err := c.FinProductService.Create(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "创建res_finance_product失败", err)
	}

	c.Success(ctx, gin.H{
		"id": finProduct.Id,
	})
}

// Update 更新res_finance_product
func (c *FinProductController) Update(ctx *gin.Context) {
	var req models.FinProductUpdateRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.FinProductService.Update(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "更新res_finance_product失败", err)
	}

	c.SuccessWithMessage(ctx, "更新成功")
}

// Delete 删除res_finance_product
func (c *FinProductController) Delete(ctx *gin.Context) {
	var req models.FinProductDeleteRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	err := c.FinProductService.Delete(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "删除res_finance_product失败", err)
	}

	c.SuccessWithMessage(ctx, "删除成功", nil)
}

// GetByID 根据ID获取res_finance_product信息
func (c *FinProductController) GetByID(ctx *gin.Context) {
	var req models.FinProductGetByIDRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	finProduct, err := c.FinProductService.GetByID(ctx, req.Id)
	if err != nil {
		c.FailAndAbort(ctx, "res_finance_product不存在", err)
	}

	c.Success(ctx, finProduct)
}

// List res_finance_product列表（分页查询）
func (c *FinProductController) List(ctx *gin.Context) {
	var req models.FinProductListRequest
	if err := req.Validate(ctx); err != nil {
		c.FailAndAbort(ctx, err.Error(), err)
	}

	finProductList, total, err := c.FinProductService.List(ctx, req)
	if err != nil {
		c.FailAndAbort(ctx, "获取res_finance_product列表失败", err)
	}

	c.Success(ctx, gin.H{
		"list":  finProductList,
		"total": total,
	})
}