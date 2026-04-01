package models

import (
	"time"
	"gin-fast/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProjectPhasePlanListRequest biz_project_phase_plan列表请求参数
type ProjectPhasePlanListRequest struct {
	models.BasePaging
	models.Validator
	Id *int64 `form:"id"` // 主键
	ProjectId *int64 `form:"projectId"` // 项目id
	ItemPhaseId *string `form:"itemPhaseId"` // 项目阶段
	PlanFinishTime *time.Time `form:"planFinishTime"` // 计划完成时间
	ActualFinishTime *time.Time `form:"actualFinishTime"` // 实际完成时间
	Tag *string `form:"tag"` // 数据标签
	UpdatedBy *int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ProjectPhasePlanListRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// Handle 获取查询条件
func (r *ProjectPhasePlanListRequest) Handle() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
        if r.Id != nil {
            // 默认等于查询
            db = db.Where("id = ?", *r.Id)
        }
        if r.ProjectId != nil {
            // 默认等于查询
            db = db.Where("project_id = ?", *r.ProjectId)
        }
        if r.ItemPhaseId != nil {
            // 默认等于查询
            db = db.Where("item_phase_id = ?", *r.ItemPhaseId)
        }
        if r.PlanFinishTime != nil {
            // 默认等于查询
            db = db.Where("plan_finish_time = ?", *r.PlanFinishTime)
        }
        if r.ActualFinishTime != nil {
            // 默认等于查询
            db = db.Where("actual_finish_time = ?", *r.ActualFinishTime)
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

// ProjectPhasePlanCreateRequest 创建biz_project_phase_plan请求参数
type ProjectPhasePlanCreateRequest struct {
	models.Validator
	ProjectId int64 `form:"projectId"` // 项目id
	ItemPhaseId string `form:"itemPhaseId" validate:"required" message:"项目阶段不能为空"` // 项目阶段
	PlanFinishTime *time.Time `form:"planFinishTime" validate:"required" message:"计划完成时间不能为空"` // 计划完成时间
	ActualFinishTime *time.Time `form:"actualFinishTime" validate:"required" message:"实际完成时间不能为空"` // 实际完成时间
	Tag string `form:"tag" validate:"required" message:"数据标签不能为空"` // 数据标签
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ProjectPhasePlanCreateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ProjectPhasePlanUpdateRequest 更新biz_project_phase_plan请求参数
type ProjectPhasePlanUpdateRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
	ProjectId int64 `form:"projectId"` // 项目id
	ItemPhaseId string `form:"itemPhaseId" validate:"required" message:"项目阶段不能为空"` // 项目阶段
	PlanFinishTime *time.Time `form:"planFinishTime" validate:"required" message:"计划完成时间不能为空"` // 计划完成时间
	ActualFinishTime *time.Time `form:"actualFinishTime" validate:"required" message:"实际完成时间不能为空"` // 实际完成时间
	Tag string `form:"tag" validate:"required" message:"数据标签不能为空"` // 数据标签
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ProjectPhasePlanUpdateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ProjectPhasePlanDeleteRequest 删除biz_project_phase_plan请求参数
type ProjectPhasePlanDeleteRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *ProjectPhasePlanDeleteRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ProjectPhasePlanGetByIDRequest 根据ID获取biz_project_phase_plan请求参数
type ProjectPhasePlanGetByIDRequest struct {
	models.Validator
	Id int64 `uri:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *ProjectPhasePlanGetByIDRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}