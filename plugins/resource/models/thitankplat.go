package models

import (
	"context"
	"time"
	"gin-fast/app/global/app"
	"gorm.io/gorm"
)

// ThitankPlat res_thinktank_plat 模型结构体
type ThitankPlat struct {
	Id int64 `gorm:"column:id;primaryKey;not null;autoIncrement" json:"id"` // 主键
	TenantId uint `gorm:"column:tenant_id;default:0;index" json:"tenantId"` // 租户ID字段
	PlatformName string `gorm:"column:platform_name" json:"platformName"` // 共性平台名称
	ThinktankId int64 `gorm:"column:thinktank_id" json:"thinktankId"` // 新研机构-数据主键id
	SupportFacilities string `gorm:"column:support_facilities" json:"supportFacilities"` // 配套设施
	PlatformType string `gorm:"column:platform_type" json:"platformType"` // 平台分类
	ServiceIndustry string `gorm:"column:service_industry" json:"serviceIndustry"` // 服务领域
	ServiceObject string `gorm:"column:service_object" json:"serviceObject"` // 服务对象
	ServiceCapability string `gorm:"column:service_capability" json:"serviceCapability"` // 服务能力
	ServiceContent string `gorm:"column:service_content" json:"serviceContent"` // 服务内容
	Qualification string `gorm:"column:qualification" json:"qualification"` // 已获资质情况
	Remark string `gorm:"column:remark" json:"remark"` // 备注
	DataSource string `gorm:"column:data_source" json:"dataSource"` // 数据来源
	IsBlacklist int `gorm:"column:is_blacklist;default:0" json:"isBlacklist"` // 是否负面清单: 0否 1是
	ServiceField string `gorm:"column:service_field" json:"serviceField"` // 服务产业
	Tag string `gorm:"column:tag" json:"tag"` // 行业数据标签
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64 `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"` // 修改时间
	UpdatedBy int64 `gorm:"column:updated_by" json:"updatedBy"` // 修改人
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// ThitankPlatList res_thinktank_plat 列表
type ThitankPlatList []ThitankPlat

// NewThitankPlat 创建res_thinktank_plat实例
func NewThitankPlat() *ThitankPlat {
	return &ThitankPlat{}
}

// NewThitankPlatList 创建res_thinktank_plat列表实例
func NewThitankPlatList() *ThitankPlatList {
	return &ThitankPlatList{}
}

// TableName 指定表名
func (ThitankPlat) TableName() string {
	return "res_thinktank_plat"
}

// GetByID 根据ID获取res_thinktank_plat
func (m *ThitankPlat) GetByID(c context.Context, id int64) error {
	return app.DB().WithContext(c).First(m, id).Error
}

// Create 创建res_thinktank_plat记录
func (m *ThitankPlat) Create(c context.Context) error {
	return app.DB().WithContext(c).Create(m).Error
}

// Update 更新res_thinktank_plat记录
func (m *ThitankPlat) Update(c context.Context) error {
	return app.DB().WithContext(c).Save(m).Error
}

// Delete 软删除res_thinktank_plat记录
func (m *ThitankPlat) Delete(c context.Context) error {
	return app.DB().WithContext(c).Delete(m).Error
}

// IsEmpty 检查模型是否为空
func (m *ThitankPlat) IsEmpty() bool {
	return m == nil || m.Id == 0
}

// Find 查询res_thinktank_plat列表
func (l *ThitankPlatList) Find(c context.Context, funcs ...func(*gorm.DB) *gorm.DB) error {
	return app.DB().WithContext(c).Model(&ThitankPlat{}).Scopes(funcs...).Find(l).Error
}

// GetTotal 获取res_thinktank_plat总数
func (l *ThitankPlatList) GetTotal(c context.Context, query ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := app.DB().WithContext(c).Model(&ThitankPlat{}).Scopes(query...).Count(&count).Error
	return count, err
}