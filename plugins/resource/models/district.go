package models

import (
	"context"
	"time"
	"gin-fast/app/global/app"
	"gorm.io/gorm"
)

// District res_district 模型结构体
type District struct {
	Id int64 `gorm:"column:id;primaryKey;not null;autoIncrement" json:"id"` // 主键
	TenantId uint `gorm:"column:tenant_id;default:0;index" json:"tenantId"` // 租户ID字段
	DistrictNo string `gorm:"column:district_no;index" json:"districtNo"` // 区县编码
	Name string `gorm:"column:name" json:"name"` // 名称
	Introduction string `gorm:"column:introduction" json:"introduction"` // 区县介绍
	Tags string `gorm:"column:tags" json:"tags"` // 产业领域
	ContactPerson string `gorm:"column:contact_person" json:"contactPerson"` // 联系人
	Phone string `gorm:"column:phone" json:"phone"` // 联系方式
	Address string `gorm:"column:address" json:"address"` // 详细地址
	Remark string `gorm:"column:remark" json:"remark"` // 备注
	DataSource string `gorm:"column:data_source" json:"dataSource"` // 数据来源
	IsBlacklist int `gorm:"column:is_blacklist;default:0" json:"isBlacklist"` // 是否负面清单: 0否 1是
	DataSourceId int64 `gorm:"column:data_source_id" json:"dataSourceId"` // 对应数据源id
	Tag string `gorm:"column:tag" json:"tag"` // 行业数据标签
	CujuChanye string `gorm:"column:cuju_chanye" json:"cujuChanye"` // 6+5+X
	Industry string `gorm:"column:industry" json:"industry"` // 所属产业
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64 `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"` // 修改时间
	UpdatedBy int64 `gorm:"column:updated_by" json:"updatedBy"` // 修改人
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// DistrictList res_district 列表
type DistrictList []District

// NewDistrict 创建res_district实例
func NewDistrict() *District {
	return &District{}
}

// NewDistrictList 创建res_district列表实例
func NewDistrictList() *DistrictList {
	return &DistrictList{}
}

// TableName 指定表名
func (District) TableName() string {
	return "res_district"
}

// GetByID 根据ID获取res_district
func (m *District) GetByID(c context.Context, id int64) error {
	return app.DB().WithContext(c).First(m, id).Error
}

// Create 创建res_district记录
func (m *District) Create(c context.Context) error {
	return app.DB().WithContext(c).Create(m).Error
}

// Update 更新res_district记录
func (m *District) Update(c context.Context) error {
	return app.DB().WithContext(c).Save(m).Error
}

// Delete 软删除res_district记录
func (m *District) Delete(c context.Context) error {
	return app.DB().WithContext(c).Delete(m).Error
}

// IsEmpty 检查模型是否为空
func (m *District) IsEmpty() bool {
	return m == nil || m.Id == 0
}

// Find 查询res_district列表
func (l *DistrictList) Find(c context.Context, funcs ...func(*gorm.DB) *gorm.DB) error {
	return app.DB().WithContext(c).Model(&District{}).Scopes(funcs...).Find(l).Error
}

// GetTotal 获取res_district总数
func (l *DistrictList) GetTotal(c context.Context, query ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := app.DB().WithContext(c).Model(&District{}).Scopes(query...).Count(&count).Error
	return count, err
}