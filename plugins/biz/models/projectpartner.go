package models

import (
	"context"
	"time"
	"gin-fast/app/global/app"
	"gorm.io/gorm"
)

// ProjectPartner biz_project_partner 模型结构体
type ProjectPartner struct {
	Id int64 `gorm:"column:id;primaryKey;not null;autoIncrement" json:"id"` // 主键
	TenantId uint `gorm:"column:tenant_id;default:0;index" json:"tenantId"` // 租户ID字段
	ProjectId int64 `gorm:"column:project_id" json:"projectId"` // 项目id
	RelationProject int64 `gorm:"column:relation_project" json:"relationProject"` // 关联项目
	RelationPartner int64 `gorm:"column:relation_partner" json:"relationPartner"` // 关联合作伙伴
	Organization string `gorm:"column:organization" json:"organization"` // 机构
	InvestmentAmount float64 `gorm:"column:investment_amount" json:"investmentAmount"` // 投资金额
	Unit string `gorm:"column:unit" json:"unit"` // 单位
	Tag string `gorm:"column:tag" json:"tag"` // 数据标签
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64 `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"` // 修改时间
	UpdatedBy int64 `gorm:"column:updated_by" json:"updatedBy"` // 修改人
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// ProjectPartnerList biz_project_partner 列表
type ProjectPartnerList []ProjectPartner

// NewProjectPartner 创建biz_project_partner实例
func NewProjectPartner() *ProjectPartner {
	return &ProjectPartner{}
}

// NewProjectPartnerList 创建biz_project_partner列表实例
func NewProjectPartnerList() *ProjectPartnerList {
	return &ProjectPartnerList{}
}

// TableName 指定表名
func (ProjectPartner) TableName() string {
	return "biz_project_partner"
}

// GetByID 根据ID获取biz_project_partner
func (m *ProjectPartner) GetByID(c context.Context, id int64) error {
	return app.DB().WithContext(c).First(m, id).Error
}

// Create 创建biz_project_partner记录
func (m *ProjectPartner) Create(c context.Context) error {
	return app.DB().WithContext(c).Create(m).Error
}

// Update 更新biz_project_partner记录
func (m *ProjectPartner) Update(c context.Context) error {
	return app.DB().WithContext(c).Save(m).Error
}

// Delete 软删除biz_project_partner记录
func (m *ProjectPartner) Delete(c context.Context) error {
	return app.DB().WithContext(c).Delete(m).Error
}

// IsEmpty 检查模型是否为空
func (m *ProjectPartner) IsEmpty() bool {
	return m == nil || m.Id == 0
}

// Find 查询biz_project_partner列表
func (l *ProjectPartnerList) Find(c context.Context, funcs ...func(*gorm.DB) *gorm.DB) error {
	return app.DB().WithContext(c).Model(&ProjectPartner{}).Scopes(funcs...).Find(l).Error
}

// GetTotal 获取biz_project_partner总数
func (l *ProjectPartnerList) GetTotal(c context.Context, query ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := app.DB().WithContext(c).Model(&ProjectPartner{}).Scopes(query...).Count(&count).Error
	return count, err
}