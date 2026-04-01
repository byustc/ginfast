package models

import (
	"context"
	"time"
	"gin-fast/app/global/app"
	"gorm.io/gorm"
)

// ProjectPhasePlan biz_project_phase_plan 模型结构体
type ProjectPhasePlan struct {
	Id int64 `gorm:"column:id;primaryKey;not null;autoIncrement" json:"id"` // 主键
	TenantId uint `gorm:"column:tenant_id;default:0;index" json:"tenantId"` // 租户ID字段
	ProjectId int64 `gorm:"column:project_id" json:"projectId"` // 项目id
	ItemPhaseId string `gorm:"column:item_phase_id" json:"itemPhaseId"` // 项目阶段
	PlanFinishTime *time.Time `gorm:"column:plan_finish_time" json:"planFinishTime"` // 计划完成时间
	ActualFinishTime *time.Time `gorm:"column:actual_finish_time" json:"actualFinishTime"` // 实际完成时间
	Tag string `gorm:"column:tag" json:"tag"` // 数据标签
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64 `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"` // 修改时间
	UpdatedBy int64 `gorm:"column:updated_by" json:"updatedBy"` // 修改人
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// ProjectPhasePlanList biz_project_phase_plan 列表
type ProjectPhasePlanList []ProjectPhasePlan

// NewProjectPhasePlan 创建biz_project_phase_plan实例
func NewProjectPhasePlan() *ProjectPhasePlan {
	return &ProjectPhasePlan{}
}

// NewProjectPhasePlanList 创建biz_project_phase_plan列表实例
func NewProjectPhasePlanList() *ProjectPhasePlanList {
	return &ProjectPhasePlanList{}
}

// TableName 指定表名
func (ProjectPhasePlan) TableName() string {
	return "biz_project_phase_plan"
}

// GetByID 根据ID获取biz_project_phase_plan
func (m *ProjectPhasePlan) GetByID(c context.Context, id int64) error {
	return app.DB().WithContext(c).First(m, id).Error
}

// Create 创建biz_project_phase_plan记录
func (m *ProjectPhasePlan) Create(c context.Context) error {
	return app.DB().WithContext(c).Create(m).Error
}

// Update 更新biz_project_phase_plan记录
func (m *ProjectPhasePlan) Update(c context.Context) error {
	return app.DB().WithContext(c).Save(m).Error
}

// Delete 软删除biz_project_phase_plan记录
func (m *ProjectPhasePlan) Delete(c context.Context) error {
	return app.DB().WithContext(c).Delete(m).Error
}

// IsEmpty 检查模型是否为空
func (m *ProjectPhasePlan) IsEmpty() bool {
	return m == nil || m.Id == 0
}

// Find 查询biz_project_phase_plan列表
func (l *ProjectPhasePlanList) Find(c context.Context, funcs ...func(*gorm.DB) *gorm.DB) error {
	return app.DB().WithContext(c).Model(&ProjectPhasePlan{}).Scopes(funcs...).Find(l).Error
}

// GetTotal 获取biz_project_phase_plan总数
func (l *ProjectPhasePlanList) GetTotal(c context.Context, query ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := app.DB().WithContext(c).Model(&ProjectPhasePlan{}).Scopes(query...).Count(&count).Error
	return count, err
}