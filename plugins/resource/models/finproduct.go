package models

import (
	"context"
	"time"
	"gin-fast/app/global/app"
	"gorm.io/gorm"
)

// FinProduct res_finance_product 模型结构体
type FinProduct struct {
	Id int64 `gorm:"column:id;primaryKey;not null;autoIncrement" json:"id"` // 主键
	TenantId uint `gorm:"column:tenant_id;default:0;index" json:"tenantId"` // 租户ID字段
	ProductName string `gorm:"column:product_name" json:"productName"` // 产品名称
	FinanceId int64 `gorm:"column:finance_id" json:"financeId"` // 所属机构
	ProductIntro string `gorm:"column:product_intro" json:"productIntro"` // 产品特点
	ProductFeatures string `gorm:"column:product_features" json:"productFeatures"` // 产品特点
	ApplyCondition string `gorm:"column:apply_condition" json:"applyCondition"` // 申请条件
	ApplyMaterial string `gorm:"column:apply_material" json:"applyMaterial"` // 申请资料
	ProductType string `gorm:"column:product_type" json:"productType"` // 金融产品类型
	ProductGroup string `gorm:"column:product_group" json:"productGroup"` // 适用客户类型
	QuotaLimit float64 `gorm:"column:quota_limit" json:"quotaLimit"` // 额度(万元)
	RateLimit float64 `gorm:"column:rate_limit" json:"rateLimit"` // 利率
	TermImit string `gorm:"column:term_imit" json:"termImit"` // 期限
	GuaranteeMethod string `gorm:"column:guarantee_method" json:"guaranteeMethod"` // 担保方式
	MaxCreditTerm string `gorm:"column:max_credit_term" json:"maxCreditTerm"` // 授信最长期限
	GuaranteeLimit string `gorm:"column:guarantee_limit" json:"guaranteeLimit"` // 担保类额度上限
	AssetCollateralRatioLimit string `gorm:"column:asset_collateral_ratio_limit" json:"assetCollateralRatioLimit"` // 资产抵押率上限
	DataSource string `gorm:"column:data_source" json:"dataSource"` // 数据来源
	IsBlacklist int `gorm:"column:is_blacklist;default:0" json:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `gorm:"column:tag" json:"tag"` // 行业数据标签
	CorrespondingSourceId float64 `gorm:"column:corresponding_source_id" json:"correspondingSourceId"` // 对应数据源id
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64 `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"` // 修改时间
	UpdatedBy int64 `gorm:"column:updated_by" json:"updatedBy"` // 修改人
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// FinProductList res_finance_product 列表
type FinProductList []FinProduct

// NewFinProduct 创建res_finance_product实例
func NewFinProduct() *FinProduct {
	return &FinProduct{}
}

// NewFinProductList 创建res_finance_product列表实例
func NewFinProductList() *FinProductList {
	return &FinProductList{}
}

// TableName 指定表名
func (FinProduct) TableName() string {
	return "res_finance_product"
}

// GetByID 根据ID获取res_finance_product
func (m *FinProduct) GetByID(c context.Context, id int64) error {
	return app.DB().WithContext(c).First(m, id).Error
}

// Create 创建res_finance_product记录
func (m *FinProduct) Create(c context.Context) error {
	return app.DB().WithContext(c).Create(m).Error
}

// Update 更新res_finance_product记录
func (m *FinProduct) Update(c context.Context) error {
	return app.DB().WithContext(c).Save(m).Error
}

// Delete 软删除res_finance_product记录
func (m *FinProduct) Delete(c context.Context) error {
	return app.DB().WithContext(c).Delete(m).Error
}

// IsEmpty 检查模型是否为空
func (m *FinProduct) IsEmpty() bool {
	return m == nil || m.Id == 0
}

// Find 查询res_finance_product列表
func (l *FinProductList) Find(c context.Context, funcs ...func(*gorm.DB) *gorm.DB) error {
	return app.DB().WithContext(c).Model(&FinProduct{}).Scopes(funcs...).Find(l).Error
}

// GetTotal 获取res_finance_product总数
func (l *FinProductList) GetTotal(c context.Context, query ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := app.DB().WithContext(c).Model(&FinProduct{}).Scopes(query...).Count(&count).Error
	return count, err
}