package models

import (
	"context"
	"time"
	"gin-fast/app/global/app"
	"gorm.io/gorm"
)

// InvFund res_invest_fund 模型结构体
type InvFund struct {
	Id int64 `gorm:"column:id;primaryKey;not null;autoIncrement" json:"id"` // 主键
	TenantId uint `gorm:"column:tenant_id;default:0;index" json:"tenantId"` // 租户ID字段
	InvestId int64 `gorm:"column:invest_id" json:"investId"` // 投资机构-数据主键id
	FundName string `gorm:"column:fund_name" json:"fundName"` // 基金名称
	FundScale string `gorm:"column:fund_scale" json:"fundScale"` // 基金规模
	InvestIndustry string `gorm:"column:invest_industry" json:"investIndustry"` // 投资行业
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
	InvestmentField string `gorm:"column:investment_field" json:"investmentField"` // 投资领域
	AvailableBalanceFund string `gorm:"column:available_balance_fund" json:"availableBalanceFund"` // 基金可投余额
	Tag string `gorm:"column:tag" json:"tag"` // 行业数据标签
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64 `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"` // 修改时间
	UpdatedBy int64 `gorm:"column:updated_by" json:"updatedBy"` // 修改人
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// InvFundList res_invest_fund 列表
type InvFundList []InvFund

// NewInvFund 创建res_invest_fund实例
func NewInvFund() *InvFund {
	return &InvFund{}
}

// NewInvFundList 创建res_invest_fund列表实例
func NewInvFundList() *InvFundList {
	return &InvFundList{}
}

// TableName 指定表名
func (InvFund) TableName() string {
	return "res_invest_fund"
}

// GetByID 根据ID获取res_invest_fund
func (m *InvFund) GetByID(c context.Context, id int64) error {
	return app.DB().WithContext(c).First(m, id).Error
}

// Create 创建res_invest_fund记录
func (m *InvFund) Create(c context.Context) error {
	return app.DB().WithContext(c).Create(m).Error
}

// Update 更新res_invest_fund记录
func (m *InvFund) Update(c context.Context) error {
	return app.DB().WithContext(c).Save(m).Error
}

// Delete 软删除res_invest_fund记录
func (m *InvFund) Delete(c context.Context) error {
	return app.DB().WithContext(c).Delete(m).Error
}

// IsEmpty 检查模型是否为空
func (m *InvFund) IsEmpty() bool {
	return m == nil || m.Id == 0
}

// Find 查询res_invest_fund列表
func (l *InvFundList) Find(c context.Context, funcs ...func(*gorm.DB) *gorm.DB) error {
	return app.DB().WithContext(c).Model(&InvFund{}).Scopes(funcs...).Find(l).Error
}

// GetTotal 获取res_invest_fund总数
func (l *InvFundList) GetTotal(c context.Context, query ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := app.DB().WithContext(c).Model(&InvFund{}).Scopes(query...).Count(&count).Error
	return count, err
}