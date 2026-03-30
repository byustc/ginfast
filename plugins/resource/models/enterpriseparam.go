package models

import (
	"time"
	"gin-fast/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// EnterpriseListRequest res_enterprise列表请求参数
type EnterpriseListRequest struct {
	models.BasePaging
	models.Validator
	Id *int64 `form:"id"` // 主键
	EnterpriseNo *string `form:"enterpriseNo"` // 企业编码
	BriefName *string `form:"briefName"` // 企业简称
	FullName *string `form:"fullName"` // 企业全称
	ComOtherName *string `form:"comOtherName"` // 其它简称
	Logo *string `form:"logo"` // 企业Logo
	ComType *string `form:"comType"` // 企业类型
	Website *string `form:"website"` // 企业网址
	BriefIntro *string `form:"briefIntro"` // 一句话简介
	DetailIntro *string `form:"detailIntro"` // 企业介绍
	EstablishTime *time.Time `form:"establishTime"` // 成立时间
	InvestRound *float64 `form:"investRound"` // 当前轮次
	OperatingStatus *float64 `form:"operatingStatus"` // 运营状态
	Staff *string `form:"staff"` // 人员规模
	BaikeLink *string `form:"baikeLink"` // 百度百科链接
	Industry *string `form:"industry"` // 一级行业标签
	IndustryName *string `form:"industryName"` // 一级行业标签名称
	MarketValuationRmb *float64 `form:"marketValuationRmb"` // 公司最新市值/估值-亿元
	Tags *string `form:"tags"` // 企业全部标签
	LatestInvestmentId *string `form:"latestInvestmentId"` // 最新融资id
	LatestInvestTime *time.Time `form:"latestInvestTime"` // 最新融资时间
	LatestExposureTime *time.Time `form:"latestExposureTime"` // 最新融资曝光时间
	LatestInvestRound *float64 `form:"latestInvestRound"` // 最新融资轮次
	LatestInvestAmountType *float64 `form:"latestInvestAmountType"` // 最新融资金额类型
	LatestInvestAmount *string `form:"latestInvestAmount"` // 最新融资金额
	UnifiedInvestAmount *float64 `form:"unifiedInvestAmount"` // 最新融资金额转化（单位：万）
	LatestInvestCurrency *float64 `form:"latestInvestCurrency"` // 最新融资金融币种
	LatestInvestor *string `form:"latestInvestor"` // 最新融资投资方
	InvestmentNum *float64 `form:"investmentNum"` // 历史融资次数
	TotalInvestors *string `form:"totalInvestors"` // 历史投资方
	TotalInvestAmount *float64 `form:"totalInvestAmount"` // 历史融资总额（统一化：万元）
	ExternalInvestmentCount *float64 `form:"externalInvestmentCount"` // 对外投资总数
	CompanyContact *string `form:"companyContact"` // 联系信息
	OrganizationalStructure *string `form:"organizationalStructure"` // 企业组织架构
	SecuritiesInfo *string `form:"securitiesInfo"` // 证券信息
	RegFullName *string `form:"regFullName"` // 注册名称
	RegTime *time.Time `form:"regTime"` // 注册时间
	RegLocation *string `form:"regLocation"` // 注册地址
	RegCapital *float64 `form:"regCapital"` // 注册资本（万）
	RegCapitalCurrency *float64 `form:"regCapitalCurrency"` // 注册资本币种
	PaidCapital *float64 `form:"paidCapital"` // 实缴资本（万）
	PaidCapitalCurrency *float64 `form:"paidCapitalCurrency"` // 实缴资本币种
	RegType *float64 `form:"regType"` // 企业注册类型
	LegalRepresent *string `form:"legalRepresent"` // 法人代表
	RegNumber *string `form:"regNumber"` // 工商注册号
	CreditNumber *string `form:"creditNumber"` // 统一社会信用代码
	BusinessPeriod *string `form:"businessPeriod"` // 营业期限
	BusinessScope *string `form:"businessScope"` // 经营范围
	RegIndustry *float64 `form:"regIndustry"` // 行业门类
	RegIndustryName *string `form:"regIndustryName"` // 行业门类名称
	RegSubIndustry *float64 `form:"regSubIndustry"` // 行业大类
	RegSubIndustryName *string `form:"regSubIndustryName"` // 行业大类名称
	RegMiddleIndustry *float64 `form:"regMiddleIndustry"` // 行业中类
	RegMiddleIndustryName *string `form:"regMiddleIndustryName"` // 行业中类名称
	RegSmallIndustry *float64 `form:"regSmallIndustry"` // 行业小类
	RegSmallIndustryName *string `form:"regSmallIndustryName"` // 行业小类名称
	RegInstitute *string `form:"regInstitute"` // 登记机关
	ApproveTime *time.Time `form:"approveTime"` // 核准日期
	InvestRoundName *string `form:"investRoundName"` // 当前轮次名称
	OperatingStatusName *string `form:"operatingStatusName"` // 运营状态名称
	LatestInvestRoundName *string `form:"latestInvestRoundName"` // 最新融资轮次名称
	RegTypeName *string `form:"regTypeName"` // 企业注册类型名称
	RegCapitalCurrencyName *string `form:"regCapitalCurrencyName"` // 注册资本币种名称
	PaidCapitalCurrencyName *string `form:"paidCapitalCurrencyName"` // 实缴资本币种名称
	LatestInvestAmountTypeName *string `form:"latestInvestAmountTypeName"` // 最新融资金额类型名称
	LatestInvestCurrencyName *string `form:"latestInvestCurrencyName"` // 最新融资金额币种名称
	Remark *string `form:"remark"` // 备注
	IndustryStatus *string `form:"industryStatus"` // 行业地位
	PecuniaryCondition *string `form:"pecuniaryCondition"` // 财务情况
	TechAdvanced *string `form:"techAdvanced"` // 技术先进性
	DataSource *string `form:"dataSource"` // 数据来源
	IsBlacklist *int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag *string `form:"tag"` // 行业数据标签
	LogoPicture *string `form:"logoPicture"` // logo图片
	OrganizationalStructureNew *string `form:"organizationalStructureNew"` // 组织架构（新）
	UpdatedBy *int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *EnterpriseListRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// Handle 获取查询条件
func (r *EnterpriseListRequest) Handle() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
        if r.Id != nil {
            // 默认等于查询
            db = db.Where("id = ?", *r.Id)
        }
        if r.EnterpriseNo != nil {
            // 默认等于查询
            db = db.Where("enterprise_no = ?", *r.EnterpriseNo)
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
        if r.OperatingStatus != nil {
            // 默认等于查询
            db = db.Where("operating_status = ?", *r.OperatingStatus)
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
        if r.LatestInvestmentId != nil {
            // 默认等于查询
            db = db.Where("latest_investment_id = ?", *r.LatestInvestmentId)
        }
        if r.LatestInvestTime != nil {
            // 默认等于查询
            db = db.Where("latest_invest_time = ?", *r.LatestInvestTime)
        }
        if r.LatestExposureTime != nil {
            // 默认等于查询
            db = db.Where("latest_exposure_time = ?", *r.LatestExposureTime)
        }
        if r.LatestInvestRound != nil {
            // 默认等于查询
            db = db.Where("latest_invest_round = ?", *r.LatestInvestRound)
        }
        if r.LatestInvestAmountType != nil {
            // 默认等于查询
            db = db.Where("latest_invest_amount_type = ?", *r.LatestInvestAmountType)
        }
        if r.LatestInvestAmount != nil {
            // 默认等于查询
            db = db.Where("latest_invest_amount = ?", *r.LatestInvestAmount)
        }
        if r.UnifiedInvestAmount != nil {
            // 默认等于查询
            db = db.Where("unified_invest_amount = ?", *r.UnifiedInvestAmount)
        }
        if r.LatestInvestCurrency != nil {
            // 默认等于查询
            db = db.Where("latest_invest_currency = ?", *r.LatestInvestCurrency)
        }
        if r.LatestInvestor != nil {
            // 默认等于查询
            db = db.Where("latest_investor = ?", *r.LatestInvestor)
        }
        if r.InvestmentNum != nil {
            // 默认等于查询
            db = db.Where("investment_num = ?", *r.InvestmentNum)
        }
        if r.TotalInvestors != nil {
            // 默认等于查询
            db = db.Where("total_investors = ?", *r.TotalInvestors)
        }
        if r.TotalInvestAmount != nil {
            // 默认等于查询
            db = db.Where("total_invest_amount = ?", *r.TotalInvestAmount)
        }
        if r.ExternalInvestmentCount != nil {
            // 默认等于查询
            db = db.Where("external_investment_count = ?", *r.ExternalInvestmentCount)
        }
        if r.CompanyContact != nil {
            // 默认等于查询
            db = db.Where("company_contact = ?", *r.CompanyContact)
        }
        if r.OrganizationalStructure != nil {
            // 默认等于查询
            db = db.Where("organizational_structure = ?", *r.OrganizationalStructure)
        }
        if r.SecuritiesInfo != nil {
            // 默认等于查询
            db = db.Where("securities_info = ?", *r.SecuritiesInfo)
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
        if r.PaidCapital != nil {
            // 默认等于查询
            db = db.Where("paid_capital = ?", *r.PaidCapital)
        }
        if r.PaidCapitalCurrency != nil {
            // 默认等于查询
            db = db.Where("paid_capital_currency = ?", *r.PaidCapitalCurrency)
        }
        if r.RegType != nil {
            // 默认等于查询
            db = db.Where("reg_type = ?", *r.RegType)
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
        if r.InvestRoundName != nil {
            // 默认等于查询
            db = db.Where("invest_round_name = ?", *r.InvestRoundName)
        }
        if r.OperatingStatusName != nil {
            // 默认等于查询
            db = db.Where("operating_status_name = ?", *r.OperatingStatusName)
        }
        if r.LatestInvestRoundName != nil {
            // 默认等于查询
            db = db.Where("latest_invest_round_name = ?", *r.LatestInvestRoundName)
        }
        if r.RegTypeName != nil {
            // 默认等于查询
            db = db.Where("reg_type_name = ?", *r.RegTypeName)
        }
        if r.RegCapitalCurrencyName != nil {
            // 默认等于查询
            db = db.Where("reg_capital_currency_name = ?", *r.RegCapitalCurrencyName)
        }
        if r.PaidCapitalCurrencyName != nil {
            // 默认等于查询
            db = db.Where("paid_capital_currency_name = ?", *r.PaidCapitalCurrencyName)
        }
        if r.LatestInvestAmountTypeName != nil {
            // 默认等于查询
            db = db.Where("latest_invest_amount_type_name = ?", *r.LatestInvestAmountTypeName)
        }
        if r.LatestInvestCurrencyName != nil {
            // 默认等于查询
            db = db.Where("latest_invest_currency_name = ?", *r.LatestInvestCurrencyName)
        }
        if r.Remark != nil {
            // 默认等于查询
            db = db.Where("remark = ?", *r.Remark)
        }
        if r.IndustryStatus != nil {
            // 默认等于查询
            db = db.Where("industry_status = ?", *r.IndustryStatus)
        }
        if r.PecuniaryCondition != nil {
            // 默认等于查询
            db = db.Where("pecuniary_condition = ?", *r.PecuniaryCondition)
        }
        if r.TechAdvanced != nil {
            // 默认等于查询
            db = db.Where("tech_advanced = ?", *r.TechAdvanced)
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
        if r.LogoPicture != nil {
            // 默认等于查询
            db = db.Where("logo_picture = ?", *r.LogoPicture)
        }
        if r.OrganizationalStructureNew != nil {
            // 默认等于查询
            db = db.Where("organizational_structure_new = ?", *r.OrganizationalStructureNew)
        }
        if r.UpdatedBy != nil {
            // 默认等于查询
            db = db.Where("updated_by = ?", *r.UpdatedBy)
        }
		return db
	}
}

// EnterpriseCreateRequest 创建res_enterprise请求参数
type EnterpriseCreateRequest struct {
	models.Validator
	EnterpriseNo string `form:"enterpriseNo" validate:"required" message:"企业编码不能为空"` // 企业编码
	BriefName string `form:"briefName" validate:"required" message:"企业简称不能为空"` // 企业简称
	FullName string `form:"fullName" validate:"required" message:"企业全称不能为空"` // 企业全称
	ComOtherName string `form:"comOtherName" validate:"required" message:"其它简称不能为空"` // 其它简称
	Logo string `form:"logo" validate:"required" message:"企业Logo不能为空"` // 企业Logo
	ComType string `form:"comType" validate:"required" message:"企业类型不能为空"` // 企业类型
	Website string `form:"website" validate:"required" message:"企业网址不能为空"` // 企业网址
	BriefIntro string `form:"briefIntro" validate:"required" message:"一句话简介不能为空"` // 一句话简介
	DetailIntro string `form:"detailIntro" validate:"required" message:"企业介绍不能为空"` // 企业介绍
	EstablishTime *time.Time `form:"establishTime" validate:"required" message:"成立时间不能为空"` // 成立时间
	InvestRound float64 `form:"investRound"` // 当前轮次
	OperatingStatus float64 `form:"operatingStatus"` // 运营状态
	Staff string `form:"staff" validate:"required" message:"人员规模不能为空"` // 人员规模
	BaikeLink string `form:"baikeLink" validate:"required" message:"百度百科链接不能为空"` // 百度百科链接
	Industry string `form:"industry" validate:"required" message:"一级行业标签不能为空"` // 一级行业标签
	IndustryName string `form:"industryName" validate:"required" message:"一级行业标签名称不能为空"` // 一级行业标签名称
	MarketValuationRmb float64 `form:"marketValuationRmb"` // 公司最新市值/估值-亿元
	Tags string `form:"tags" validate:"required" message:"企业全部标签不能为空"` // 企业全部标签
	LatestInvestmentId string `form:"latestInvestmentId" validate:"required" message:"最新融资id不能为空"` // 最新融资id
	LatestInvestTime *time.Time `form:"latestInvestTime" validate:"required" message:"最新融资时间不能为空"` // 最新融资时间
	LatestExposureTime *time.Time `form:"latestExposureTime" validate:"required" message:"最新融资曝光时间不能为空"` // 最新融资曝光时间
	LatestInvestRound float64 `form:"latestInvestRound"` // 最新融资轮次
	LatestInvestAmountType float64 `form:"latestInvestAmountType"` // 最新融资金额类型
	LatestInvestAmount string `form:"latestInvestAmount" validate:"required" message:"最新融资金额不能为空"` // 最新融资金额
	UnifiedInvestAmount float64 `form:"unifiedInvestAmount"` // 最新融资金额转化（单位：万）
	LatestInvestCurrency float64 `form:"latestInvestCurrency"` // 最新融资金融币种
	LatestInvestor string `form:"latestInvestor" validate:"required" message:"最新融资投资方不能为空"` // 最新融资投资方
	InvestmentNum float64 `form:"investmentNum"` // 历史融资次数
	TotalInvestors string `form:"totalInvestors" validate:"required" message:"历史投资方不能为空"` // 历史投资方
	TotalInvestAmount float64 `form:"totalInvestAmount"` // 历史融资总额（统一化：万元）
	ExternalInvestmentCount float64 `form:"externalInvestmentCount"` // 对外投资总数
	CompanyContact string `form:"companyContact" validate:"required" message:"联系信息不能为空"` // 联系信息
	OrganizationalStructure string `form:"organizationalStructure" validate:"required" message:"企业组织架构不能为空"` // 企业组织架构
	SecuritiesInfo string `form:"securitiesInfo" validate:"required" message:"证券信息不能为空"` // 证券信息
	RegFullName string `form:"regFullName" validate:"required" message:"注册名称不能为空"` // 注册名称
	RegTime *time.Time `form:"regTime" validate:"required" message:"注册时间不能为空"` // 注册时间
	RegLocation string `form:"regLocation" validate:"required" message:"注册地址不能为空"` // 注册地址
	RegCapital float64 `form:"regCapital"` // 注册资本（万）
	RegCapitalCurrency float64 `form:"regCapitalCurrency"` // 注册资本币种
	PaidCapital float64 `form:"paidCapital"` // 实缴资本（万）
	PaidCapitalCurrency float64 `form:"paidCapitalCurrency"` // 实缴资本币种
	RegType float64 `form:"regType"` // 企业注册类型
	LegalRepresent string `form:"legalRepresent" validate:"required" message:"法人代表不能为空"` // 法人代表
	RegNumber string `form:"regNumber" validate:"required" message:"工商注册号不能为空"` // 工商注册号
	CreditNumber string `form:"creditNumber" validate:"required" message:"统一社会信用代码不能为空"` // 统一社会信用代码
	BusinessPeriod string `form:"businessPeriod" validate:"required" message:"营业期限不能为空"` // 营业期限
	BusinessScope string `form:"businessScope" validate:"required" message:"经营范围不能为空"` // 经营范围
	RegIndustry float64 `form:"regIndustry"` // 行业门类
	RegIndustryName string `form:"regIndustryName" validate:"required" message:"行业门类名称不能为空"` // 行业门类名称
	RegSubIndustry float64 `form:"regSubIndustry"` // 行业大类
	RegSubIndustryName string `form:"regSubIndustryName" validate:"required" message:"行业大类名称不能为空"` // 行业大类名称
	RegMiddleIndustry float64 `form:"regMiddleIndustry"` // 行业中类
	RegMiddleIndustryName string `form:"regMiddleIndustryName" validate:"required" message:"行业中类名称不能为空"` // 行业中类名称
	RegSmallIndustry float64 `form:"regSmallIndustry"` // 行业小类
	RegSmallIndustryName string `form:"regSmallIndustryName" validate:"required" message:"行业小类名称不能为空"` // 行业小类名称
	RegInstitute string `form:"regInstitute" validate:"required" message:"登记机关不能为空"` // 登记机关
	ApproveTime *time.Time `form:"approveTime" validate:"required" message:"核准日期不能为空"` // 核准日期
	InvestRoundName string `form:"investRoundName" validate:"required" message:"当前轮次名称不能为空"` // 当前轮次名称
	OperatingStatusName string `form:"operatingStatusName" validate:"required" message:"运营状态名称不能为空"` // 运营状态名称
	LatestInvestRoundName string `form:"latestInvestRoundName" validate:"required" message:"最新融资轮次名称不能为空"` // 最新融资轮次名称
	RegTypeName string `form:"regTypeName" validate:"required" message:"企业注册类型名称不能为空"` // 企业注册类型名称
	RegCapitalCurrencyName string `form:"regCapitalCurrencyName" validate:"required" message:"注册资本币种名称不能为空"` // 注册资本币种名称
	PaidCapitalCurrencyName string `form:"paidCapitalCurrencyName" validate:"required" message:"实缴资本币种名称不能为空"` // 实缴资本币种名称
	LatestInvestAmountTypeName string `form:"latestInvestAmountTypeName" validate:"required" message:"最新融资金额类型名称不能为空"` // 最新融资金额类型名称
	LatestInvestCurrencyName string `form:"latestInvestCurrencyName" validate:"required" message:"最新融资金额币种名称不能为空"` // 最新融资金额币种名称
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	IndustryStatus string `form:"industryStatus" validate:"required" message:"行业地位不能为空"` // 行业地位
	PecuniaryCondition string `form:"pecuniaryCondition" validate:"required" message:"财务情况不能为空"` // 财务情况
	TechAdvanced string `form:"techAdvanced" validate:"required" message:"技术先进性不能为空"` // 技术先进性
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	LogoPicture string `form:"logoPicture" validate:"required" message:"logo图片不能为空"` // logo图片
	OrganizationalStructureNew string `form:"organizationalStructureNew" validate:"required" message:"组织架构（新）不能为空"` // 组织架构（新）
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *EnterpriseCreateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// EnterpriseUpdateRequest 更新res_enterprise请求参数
type EnterpriseUpdateRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
	EnterpriseNo string `form:"enterpriseNo" validate:"required" message:"企业编码不能为空"` // 企业编码
	BriefName string `form:"briefName" validate:"required" message:"企业简称不能为空"` // 企业简称
	FullName string `form:"fullName" validate:"required" message:"企业全称不能为空"` // 企业全称
	ComOtherName string `form:"comOtherName" validate:"required" message:"其它简称不能为空"` // 其它简称
	Logo string `form:"logo" validate:"required" message:"企业Logo不能为空"` // 企业Logo
	ComType string `form:"comType" validate:"required" message:"企业类型不能为空"` // 企业类型
	Website string `form:"website" validate:"required" message:"企业网址不能为空"` // 企业网址
	BriefIntro string `form:"briefIntro" validate:"required" message:"一句话简介不能为空"` // 一句话简介
	DetailIntro string `form:"detailIntro" validate:"required" message:"企业介绍不能为空"` // 企业介绍
	EstablishTime *time.Time `form:"establishTime" validate:"required" message:"成立时间不能为空"` // 成立时间
	InvestRound float64 `form:"investRound"` // 当前轮次
	OperatingStatus float64 `form:"operatingStatus"` // 运营状态
	Staff string `form:"staff" validate:"required" message:"人员规模不能为空"` // 人员规模
	BaikeLink string `form:"baikeLink" validate:"required" message:"百度百科链接不能为空"` // 百度百科链接
	Industry string `form:"industry" validate:"required" message:"一级行业标签不能为空"` // 一级行业标签
	IndustryName string `form:"industryName" validate:"required" message:"一级行业标签名称不能为空"` // 一级行业标签名称
	MarketValuationRmb float64 `form:"marketValuationRmb"` // 公司最新市值/估值-亿元
	Tags string `form:"tags" validate:"required" message:"企业全部标签不能为空"` // 企业全部标签
	LatestInvestmentId string `form:"latestInvestmentId" validate:"required" message:"最新融资id不能为空"` // 最新融资id
	LatestInvestTime *time.Time `form:"latestInvestTime" validate:"required" message:"最新融资时间不能为空"` // 最新融资时间
	LatestExposureTime *time.Time `form:"latestExposureTime" validate:"required" message:"最新融资曝光时间不能为空"` // 最新融资曝光时间
	LatestInvestRound float64 `form:"latestInvestRound"` // 最新融资轮次
	LatestInvestAmountType float64 `form:"latestInvestAmountType"` // 最新融资金额类型
	LatestInvestAmount string `form:"latestInvestAmount" validate:"required" message:"最新融资金额不能为空"` // 最新融资金额
	UnifiedInvestAmount float64 `form:"unifiedInvestAmount"` // 最新融资金额转化（单位：万）
	LatestInvestCurrency float64 `form:"latestInvestCurrency"` // 最新融资金融币种
	LatestInvestor string `form:"latestInvestor" validate:"required" message:"最新融资投资方不能为空"` // 最新融资投资方
	InvestmentNum float64 `form:"investmentNum"` // 历史融资次数
	TotalInvestors string `form:"totalInvestors" validate:"required" message:"历史投资方不能为空"` // 历史投资方
	TotalInvestAmount float64 `form:"totalInvestAmount"` // 历史融资总额（统一化：万元）
	ExternalInvestmentCount float64 `form:"externalInvestmentCount"` // 对外投资总数
	CompanyContact string `form:"companyContact" validate:"required" message:"联系信息不能为空"` // 联系信息
	OrganizationalStructure string `form:"organizationalStructure" validate:"required" message:"企业组织架构不能为空"` // 企业组织架构
	SecuritiesInfo string `form:"securitiesInfo" validate:"required" message:"证券信息不能为空"` // 证券信息
	RegFullName string `form:"regFullName" validate:"required" message:"注册名称不能为空"` // 注册名称
	RegTime *time.Time `form:"regTime" validate:"required" message:"注册时间不能为空"` // 注册时间
	RegLocation string `form:"regLocation" validate:"required" message:"注册地址不能为空"` // 注册地址
	RegCapital float64 `form:"regCapital"` // 注册资本（万）
	RegCapitalCurrency float64 `form:"regCapitalCurrency"` // 注册资本币种
	PaidCapital float64 `form:"paidCapital"` // 实缴资本（万）
	PaidCapitalCurrency float64 `form:"paidCapitalCurrency"` // 实缴资本币种
	RegType float64 `form:"regType"` // 企业注册类型
	LegalRepresent string `form:"legalRepresent" validate:"required" message:"法人代表不能为空"` // 法人代表
	RegNumber string `form:"regNumber" validate:"required" message:"工商注册号不能为空"` // 工商注册号
	CreditNumber string `form:"creditNumber" validate:"required" message:"统一社会信用代码不能为空"` // 统一社会信用代码
	BusinessPeriod string `form:"businessPeriod" validate:"required" message:"营业期限不能为空"` // 营业期限
	BusinessScope string `form:"businessScope" validate:"required" message:"经营范围不能为空"` // 经营范围
	RegIndustry float64 `form:"regIndustry"` // 行业门类
	RegIndustryName string `form:"regIndustryName" validate:"required" message:"行业门类名称不能为空"` // 行业门类名称
	RegSubIndustry float64 `form:"regSubIndustry"` // 行业大类
	RegSubIndustryName string `form:"regSubIndustryName" validate:"required" message:"行业大类名称不能为空"` // 行业大类名称
	RegMiddleIndustry float64 `form:"regMiddleIndustry"` // 行业中类
	RegMiddleIndustryName string `form:"regMiddleIndustryName" validate:"required" message:"行业中类名称不能为空"` // 行业中类名称
	RegSmallIndustry float64 `form:"regSmallIndustry"` // 行业小类
	RegSmallIndustryName string `form:"regSmallIndustryName" validate:"required" message:"行业小类名称不能为空"` // 行业小类名称
	RegInstitute string `form:"regInstitute" validate:"required" message:"登记机关不能为空"` // 登记机关
	ApproveTime *time.Time `form:"approveTime" validate:"required" message:"核准日期不能为空"` // 核准日期
	InvestRoundName string `form:"investRoundName" validate:"required" message:"当前轮次名称不能为空"` // 当前轮次名称
	OperatingStatusName string `form:"operatingStatusName" validate:"required" message:"运营状态名称不能为空"` // 运营状态名称
	LatestInvestRoundName string `form:"latestInvestRoundName" validate:"required" message:"最新融资轮次名称不能为空"` // 最新融资轮次名称
	RegTypeName string `form:"regTypeName" validate:"required" message:"企业注册类型名称不能为空"` // 企业注册类型名称
	RegCapitalCurrencyName string `form:"regCapitalCurrencyName" validate:"required" message:"注册资本币种名称不能为空"` // 注册资本币种名称
	PaidCapitalCurrencyName string `form:"paidCapitalCurrencyName" validate:"required" message:"实缴资本币种名称不能为空"` // 实缴资本币种名称
	LatestInvestAmountTypeName string `form:"latestInvestAmountTypeName" validate:"required" message:"最新融资金额类型名称不能为空"` // 最新融资金额类型名称
	LatestInvestCurrencyName string `form:"latestInvestCurrencyName" validate:"required" message:"最新融资金额币种名称不能为空"` // 最新融资金额币种名称
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	IndustryStatus string `form:"industryStatus" validate:"required" message:"行业地位不能为空"` // 行业地位
	PecuniaryCondition string `form:"pecuniaryCondition" validate:"required" message:"财务情况不能为空"` // 财务情况
	TechAdvanced string `form:"techAdvanced" validate:"required" message:"技术先进性不能为空"` // 技术先进性
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	LogoPicture string `form:"logoPicture" validate:"required" message:"logo图片不能为空"` // logo图片
	OrganizationalStructureNew string `form:"organizationalStructureNew" validate:"required" message:"组织架构（新）不能为空"` // 组织架构（新）
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *EnterpriseUpdateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// EnterpriseDeleteRequest 删除res_enterprise请求参数
type EnterpriseDeleteRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *EnterpriseDeleteRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// EnterpriseGetByIDRequest 根据ID获取res_enterprise请求参数
type EnterpriseGetByIDRequest struct {
	models.Validator
	Id int64 `uri:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *EnterpriseGetByIDRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}