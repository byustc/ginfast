package models

import (
	"time"
	"gin-fast/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ThinktankListRequest res_thinktank列表请求参数
type ThinktankListRequest struct {
	models.BasePaging
	models.Validator
	Id *int64 `form:"id"` // 主键
	ThinktankNo *string `form:"thinktankNo"` // 新研机构编码
	BriefName *string `form:"briefName"` // 企业简称
	FullName *string `form:"fullName"` // 企业全称
	ComOtherName *string `form:"comOtherName"` // 其他简称
	Logo *string `form:"logo"` // 企业logo
	ComType *string `form:"comType"` // 企业类型
	Website *string `form:"website"` // 企业网址
	BriefIntro *string `form:"briefIntro"` // 一句话简介
	DetailIntro *string `form:"detailIntro"` // 企业介绍
	EstablishTime *time.Time `form:"establishTime"` // 成立时间
	InvestRound *float64 `form:"investRound"` // 当前轮次
	InvestRoundName *string `form:"investRoundName"` // 当前轮次名称
	OperatingStatus *float64 `form:"operatingStatus"` // 运营状态
	OperatingStatusName *string `form:"operatingStatusName"` // 运营状态名称
	Staff *string `form:"staff"` // 人员规模
	BaikeLink *string `form:"baikeLink"` // 百度百科链接
	Industry *string `form:"industry"` // 一级行业标签
	IndustryName *string `form:"industryName"` // 一级行业标签名称
	MarketValuationRmb *float64 `form:"marketValuationRmb"` // 公司最新市值/估值-亿元
	Tags *string `form:"tags"` // 企业全部标签
	Location *string `form:"location"` // 所在地
	Address *string `form:"address"` // 详细地址
	CooperationType *string `form:"cooperationType"` // 合作类别
	CooperationForm *string `form:"cooperationForm"` // 合作形式
	ResourceAnalysis *string `form:"resourceAnalysis"` // 资源方穿透分析
	ConstructionContent *string `form:"constructionContent"` // 建设内容
	IndustryFocus *string `form:"industryFocus"` // 聚焦产业领域
	CooperationUnits *string `form:"cooperationUnits"` // 合作单位
	BusinessFocus *string `form:"businessFocus"` // 业务侧重
	OperationMode *string `form:"operationMode"` // 运营模式
	ResponsibleDepartment *string `form:"responsibleDepartment"` // 主责部门
	ResearchTopic *string `form:"researchTopic"` // 在研课题及相关进度
	Remark *string `form:"remark"` // 备注
	RegFullName *string `form:"regFullName"` // 注册名称
	RegTime *time.Time `form:"regTime"` // 注册时间
	RegLocation *string `form:"regLocation"` // 注册地址
	RegCapital *float64 `form:"regCapital"` // 注册资本（万）
	RegCapitalCurrency *float64 `form:"regCapitalCurrency"` // 注册资本币种
	RegCapitalCurrencyName *string `form:"regCapitalCurrencyName"` // 注册资本币种名称
	PaidCapital *float64 `form:"paidCapital"` // 实缴资本（万）
	PaidCapitalCurrency *float64 `form:"paidCapitalCurrency"` // 实缴资本币种
	PaidCapitalCurrencyName *string `form:"paidCapitalCurrencyName"` // 实缴资本币种名称
	RegType *float64 `form:"regType"` // 企业注册类型
	RegTypeName *string `form:"regTypeName"` // 企业注册类型名称
	LegalRepresent *string `form:"legalRepresent"` // 法人代表
	RegNumber *string `form:"regNumber"` // 工商注册号
	CreditNumber *string `form:"creditNumber"` // 统一社会信用代码
	BusinessPeriod *string `form:"businessPeriod"` // 营业期限
	BusinessScope *string `form:"businessScope"` // 经营范围
	RegIndustry *float64 `form:"regIndustry"` // 行业名称
	RegIndustryName *string `form:"regIndustryName"` // 行业门类名称
	RegSubIndustry *float64 `form:"regSubIndustry"` // 行业大类
	RegSubIndustryName *string `form:"regSubIndustryName"` // 行业大类名称
	RegMiddleIndustry *float64 `form:"regMiddleIndustry"` // 行业中类
	RegMiddleIndustryName *string `form:"regMiddleIndustryName"` // 行业中类名称
	RegSmallIndustry *float64 `form:"regSmallIndustry"` // 行业小类
	RegSmallIndustryName *string `form:"regSmallIndustryName"` // 行业小类名称
	RegInstitute *string `form:"regInstitute"` // 登记机关
	ApproveTime *time.Time `form:"approveTime"` // 核准日期
	DataSource *string `form:"dataSource"` // 数据来源
	IsBlacklist *int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag *string `form:"tag"` // 行业数据标签
	Partner *int `form:"partner"` // 是否合作伙伴: 0否 1是
	PlatformResources *string `form:"platformResources"` // 平台资源
	Contact *string `form:"contact"` // 联系人
	ContactNumber *string `form:"contactNumber"` // 联系电话
	Duties *string `form:"duties"` // 职务
	LocationChina *string `form:"locationChina"` // 所在地
	UpdatedBy *int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ThinktankListRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// Handle 获取查询条件
func (r *ThinktankListRequest) Handle() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
        if r.Id != nil {
            // 默认等于查询
            db = db.Where("id = ?", *r.Id)
        }
        if r.ThinktankNo != nil {
            // 默认等于查询
            db = db.Where("thinktank_no = ?", *r.ThinktankNo)
        }
        if r.BriefName != nil {
            // 默认等于查询
            db = db.Where("brief_name = ?", *r.BriefName)
        }
        if r.FullName != nil {
            // 默认等于查询
            db = db.Where("full_name = ?", *r.FullName)
        }
        if r.ComOtherName != nil {
            // 默认等于查询
            db = db.Where("com_other_name = ?", *r.ComOtherName)
        }
        if r.Logo != nil {
            // 默认等于查询
            db = db.Where("logo = ?", *r.Logo)
        }
        if r.ComType != nil {
            // 默认等于查询
            db = db.Where("com_type = ?", *r.ComType)
        }
        if r.Website != nil {
            // 默认等于查询
            db = db.Where("website = ?", *r.Website)
        }
        if r.BriefIntro != nil {
            // 默认等于查询
            db = db.Where("brief_intro = ?", *r.BriefIntro)
        }
        if r.DetailIntro != nil {
            // 默认等于查询
            db = db.Where("detail_intro = ?", *r.DetailIntro)
        }
        if r.EstablishTime != nil {
            // 默认等于查询
            db = db.Where("establish_time = ?", *r.EstablishTime)
        }
        if r.InvestRound != nil {
            // 默认等于查询
            db = db.Where("invest_round = ?", *r.InvestRound)
        }
        if r.InvestRoundName != nil {
            // 默认等于查询
            db = db.Where("invest_round_name = ?", *r.InvestRoundName)
        }
        if r.OperatingStatus != nil {
            // 默认等于查询
            db = db.Where("operating_status = ?", *r.OperatingStatus)
        }
        if r.OperatingStatusName != nil {
            // 默认等于查询
            db = db.Where("operating_status_name = ?", *r.OperatingStatusName)
        }
        if r.Staff != nil {
            // 默认等于查询
            db = db.Where("staff = ?", *r.Staff)
        }
        if r.BaikeLink != nil {
            // 默认等于查询
            db = db.Where("baike_link = ?", *r.BaikeLink)
        }
        if r.Industry != nil {
            // 默认等于查询
            db = db.Where("industry = ?", *r.Industry)
        }
        if r.IndustryName != nil {
            // 默认等于查询
            db = db.Where("industry_name = ?", *r.IndustryName)
        }
        if r.MarketValuationRmb != nil {
            // 默认等于查询
            db = db.Where("market_valuation_rmb = ?", *r.MarketValuationRmb)
        }
        if r.Tags != nil {
            // 默认等于查询
            db = db.Where("tags = ?", *r.Tags)
        }
        if r.Location != nil {
            // 默认等于查询
            db = db.Where("location = ?", *r.Location)
        }
        if r.Address != nil {
            // 默认等于查询
            db = db.Where("address = ?", *r.Address)
        }
        if r.CooperationType != nil {
            // 默认等于查询
            db = db.Where("cooperation_type = ?", *r.CooperationType)
        }
        if r.CooperationForm != nil {
            // 默认等于查询
            db = db.Where("cooperation_form = ?", *r.CooperationForm)
        }
        if r.ResourceAnalysis != nil {
            // 默认等于查询
            db = db.Where("resource_analysis = ?", *r.ResourceAnalysis)
        }
        if r.ConstructionContent != nil {
            // 默认等于查询
            db = db.Where("construction_content = ?", *r.ConstructionContent)
        }
        if r.IndustryFocus != nil {
            // 默认等于查询
            db = db.Where("industry_focus = ?", *r.IndustryFocus)
        }
        if r.CooperationUnits != nil {
            // 默认等于查询
            db = db.Where("cooperation_units = ?", *r.CooperationUnits)
        }
        if r.BusinessFocus != nil {
            // 默认等于查询
            db = db.Where("business_focus = ?", *r.BusinessFocus)
        }
        if r.OperationMode != nil {
            // 默认等于查询
            db = db.Where("operation_mode = ?", *r.OperationMode)
        }
        if r.ResponsibleDepartment != nil {
            // 默认等于查询
            db = db.Where("responsible_department = ?", *r.ResponsibleDepartment)
        }
        if r.ResearchTopic != nil {
            // 默认等于查询
            db = db.Where("research_topic = ?", *r.ResearchTopic)
        }
        if r.Remark != nil {
            // 默认等于查询
            db = db.Where("remark = ?", *r.Remark)
        }
        if r.RegFullName != nil {
            // 默认等于查询
            db = db.Where("reg_full_name = ?", *r.RegFullName)
        }
        if r.RegTime != nil {
            // 默认等于查询
            db = db.Where("reg_time = ?", *r.RegTime)
        }
        if r.RegLocation != nil {
            // 默认等于查询
            db = db.Where("reg_location = ?", *r.RegLocation)
        }
        if r.RegCapital != nil {
            // 默认等于查询
            db = db.Where("reg_capital = ?", *r.RegCapital)
        }
        if r.RegCapitalCurrency != nil {
            // 默认等于查询
            db = db.Where("reg_capital_currency = ?", *r.RegCapitalCurrency)
        }
        if r.RegCapitalCurrencyName != nil {
            // 默认等于查询
            db = db.Where("reg_capital_currency_name = ?", *r.RegCapitalCurrencyName)
        }
        if r.PaidCapital != nil {
            // 默认等于查询
            db = db.Where("paid_capital = ?", *r.PaidCapital)
        }
        if r.PaidCapitalCurrency != nil {
            // 默认等于查询
            db = db.Where("paid_capital_currency = ?", *r.PaidCapitalCurrency)
        }
        if r.PaidCapitalCurrencyName != nil {
            // 默认等于查询
            db = db.Where("paid_capital_currency_name = ?", *r.PaidCapitalCurrencyName)
        }
        if r.RegType != nil {
            // 默认等于查询
            db = db.Where("reg_type = ?", *r.RegType)
        }
        if r.RegTypeName != nil {
            // 默认等于查询
            db = db.Where("reg_type_name = ?", *r.RegTypeName)
        }
        if r.LegalRepresent != nil {
            // 默认等于查询
            db = db.Where("legal_represent = ?", *r.LegalRepresent)
        }
        if r.RegNumber != nil {
            // 默认等于查询
            db = db.Where("reg_number = ?", *r.RegNumber)
        }
        if r.CreditNumber != nil {
            // 默认等于查询
            db = db.Where("credit_number = ?", *r.CreditNumber)
        }
        if r.BusinessPeriod != nil {
            // 默认等于查询
            db = db.Where("business_period = ?", *r.BusinessPeriod)
        }
        if r.BusinessScope != nil {
            // 默认等于查询
            db = db.Where("business_scope = ?", *r.BusinessScope)
        }
        if r.RegIndustry != nil {
            // 默认等于查询
            db = db.Where("reg_industry = ?", *r.RegIndustry)
        }
        if r.RegIndustryName != nil {
            // 默认等于查询
            db = db.Where("reg_industry_name = ?", *r.RegIndustryName)
        }
        if r.RegSubIndustry != nil {
            // 默认等于查询
            db = db.Where("reg_sub_industry = ?", *r.RegSubIndustry)
        }
        if r.RegSubIndustryName != nil {
            // 默认等于查询
            db = db.Where("reg_sub_industry_name = ?", *r.RegSubIndustryName)
        }
        if r.RegMiddleIndustry != nil {
            // 默认等于查询
            db = db.Where("reg_middle_industry = ?", *r.RegMiddleIndustry)
        }
        if r.RegMiddleIndustryName != nil {
            // 默认等于查询
            db = db.Where("reg_middle_industry_name = ?", *r.RegMiddleIndustryName)
        }
        if r.RegSmallIndustry != nil {
            // 默认等于查询
            db = db.Where("reg_small_industry = ?", *r.RegSmallIndustry)
        }
        if r.RegSmallIndustryName != nil {
            // 默认等于查询
            db = db.Where("reg_small_industry_name = ?", *r.RegSmallIndustryName)
        }
        if r.RegInstitute != nil {
            // 默认等于查询
            db = db.Where("reg_institute = ?", *r.RegInstitute)
        }
        if r.ApproveTime != nil {
            // 默认等于查询
            db = db.Where("approve_time = ?", *r.ApproveTime)
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
        if r.Partner != nil {
            // 默认等于查询
            db = db.Where("partner = ?", *r.Partner)
        }
        if r.PlatformResources != nil {
            // 默认等于查询
            db = db.Where("platform_resources = ?", *r.PlatformResources)
        }
        if r.Contact != nil {
            // 默认等于查询
            db = db.Where("contact = ?", *r.Contact)
        }
        if r.ContactNumber != nil {
            // 默认等于查询
            db = db.Where("contact_number = ?", *r.ContactNumber)
        }
        if r.Duties != nil {
            // 默认等于查询
            db = db.Where("duties = ?", *r.Duties)
        }
        if r.LocationChina != nil {
            // 默认等于查询
            db = db.Where("location_china = ?", *r.LocationChina)
        }
        if r.UpdatedBy != nil {
            // 默认等于查询
            db = db.Where("updated_by = ?", *r.UpdatedBy)
        }
		return db
	}
}

// ThinktankCreateRequest 创建res_thinktank请求参数
type ThinktankCreateRequest struct {
	models.Validator
	ThinktankNo string `form:"thinktankNo" validate:"required" message:"新研机构编码不能为空"` // 新研机构编码
	BriefName string `form:"briefName" validate:"required" message:"企业简称不能为空"` // 企业简称
	FullName string `form:"fullName" validate:"required" message:"企业全称不能为空"` // 企业全称
	ComOtherName string `form:"comOtherName" validate:"required" message:"其他简称不能为空"` // 其他简称
	Logo string `form:"logo" validate:"required" message:"企业logo不能为空"` // 企业logo
	ComType string `form:"comType" validate:"required" message:"企业类型不能为空"` // 企业类型
	Website string `form:"website" validate:"required" message:"企业网址不能为空"` // 企业网址
	BriefIntro string `form:"briefIntro" validate:"required" message:"一句话简介不能为空"` // 一句话简介
	DetailIntro string `form:"detailIntro" validate:"required" message:"企业介绍不能为空"` // 企业介绍
	EstablishTime *time.Time `form:"establishTime" validate:"required" message:"成立时间不能为空"` // 成立时间
	InvestRound float64 `form:"investRound"` // 当前轮次
	InvestRoundName string `form:"investRoundName" validate:"required" message:"当前轮次名称不能为空"` // 当前轮次名称
	OperatingStatus float64 `form:"operatingStatus"` // 运营状态
	OperatingStatusName string `form:"operatingStatusName" validate:"required" message:"运营状态名称不能为空"` // 运营状态名称
	Staff string `form:"staff" validate:"required" message:"人员规模不能为空"` // 人员规模
	BaikeLink string `form:"baikeLink" validate:"required" message:"百度百科链接不能为空"` // 百度百科链接
	Industry string `form:"industry" validate:"required" message:"一级行业标签不能为空"` // 一级行业标签
	IndustryName string `form:"industryName" validate:"required" message:"一级行业标签名称不能为空"` // 一级行业标签名称
	MarketValuationRmb float64 `form:"marketValuationRmb"` // 公司最新市值/估值-亿元
	Tags string `form:"tags" validate:"required" message:"企业全部标签不能为空"` // 企业全部标签
	Location string `form:"location" validate:"required" message:"所在地不能为空"` // 所在地
	Address string `form:"address" validate:"required" message:"详细地址不能为空"` // 详细地址
	CooperationType string `form:"cooperationType" validate:"required" message:"合作类别不能为空"` // 合作类别
	CooperationForm string `form:"cooperationForm" validate:"required" message:"合作形式不能为空"` // 合作形式
	ResourceAnalysis string `form:"resourceAnalysis" validate:"required" message:"资源方穿透分析不能为空"` // 资源方穿透分析
	ConstructionContent string `form:"constructionContent" validate:"required" message:"建设内容不能为空"` // 建设内容
	IndustryFocus string `form:"industryFocus" validate:"required" message:"聚焦产业领域不能为空"` // 聚焦产业领域
	CooperationUnits string `form:"cooperationUnits" validate:"required" message:"合作单位不能为空"` // 合作单位
	BusinessFocus string `form:"businessFocus" validate:"required" message:"业务侧重不能为空"` // 业务侧重
	OperationMode string `form:"operationMode" validate:"required" message:"运营模式不能为空"` // 运营模式
	ResponsibleDepartment string `form:"responsibleDepartment" validate:"required" message:"主责部门不能为空"` // 主责部门
	ResearchTopic string `form:"researchTopic" validate:"required" message:"在研课题及相关进度不能为空"` // 在研课题及相关进度
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	RegFullName string `form:"regFullName" validate:"required" message:"注册名称不能为空"` // 注册名称
	RegTime *time.Time `form:"regTime" validate:"required" message:"注册时间不能为空"` // 注册时间
	RegLocation string `form:"regLocation" validate:"required" message:"注册地址不能为空"` // 注册地址
	RegCapital float64 `form:"regCapital"` // 注册资本（万）
	RegCapitalCurrency float64 `form:"regCapitalCurrency"` // 注册资本币种
	RegCapitalCurrencyName string `form:"regCapitalCurrencyName" validate:"required" message:"注册资本币种名称不能为空"` // 注册资本币种名称
	PaidCapital float64 `form:"paidCapital"` // 实缴资本（万）
	PaidCapitalCurrency float64 `form:"paidCapitalCurrency"` // 实缴资本币种
	PaidCapitalCurrencyName string `form:"paidCapitalCurrencyName" validate:"required" message:"实缴资本币种名称不能为空"` // 实缴资本币种名称
	RegType float64 `form:"regType"` // 企业注册类型
	RegTypeName string `form:"regTypeName" validate:"required" message:"企业注册类型名称不能为空"` // 企业注册类型名称
	LegalRepresent string `form:"legalRepresent" validate:"required" message:"法人代表不能为空"` // 法人代表
	RegNumber string `form:"regNumber" validate:"required" message:"工商注册号不能为空"` // 工商注册号
	CreditNumber string `form:"creditNumber" validate:"required" message:"统一社会信用代码不能为空"` // 统一社会信用代码
	BusinessPeriod string `form:"businessPeriod" validate:"required" message:"营业期限不能为空"` // 营业期限
	BusinessScope string `form:"businessScope" validate:"required" message:"经营范围不能为空"` // 经营范围
	RegIndustry float64 `form:"regIndustry"` // 行业名称
	RegIndustryName string `form:"regIndustryName" validate:"required" message:"行业门类名称不能为空"` // 行业门类名称
	RegSubIndustry float64 `form:"regSubIndustry"` // 行业大类
	RegSubIndustryName string `form:"regSubIndustryName" validate:"required" message:"行业大类名称不能为空"` // 行业大类名称
	RegMiddleIndustry float64 `form:"regMiddleIndustry"` // 行业中类
	RegMiddleIndustryName string `form:"regMiddleIndustryName" validate:"required" message:"行业中类名称不能为空"` // 行业中类名称
	RegSmallIndustry float64 `form:"regSmallIndustry"` // 行业小类
	RegSmallIndustryName string `form:"regSmallIndustryName" validate:"required" message:"行业小类名称不能为空"` // 行业小类名称
	RegInstitute string `form:"regInstitute" validate:"required" message:"登记机关不能为空"` // 登记机关
	ApproveTime *time.Time `form:"approveTime" validate:"required" message:"核准日期不能为空"` // 核准日期
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	Partner int `form:"partner"` // 是否合作伙伴: 0否 1是
	PlatformResources string `form:"platformResources" validate:"required" message:"平台资源不能为空"` // 平台资源
	Contact string `form:"contact" validate:"required" message:"联系人不能为空"` // 联系人
	ContactNumber string `form:"contactNumber" validate:"required" message:"联系电话不能为空"` // 联系电话
	Duties string `form:"duties" validate:"required" message:"职务不能为空"` // 职务
	LocationChina string `form:"locationChina" validate:"required" message:"所在地不能为空"` // 所在地
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ThinktankCreateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ThinktankUpdateRequest 更新res_thinktank请求参数
type ThinktankUpdateRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
	ThinktankNo string `form:"thinktankNo" validate:"required" message:"新研机构编码不能为空"` // 新研机构编码
	BriefName string `form:"briefName" validate:"required" message:"企业简称不能为空"` // 企业简称
	FullName string `form:"fullName" validate:"required" message:"企业全称不能为空"` // 企业全称
	ComOtherName string `form:"comOtherName" validate:"required" message:"其他简称不能为空"` // 其他简称
	Logo string `form:"logo" validate:"required" message:"企业logo不能为空"` // 企业logo
	ComType string `form:"comType" validate:"required" message:"企业类型不能为空"` // 企业类型
	Website string `form:"website" validate:"required" message:"企业网址不能为空"` // 企业网址
	BriefIntro string `form:"briefIntro" validate:"required" message:"一句话简介不能为空"` // 一句话简介
	DetailIntro string `form:"detailIntro" validate:"required" message:"企业介绍不能为空"` // 企业介绍
	EstablishTime *time.Time `form:"establishTime" validate:"required" message:"成立时间不能为空"` // 成立时间
	InvestRound float64 `form:"investRound"` // 当前轮次
	InvestRoundName string `form:"investRoundName" validate:"required" message:"当前轮次名称不能为空"` // 当前轮次名称
	OperatingStatus float64 `form:"operatingStatus"` // 运营状态
	OperatingStatusName string `form:"operatingStatusName" validate:"required" message:"运营状态名称不能为空"` // 运营状态名称
	Staff string `form:"staff" validate:"required" message:"人员规模不能为空"` // 人员规模
	BaikeLink string `form:"baikeLink" validate:"required" message:"百度百科链接不能为空"` // 百度百科链接
	Industry string `form:"industry" validate:"required" message:"一级行业标签不能为空"` // 一级行业标签
	IndustryName string `form:"industryName" validate:"required" message:"一级行业标签名称不能为空"` // 一级行业标签名称
	MarketValuationRmb float64 `form:"marketValuationRmb"` // 公司最新市值/估值-亿元
	Tags string `form:"tags" validate:"required" message:"企业全部标签不能为空"` // 企业全部标签
	Location string `form:"location" validate:"required" message:"所在地不能为空"` // 所在地
	Address string `form:"address" validate:"required" message:"详细地址不能为空"` // 详细地址
	CooperationType string `form:"cooperationType" validate:"required" message:"合作类别不能为空"` // 合作类别
	CooperationForm string `form:"cooperationForm" validate:"required" message:"合作形式不能为空"` // 合作形式
	ResourceAnalysis string `form:"resourceAnalysis" validate:"required" message:"资源方穿透分析不能为空"` // 资源方穿透分析
	ConstructionContent string `form:"constructionContent" validate:"required" message:"建设内容不能为空"` // 建设内容
	IndustryFocus string `form:"industryFocus" validate:"required" message:"聚焦产业领域不能为空"` // 聚焦产业领域
	CooperationUnits string `form:"cooperationUnits" validate:"required" message:"合作单位不能为空"` // 合作单位
	BusinessFocus string `form:"businessFocus" validate:"required" message:"业务侧重不能为空"` // 业务侧重
	OperationMode string `form:"operationMode" validate:"required" message:"运营模式不能为空"` // 运营模式
	ResponsibleDepartment string `form:"responsibleDepartment" validate:"required" message:"主责部门不能为空"` // 主责部门
	ResearchTopic string `form:"researchTopic" validate:"required" message:"在研课题及相关进度不能为空"` // 在研课题及相关进度
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	RegFullName string `form:"regFullName" validate:"required" message:"注册名称不能为空"` // 注册名称
	RegTime *time.Time `form:"regTime" validate:"required" message:"注册时间不能为空"` // 注册时间
	RegLocation string `form:"regLocation" validate:"required" message:"注册地址不能为空"` // 注册地址
	RegCapital float64 `form:"regCapital"` // 注册资本（万）
	RegCapitalCurrency float64 `form:"regCapitalCurrency"` // 注册资本币种
	RegCapitalCurrencyName string `form:"regCapitalCurrencyName" validate:"required" message:"注册资本币种名称不能为空"` // 注册资本币种名称
	PaidCapital float64 `form:"paidCapital"` // 实缴资本（万）
	PaidCapitalCurrency float64 `form:"paidCapitalCurrency"` // 实缴资本币种
	PaidCapitalCurrencyName string `form:"paidCapitalCurrencyName" validate:"required" message:"实缴资本币种名称不能为空"` // 实缴资本币种名称
	RegType float64 `form:"regType"` // 企业注册类型
	RegTypeName string `form:"regTypeName" validate:"required" message:"企业注册类型名称不能为空"` // 企业注册类型名称
	LegalRepresent string `form:"legalRepresent" validate:"required" message:"法人代表不能为空"` // 法人代表
	RegNumber string `form:"regNumber" validate:"required" message:"工商注册号不能为空"` // 工商注册号
	CreditNumber string `form:"creditNumber" validate:"required" message:"统一社会信用代码不能为空"` // 统一社会信用代码
	BusinessPeriod string `form:"businessPeriod" validate:"required" message:"营业期限不能为空"` // 营业期限
	BusinessScope string `form:"businessScope" validate:"required" message:"经营范围不能为空"` // 经营范围
	RegIndustry float64 `form:"regIndustry"` // 行业名称
	RegIndustryName string `form:"regIndustryName" validate:"required" message:"行业门类名称不能为空"` // 行业门类名称
	RegSubIndustry float64 `form:"regSubIndustry"` // 行业大类
	RegSubIndustryName string `form:"regSubIndustryName" validate:"required" message:"行业大类名称不能为空"` // 行业大类名称
	RegMiddleIndustry float64 `form:"regMiddleIndustry"` // 行业中类
	RegMiddleIndustryName string `form:"regMiddleIndustryName" validate:"required" message:"行业中类名称不能为空"` // 行业中类名称
	RegSmallIndustry float64 `form:"regSmallIndustry"` // 行业小类
	RegSmallIndustryName string `form:"regSmallIndustryName" validate:"required" message:"行业小类名称不能为空"` // 行业小类名称
	RegInstitute string `form:"regInstitute" validate:"required" message:"登记机关不能为空"` // 登记机关
	ApproveTime *time.Time `form:"approveTime" validate:"required" message:"核准日期不能为空"` // 核准日期
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	Partner int `form:"partner"` // 是否合作伙伴: 0否 1是
	PlatformResources string `form:"platformResources" validate:"required" message:"平台资源不能为空"` // 平台资源
	Contact string `form:"contact" validate:"required" message:"联系人不能为空"` // 联系人
	ContactNumber string `form:"contactNumber" validate:"required" message:"联系电话不能为空"` // 联系电话
	Duties string `form:"duties" validate:"required" message:"职务不能为空"` // 职务
	LocationChina string `form:"locationChina" validate:"required" message:"所在地不能为空"` // 所在地
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *ThinktankUpdateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ThinktankDeleteRequest 删除res_thinktank请求参数
type ThinktankDeleteRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *ThinktankDeleteRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ThinktankGetByIDRequest 根据ID获取res_thinktank请求参数
type ThinktankGetByIDRequest struct {
	models.Validator
	Id int64 `uri:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *ThinktankGetByIDRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}