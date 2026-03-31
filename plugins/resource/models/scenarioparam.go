package models

import (
	"time"
	"gin-fast/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ScenarioListRequest res_scenario列表请求参数
type ScenarioListRequest struct {
	models.BasePaging
	models.Validator
	Id *int64 `form:"id"` // 主键
	ScenarioNo *string `form:"scenarioNo"` // 场景编码
	PublishTime *time.Time `form:"publishTime"` // 发布时间
	OwnerUnit *string `form:"ownerUnit"` // 业主单位
	SceneName *string `form:"sceneName"` // 场景名称
	Industry *string `form:"industry"` // 所属产业
	OpportunityZone *string `form:"opportunityZone"` // 机会区域
	ValidityPeriod *string `form:"validityPeriod"` // 机会有效期
	BudgetGuarantee *string `form:"budgetGuarantee"` // 预算保障情况
	OtherSituations *string `form:"otherSituations"` // 其他情况
	OpportunityDesc *string `form:"opportunityDesc"` // 机会描述
	BuildBasement *string `form:"buildBasement"` // 建设基础
	CooperationRequirements *string `form:"cooperationRequirements"` // 合作需求
	CooperationForm *string `form:"cooperationForm"` // 合作形式
	ContactPerson *string `form:"contactPerson"` // 联系人
	Position *string `form:"position"` // 职务
	Phone *string `form:"phone"` // 联系方式
	Remark *string `form:"remark"` // 备注
	DataSource *string `form:"dataSource"` // 数据来源
	IsBlacklist *int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag *string `form:"tag"` // 行业数据标签
	IndustrialFields *string `form:"industrialFields"` // 产业领域
	XcujuChanye *string `form:"xcujuChanye"` // 6+5+X
	UpdatedBy *int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ScenarioListRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// Handle 获取查询条件
func (r *ScenarioListRequest) Handle() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
        if r.Id != nil {
            // 默认等于查询
            db = db.Where("id = ?", *r.Id)
        }
        if r.ScenarioNo != nil {
            // 默认等于查询
            db = db.Where("scenario_no = ?", *r.ScenarioNo)
        }
        if r.PublishTime != nil {
            // 默认等于查询
            db = db.Where("publish_time = ?", *r.PublishTime)
        }
        if r.OwnerUnit != nil {
            // 默认等于查询
            db = db.Where("owner_unit = ?", *r.OwnerUnit)
        }
        if r.SceneName != nil {
            // 默认等于查询
            db = db.Where("scene_name = ?", *r.SceneName)
        }
        if r.Industry != nil {
            // 默认等于查询
            db = db.Where("industry = ?", *r.Industry)
        }
        if r.OpportunityZone != nil {
            // 默认等于查询
            db = db.Where("opportunity_zone = ?", *r.OpportunityZone)
        }
        if r.ValidityPeriod != nil {
            // 默认等于查询
            db = db.Where("validity_period = ?", *r.ValidityPeriod)
        }
        if r.BudgetGuarantee != nil {
            // 默认等于查询
            db = db.Where("budget_guarantee = ?", *r.BudgetGuarantee)
        }
        if r.OtherSituations != nil {
            // 默认等于查询
            db = db.Where("other_situations = ?", *r.OtherSituations)
        }
        if r.OpportunityDesc != nil {
            // 默认等于查询
            db = db.Where("opportunity_desc = ?", *r.OpportunityDesc)
        }
        if r.BuildBasement != nil {
            // 默认等于查询
            db = db.Where("build_basement = ?", *r.BuildBasement)
        }
        if r.CooperationRequirements != nil {
            // 默认等于查询
            db = db.Where("cooperation_requirements = ?", *r.CooperationRequirements)
        }
        if r.CooperationForm != nil {
            // 默认等于查询
            db = db.Where("cooperation_form = ?", *r.CooperationForm)
        }
        if r.ContactPerson != nil {
            // 默认等于查询
            db = db.Where("contact_person = ?", *r.ContactPerson)
        }
        if r.Position != nil {
            // 默认等于查询
            db = db.Where("position = ?", *r.Position)
        }
        if r.Phone != nil {
            // 默认等于查询
            db = db.Where("phone = ?", *r.Phone)
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
        if r.Tag != nil {
            // 默认等于查询
            db = db.Where("tag = ?", *r.Tag)
        }
        if r.IndustrialFields != nil {
            // 默认等于查询
            db = db.Where("industrial_fields = ?", *r.IndustrialFields)
        }
        if r.XcujuChanye != nil {
            // 默认等于查询
            db = db.Where("xcuju_chanye = ?", *r.XcujuChanye)
        }
        if r.UpdatedBy != nil {
            // 默认等于查询
            db = db.Where("updated_by = ?", *r.UpdatedBy)
        }
		return db
	}
}

// ScenarioCreateRequest 创建res_scenario请求参数
type ScenarioCreateRequest struct {
	models.Validator
	ScenarioNo string `form:"scenarioNo" validate:"required" message:"场景编码不能为空"` // 场景编码
	PublishTime *time.Time `form:"publishTime" validate:"required" message:"发布时间不能为空"` // 发布时间
	OwnerUnit string `form:"ownerUnit" validate:"required" message:"业主单位不能为空"` // 业主单位
	SceneName string `form:"sceneName" validate:"required" message:"场景名称不能为空"` // 场景名称
	Industry string `form:"industry" validate:"required" message:"所属产业不能为空"` // 所属产业
	OpportunityZone string `form:"opportunityZone" validate:"required" message:"机会区域不能为空"` // 机会区域
	ValidityPeriod string `form:"validityPeriod" validate:"required" message:"机会有效期不能为空"` // 机会有效期
	BudgetGuarantee string `form:"budgetGuarantee" validate:"required" message:"预算保障情况不能为空"` // 预算保障情况
	OtherSituations string `form:"otherSituations" validate:"required" message:"其他情况不能为空"` // 其他情况
	OpportunityDesc string `form:"opportunityDesc" validate:"required" message:"机会描述不能为空"` // 机会描述
	BuildBasement string `form:"buildBasement" validate:"required" message:"建设基础不能为空"` // 建设基础
	CooperationRequirements string `form:"cooperationRequirements" validate:"required" message:"合作需求不能为空"` // 合作需求
	CooperationForm string `form:"cooperationForm" validate:"required" message:"合作形式不能为空"` // 合作形式
	ContactPerson string `form:"contactPerson" validate:"required" message:"联系人不能为空"` // 联系人
	Position string `form:"position" validate:"required" message:"职务不能为空"` // 职务
	Phone string `form:"phone" validate:"required" message:"联系方式不能为空"` // 联系方式
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	IndustrialFields string `form:"industrialFields" validate:"required" message:"产业领域不能为空"` // 产业领域
	XcujuChanye string `form:"xcujuChanye" validate:"required" message:"6+5+X不能为空"` // 6+5+X
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ScenarioCreateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ScenarioUpdateRequest 更新res_scenario请求参数
type ScenarioUpdateRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
	ScenarioNo string `form:"scenarioNo" validate:"required" message:"场景编码不能为空"` // 场景编码
	PublishTime *time.Time `form:"publishTime" validate:"required" message:"发布时间不能为空"` // 发布时间
	OwnerUnit string `form:"ownerUnit" validate:"required" message:"业主单位不能为空"` // 业主单位
	SceneName string `form:"sceneName" validate:"required" message:"场景名称不能为空"` // 场景名称
	Industry string `form:"industry" validate:"required" message:"所属产业不能为空"` // 所属产业
	OpportunityZone string `form:"opportunityZone" validate:"required" message:"机会区域不能为空"` // 机会区域
	ValidityPeriod string `form:"validityPeriod" validate:"required" message:"机会有效期不能为空"` // 机会有效期
	BudgetGuarantee string `form:"budgetGuarantee" validate:"required" message:"预算保障情况不能为空"` // 预算保障情况
	OtherSituations string `form:"otherSituations" validate:"required" message:"其他情况不能为空"` // 其他情况
	OpportunityDesc string `form:"opportunityDesc" validate:"required" message:"机会描述不能为空"` // 机会描述
	BuildBasement string `form:"buildBasement" validate:"required" message:"建设基础不能为空"` // 建设基础
	CooperationRequirements string `form:"cooperationRequirements" validate:"required" message:"合作需求不能为空"` // 合作需求
	CooperationForm string `form:"cooperationForm" validate:"required" message:"合作形式不能为空"` // 合作形式
	ContactPerson string `form:"contactPerson" validate:"required" message:"联系人不能为空"` // 联系人
	Position string `form:"position" validate:"required" message:"职务不能为空"` // 职务
	Phone string `form:"phone" validate:"required" message:"联系方式不能为空"` // 联系方式
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	IndustrialFields string `form:"industrialFields" validate:"required" message:"产业领域不能为空"` // 产业领域
	XcujuChanye string `form:"xcujuChanye" validate:"required" message:"6+5+X不能为空"` // 6+5+X
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ScenarioUpdateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ScenarioDeleteRequest 删除res_scenario请求参数
type ScenarioDeleteRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *ScenarioDeleteRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ScenarioGetByIDRequest 根据ID获取res_scenario请求参数
type ScenarioGetByIDRequest struct {
	models.Validator
	Id int64 `uri:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *ScenarioGetByIDRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}