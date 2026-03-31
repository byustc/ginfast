package models

import (
	"gin-fast/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ThitankPlatListRequest res_thinktank_plat列表请求参数
type ThitankPlatListRequest struct {
	models.BasePaging
	models.Validator
	Id *int64 `form:"id"` // 主键
	PlatformName *string `form:"platformName"` // 共性平台名称
	ThinktankId *int64 `form:"thinktankId"` // 新研机构-数据主键id
	SupportFacilities *string `form:"supportFacilities"` // 配套设施
	PlatformType *string `form:"platformType"` // 平台分类
	ServiceIndustry *string `form:"serviceIndustry"` // 服务领域
	ServiceObject *string `form:"serviceObject"` // 服务对象
	ServiceCapability *string `form:"serviceCapability"` // 服务能力
	ServiceContent *string `form:"serviceContent"` // 服务内容
	Qualification *string `form:"qualification"` // 已获资质情况
	Remark *string `form:"remark"` // 备注
	DataSource *string `form:"dataSource"` // 数据来源
	IsBlacklist *int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	ServiceField *string `form:"serviceField"` // 服务产业
	Tag *string `form:"tag"` // 行业数据标签
	UpdatedBy *int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ThitankPlatListRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// Handle 获取查询条件
func (r *ThitankPlatListRequest) Handle() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
        if r.Id != nil {
            // 默认等于查询
            db = db.Where("id = ?", *r.Id)
        }
        if r.PlatformName != nil {
            // 默认等于查询
            db = db.Where("platform_name = ?", *r.PlatformName)
        }
        if r.ThinktankId != nil {
            // 默认等于查询
            db = db.Where("thinktank_id = ?", *r.ThinktankId)
        }
        if r.SupportFacilities != nil {
            // 默认等于查询
            db = db.Where("support_facilities = ?", *r.SupportFacilities)
        }
        if r.PlatformType != nil {
            // 默认等于查询
            db = db.Where("platform_type = ?", *r.PlatformType)
        }
        if r.ServiceIndustry != nil {
            // 默认等于查询
            db = db.Where("service_industry = ?", *r.ServiceIndustry)
        }
        if r.ServiceObject != nil {
            // 默认等于查询
            db = db.Where("service_object = ?", *r.ServiceObject)
        }
        if r.ServiceCapability != nil {
            // 默认等于查询
            db = db.Where("service_capability = ?", *r.ServiceCapability)
        }
        if r.ServiceContent != nil {
            // 默认等于查询
            db = db.Where("service_content = ?", *r.ServiceContent)
        }
        if r.Qualification != nil {
            // 默认等于查询
            db = db.Where("qualification = ?", *r.Qualification)
        }
        if r.Remark != nil {
            // 默认等于查询
            db = db.Where("remark = ?", *r.Remark)
        }
        if r.DataSource != nil {
            // 默认等于查询
            db = db.Where("data_source = ?", *r.DataSource)
        }
        if r.IsBlacklist != nil {
            // 默认等于查询
            db = db.Where("is_blacklist = ?", *r.IsBlacklist)
        }
        if r.ServiceField != nil {
            // 默认等于查询
            db = db.Where("service_field = ?", *r.ServiceField)
        }
        if r.Tag != nil {
            // 默认等于查询
            db = db.Where("tag = ?", *r.Tag)
        }
        if r.UpdatedBy != nil {
            // 默认等于查询
            db = db.Where("updated_by = ?", *r.UpdatedBy)
        }
		return db
	}
}

// ThitankPlatCreateRequest 创建res_thinktank_plat请求参数
type ThitankPlatCreateRequest struct {
	models.Validator
	PlatformName string `form:"platformName" validate:"required" message:"共性平台名称不能为空"` // 共性平台名称
	ThinktankId int64 `form:"thinktankId"` // 新研机构-数据主键id
	SupportFacilities string `form:"supportFacilities" validate:"required" message:"配套设施不能为空"` // 配套设施
	PlatformType string `form:"platformType" validate:"required" message:"平台分类不能为空"` // 平台分类
	ServiceIndustry string `form:"serviceIndustry" validate:"required" message:"服务领域不能为空"` // 服务领域
	ServiceObject string `form:"serviceObject" validate:"required" message:"服务对象不能为空"` // 服务对象
	ServiceCapability string `form:"serviceCapability" validate:"required" message:"服务能力不能为空"` // 服务能力
	ServiceContent string `form:"serviceContent" validate:"required" message:"服务内容不能为空"` // 服务内容
	Qualification string `form:"qualification" validate:"required" message:"已获资质情况不能为空"` // 已获资质情况
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	ServiceField string `form:"serviceField" validate:"required" message:"服务产业不能为空"` // 服务产业
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ThitankPlatCreateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ThitankPlatUpdateRequest 更新res_thinktank_plat请求参数
type ThitankPlatUpdateRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
	PlatformName string `form:"platformName" validate:"required" message:"共性平台名称不能为空"` // 共性平台名称
	ThinktankId int64 `form:"thinktankId"` // 新研机构-数据主键id
	SupportFacilities string `form:"supportFacilities" validate:"required" message:"配套设施不能为空"` // 配套设施
	PlatformType string `form:"platformType" validate:"required" message:"平台分类不能为空"` // 平台分类
	ServiceIndustry string `form:"serviceIndustry" validate:"required" message:"服务领域不能为空"` // 服务领域
	ServiceObject string `form:"serviceObject" validate:"required" message:"服务对象不能为空"` // 服务对象
	ServiceCapability string `form:"serviceCapability" validate:"required" message:"服务能力不能为空"` // 服务能力
	ServiceContent string `form:"serviceContent" validate:"required" message:"服务内容不能为空"` // 服务内容
	Qualification string `form:"qualification" validate:"required" message:"已获资质情况不能为空"` // 已获资质情况
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	ServiceField string `form:"serviceField" validate:"required" message:"服务产业不能为空"` // 服务产业
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ThitankPlatUpdateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ThitankPlatDeleteRequest 删除res_thinktank_plat请求参数
type ThitankPlatDeleteRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *ThitankPlatDeleteRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ThitankPlatGetByIDRequest 根据ID获取res_thinktank_plat请求参数
type ThitankPlatGetByIDRequest struct {
	models.Validator
	Id int64 `uri:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *ThitankPlatGetByIDRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}