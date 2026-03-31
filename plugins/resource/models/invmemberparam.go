package models

import (
	"gin-fast/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InvMemberListRequest res_invest_member列表请求参数
type InvMemberListRequest struct {
	models.BasePaging
	models.Validator
	Id *int64 `form:"id"` // 主键
	InvestId *int64 `form:"investId"` // 投资机构-数据主键id
	PersonName *string `form:"personName"` // 姓名
	Position *string `form:"position"` // 职位
	PersonIntro *string `form:"personIntro"` // 个人介绍
	RankNum *float64 `form:"rankNum"` // 排序字段
	Phone *string `form:"phone"` // 联系方式
	EduBackground *string `form:"eduBackground"` // 学历
	DataSource *string `form:"dataSource"` // 数据来源
	IsBlacklist *int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag *string `form:"tag"` // 数据标签
	UpdatedBy *int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *InvMemberListRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// Handle 获取查询条件
func (r *InvMemberListRequest) Handle() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
        if r.Id != nil {
            // 默认等于查询
            db = db.Where("id = ?", *r.Id)
        }
        if r.InvestId != nil {
            // 默认等于查询
            db = db.Where("invest_id = ?", *r.InvestId)
        }
        if r.PersonName != nil {
            // 默认等于查询
            db = db.Where("person_name = ?", *r.PersonName)
        }
        if r.Position != nil {
            // 默认等于查询
            db = db.Where("position = ?", *r.Position)
        }
        if r.PersonIntro != nil {
            // 默认等于查询
            db = db.Where("person_intro = ?", *r.PersonIntro)
        }
        if r.RankNum != nil {
            // 默认等于查询
            db = db.Where("rank_num = ?", *r.RankNum)
        }
        if r.Phone != nil {
            // 默认等于查询
            db = db.Where("phone = ?", *r.Phone)
        }
        if r.EduBackground != nil {
            // 默认等于查询
            db = db.Where("edu_background = ?", *r.EduBackground)
        }
        if r.DataSource != nil {
            // 默认等于查询
            db = db.Where("data_source = ?", *r.DataSource)
        }
        if r.IsBlacklist != nil {
            // 默认等于查询
            db = db.Where("is_blacklist = ?", *r.IsBlacklist)
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

// InvMemberCreateRequest 创建res_invest_member请求参数
type InvMemberCreateRequest struct {
	models.Validator
	InvestId int64 `form:"investId"` // 投资机构-数据主键id
	PersonName string `form:"personName" validate:"required" message:"姓名不能为空"` // 姓名
	Position string `form:"position" validate:"required" message:"职位不能为空"` // 职位
	PersonIntro string `form:"personIntro" validate:"required" message:"个人介绍不能为空"` // 个人介绍
	RankNum float64 `form:"rankNum"` // 排序字段
	Phone string `form:"phone" validate:"required" message:"联系方式不能为空"` // 联系方式
	EduBackground string `form:"eduBackground" validate:"required" message:"学历不能为空"` // 学历
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `form:"tag" validate:"required" message:"数据标签不能为空"` // 数据标签
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *InvMemberCreateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// InvMemberUpdateRequest 更新res_invest_member请求参数
type InvMemberUpdateRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
	InvestId int64 `form:"investId"` // 投资机构-数据主键id
	PersonName string `form:"personName" validate:"required" message:"姓名不能为空"` // 姓名
	Position string `form:"position" validate:"required" message:"职位不能为空"` // 职位
	PersonIntro string `form:"personIntro" validate:"required" message:"个人介绍不能为空"` // 个人介绍
	RankNum float64 `form:"rankNum"` // 排序字段
	Phone string `form:"phone" validate:"required" message:"联系方式不能为空"` // 联系方式
	EduBackground string `form:"eduBackground" validate:"required" message:"学历不能为空"` // 学历
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `form:"tag" validate:"required" message:"数据标签不能为空"` // 数据标签
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *InvMemberUpdateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// InvMemberDeleteRequest 删除res_invest_member请求参数
type InvMemberDeleteRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *InvMemberDeleteRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// InvMemberGetByIDRequest 根据ID获取res_invest_member请求参数
type InvMemberGetByIDRequest struct {
	models.Validator
	Id int64 `uri:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *InvMemberGetByIDRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}