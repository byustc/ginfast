package models

import (
	"gin-fast/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IndustryChainListRequest res_industry_chain列表请求参数
type IndustryChainListRequest struct {
	models.BasePaging
	models.Validator
	Id *int64 `form:"id"` // 主键
	IndustryChainNo *string `form:"industryChainNo"` // 产业链编码
	IndustryChainName *string `form:"industryChainName"` // 产业链名称
	IndustryChainIntro *string `form:"industryChainIntro"` // 产业链介绍
	IndustryChainMap *string `form:"industryChainMap"` // 产业链图谱
	DataSource *string `form:"dataSource"` // 数据来源
	IsBlacklist *int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	GraphLayer *string `form:"graphLayer"` // 图谱配置
	Tag *string `form:"tag"` // 数据标签
	UpdatedBy *int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *IndustryChainListRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// Handle 获取查询条件
func (r *IndustryChainListRequest) Handle() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
        if r.Id != nil {
            // 默认等于查询
            db = db.Where("id = ?", *r.Id)
        }
        if r.IndustryChainNo != nil {
            // 默认等于查询
            db = db.Where("industry_chain_no = ?", *r.IndustryChainNo)
        }
        if r.IndustryChainName != nil {
            // 默认等于查询
            db = db.Where("industry_chain_name = ?", *r.IndustryChainName)
        }
        if r.IndustryChainIntro != nil {
            // 默认等于查询
            db = db.Where("industry_chain_intro = ?", *r.IndustryChainIntro)
        }
        if r.IndustryChainMap != nil {
            // 默认等于查询
            db = db.Where("industry_chain_map = ?", *r.IndustryChainMap)
        }
        if r.DataSource != nil {
            // 默认等于查询
            db = db.Where("data_source = ?", *r.DataSource)
        }
        if r.IsBlacklist != nil {
            // 默认等于查询
            db = db.Where("is_blacklist = ?", *r.IsBlacklist)
        }
        if r.GraphLayer != nil {
            // 默认等于查询
            db = db.Where("graph_layer = ?", *r.GraphLayer)
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

// IndustryChainCreateRequest 创建res_industry_chain请求参数
type IndustryChainCreateRequest struct {
	models.Validator
	IndustryChainNo string `form:"industryChainNo" validate:"required" message:"产业链编码不能为空"` // 产业链编码
	IndustryChainName string `form:"industryChainName" validate:"required" message:"产业链名称不能为空"` // 产业链名称
	IndustryChainIntro string `form:"industryChainIntro" validate:"required" message:"产业链介绍不能为空"` // 产业链介绍
	IndustryChainMap string `form:"industryChainMap" validate:"required" message:"产业链图谱不能为空"` // 产业链图谱
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	GraphLayer string `form:"graphLayer" validate:"required" message:"图谱配置不能为空"` // 图谱配置
	Tag string `form:"tag" validate:"required" message:"数据标签不能为空"` // 数据标签
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *IndustryChainCreateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// IndustryChainUpdateRequest 更新res_industry_chain请求参数
type IndustryChainUpdateRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
	IndustryChainNo string `form:"industryChainNo" validate:"required" message:"产业链编码不能为空"` // 产业链编码
	IndustryChainName string `form:"industryChainName" validate:"required" message:"产业链名称不能为空"` // 产业链名称
	IndustryChainIntro string `form:"industryChainIntro" validate:"required" message:"产业链介绍不能为空"` // 产业链介绍
	IndustryChainMap string `form:"industryChainMap" validate:"required" message:"产业链图谱不能为空"` // 产业链图谱
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	GraphLayer string `form:"graphLayer" validate:"required" message:"图谱配置不能为空"` // 图谱配置
	Tag string `form:"tag" validate:"required" message:"数据标签不能为空"` // 数据标签
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *IndustryChainUpdateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// IndustryChainDeleteRequest 删除res_industry_chain请求参数
type IndustryChainDeleteRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *IndustryChainDeleteRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// IndustryChainGetByIDRequest 根据ID获取res_industry_chain请求参数
type IndustryChainGetByIDRequest struct {
	models.Validator
	Id int64 `uri:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *IndustryChainGetByIDRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}