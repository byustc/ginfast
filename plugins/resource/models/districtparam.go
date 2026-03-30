package models

import (
	"gin-fast/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DistrictListRequest res_district列表请求参数
type DistrictListRequest struct {
	models.BasePaging
	models.Validator
	Id *int64 `form:"id"` // 主键
	DistrictNo *string `form:"districtNo"` // 区县编码
	Name *string `form:"name"` // 名称
	Introduction *string `form:"introduction"` // 区县介绍
	Tags *string `form:"tags"` // 产业领域
	ContactPerson *string `form:"contactPerson"` // 联系人
	Phone *string `form:"phone"` // 联系方式
	Address *string `form:"address"` // 详细地址
	Remark *string `form:"remark"` // 备注
	DataSource *string `form:"dataSource"` // 数据来源
	IsBlacklist *int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	DataSourceId *int64 `form:"dataSourceId"` // 对应数据源id
	Tag *string `form:"tag"` // 行业数据标签
	CujuChanye *string `form:"cujuChanye"` // 6+5+X
	Industry *string `form:"industry"` // 所属产业
	UpdatedBy *int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *DistrictListRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// Handle 获取查询条件
func (r *DistrictListRequest) Handle() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
        if r.Id != nil {
            // 默认等于查询
            db = db.Where("id = ?", *r.Id)
        }
        if r.DistrictNo != nil {
            // 默认等于查询
            db = db.Where("district_no = ?", *r.DistrictNo)
        }
        if r.Name != nil {
            // 默认等于查询
            db = db.Where("name = ?", *r.Name)
        }
        if r.Introduction != nil {
            // 默认等于查询
            db = db.Where("introduction = ?", *r.Introduction)
        }
        if r.Tags != nil {
            // 默认等于查询
            db = db.Where("tags = ?", *r.Tags)
        }
        if r.ContactPerson != nil {
            // 默认等于查询
            db = db.Where("contact_person = ?", *r.ContactPerson)
        }
        if r.Phone != nil {
            // 默认等于查询
            db = db.Where("phone = ?", *r.Phone)
        }
        if r.Address != nil {
            // 默认等于查询
            db = db.Where("address = ?", *r.Address)
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
        if r.DataSourceId != nil {
            // 默认等于查询
            db = db.Where("data_source_id = ?", *r.DataSourceId)
        }
        if r.Tag != nil {
            // 默认等于查询
            db = db.Where("tag = ?", *r.Tag)
        }
        if r.CujuChanye != nil {
            // 默认等于查询
            db = db.Where("cuju_chanye = ?", *r.CujuChanye)
        }
        if r.Industry != nil {
            // 默认等于查询
            db = db.Where("industry = ?", *r.Industry)
        }
        if r.UpdatedBy != nil {
            // 默认等于查询
            db = db.Where("updated_by = ?", *r.UpdatedBy)
        }
		return db
	}
}

// DistrictCreateRequest 创建res_district请求参数
type DistrictCreateRequest struct {
	models.Validator
	DistrictNo string `form:"districtNo" validate:"required" message:"区县编码不能为空"` // 区县编码
	Name string `form:"name" validate:"required" message:"名称不能为空"` // 名称
	Introduction string `form:"introduction" validate:"required" message:"区县介绍不能为空"` // 区县介绍
	Tags string `form:"tags" validate:"required" message:"产业领域不能为空"` // 产业领域
	ContactPerson string `form:"contactPerson" validate:"required" message:"联系人不能为空"` // 联系人
	Phone string `form:"phone" validate:"required" message:"联系方式不能为空"` // 联系方式
	Address string `form:"address" validate:"required" message:"详细地址不能为空"` // 详细地址
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	DataSourceId int64 `form:"dataSourceId"` // 对应数据源id
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	CujuChanye string `form:"cujuChanye" validate:"required" message:"6+5+X不能为空"` // 6+5+X
	Industry string `form:"industry" validate:"required" message:"所属产业不能为空"` // 所属产业
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *DistrictCreateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// DistrictUpdateRequest 更新res_district请求参数
type DistrictUpdateRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
	DistrictNo string `form:"districtNo" validate:"required" message:"区县编码不能为空"` // 区县编码
	Name string `form:"name" validate:"required" message:"名称不能为空"` // 名称
	Introduction string `form:"introduction" validate:"required" message:"区县介绍不能为空"` // 区县介绍
	Tags string `form:"tags" validate:"required" message:"产业领域不能为空"` // 产业领域
	ContactPerson string `form:"contactPerson" validate:"required" message:"联系人不能为空"` // 联系人
	Phone string `form:"phone" validate:"required" message:"联系方式不能为空"` // 联系方式
	Address string `form:"address" validate:"required" message:"详细地址不能为空"` // 详细地址
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	DataSourceId int64 `form:"dataSourceId"` // 对应数据源id
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	CujuChanye string `form:"cujuChanye" validate:"required" message:"6+5+X不能为空"` // 6+5+X
	Industry string `form:"industry" validate:"required" message:"所属产业不能为空"` // 所属产业
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *DistrictUpdateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// DistrictDeleteRequest 删除res_district请求参数
type DistrictDeleteRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *DistrictDeleteRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// DistrictGetByIDRequest 根据ID获取res_district请求参数
type DistrictGetByIDRequest struct {
	models.Validator
	Id int64 `uri:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *DistrictGetByIDRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}