package service

import (
	"gin-fast/plugins/resource/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// FinProductService res_finance_product服务
type FinProductService struct{}

// NewFinProductService 创建res_finance_product服务
func NewFinProductService() *FinProductService {
	return &FinProductService{}
}

// Create 创建res_finance_product
func (s *FinProductService) Create(c *gin.Context, req models.FinProductCreateRequest) (*models.FinProduct, error) {
	// 创建res_finance_product记录
	finProduct := models.NewFinProduct()
    finProduct.ProductName = req.ProductName
    finProduct.FinanceId = req.FinanceId
    finProduct.ProductIntro = req.ProductIntro
    finProduct.ProductFeatures = req.ProductFeatures
    finProduct.ApplyCondition = req.ApplyCondition
    finProduct.ApplyMaterial = req.ApplyMaterial
    finProduct.ProductType = req.ProductType
    finProduct.ProductGroup = req.ProductGroup
    finProduct.QuotaLimit = req.QuotaLimit
    finProduct.RateLimit = req.RateLimit
    finProduct.TermImit = req.TermImit
    finProduct.GuaranteeMethod = req.GuaranteeMethod
    finProduct.MaxCreditTerm = req.MaxCreditTerm
    finProduct.GuaranteeLimit = req.GuaranteeLimit
    finProduct.AssetCollateralRatioLimit = req.AssetCollateralRatioLimit
    finProduct.DataSource = req.DataSource
    finProduct.IsBlacklist = req.IsBlacklist
    finProduct.Tag = req.Tag
    finProduct.CorrespondingSourceId = req.CorrespondingSourceId
    finProduct.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := finProduct.Create(c); err != nil {
		return nil, err
	}

	return finProduct, nil
}

// Update 更新res_finance_product
func (s *FinProductService) Update(c *gin.Context, req models.FinProductUpdateRequest) error {
	// 查找res_finance_product记录
	finProduct := models.NewFinProduct()
	if err := finProduct.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新res_finance_product信息
    finProduct.ProductName = req.ProductName
    finProduct.FinanceId = req.FinanceId
    finProduct.ProductIntro = req.ProductIntro
    finProduct.ProductFeatures = req.ProductFeatures
    finProduct.ApplyCondition = req.ApplyCondition
    finProduct.ApplyMaterial = req.ApplyMaterial
    finProduct.ProductType = req.ProductType
    finProduct.ProductGroup = req.ProductGroup
    finProduct.QuotaLimit = req.QuotaLimit
    finProduct.RateLimit = req.RateLimit
    finProduct.TermImit = req.TermImit
    finProduct.GuaranteeMethod = req.GuaranteeMethod
    finProduct.MaxCreditTerm = req.MaxCreditTerm
    finProduct.GuaranteeLimit = req.GuaranteeLimit
    finProduct.AssetCollateralRatioLimit = req.AssetCollateralRatioLimit
    finProduct.DataSource = req.DataSource
    finProduct.IsBlacklist = req.IsBlacklist
    finProduct.Tag = req.Tag
    finProduct.CorrespondingSourceId = req.CorrespondingSourceId
    finProduct.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := finProduct.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除res_finance_product
func (s *FinProductService) Delete(c *gin.Context, id int64) error {
	// 查找res_finance_product记录
	finProduct := models.NewFinProduct()
	if err := finProduct.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := finProduct.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取res_finance_product
func (s *FinProductService) GetByID(c *gin.Context, id int64) (*models.FinProduct, error) {
	// 查找res_finance_product记录
	finProduct := models.NewFinProduct()
	if err := finProduct.GetByID(c, id); err != nil {
		return nil, err
	}

	return finProduct, nil
}

// List res_finance_product列表（分页查询）
func (s *FinProductService) List(c *gin.Context, req models.FinProductListRequest) (*models.FinProductList, int64, error) {
	// 获取总数
	finProductList := models.NewFinProductList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := finProductList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = finProductList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return finProductList, total, nil
}