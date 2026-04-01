package models

import (
	"gin-fast/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProjectPartnerListRequest biz_project_partner列表请求参数
type ProjectPartnerListRequest struct {
	models.BasePaging
	models.Validator
	Id *int64 `form:"id"` // 主键
	ProjectId *int64 `form:"projectId"` // 项目id
	RelationProject *int64 `form:"relationProject"` // 关联项目
	RelationPartner *int64 `form:"relationPartner"` // 关联合作伙伴
	Organization *string `form:"organization"` // 机构
	InvestmentAmount *float64 `form:"investmentAmount"` // 投资金额
	Unit *string `form:"unit"` // 单位
	Tag *string `form:"tag"` // 数据标签
	UpdatedBy *int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ProjectPartnerListRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// Handle 获取查询条件
func (r *ProjectPartnerListRequest) Handle() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
        if r.Id != nil {
            // 默认等于查询
            db = db.Where("id = ?", *r.Id)
        }
        if r.ProjectId != nil {
            // 默认等于查询
            db = db.Where("project_id = ?", *r.ProjectId)
        }
        if r.RelationProject != nil {
            // 默认等于查询
            db = db.Where("relation_project = ?", *r.RelationProject)
        }
        if r.RelationPartner != nil {
            // 默认等于查询
            db = db.Where("relation_partner = ?", *r.RelationPartner)
        }
        if r.Organization != nil {
            // 默认等于查询
            db = db.Where("organization = ?", *r.Organization)
        }
        if r.InvestmentAmount != nil {
            // 默认等于查询
            db = db.Where("investment_amount = ?", *r.InvestmentAmount)
        }
        if r.Unit != nil {
            // 默认等于查询
            db = db.Where("unit = ?", *r.Unit)
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

// ProjectPartnerCreateRequest 创建biz_project_partner请求参数
type ProjectPartnerCreateRequest struct {
	models.Validator
	ProjectId int64 `form:"projectId"` // 项目id
	RelationProject int64 `form:"relationProject"` // 关联项目
	RelationPartner int64 `form:"relationPartner"` // 关联合作伙伴
	Organization string `form:"organization" validate:"required" message:"机构不能为空"` // 机构
	InvestmentAmount float64 `form:"investmentAmount"` // 投资金额
	Unit string `form:"unit" validate:"required" message:"单位不能为空"` // 单位
	Tag string `form:"tag" validate:"required" message:"数据标签不能为空"` // 数据标签
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ProjectPartnerCreateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ProjectPartnerUpdateRequest 更新biz_project_partner请求参数
type ProjectPartnerUpdateRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
	ProjectId int64 `form:"projectId"` // 项目id
	RelationProject int64 `form:"relationProject"` // 关联项目
	RelationPartner int64 `form:"relationPartner"` // 关联合作伙伴
	Organization string `form:"organization" validate:"required" message:"机构不能为空"` // 机构
	InvestmentAmount float64 `form:"investmentAmount"` // 投资金额
	Unit string `form:"unit" validate:"required" message:"单位不能为空"` // 单位
	Tag string `form:"tag" validate:"required" message:"数据标签不能为空"` // 数据标签
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ProjectPartnerUpdateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ProjectPartnerDeleteRequest 删除biz_project_partner请求参数
type ProjectPartnerDeleteRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *ProjectPartnerDeleteRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ProjectPartnerGetByIDRequest 根据ID获取biz_project_partner请求参数
type ProjectPartnerGetByIDRequest struct {
	models.Validator
	Id int64 `uri:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *ProjectPartnerGetByIDRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}