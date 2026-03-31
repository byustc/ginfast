package models

import (
	"context"
	"time"
	"gin-fast/app/global/app"
	"gorm.io/gorm"
)

// ThitankFund res_thinktank_fund 模型结构体
type ThitankFund struct {
	Id int64 `gorm:"column:id;primaryKey;not null;autoIncrement" json:"id"` // 主键
	TenantId uint `gorm:"column:tenant_id;default:0;index" json:"tenantId"` // 租户ID字段
	FundName string `gorm:"column:fund_name" json:"fundName"` // 基金名称
	FundScale string `gorm:"column:fund_scale" json:"fundScale"` // 基金规模
	ThinktankId int64 `gorm:"column:thinktank_id" json:"thinktankId"` // 新研机构-数据主键id
	InvestIndustry string `gorm:"column:invest_industry" json:"investIndustry"` // 投资行业偏好
	RoundPrefer string `gorm:"column:round_prefer" json:"roundPrefer"` // 投资阶段偏好
	SingleInvestAmount string `gorm:"column:single_invest_amount" json:"singleInvestAmount"` // 单笔投资金额
	InvestMethod string `gorm:"column:invest_method" json:"investMethod"` // 投资方式
	InvestableBalance float64 `gorm:"column:investable_balance" json:"investableBalance"` // 可投余额
	ReturnRequirements string `gorm:"column:return_requirements" json:"returnRequirements"` // 返投要求
	CompletionStatus string `gorm:"column:completion_status" json:"completionStatus"` // 返投完成情况
	FundType string `gorm:"column:fund_type" json:"fundType"` // 基金类型
	EstablishTime *time.Time `gorm:"column:establish_time" json:"establishTime"` // 成立时间
	DataSource string `gorm:"column:data_source" json:"dataSource"` // 数据来源
	IsBlacklist int `gorm:"column:is_blacklist;default:0" json:"isBlacklist"` // 是否负面清单: 0否 1是
	InvestmentIndustries string `gorm:"column:investment_industries" json:"investmentIndustries"` // 投资产业
	AvailableBalanceOfFund string `gorm:"column:available_balance_of_fund" json:"availableBalanceOfFund"` // 基金可投余额
	Tag string `gorm:"column:tag" json:"tag"` // 数据标签
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64 `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"` // 修改时间
	UpdatedBy int64 `gorm:"column:updated_by" json:"updatedBy"` // 修改人
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// ThitankFundList res_thinktank_fund 列表
type ThitankFundList []ThitankFund

// NewThitankFund 创建res_thinktank_fund实例
func NewThitankFund() *ThitankFund {
	return &ThitankFund{}
}

// NewThitankFundList 创建res_thinktank_fund列表实例
func NewThitankFundList() *ThitankFundList {
	return &ThitankFundList{}
}

// TableName 指定表名
func (ThitankFund) TableName() string {
	return "res_thinktank_fund"
}

// GetByID 根据ID获取res_thinktank_fund
func (m *ThitankFund) GetByID(c context.Context, id int64) error {
	return app.DB().WithContext(c).First(m, id).Error
}

// Create 创建res_thinktank_fund记录
func (m *ThitankFund) Create(c context.Context) error {
	return app.DB().WithContext(c).Create(m).Error
}

// Update 更新res_thinktank_fund记录
func (m *ThitankFund) Update(c context.Context) error {
	return app.DB().WithContext(c).Save(m).Error
}

// Delete 软删除res_thinktank_fund记录
func (m *ThitankFund) Delete(c context.Context) error {
	return app.DB().WithContext(c).Delete(m).Error
}

// IsEmpty 检查模型是否为空
func (m *ThitankFund) IsEmpty() bool {
	return m == nil || m.Id == 0
}

// Find 查询res_thinktank_fund列表
func (l *ThitankFundList) Find(c context.Context, funcs ...func(*gorm.DB) *gorm.DB) error {
	return app.DB().WithContext(c).Model(&ThitankFund{}).Scopes(funcs...).Find(l).Error
}

// GetTotal 获取res_thinktank_fund总数
func (l *ThitankFundList) GetTotal(c context.Context, query ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := app.DB().WithContext(c).Model(&ThitankFund{}).Scopes(query...).Count(&count).Error
	return count, err
}