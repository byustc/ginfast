package models

import (
	"context"
	"time"
	"gin-fast/app/global/app"
	"gorm.io/gorm"
)

// Carrier res_carrier 模型结构体
type Carrier struct {
	Id int64 `gorm:"column:id;primaryKey;not null;autoIncrement" json:"id"` // 主键
	TenantId uint `gorm:"column:tenant_id;default:0;index" json:"tenantId"` // 租户ID字段
	CarrierNo string `gorm:"column:carrier_no;index" json:"carrierNo"` // 载体编码
	CarrierType string `gorm:"column:carrier_type" json:"carrierType"` // 载体类型
	SecType string `gorm:"column:sec_type" json:"secType"` // 二级分类
	Location string `gorm:"column:location" json:"location"` // 载体位置
	Address string `gorm:"column:address" json:"address"` // 详细地址
	IndustryPark string `gorm:"column:industry_park" json:"industryPark"` // 园区等级
	CarrierName string `gorm:"column:carrier_name" json:"carrierName"` // 名称
	CarrierPicture string `gorm:"column:carrier_picture" json:"carrierPicture"` // 载体图片
	CarrierIntro string `gorm:"column:carrier_intro" json:"carrierIntro"` // 介绍
	Contact string `gorm:"column:contact" json:"contact"` // 联系人
	Phone string `gorm:"column:phone" json:"phone"` // 联系方式
	StarCom string `gorm:"column:star_com" json:"starCom"` // 明星企业
	Area string `gorm:"column:area" json:"area"` // 面积
	FloorHeight int64 `gorm:"column:floor_height" json:"floorHeight"` // 层高
	Industry string `gorm:"column:industry" json:"industry"` // 产业领域
	AvailableLand string `gorm:"column:available_land" json:"availableLand"` // 可用地情况
	InvestmentArea string `gorm:"column:investment_area" json:"investmentArea"` // 可招商面积
	Rent string `gorm:"column:rent" json:"rent"` // 租金
	Remark string `gorm:"column:remark" json:"remark"` // 备注
	DataSource string `gorm:"column:data_source" json:"dataSource"` // 数据来源
	IsBlacklist int `gorm:"column:is_blacklist;default:0" json:"isBlacklist"` // 是否负面清单: 0否 1是
	ParkType string `gorm:"column:park_type" json:"parkType"` // 园区类型
	EndOfFloor int64 `gorm:"column:end_of_floor" json:"endOfFloor"` // 层高结束
	AvailabilityType string `gorm:"column:availability_type" json:"availabilityType"` // 可用地情况
	LeadingIndustries string `gorm:"column:leading_industries" json:"leadingIndustries"` // 主导产业
	TotalArea int64 `gorm:"column:total_area" json:"totalArea"` // 总面积
	DataSourceId int64 `gorm:"column:data_source_id" json:"dataSourceId"` // 对应数据源id
	ParcelArea int64 `gorm:"column:parcel_area" json:"parcelArea"` // 地块面积
	Tag string `gorm:"column:tag" json:"tag"` // 行业数据标签
	RentType string `gorm:"column:rent_type" json:"rentType"` // 租金类型
	MinimumRent int64 `gorm:"column:minimum_rent" json:"minimumRent"` // 租金最小值
	MaximumRent int64 `gorm:"column:maximum_rent" json:"maximumRent"` // 租金最大值
	LettableArea int64 `gorm:"column:lettable_area" json:"lettableArea"` // 可租用面积
	DistrictCounty string `gorm:"column:district_county" json:"districtCounty"` // 所在区县
	LandUse string `gorm:"column:land_use" json:"landUse"` // 土地用途
	FloorHeightRange string `gorm:"column:floor_height_range" json:"floorHeightRange"` // 层高范围
	ParkLevel string `gorm:"column:park_level" json:"parkLevel"` // 园区级别
	CujuChanye string `gorm:"column:cuju_chanye" json:"cujuChanye"` // 6+5+X
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64 `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"` // 修改时间
	UpdatedBy int64 `gorm:"column:updated_by" json:"updatedBy"` // 修改人
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// CarrierList res_carrier 列表
type CarrierList []Carrier

// NewCarrier 创建res_carrier实例
func NewCarrier() *Carrier {
	return &Carrier{}
}

// NewCarrierList 创建res_carrier列表实例
func NewCarrierList() *CarrierList {
	return &CarrierList{}
}

// TableName 指定表名
func (Carrier) TableName() string {
	return "res_carrier"
}

// GetByID 根据ID获取res_carrier
func (m *Carrier) GetByID(c context.Context, id int64) error {
	return app.DB().WithContext(c).First(m, id).Error
}

// Create 创建res_carrier记录
func (m *Carrier) Create(c context.Context) error {
	return app.DB().WithContext(c).Create(m).Error
}

// Update 更新res_carrier记录
func (m *Carrier) Update(c context.Context) error {
	return app.DB().WithContext(c).Save(m).Error
}

// Delete 软删除res_carrier记录
func (m *Carrier) Delete(c context.Context) error {
	return app.DB().WithContext(c).Delete(m).Error
}

// IsEmpty 检查模型是否为空
func (m *Carrier) IsEmpty() bool {
	return m == nil || m.Id == 0
}

// Find 查询res_carrier列表
func (l *CarrierList) Find(c context.Context, funcs ...func(*gorm.DB) *gorm.DB) error {
	return app.DB().WithContext(c).Model(&Carrier{}).Scopes(funcs...).Find(l).Error
}

// GetTotal 获取res_carrier总数
func (l *CarrierList) GetTotal(c context.Context, query ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := app.DB().WithContext(c).Model(&Carrier{}).Scopes(query...).Count(&count).Error
	return count, err
}