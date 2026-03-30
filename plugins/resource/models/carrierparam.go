package models

import (
	"gin-fast/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CarrierListRequest res_carrier列表请求参数
type CarrierListRequest struct {
	models.BasePaging
	models.Validator
	CarrierNo *string `form:"carrierNo"` // 载体编码
	CarrierType *string `form:"carrierType"` // 载体类型
	CarrierName *string `form:"carrierName"` // 名称
	ParkLevel *string `form:"parkLevel"` // 园区级别
	CujuChanye *string `form:"cujuChanye"` // 6+5+X
	UpdatedBy *int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *CarrierListRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// Handle 获取查询条件
func (r *CarrierListRequest) Handle() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
        if r.CarrierNo != nil {
            // 默认等于查询
            db = db.Where("carrier_no = ?", *r.CarrierNo)
        }
        if r.CarrierType != nil {
            // 默认等于查询
            db = db.Where("carrier_type = ?", *r.CarrierType)
        }
        if r.CarrierName != nil {
            // 默认等于查询
            db = db.Where("carrier_name = ?", *r.CarrierName)
        }
        if r.ParkLevel != nil {
            // 默认等于查询
            db = db.Where("park_level = ?", *r.ParkLevel)
        }
        if r.CujuChanye != nil {
            // 默认等于查询
            db = db.Where("cuju_chanye = ?", *r.CujuChanye)
        }
        if r.UpdatedBy != nil {
            // 默认等于查询
            db = db.Where("updated_by = ?", *r.UpdatedBy)
        }
		return db
	}
}

