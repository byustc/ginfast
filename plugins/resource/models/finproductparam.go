package models

import (
	"gin-fast/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FinProductListRequest res_finance_product列表请求参数
type FinProductListRequest struct {
	models.BasePaging
	models.Validator
	Id *int64 `form:"id"` // 主键
	ProductName *string `form:"productName"` // 产品名称
	FinanceId *int64 `form:"financeId"` // 所属机构
	ProductIntro *string `form:"productIntro"` // 产品特点
	ProductFeatures *string `form:"productFeatures"` // 产品特点
	ApplyCondition *string `form:"applyCondition"` // 申请条件
	ApplyMaterial *string `form:"applyMaterial"` // 申请资料
	ProductType *string `form:"productType"` // 金融产品类型
	ProductGroup *string `form:"productGroup"` // 适用客户类型
	QuotaLimit *float64 `form:"quotaLimit"` // 额度(万元)
	RateLimit *float64 `form:"rateLimit"` // 利率
	TermImit *string `form:"termImit"` // 期限
	GuaranteeMethod *string `form:"guaranteeMethod"` // 担保方式
	MaxCreditTerm *string `form:"maxCreditTerm"` // 授信最长期限
	GuaranteeLimit *string `form:"guaranteeLimit"` // 担保类额度上限
	AssetCollateralRatioLimit *string `form:"assetCollateralRatioLimit"` // 资产抵押率上限
	DataSource *string `form:"dataSource"` // 数据来源
	IsBlacklist *int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag *string `form:"tag"` // 行业数据标签
	CorrespondingSourceId *float64 `form:"correspondingSourceId"` // 对应数据源id
	UpdatedBy *int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *FinProductListRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// Handle 获取查询条件
func (r *FinProductListRequest) Handle() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
        if r.Id != nil {
            // 默认等于查询
            db = db.Where("id = ?", *r.Id)
        }
        if r.ProductName != nil {
            // 默认等于查询
            db = db.Where("product_name = ?", *r.ProductName)
        }
        if r.FinanceId != nil {
            // 默认等于查询
            db = db.Where("finance_id = ?", *r.FinanceId)
        }
        if r.ProductIntro != nil {
            // 默认等于查询
            db = db.Where("product_intro = ?", *r.ProductIntro)
        }
        if r.ProductFeatures != nil {
            // 默认等于查询
            db = db.Where("product_features = ?", *r.ProductFeatures)
        }
        if r.ApplyCondition != nil {
            // 默认等于查询
            db = db.Where("apply_condition = ?", *r.ApplyCondition)
        }
        if r.ApplyMaterial != nil {
            // 默认等于查询
            db = db.Where("apply_material = ?", *r.ApplyMaterial)
        }
        if r.ProductType != nil {
            // 默认等于查询
            db = db.Where("product_type = ?", *r.ProductType)
        }
        if r.ProductGroup != nil {
            // 默认等于查询
            db = db.Where("product_group = ?", *r.ProductGroup)
        }
        if r.QuotaLimit != nil {
            // 默认等于查询
            db = db.Where("quota_limit = ?", *r.QuotaLimit)
        }
        if r.RateLimit != nil {
            // 默认等于查询
            db = db.Where("rate_limit = ?", *r.RateLimit)
        }
        if r.TermImit != nil {
            // 默认等于查询
            db = db.Where("term_imit = ?", *r.TermImit)
        }
        if r.GuaranteeMethod != nil {
            // 默认等于查询
            db = db.Where("guarantee_method = ?", *r.GuaranteeMethod)
        }
        if r.MaxCreditTerm != nil {
            // 默认等于查询
            db = db.Where("max_credit_term = ?", *r.MaxCreditTerm)
        }
        if r.GuaranteeLimit != nil {
            // 默认等于查询
            db = db.Where("guarantee_limit = ?", *r.GuaranteeLimit)
        }
        if r.AssetCollateralRatioLimit != nil {
            // 默认等于查询
            db = db.Where("asset_collateral_ratio_limit = ?", *r.AssetCollateralRatioLimit)
        }
        if r.DataSource != nil {
            // 默认等于查询
            db = db.Where("data_source = ?", *r.DataSource)
        }
        if r.IsBlacklist != nil {
            // 默认等于查询
            db = db.Where("is_blacklist = ?", *r.IsBlacklist)
        }
        if r.Tag != nil {
            // 默认等于查询
            db = db.Where("tag = ?", *r.Tag)
        }
        if r.CorrespondingSourceId != nil {
            // 默认等于查询
            db = db.Where("corresponding_source_id = ?", *r.CorrespondingSourceId)
        }
        if r.UpdatedBy != nil {
            // 默认等于查询
            db = db.Where("updated_by = ?", *r.UpdatedBy)
        }
		return db
	}
}

