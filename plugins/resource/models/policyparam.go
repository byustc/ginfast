package models

import (
	"time"
	"gin-fast/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PolicyListRequest res_policy列表请求参数
type PolicyListRequest struct {
	models.BasePaging
	models.Validator
	Id *int64 `form:"id"` // 主键
	PolicyNo *string `form:"policyNo"` // 政策编码
	PolicyTitle *string `form:"policyTitle"` // 政策标题
	PolicyAbstract *string `form:"policyAbstract"` // 政策摘要
	MainBody *string `form:"mainBody"` // 正文
	PublishTime *time.Time `form:"publishTime"` // 发布时间
	PolicyDepartment *string `form:"policyDepartment"` // 发布部门
	Source *string `form:"source"` // 来源
	Link *string `form:"link"` // 原文链接
	EffectDate *time.Time `form:"effectDate"` // 生效日期
	FailureDate *time.Time `form:"failureDate"` // 失效日期
	Status *string `form:"status"` // 政策状态
	PolicyTimeliness *string `form:"policyTimeliness"` // 时效性
	ContactName *string `form:"contactName"` // 联系人
	ContactPhone *string `form:"contactPhone"` // 联系方式
	PolicyRegion *string `form:"policyRegion"` // 政策地区
	PolicyTag *string `form:"policyTag"` // 目标产业
	PolicyComType *string `form:"policyComType"` // 适用企业类型
	Level *string `form:"level"` // 政策级别
	Attachment *string `form:"attachment"` // 附件
	PolicyNumber *string `form:"policyNumber"` // 文号
	Remark *string `form:"remark"` // 备注
	Title *string `form:"title"` // 标题
	DataSource *string `form:"dataSource"` // 数据来源
	IsBlacklist *int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag *string `form:"tag"` // 行业数据标签
	UpdatedBy *int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *PolicyListRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// Handle 获取查询条件
func (r *PolicyListRequest) Handle() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
        if r.Id != nil {
            // 默认等于查询
            db = db.Where("id = ?", *r.Id)
        }
        if r.PolicyNo != nil {
            // 默认等于查询
            db = db.Where("policy_no = ?", *r.PolicyNo)
        }
        if r.PolicyTitle != nil {
            // 默认等于查询
            db = db.Where("policy_title = ?", *r.PolicyTitle)
        }
        if r.PolicyAbstract != nil {
            // 默认等于查询
            db = db.Where("policy_abstract = ?", *r.PolicyAbstract)
        }
        if r.MainBody != nil {
            // 默认等于查询
            db = db.Where("main_body = ?", *r.MainBody)
        }
        if r.PublishTime != nil {
            // 默认等于查询
            db = db.Where("publish_time = ?", *r.PublishTime)
        }
        if r.PolicyDepartment != nil {
            // 默认等于查询
            db = db.Where("policy_department = ?", *r.PolicyDepartment)
        }
        if r.Source != nil {
            // 默认等于查询
            db = db.Where("source = ?", *r.Source)
        }
        if r.Link != nil {
            // 默认等于查询
            db = db.Where("link = ?", *r.Link)
        }
        if r.EffectDate != nil {
            // 默认等于查询
            db = db.Where("effect_date = ?", *r.EffectDate)
        }
        if r.FailureDate != nil {
            // 默认等于查询
            db = db.Where("failure_date = ?", *r.FailureDate)
        }
        if r.Status != nil {
            // 默认等于查询
            db = db.Where("status = ?", *r.Status)
        }
        if r.PolicyTimeliness != nil {
            // 默认等于查询
            db = db.Where("policy_timeliness = ?", *r.PolicyTimeliness)
        }
        if r.ContactName != nil {
            // 默认等于查询
            db = db.Where("contact_name = ?", *r.ContactName)
        }
        if r.ContactPhone != nil {
            // 默认等于查询
            db = db.Where("contact_phone = ?", *r.ContactPhone)
        }
        if r.PolicyRegion != nil {
            // 默认等于查询
            db = db.Where("policy_region = ?", *r.PolicyRegion)
        }
        if r.PolicyTag != nil {
            // 默认等于查询
            db = db.Where("policy_tag = ?", *r.PolicyTag)
        }
        if r.PolicyComType != nil {
            // 默认等于查询
            db = db.Where("policy_com_type = ?", *r.PolicyComType)
        }
        if r.Level != nil {
            // 默认等于查询
            db = db.Where("level = ?", *r.Level)
        }
        if r.Attachment != nil {
            // 默认等于查询
            db = db.Where("attachment = ?", *r.Attachment)
        }
        if r.PolicyNumber != nil {
            // 默认等于查询
            db = db.Where("policy_number = ?", *r.PolicyNumber)
        }
        if r.Remark != nil {
            // 默认等于查询
            db = db.Where("remark = ?", *r.Remark)
        }
        if r.Title != nil {
            // 默认等于查询
            db = db.Where("title = ?", *r.Title)
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

// PolicyCreateRequest 创建res_policy请求参数
type PolicyCreateRequest struct {
	models.Validator
	PolicyNo string `form:"policyNo" validate:"required" message:"政策编码不能为空"` // 政策编码
	PolicyTitle string `form:"policyTitle" validate:"required" message:"政策标题不能为空"` // 政策标题
	PolicyAbstract string `form:"policyAbstract" validate:"required" message:"政策摘要不能为空"` // 政策摘要
	MainBody string `form:"mainBody" validate:"required" message:"正文不能为空"` // 正文
	PublishTime *time.Time `form:"publishTime" validate:"required" message:"发布时间不能为空"` // 发布时间
	PolicyDepartment string `form:"policyDepartment" validate:"required" message:"发布部门不能为空"` // 发布部门
	Source string `form:"source" validate:"required" message:"来源不能为空"` // 来源
	Link string `form:"link" validate:"required" message:"原文链接不能为空"` // 原文链接
	EffectDate *time.Time `form:"effectDate" validate:"required" message:"生效日期不能为空"` // 生效日期
	FailureDate *time.Time `form:"failureDate" validate:"required" message:"失效日期不能为空"` // 失效日期
	Status string `form:"status" validate:"required" message:"政策状态不能为空"` // 政策状态
	PolicyTimeliness string `form:"policyTimeliness" validate:"required" message:"时效性不能为空"` // 时效性
	ContactName string `form:"contactName" validate:"required" message:"联系人不能为空"` // 联系人
	ContactPhone string `form:"contactPhone" validate:"required" message:"联系方式不能为空"` // 联系方式
	PolicyRegion string `form:"policyRegion" validate:"required" message:"政策地区不能为空"` // 政策地区
	PolicyTag string `form:"policyTag" validate:"required" message:"目标产业不能为空"` // 目标产业
	PolicyComType string `form:"policyComType" validate:"required" message:"适用企业类型不能为空"` // 适用企业类型
	Level string `form:"level" validate:"required" message:"政策级别不能为空"` // 政策级别
	Attachment string `form:"attachment" validate:"required" message:"附件不能为空"` // 附件
	PolicyNumber string `form:"policyNumber" validate:"required" message:"文号不能为空"` // 文号
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	Title string `form:"title" validate:"required" message:"标题不能为空"` // 标题
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *PolicyCreateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// PolicyUpdateRequest 更新res_policy请求参数
type PolicyUpdateRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
	PolicyNo string `form:"policyNo" validate:"required" message:"政策编码不能为空"` // 政策编码
	PolicyTitle string `form:"policyTitle" validate:"required" message:"政策标题不能为空"` // 政策标题
	PolicyAbstract string `form:"policyAbstract" validate:"required" message:"政策摘要不能为空"` // 政策摘要
	MainBody string `form:"mainBody" validate:"required" message:"正文不能为空"` // 正文
	PublishTime *time.Time `form:"publishTime" validate:"required" message:"发布时间不能为空"` // 发布时间
	PolicyDepartment string `form:"policyDepartment" validate:"required" message:"发布部门不能为空"` // 发布部门
	Source string `form:"source" validate:"required" message:"来源不能为空"` // 来源
	Link string `form:"link" validate:"required" message:"原文链接不能为空"` // 原文链接
	EffectDate *time.Time `form:"effectDate" validate:"required" message:"生效日期不能为空"` // 生效日期
	FailureDate *time.Time `form:"failureDate" validate:"required" message:"失效日期不能为空"` // 失效日期
	Status string `form:"status" validate:"required" message:"政策状态不能为空"` // 政策状态
	PolicyTimeliness string `form:"policyTimeliness" validate:"required" message:"时效性不能为空"` // 时效性
	ContactName string `form:"contactName" validate:"required" message:"联系人不能为空"` // 联系人
	ContactPhone string `form:"contactPhone" validate:"required" message:"联系方式不能为空"` // 联系方式
	PolicyRegion string `form:"policyRegion" validate:"required" message:"政策地区不能为空"` // 政策地区
	PolicyTag string `form:"policyTag" validate:"required" message:"目标产业不能为空"` // 目标产业
	PolicyComType string `form:"policyComType" validate:"required" message:"适用企业类型不能为空"` // 适用企业类型
	Level string `form:"level" validate:"required" message:"政策级别不能为空"` // 政策级别
	Attachment string `form:"attachment" validate:"required" message:"附件不能为空"` // 附件
	PolicyNumber string `form:"policyNumber" validate:"required" message:"文号不能为空"` // 文号
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	Title string `form:"title" validate:"required" message:"标题不能为空"` // 标题
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *PolicyUpdateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// PolicyDeleteRequest 删除res_policy请求参数
type PolicyDeleteRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *PolicyDeleteRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// PolicyGetByIDRequest 根据ID获取res_policy请求参数
type PolicyGetByIDRequest struct {
	models.Validator
	Id int64 `uri:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *PolicyGetByIDRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}