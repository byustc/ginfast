package models

import (
	"context"
	"time"
	"gin-fast/app/global/app"
	"gorm.io/gorm"
)

// ThitankMember res_thinktank_member 模型结构体
type ThitankMember struct {
	Id int64 `gorm:"column:id;primaryKey;not null;autoIncrement" json:"id"` // 主键
	TenantId uint `gorm:"column:tenant_id;default:0;index" json:"tenantId"` // 租户ID字段
	ThinktankId int64 `gorm:"column:thinktank_id" json:"thinktankId"` // 新研机构-数据主键id
	PersonName string `gorm:"column:person_name" json:"personName"` // 姓名
	Position string `gorm:"column:position" json:"position"` // 职位
	PersonIntro string `gorm:"column:person_intro" json:"personIntro"` // 个人介绍
	Phone string `gorm:"column:phone" json:"phone"` // 联系方式
	EduBackground string `gorm:"column:edu_background" json:"eduBackground"` // 学历
	EnterpriseId int64 `gorm:"column:enterprise_id" json:"enterpriseId"` // 企业-数据主键id
	RankNum float64 `gorm:"column:rank_num" json:"rankNum"` // 排序字段
	DataSource string `gorm:"column:data_source" json:"dataSource"` // 数据来源
	IsBlacklist int `gorm:"column:is_blacklist;default:0" json:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `gorm:"column:tag" json:"tag"` // 数据标签
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64 `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"` // 修改时间
	UpdatedBy int64 `gorm:"column:updated_by" json:"updatedBy"` // 修改人
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// ThitankMemberList res_thinktank_member 列表
type ThitankMemberList []ThitankMember

// NewThitankMember 创建res_thinktank_member实例
func NewThitankMember() *ThitankMember {
	return &ThitankMember{}
}

// NewThitankMemberList 创建res_thinktank_member列表实例
func NewThitankMemberList() *ThitankMemberList {
	return &ThitankMemberList{}
}

// TableName 指定表名
func (ThitankMember) TableName() string {
	return "res_thinktank_member"
}

// GetByID 根据ID获取res_thinktank_member
func (m *ThitankMember) GetByID(c context.Context, id int64) error {
	return app.DB().WithContext(c).First(m, id).Error
}

// Create 创建res_thinktank_member记录
func (m *ThitankMember) Create(c context.Context) error {
	return app.DB().WithContext(c).Create(m).Error
}

// Update 更新res_thinktank_member记录
func (m *ThitankMember) Update(c context.Context) error {
	return app.DB().WithContext(c).Save(m).Error
}

// Delete 软删除res_thinktank_member记录
func (m *ThitankMember) Delete(c context.Context) error {
	return app.DB().WithContext(c).Delete(m).Error
}

// IsEmpty 检查模型是否为空
func (m *ThitankMember) IsEmpty() bool {
	return m == nil || m.Id == 0
}

// Find 查询res_thinktank_member列表
func (l *ThitankMemberList) Find(c context.Context, funcs ...func(*gorm.DB) *gorm.DB) error {
	return app.DB().WithContext(c).Model(&ThitankMember{}).Scopes(funcs...).Find(l).Error
}

// GetTotal 获取res_thinktank_member总数
func (l *ThitankMemberList) GetTotal(c context.Context, query ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := app.DB().WithContext(c).Model(&ThitankMember{}).Scopes(query...).Count(&count).Error
	return count, err
}