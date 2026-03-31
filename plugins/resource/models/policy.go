package models

import (
	"context"
	"time"
	"gin-fast/app/global/app"
	"gorm.io/gorm"
)

// Policy res_policy 模型结构体
type Policy struct {
	Id int64 `gorm:"column:id;primaryKey;not null;autoIncrement" json:"id"` // 主键
	TenantId uint `gorm:"column:tenant_id;default:0;index" json:"tenantId"` // 租户ID字段
	PolicyNo string `gorm:"column:policy_no;index" json:"policyNo"` // 政策编码
	PolicyTitle string `gorm:"column:policy_title" json:"policyTitle"` // 政策标题
	PolicyAbstract string `gorm:"column:policy_abstract" json:"policyAbstract"` // 政策摘要
	MainBody string `gorm:"column:main_body" json:"mainBody"` // 正文
	PublishTime *time.Time `gorm:"column:publish_time" json:"publishTime"` // 发布时间
	PolicyDepartment string `gorm:"column:policy_department" json:"policyDepartment"` // 发布部门
	Source string `gorm:"column:source" json:"source"` // 来源
	Link string `gorm:"column:link" json:"link"` // 原文链接
	EffectDate *time.Time `gorm:"column:effect_date" json:"effectDate"` // 生效日期
	FailureDate *time.Time `gorm:"column:failure_date" json:"failureDate"` // 失效日期
	Status string `gorm:"column:status" json:"status"` // 政策状态
	PolicyTimeliness string `gorm:"column:policy_timeliness" json:"policyTimeliness"` // 时效性
	ContactName string `gorm:"column:contact_name" json:"contactName"` // 联系人
	ContactPhone string `gorm:"column:contact_phone" json:"contactPhone"` // 联系方式
	PolicyRegion string `gorm:"column:policy_region" json:"policyRegion"` // 政策地区
	PolicyTag string `gorm:"column:policy_tag" json:"policyTag"` // 目标产业
	PolicyComType string `gorm:"column:policy_com_type" json:"policyComType"` // 适用企业类型
	Level string `gorm:"column:level" json:"level"` // 政策级别
	Attachment string `gorm:"column:attachment" json:"attachment"` // 附件
	PolicyNumber string `gorm:"column:policy_number" json:"policyNumber"` // 文号
	Remark string `gorm:"column:remark" json:"remark"` // 备注
	Title string `gorm:"column:title" json:"title"` // 标题
	DataSource string `gorm:"column:data_source" json:"dataSource"` // 数据来源
	IsBlacklist int `gorm:"column:is_blacklist;default:0" json:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `gorm:"column:tag" json:"tag"` // 行业数据标签
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64 `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"` // 修改时间
	UpdatedBy int64 `gorm:"column:updated_by" json:"updatedBy"` // 修改人
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// PolicyList res_policy 列表
type PolicyList []Policy

// NewPolicy 创建res_policy实例
func NewPolicy() *Policy {
	return &Policy{}
}

// NewPolicyList 创建res_policy列表实例
func NewPolicyList() *PolicyList {
	return &PolicyList{}
}

// TableName 指定表名
func (Policy) TableName() string {
	return "res_policy"
}

// GetByID 根据ID获取res_policy
func (m *Policy) GetByID(c context.Context, id int64) error {
	return app.DB().WithContext(c).First(m, id).Error
}

// Create 创建res_policy记录
func (m *Policy) Create(c context.Context) error {
	return app.DB().WithContext(c).Create(m).Error
}

// Update 更新res_policy记录
func (m *Policy) Update(c context.Context) error {
	return app.DB().WithContext(c).Save(m).Error
}

// Delete 软删除res_policy记录
func (m *Policy) Delete(c context.Context) error {
	return app.DB().WithContext(c).Delete(m).Error
}

// IsEmpty 检查模型是否为空
func (m *Policy) IsEmpty() bool {
	return m == nil || m.Id == 0
}

// Find 查询res_policy列表
func (l *PolicyList) Find(c context.Context, funcs ...func(*gorm.DB) *gorm.DB) error {
	return app.DB().WithContext(c).Model(&Policy{}).Scopes(funcs...).Find(l).Error
}

// GetTotal 获取res_policy总数
func (l *PolicyList) GetTotal(c context.Context, query ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := app.DB().WithContext(c).Model(&Policy{}).Scopes(query...).Count(&count).Error
	return count, err
}