// CarrierCreateRequest 创建res_carrier请求参数
type CarrierCreateRequest struct {
	models.Validator
	CarrierNo string `form:"carrierNo" validate:"required" message:"载体编码不能为空"` // 载体编码
	CarrierType string `form:"carrierType" validate:"required" message:"载体类型不能为空"` // 载体类型
	SecType string `form:"secType" validate:"required" message:"二级分类不能为空"` // 二级分类
	Location string `form:"location" validate:"required" message:"载体位置不能为空"` // 载体位置
	Address string `form:"address" validate:"required" message:"详细地址不能为空"` // 详细地址
	IndustryPark string `form:"industryPark" validate:"required" message:"园区等级不能为空"` // 园区等级
	CarrierName string `form:"carrierName" validate:"required" message:"名称不能为空"` // 名称
	CarrierPicture string `form:"carrierPicture" validate:"required" message:"载体图片不能为空"` // 载体图片
	CarrierIntro string `form:"carrierIntro" validate:"required" message:"介绍不能为空"` // 介绍
	Contact string `form:"contact" validate:"required" message:"联系人不能为空"` // 联系人
	Phone string `form:"phone" validate:"required" message:"联系方式不能为空"` // 联系方式
	StarCom string `form:"starCom" validate:"required" message:"明星企业不能为空"` // 明星企业
	Area string `form:"area" validate:"required" message:"面积不能为空"` // 面积
	FloorHeight int64 `form:"floorHeight"` // 层高
	Industry string `form:"industry" validate:"required" message:"产业领域不能为空"` // 产业领域
	AvailableLand string `form:"availableLand" validate:"required" message:"可用地情况不能为空"` // 可用地情况
	InvestmentArea string `form:"investmentArea" validate:"required" message:"可招商面积不能为空"` // 可招商面积
	Rent string `form:"rent" validate:"required" message:"租金不能为空"` // 租金
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	ParkType string `form:"parkType" validate:"required" message:"园区类型不能为空"` // 园区类型
	EndOfFloor int64 `form:"endOfFloor"` // 层高结束
	AvailabilityType string `form:"availabilityType" validate:"required" message:"可用地情况不能为空"` // 可用地情况
	LeadingIndustries string `form:"leadingIndustries" validate:"required" message:"主导产业不能为空"` // 主导产业
	TotalArea int64 `form:"totalArea"` // 总面积
	DataSourceId int64 `form:"dataSourceId"` // 对应数据源id
	ParcelArea int64 `form:"parcelArea"` // 地块面积
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	RentType string `form:"rentType" validate:"required" message:"租金类型不能为空"` // 租金类型
	MinimumRent int64 `form:"minimumRent"` // 租金最小值
	MaximumRent int64 `form:"maximumRent"` // 租金最大值
	LettableArea int64 `form:"lettableArea"` // 可租用面积
	DistrictCounty string `form:"districtCounty" validate:"required" message:"所在区县不能为空"` // 所在区县
	LandUse string `form:"landUse" validate:"required" message:"土地用途不能为空"` // 土地用途
	FloorHeightRange string `form:"floorHeightRange" validate:"required" message:"层高范围不能为空"` // 层高范围
	ParkLevel string `form:"parkLevel" validate:"required" message:"园区级别不能为空"` // 园区级别
	CujuChanye string `form:"cujuChanye" validate:"required" message:"6+5+X不能为空"` // 6+5+X
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *CarrierCreateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// CarrierUpdateRequest 更新res_carrier请求参数
type CarrierUpdateRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
	CarrierNo string `form:"carrierNo" validate:"required" message:"载体编码不能为空"` // 载体编码
	CarrierType string `form:"carrierType" validate:"required" message:"载体类型不能为空"` // 载体类型
	SecType string `form:"secType" validate:"required" message:"二级分类不能为空"` // 二级分类
	Location string `form:"location" validate:"required" message:"载体位置不能为空"` // 载体位置
	Address string `form:"address" validate:"required" message:"详细地址不能为空"` // 详细地址
	IndustryPark string `form:"industryPark" validate:"required" message:"园区等级不能为空"` // 园区等级
	CarrierName string `form:"carrierName" validate:"required" message:"名称不能为空"` // 名称
	CarrierPicture string `form:"carrierPicture" validate:"required" message:"载体图片不能为空"` // 载体图片
	CarrierIntro string `form:"carrierIntro" validate:"required" message:"介绍不能为空"` // 介绍
	Contact string `form:"contact" validate:"required" message:"联系人不能为空"` // 联系人
	Phone string `form:"phone" validate:"required" message:"联系方式不能为空"` // 联系方式
	StarCom string `form:"starCom" validate:"required" message:"明星企业不能为空"` // 明星企业
	Area string `form:"area" validate:"required" message:"面积不能为空"` // 面积
	FloorHeight int64 `form:"floorHeight"` // 层高
	Industry string `form:"industry" validate:"required" message:"产业领域不能为空"` // 产业领域
	AvailableLand string `form:"availableLand" validate:"required" message:"可用地情况不能为空"` // 可用地情况
	InvestmentArea string `form:"investmentArea" validate:"required" message:"可招商面积不能为空"` // 可招商面积
	Rent string `form:"rent" validate:"required" message:"租金不能为空"` // 租金
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	ParkType string `form:"parkType" validate:"required" message:"园区类型不能为空"` // 园区类型
	EndOfFloor int64 `form:"endOfFloor"` // 层高结束
	AvailabilityType string `form:"availabilityType" validate:"required" message:"可用地情况不能为空"` // 可用地情况
	LeadingIndustries string `form:"leadingIndustries" validate:"required" message:"主导产业不能为空"` // 主导产业
	TotalArea int64 `form:"totalArea"` // 总面积
	DataSourceId int64 `form:"dataSourceId"` // 对应数据源id
	ParcelArea int64 `form:"parcelArea"` // 地块面积
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	RentType string `form:"rentType" validate:"required" message:"租金类型不能为空"` // 租金类型
	MinimumRent int64 `form:"minimumRent"` // 租金最小值
	MaximumRent int64 `form:"maximumRent"` // 租金最大值
	LettableArea int64 `form:"lettableArea"` // 可租用面积
	DistrictCounty string `form:"districtCounty" validate:"required" message:"所在区县不能为空"` // 所在区县
	LandUse string `form:"landUse" validate:"required" message:"土地用途不能为空"` // 土地用途
	FloorHeightRange string `form:"floorHeightRange" validate:"required" message:"层高范围不能为空"` // 层高范围
	ParkLevel string `form:"parkLevel" validate:"required" message:"园区级别不能为空"` // 园区级别
	CujuChanye string `form:"cujuChanye" validate:"required" message:"6+5+X不能为空"` // 6+5+X
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *CarrierUpdateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// CarrierDeleteRequest 删除res_carrier请求参数
type CarrierDeleteRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *CarrierDeleteRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// CarrierGetByIDRequest 根据ID获取res_carrier请求参数
type CarrierGetByIDRequest struct {
	models.Validator
	Id int64 `uri:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *CarrierGetByIDRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}