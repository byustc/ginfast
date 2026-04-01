package models

import (
	"gin-fast/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProjectResListRequest biz_project_res列表请求参数
type ProjectResListRequest struct {
	models.BasePaging
	models.Validator
	Id *int64 `form:"id"` // 主键
	ProjectId *int64 `form:"projectId"` // 项目id
	AgencyType *string `form:"agencyType"` // 资源类型
	FinanceId *int64 `form:"financeId"` // 金融机构
	InvestId *int64 `form:"investId"` // 投资机构
	ThinktankId *int64 `form:"thinktankId"` // 新研机构
	SceneId *int64 `form:"sceneId"` // 场景
	DistrictId *int64 `form:"districtId"` // 区县id
	CarrierId *int64 `form:"carrierId"` // 载体id
	CarrierInfoId *int64 `form:"carrierInfoId"` // 载体数据id
	Tag *string `form:"tag"` // 数据标签
	PolicyId *int64 `form:"policyId"` // 政策id
	UpdatedBy *int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ProjectResListRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// Handle 获取查询条件
func (r *ProjectResListRequest) Handle() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
        if r.Id != nil {
            // 默认等于查询
            db = db.Where("id = ?", *r.Id)
        }
        if r.ProjectId != nil {
            // 默认等于查询
            db = db.Where("project_id = ?", *r.ProjectId)
        }
        if r.AgencyType != nil {
            // 默认等于查询
            db = db.Where("agency_type = ?", *r.AgencyType)
        }
        if r.FinanceId != nil {
            // 默认等于查询
            db = db.Where("finance_id = ?", *r.FinanceId)
        }
        if r.InvestId != nil {
            // 默认等于查询
            db = db.Where("invest_id = ?", *r.InvestId)
        }
        if r.ThinktankId != nil {
            // 默认等于查询
            db = db.Where("thinktank_id = ?", *r.ThinktankId)
        }
        if r.SceneId != nil {
            // 默认等于查询
            db = db.Where("scene_id = ?", *r.SceneId)
        }
        if r.DistrictId != nil {
            // 默认等于查询
            db = db.Where("district_id = ?", *r.DistrictId)
        }
        if r.CarrierId != nil {
            // 默认等于查询
            db = db.Where("carrier_id = ?", *r.CarrierId)
        }
        if r.CarrierInfoId != nil {
            // 默认等于查询
            db = db.Where("carrier_info_id = ?", *r.CarrierInfoId)
        }
        if r.Tag != nil {
            // 默认等于查询
            db = db.Where("tag = ?", *r.Tag)
        }
        if r.PolicyId != nil {
            // 默认等于查询
            db = db.Where("policy_id = ?", *r.PolicyId)
        }
        if r.UpdatedBy != nil {
            // 默认等于查询
            db = db.Where("updated_by = ?", *r.UpdatedBy)
        }
		return db
	}
}

// ProjectResCreateRequest 创建biz_project_res请求参数
type ProjectResCreateRequest struct {
	models.Validator
	ProjectId int64 `form:"projectId"` // 项目id
	AgencyType string `form:"agencyType" validate:"required" message:"资源类型不能为空"` // 资源类型
	FinanceId int64 `form:"financeId"` // 金融机构
	InvestId int64 `form:"investId"` // 投资机构
	ThinktankId int64 `form:"thinktankId"` // 新研机构
	SceneId int64 `form:"sceneId"` // 场景
	DistrictId int64 `form:"districtId"` // 区县id
	CarrierId int64 `form:"carrierId"` // 载体id
	CarrierInfoId int64 `form:"carrierInfoId"` // 载体数据id
	Tag string `form:"tag" validate:"required" message:"数据标签不能为空"` // 数据标签
	PolicyId int64 `form:"policyId"` // 政策id
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ProjectResCreateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ProjectResUpdateRequest 更新biz_project_res请求参数
type ProjectResUpdateRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
	ProjectId int64 `form:"projectId"` // 项目id
	AgencyType string `form:"agencyType" validate:"required" message:"资源类型不能为空"` // 资源类型
	FinanceId int64 `form:"financeId"` // 金融机构
	InvestId int64 `form:"investId"` // 投资机构
	ThinktankId int64 `form:"thinktankId"` // 新研机构
	SceneId int64 `form:"sceneId"` // 场景
	DistrictId int64 `form:"districtId"` // 区县id
	CarrierId int64 `form:"carrierId"` // 载体id
	CarrierInfoId int64 `form:"carrierInfoId"` // 载体数据id
	Tag string `form:"tag" validate:"required" message:"数据标签不能为空"` // 数据标签
	PolicyId int64 `form:"policyId"` // 政策id
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ProjectResUpdateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ProjectResDeleteRequest 删除biz_project_res请求参数
type ProjectResDeleteRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *ProjectResDeleteRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ProjectResGetByIDRequest 根据ID获取biz_project_res请求参数
type ProjectResGetByIDRequest struct {
	models.Validator
	Id int64 `uri:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *ProjectResGetByIDRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}