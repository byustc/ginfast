package models

import (
	"gin-fast/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ExpertListRequest res_expert列表请求参数
type ExpertListRequest struct {
	models.BasePaging
	models.Validator
	Id *int64 `form:"id"` // 主键
	ExpertNo *string `form:"expertNo"` // 专家编码
	Name *string `form:"name"` // 专家姓名
	Phone *string `form:"phone"` // 联系方式
	Location *string `form:"location"` // 所在地
	Degree *string `form:"degree"` // 学历学位
	Major *string `form:"major"` // 专业
	MajorLevel *string `form:"majorLevel"` // 专业职称等级
	MajorType *string `form:"majorType"` // 专业类别
	ResearchIndustry *string `form:"researchIndustry"` // 研究方向
	WorkYears *int `form:"workYears"` // 工作年限
	WorkCom *string `form:"workCom"` // 工作单位
	Position *string `form:"position"` // 职位
	Industry *string `form:"industry"` // 所属产业链
	IsLecturer *string `form:"isLecturer"` // 是否能纳入培训讲师
	ExpertType *string `form:"expertType"` // 专家类型
	ExpertLevel *string `form:"expertLevel"` // 专家级别
	Source *string `form:"source"` // 专家来源
	Remark *string `form:"remark"` // 备注
	DataSource *string `form:"dataSource"` // 数据来源
	IsBlacklist *int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag *string `form:"tag"` // 行业数据标签
	IndustrySj *string `form:"industrySj"` // 所属产业
	IsPartner *int `form:"isPartner"` // 是否合作伙伴: 0否 1是
	Profile *string `form:"profile"` // 简介
	InstitutionType *string `form:"institutionType"` // 机构类型
	ContactNumber *string `form:"contactNumber"` // 联系电话
	AffiliatedInstitutions *string `form:"affiliatedInstitutions"` // 所属机构
	EntryPerson *int64 `form:"entryPerson"` // 录入人
	UpdatedBy *int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ExpertListRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// Handle 获取查询条件
func (r *ExpertListRequest) Handle() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
        if r.Id != nil {
            // 默认等于查询
            db = db.Where("id = ?", *r.Id)
        }
        if r.ExpertNo != nil {
            // 默认等于查询
            db = db.Where("expert_no = ?", *r.ExpertNo)
        }
        if r.Name != nil {
            // 默认等于查询
            db = db.Where("name = ?", *r.Name)
        }
        if r.Phone != nil {
            // 默认等于查询
            db = db.Where("phone = ?", *r.Phone)
        }
        if r.Location != nil {
            // 默认等于查询
            db = db.Where("location = ?", *r.Location)
        }
        if r.Degree != nil {
            // 默认等于查询
            db = db.Where("degree = ?", *r.Degree)
        }
        if r.Major != nil {
            // 默认等于查询
            db = db.Where("major = ?", *r.Major)
        }
        if r.MajorLevel != nil {
            // 默认等于查询
            db = db.Where("major_level = ?", *r.MajorLevel)
        }
        if r.MajorType != nil {
            // 默认等于查询
            db = db.Where("major_type = ?", *r.MajorType)
        }
        if r.ResearchIndustry != nil {
            // 默认等于查询
            db = db.Where("research_industry = ?", *r.ResearchIndustry)
        }
        if r.WorkYears != nil {
            // 默认等于查询
            db = db.Where("work_years = ?", *r.WorkYears)
        }
        if r.WorkCom != nil {
            // 默认等于查询
            db = db.Where("work_com = ?", *r.WorkCom)
        }
        if r.Position != nil {
            // 默认等于查询
            db = db.Where("position = ?", *r.Position)
        }
        if r.Industry != nil {
            // 默认等于查询
            db = db.Where("industry = ?", *r.Industry)
        }
        if r.IsLecturer != nil {
            // 默认等于查询
            db = db.Where("is_lecturer = ?", *r.IsLecturer)
        }
        if r.ExpertType != nil {
            // 默认等于查询
            db = db.Where("expert_type = ?", *r.ExpertType)
        }
        if r.ExpertLevel != nil {
            // 默认等于查询
            db = db.Where("expert_level = ?", *r.ExpertLevel)
        }
        if r.Source != nil {
            // 默认等于查询
            db = db.Where("source = ?", *r.Source)
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
        if r.IndustrySj != nil {
            // 默认等于查询
            db = db.Where("industry_sj = ?", *r.IndustrySj)
        }
        if r.IsPartner != nil {
            // 默认等于查询
            db = db.Where("is_partner = ?", *r.IsPartner)
        }
        if r.Profile != nil {
            // 默认等于查询
            db = db.Where("profile = ?", *r.Profile)
        }
        if r.InstitutionType != nil {
            // 默认等于查询
            db = db.Where("institution_type = ?", *r.InstitutionType)
        }
        if r.ContactNumber != nil {
            // 默认等于查询
            db = db.Where("contact_number = ?", *r.ContactNumber)
        }
        if r.AffiliatedInstitutions != nil {
            // 默认等于查询
            db = db.Where("affiliated_institutions = ?", *r.AffiliatedInstitutions)
        }
        if r.EntryPerson != nil {
            // 默认等于查询
            db = db.Where("entry_person = ?", *r.EntryPerson)
        }
        if r.UpdatedBy != nil {
            // 默认等于查询
            db = db.Where("updated_by = ?", *r.UpdatedBy)
        }
		return db
	}
}

// ExpertCreateRequest 创建res_expert请求参数
type ExpertCreateRequest struct {
	models.Validator
	ExpertNo string `form:"expertNo" validate:"required" message:"专家编码不能为空"` // 专家编码
	Name string `form:"name" validate:"required" message:"专家姓名不能为空"` // 专家姓名
	Phone string `form:"phone" validate:"required" message:"联系方式不能为空"` // 联系方式
	Location string `form:"location" validate:"required" message:"所在地不能为空"` // 所在地
	Degree string `form:"degree" validate:"required" message:"学历学位不能为空"` // 学历学位
	Major string `form:"major" validate:"required" message:"专业不能为空"` // 专业
	MajorLevel string `form:"majorLevel" validate:"required" message:"专业职称等级不能为空"` // 专业职称等级
	MajorType string `form:"majorType" validate:"required" message:"专业类别不能为空"` // 专业类别
	ResearchIndustry string `form:"researchIndustry" validate:"required" message:"研究方向不能为空"` // 研究方向
	WorkYears int `form:"workYears"` // 工作年限
	WorkCom string `form:"workCom" validate:"required" message:"工作单位不能为空"` // 工作单位
	Position string `form:"position" validate:"required" message:"职位不能为空"` // 职位
	Industry string `form:"industry" validate:"required" message:"所属产业链不能为空"` // 所属产业链
	IsLecturer string `form:"isLecturer" validate:"required" message:"是否能纳入培训讲师不能为空"` // 是否能纳入培训讲师
	ExpertType string `form:"expertType" validate:"required" message:"专家类型不能为空"` // 专家类型
	ExpertLevel string `form:"expertLevel" validate:"required" message:"专家级别不能为空"` // 专家级别
	Source string `form:"source" validate:"required" message:"专家来源不能为空"` // 专家来源
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	IndustrySj string `form:"industrySj" validate:"required" message:"所属产业不能为空"` // 所属产业
	IsPartner int `form:"isPartner"` // 是否合作伙伴: 0否 1是
	Profile string `form:"profile" validate:"required" message:"简介不能为空"` // 简介
	InstitutionType string `form:"institutionType" validate:"required" message:"机构类型不能为空"` // 机构类型
	ContactNumber string `form:"contactNumber" validate:"required" message:"联系电话不能为空"` // 联系电话
	AffiliatedInstitutions string `form:"affiliatedInstitutions" validate:"required" message:"所属机构不能为空"` // 所属机构
	EntryPerson int64 `form:"entryPerson"` // 录入人
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ExpertCreateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ExpertUpdateRequest 更新res_expert请求参数
type ExpertUpdateRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
	ExpertNo string `form:"expertNo" validate:"required" message:"专家编码不能为空"` // 专家编码
	Name string `form:"name" validate:"required" message:"专家姓名不能为空"` // 专家姓名
	Phone string `form:"phone" validate:"required" message:"联系方式不能为空"` // 联系方式
	Location string `form:"location" validate:"required" message:"所在地不能为空"` // 所在地
	Degree string `form:"degree" validate:"required" message:"学历学位不能为空"` // 学历学位
	Major string `form:"major" validate:"required" message:"专业不能为空"` // 专业
	MajorLevel string `form:"majorLevel" validate:"required" message:"专业职称等级不能为空"` // 专业职称等级
	MajorType string `form:"majorType" validate:"required" message:"专业类别不能为空"` // 专业类别
	ResearchIndustry string `form:"researchIndustry" validate:"required" message:"研究方向不能为空"` // 研究方向
	WorkYears int `form:"workYears"` // 工作年限
	WorkCom string `form:"workCom" validate:"required" message:"工作单位不能为空"` // 工作单位
	Position string `form:"position" validate:"required" message:"职位不能为空"` // 职位
	Industry string `form:"industry" validate:"required" message:"所属产业链不能为空"` // 所属产业链
	IsLecturer string `form:"isLecturer" validate:"required" message:"是否能纳入培训讲师不能为空"` // 是否能纳入培训讲师
	ExpertType string `form:"expertType" validate:"required" message:"专家类型不能为空"` // 专家类型
	ExpertLevel string `form:"expertLevel" validate:"required" message:"专家级别不能为空"` // 专家级别
	Source string `form:"source" validate:"required" message:"专家来源不能为空"` // 专家来源
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	IndustrySj string `form:"industrySj" validate:"required" message:"所属产业不能为空"` // 所属产业
	IsPartner int `form:"isPartner"` // 是否合作伙伴: 0否 1是
	Profile string `form:"profile" validate:"required" message:"简介不能为空"` // 简介
	InstitutionType string `form:"institutionType" validate:"required" message:"机构类型不能为空"` // 机构类型
	ContactNumber string `form:"contactNumber" validate:"required" message:"联系电话不能为空"` // 联系电话
	AffiliatedInstitutions string `form:"affiliatedInstitutions" validate:"required" message:"所属机构不能为空"` // 所属机构
	EntryPerson int64 `form:"entryPerson"` // 录入人
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ExpertUpdateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ExpertDeleteRequest 删除res_expert请求参数
type ExpertDeleteRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *ExpertDeleteRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ExpertGetByIDRequest 根据ID获取res_expert请求参数
type ExpertGetByIDRequest struct {
	models.Validator
	Id int64 `uri:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *ExpertGetByIDRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}