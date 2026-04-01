package models

import (
	"context"
	"time"
	"gin-fast/app/global/app"
	"gorm.io/gorm"
)

// ProjectRes biz_project_res 模型结构体
type ProjectRes struct {
	Id int64 `gorm:"column:id;primaryKey;not null;autoIncrement" json:"id"` // 主键
	TenantId uint `gorm:"column:tenant_id;default:0;index" json:"tenantId"` // 租户ID字段
	ProjectId int64 `gorm:"column:project_id" json:"projectId"` // 项目id
	AgencyType string `gorm:"column:agency_type" json:"agencyType"` // 资源类型
	FinanceId int64 `gorm:"column:finance_id" json:"financeId"` // 金融机构
	InvestId int64 `gorm:"column:invest_id" json:"investId"` // 投资机构
	ThinktankId int64 `gorm:"column:thinktank_id" json:"thinktankId"` // 新研机构
	SceneId int64 `gorm:"column:scene_id" json:"sceneId"` // 场景
	DistrictId int64 `gorm:"column:district_id" json:"districtId"` // 区县id
	CarrierId int64 `gorm:"column:carrier_id" json:"carrierId"` // 载体id
	CarrierInfoId int64 `gorm:"column:carrier_info_id" json:"carrierInfoId"` // 载体数据id
	Tag string `gorm:"column:tag" json:"tag"` // 数据标签
	PolicyId int64 `gorm:"column:policy_id" json:"policyId"` // 政策id
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64 `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"` // 修改时间
	UpdatedBy int64 `gorm:"column:updated_by" json:"updatedBy"` // 修改人
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// ProjectResList biz_project_res 列表
type ProjectResList []ProjectRes

// NewProjectRes 创建biz_project_res实例
func NewProjectRes() *ProjectRes {
	return &ProjectRes{}
}

// NewProjectResList 创建biz_project_res列表实例
func NewProjectResList() *ProjectResList {
	return &ProjectResList{}
}

// TableName 指定表名
func (ProjectRes) TableName() string {
	return "biz_project_res"
}

// GetByID 根据ID获取biz_project_res
func (m *ProjectRes) GetByID(c context.Context, id int64) error {
	return app.DB().WithContext(c).First(m, id).Error
}

// Create 创建biz_project_res记录
func (m *ProjectRes) Create(c context.Context) error {
	return app.DB().WithContext(c).Create(m).Error
}

// Update 更新biz_project_res记录
func (m *ProjectRes) Update(c context.Context) error {
	return app.DB().WithContext(c).Save(m).Error
}

// Delete 软删除biz_project_res记录
func (m *ProjectRes) Delete(c context.Context) error {
	return app.DB().WithContext(c).Delete(m).Error
}

// IsEmpty 检查模型是否为空
func (m *ProjectRes) IsEmpty() bool {
	return m == nil || m.Id == 0
}

// Find 查询biz_project_res列表
func (l *ProjectResList) Find(c context.Context, funcs ...func(*gorm.DB) *gorm.DB) error {
	return app.DB().WithContext(c).Model(&ProjectRes{}).Scopes(funcs...).Find(l).Error
}

// GetTotal 获取biz_project_res总数
func (l *ProjectResList) GetTotal(c context.Context, query ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := app.DB().WithContext(c).Model(&ProjectRes{}).Scopes(query...).Count(&count).Error
	return count, err
}