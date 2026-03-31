package models

import (
	"context"
	"time"
	"gin-fast/app/global/app"
	"gorm.io/gorm"
)

// IndustryChain res_industry_chain 模型结构体
type IndustryChain struct {
	Id int64 `gorm:"column:id;primaryKey;not null;autoIncrement" json:"id"` // 主键
	TenantId uint `gorm:"column:tenant_id;default:0;index" json:"tenantId"` // 租户ID字段
	IndustryChainNo string `gorm:"column:industry_chain_no;index" json:"industryChainNo"` // 产业链编码
	IndustryChainName string `gorm:"column:industry_chain_name" json:"industryChainName"` // 产业链名称
	IndustryChainIntro string `gorm:"column:industry_chain_intro" json:"industryChainIntro"` // 产业链介绍
	IndustryChainMap string `gorm:"column:industry_chain_map" json:"industryChainMap"` // 产业链图谱
	DataSource string `gorm:"column:data_source" json:"dataSource"` // 数据来源
	IsBlacklist int `gorm:"column:is_blacklist;default:0" json:"isBlacklist"` // 是否负面清单: 0否 1是
	GraphLayer string `gorm:"column:graph_layer" json:"graphLayer"` // 图谱配置
	Tag string `gorm:"column:tag" json:"tag"` // 数据标签
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64 `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"` // 修改时间
	UpdatedBy int64 `gorm:"column:updated_by" json:"updatedBy"` // 修改人
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// IndustryChainList res_industry_chain 列表
type IndustryChainList []IndustryChain

// NewIndustryChain 创建res_industry_chain实例
func NewIndustryChain() *IndustryChain {
	return &IndustryChain{}
}

// NewIndustryChainList 创建res_industry_chain列表实例
func NewIndustryChainList() *IndustryChainList {
	return &IndustryChainList{}
}

// TableName 指定表名
func (IndustryChain) TableName() string {
	return "res_industry_chain"
}

// GetByID 根据ID获取res_industry_chain
func (m *IndustryChain) GetByID(c context.Context, id int64) error {
	return app.DB().WithContext(c).First(m, id).Error
}

// Create 创建res_industry_chain记录
func (m *IndustryChain) Create(c context.Context) error {
	return app.DB().WithContext(c).Create(m).Error
}

// Update 更新res_industry_chain记录
func (m *IndustryChain) Update(c context.Context) error {
	return app.DB().WithContext(c).Save(m).Error
}

// Delete 软删除res_industry_chain记录
func (m *IndustryChain) Delete(c context.Context) error {
	return app.DB().WithContext(c).Delete(m).Error
}

// IsEmpty 检查模型是否为空
func (m *IndustryChain) IsEmpty() bool {
	return m == nil || m.Id == 0
}

// Find 查询res_industry_chain列表
func (l *IndustryChainList) Find(c context.Context, funcs ...func(*gorm.DB) *gorm.DB) error {
	return app.DB().WithContext(c).Model(&IndustryChain{}).Scopes(funcs...).Find(l).Error
}

// GetTotal 获取res_industry_chain总数
func (l *IndustryChainList) GetTotal(c context.Context, query ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := app.DB().WithContext(c).Model(&IndustryChain{}).Scopes(query...).Count(&count).Error
	return count, err
}