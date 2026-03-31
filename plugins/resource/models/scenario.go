package models

import (
	"context"
	"time"
	"gin-fast/app/global/app"
	"gorm.io/gorm"
)

// Scenario res_scenario 模型结构体
type Scenario struct {
	Id int64 `gorm:"column:id;primaryKey;not null;autoIncrement" json:"id"` // 主键
	TenantId uint `gorm:"column:tenant_id;default:0;index" json:"tenantId"` // 租户ID字段
	ScenarioNo string `gorm:"column:scenario_no;index" json:"scenarioNo"` // 场景编码
	PublishTime *time.Time `gorm:"column:publish_time" json:"publishTime"` // 发布时间
	OwnerUnit string `gorm:"column:owner_unit" json:"ownerUnit"` // 业主单位
	SceneName string `gorm:"column:scene_name" json:"sceneName"` // 场景名称
	Industry string `gorm:"column:industry" json:"industry"` // 所属产业
	OpportunityZone string `gorm:"column:opportunity_zone" json:"opportunityZone"` // 机会区域
	ValidityPeriod string `gorm:"column:validity_period" json:"validityPeriod"` // 机会有效期
	BudgetGuarantee string `gorm:"column:budget_guarantee" json:"budgetGuarantee"` // 预算保障情况
	OtherSituations string `gorm:"column:other_situations" json:"otherSituations"` // 其他情况
	OpportunityDesc string `gorm:"column:opportunity_desc" json:"opportunityDesc"` // 机会描述
	BuildBasement string `gorm:"column:build_basement" json:"buildBasement"` // 建设基础
	CooperationRequirements string `gorm:"column:cooperation_requirements" json:"cooperationRequirements"` // 合作需求
	CooperationForm string `gorm:"column:cooperation_form" json:"cooperationForm"` // 合作形式
	ContactPerson string `gorm:"column:contact_person" json:"contactPerson"` // 联系人
	Position string `gorm:"column:position" json:"position"` // 职务
	Phone string `gorm:"column:phone" json:"phone"` // 联系方式
	Remark string `gorm:"column:remark" json:"remark"` // 备注
	DataSource string `gorm:"column:data_source" json:"dataSource"` // 数据来源
	IsBlacklist int `gorm:"column:is_blacklist;default:0" json:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `gorm:"column:tag" json:"tag"` // 行业数据标签
	IndustrialFields string `gorm:"column:industrial_fields" json:"industrialFields"` // 产业领域
	XcujuChanye string `gorm:"column:xcuju_chanye" json:"xcujuChanye"` // 6+5+X
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64 `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"` // 修改时间
	UpdatedBy int64 `gorm:"column:updated_by" json:"updatedBy"` // 修改人
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// ScenarioList res_scenario 列表
type ScenarioList []Scenario

// NewScenario 创建res_scenario实例
func NewScenario() *Scenario {
	return &Scenario{}
}

// NewScenarioList 创建res_scenario列表实例
func NewScenarioList() *ScenarioList {
	return &ScenarioList{}
}

// TableName 指定表名
func (Scenario) TableName() string {
	return "res_scenario"
}

// GetByID 根据ID获取res_scenario
func (m *Scenario) GetByID(c context.Context, id int64) error {
	return app.DB().WithContext(c).First(m, id).Error
}

// Create 创建res_scenario记录
func (m *Scenario) Create(c context.Context) error {
	return app.DB().WithContext(c).Create(m).Error
}

// Update 更新res_scenario记录
func (m *Scenario) Update(c context.Context) error {
	return app.DB().WithContext(c).Save(m).Error
}

// Delete 软删除res_scenario记录
func (m *Scenario) Delete(c context.Context) error {
	return app.DB().WithContext(c).Delete(m).Error
}

// IsEmpty 检查模型是否为空
func (m *Scenario) IsEmpty() bool {
	return m == nil || m.Id == 0
}

// Find 查询res_scenario列表
func (l *ScenarioList) Find(c context.Context, funcs ...func(*gorm.DB) *gorm.DB) error {
	return app.DB().WithContext(c).Model(&Scenario{}).Scopes(funcs...).Find(l).Error
}

// GetTotal 获取res_scenario总数
func (l *ScenarioList) GetTotal(c context.Context, query ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := app.DB().WithContext(c).Model(&Scenario{}).Scopes(query...).Count(&count).Error
	return count, err
}