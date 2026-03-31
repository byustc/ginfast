package models

import (
	"context"
	"time"
	"gin-fast/app/global/app"
	"gorm.io/gorm"
)

// Expert res_expert 模型结构体
type Expert struct {
	Id int64 `gorm:"column:id;primaryKey;not null;autoIncrement" json:"id"` // 主键
	TenantId uint `gorm:"column:tenant_id;default:0;index" json:"tenantId"` // 租户ID字段
	ExpertNo string `gorm:"column:expert_no;index" json:"expertNo"` // 专家编码
	Name string `gorm:"column:name" json:"name"` // 专家姓名
	Phone string `gorm:"column:phone" json:"phone"` // 联系方式
	Location string `gorm:"column:location" json:"location"` // 所在地
	Degree string `gorm:"column:degree" json:"degree"` // 学历学位
	Major string `gorm:"column:major" json:"major"` // 专业
	MajorLevel string `gorm:"column:major_level" json:"majorLevel"` // 专业职称等级
	MajorType string `gorm:"column:major_type" json:"majorType"` // 专业类别
	ResearchIndustry string `gorm:"column:research_industry" json:"researchIndustry"` // 研究方向
	WorkYears int `gorm:"column:work_years" json:"workYears"` // 工作年限
	WorkCom string `gorm:"column:work_com" json:"workCom"` // 工作单位
	Position string `gorm:"column:position" json:"position"` // 职位
	Industry string `gorm:"column:industry" json:"industry"` // 所属产业链
	IsLecturer string `gorm:"column:is_lecturer" json:"isLecturer"` // 是否能纳入培训讲师
	ExpertType string `gorm:"column:expert_type" json:"expertType"` // 专家类型
	ExpertLevel string `gorm:"column:expert_level" json:"expertLevel"` // 专家级别
	Source string `gorm:"column:source" json:"source"` // 专家来源
	Remark string `gorm:"column:remark" json:"remark"` // 备注
	DataSource string `gorm:"column:data_source" json:"dataSource"` // 数据来源
	IsBlacklist int `gorm:"column:is_blacklist;default:0" json:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `gorm:"column:tag" json:"tag"` // 行业数据标签
	IndustrySj string `gorm:"column:industry_sj" json:"industrySj"` // 所属产业
	IsPartner int `gorm:"column:is_partner" json:"isPartner"` // 是否合作伙伴: 0否 1是
	Profile string `gorm:"column:profile" json:"profile"` // 简介
	InstitutionType string `gorm:"column:institution_type" json:"institutionType"` // 机构类型
	ContactNumber string `gorm:"column:contact_number" json:"contactNumber"` // 联系电话
	AffiliatedInstitutions string `gorm:"column:affiliated_institutions" json:"affiliatedInstitutions"` // 所属机构
	EntryPerson int64 `gorm:"column:entry_person" json:"entryPerson"` // 录入人
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64 `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"` // 修改时间
	UpdatedBy int64 `gorm:"column:updated_by" json:"updatedBy"` // 修改人
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// ExpertList res_expert 列表
type ExpertList []Expert

// NewExpert 创建res_expert实例
func NewExpert() *Expert {
	return &Expert{}
}

// NewExpertList 创建res_expert列表实例
func NewExpertList() *ExpertList {
	return &ExpertList{}
}

// TableName 指定表名
func (Expert) TableName() string {
	return "res_expert"
}

// GetByID 根据ID获取res_expert
func (m *Expert) GetByID(c context.Context, id int64) error {
	return app.DB().WithContext(c).First(m, id).Error
}

// Create 创建res_expert记录
func (m *Expert) Create(c context.Context) error {
	return app.DB().WithContext(c).Create(m).Error
}

// Update 更新res_expert记录
func (m *Expert) Update(c context.Context) error {
	return app.DB().WithContext(c).Save(m).Error
}

// Delete 软删除res_expert记录
func (m *Expert) Delete(c context.Context) error {
	return app.DB().WithContext(c).Delete(m).Error
}

// IsEmpty 检查模型是否为空
func (m *Expert) IsEmpty() bool {
	return m == nil || m.Id == 0
}

// Find 查询res_expert列表
func (l *ExpertList) Find(c context.Context, funcs ...func(*gorm.DB) *gorm.DB) error {
	return app.DB().WithContext(c).Model(&Expert{}).Scopes(funcs...).Find(l).Error
}

// GetTotal 获取res_expert总数
func (l *ExpertList) GetTotal(c context.Context, query ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := app.DB().WithContext(c).Model(&Expert{}).Scopes(query...).Count(&count).Error
	return count, err
}