// FinProductCreateRequest 创建res_finance_product请求参数
type FinProductCreateRequest struct {
	models.Validator
	ProductName string `form:"productName" validate:"required" message:"产品名称不能为空"` // 产品名称
	FinanceId int64 `form:"financeId"` // 所属机构
	ProductIntro string `form:"productIntro" validate:"required" message:"产品特点不能为空"` // 产品特点
	ProductFeatures string `form:"productFeatures" validate:"required" message:"产品特点不能为空"` // 产品特点
	ApplyCondition string `form:"applyCondition" validate:"required" message:"申请条件不能为空"` // 申请条件
	ApplyMaterial string `form:"applyMaterial" validate:"required" message:"申请资料不能为空"` // 申请资料
	ProductType string `form:"productType" validate:"required" message:"金融产品类型不能为空"` // 金融产品类型
	ProductGroup string `form:"productGroup" validate:"required" message:"适用客户类型不能为空"` // 适用客户类型
	QuotaLimit float64 `form:"quotaLimit"` // 额度(万元)
	RateLimit float64 `form:"rateLimit"` // 利率
	TermImit string `form:"termImit" validate:"required" message:"期限不能为空"` // 期限
	GuaranteeMethod string `form:"guaranteeMethod" validate:"required" message:"担保方式不能为空"` // 担保方式
	MaxCreditTerm string `form:"maxCreditTerm" validate:"required" message:"授信最长期限不能为空"` // 授信最长期限
	GuaranteeLimit string `form:"guaranteeLimit" validate:"required" message:"担保类额度上限不能为空"` // 担保类额度上限
	AssetCollateralRatioLimit string `form:"assetCollateralRatioLimit" validate:"required" message:"资产抵押率上限不能为空"` // 资产抵押率上限
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	CorrespondingSourceId float64 `form:"correspondingSourceId"` // 对应数据源id
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *FinProductCreateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// FinProductUpdateRequest 更新res_finance_product请求参数
type FinProductUpdateRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
	ProductName string `form:"productName" validate:"required" message:"产品名称不能为空"` // 产品名称
	FinanceId int64 `form:"financeId"` // 所属机构
	ProductIntro string `form:"productIntro" validate:"required" message:"产品特点不能为空"` // 产品特点
	ProductFeatures string `form:"productFeatures" validate:"required" message:"产品特点不能为空"` // 产品特点
	ApplyCondition string `form:"applyCondition" validate:"required" message:"申请条件不能为空"` // 申请条件
	ApplyMaterial string `form:"applyMaterial" validate:"required" message:"申请资料不能为空"` // 申请资料
	ProductType string `form:"productType" validate:"required" message:"金融产品类型不能为空"` // 金融产品类型
	ProductGroup string `form:"productGroup" validate:"required" message:"适用客户类型不能为空"` // 适用客户类型
	QuotaLimit float64 `form:"quotaLimit"` // 额度(万元)
	RateLimit float64 `form:"rateLimit"` // 利率
	TermImit string `form:"termImit" validate:"required" message:"期限不能为空"` // 期限
	GuaranteeMethod string `form:"guaranteeMethod" validate:"required" message:"担保方式不能为空"` // 担保方式
	MaxCreditTerm string `form:"maxCreditTerm" validate:"required" message:"授信最长期限不能为空"` // 授信最长期限
	GuaranteeLimit string `form:"guaranteeLimit" validate:"required" message:"担保类额度上限不能为空"` // 担保类额度上限
	AssetCollateralRatioLimit string `form:"assetCollateralRatioLimit" validate:"required" message:"资产抵押率上限不能为空"` // 资产抵押率上限
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	CorrespondingSourceId float64 `form:"correspondingSourceId"` // 对应数据源id
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *FinProductUpdateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// FinProductDeleteRequest 删除res_finance_product请求参数
type FinProductDeleteRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *FinProductDeleteRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// FinProductGetByIDRequest 根据ID获取res_finance_product请求参数
type FinProductGetByIDRequest struct {
	models.Validator
	Id int64 `uri:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *FinProductGetByIDRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}