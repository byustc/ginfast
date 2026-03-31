package models

import (
	"time"
	"gin-fast/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FinanceListRequest res_finance列表请求参数
type FinanceListRequest struct {
	models.BasePaging
	models.Validator
	Id *int64 `form:"id"` // 主键
	FinanceNo *string `form:"financeNo"` // 金融编码
	BriefName *string `form:"briefName"` // 机构简称
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
	Location *string `form:"location"` // 所在地
	FinType *string `form:"finType"` // 金融机构类型
	FinNature *string `form:"finNature"` // 机构性质
	IsConfidential *string `form:"isConfidential"` // 是否签订保密协议
	CooperationTime *time.Time `form:"cooperationTime"` // 合作建立时间
	Cooperativeness *string `form:"cooperativeness"` // 合作程度
	CooperationScope *string `form:"cooperationScope"` // 合作范围
	Remark *string `form:"remark"` // 备注
	InvestRoundName *string `form:"investRoundName"` // 当前轮次名称
	OperatingStatusName *string `form:"operatingStatusName"` // 运营状态名称
	RegCapitalCurrencyName *string `form:"regCapitalCurrencyName"` // 注册资本币种名称
	PaidCapitalCurrencyName *string `form:"paidCapitalCurrencyName"` // 实缴资本币种名称
	RegTypeName *string `form:"regTypeName"` // 企业注册类型名称
	DataSource *string `form:"dataSource"` // 数据来源
	IsBlacklist *int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	PositionNum *float64 `form:"positionNum"` // 挂职人数
	PositionAdvancedNum *float64 `form:"positionAdvancedNum"` // 高级挂职人数
	Tag *string `form:"tag"` // 行业数据标签
	DockingPerson *string `form:"dockingPerson"` // 对接人
	CorrespondingSourceId *float64 `form:"correspondingSourceId"` // 对应数据源id
	DockingPersonnel *string `form:"dockingPersonnel"` // 对接人员
	Level *string `form:"level"` // 级别
	Partner *int `form:"partner"` // 是否合作伙伴: 0否 1是
	WarehousingTime *time.Time `form:"warehousingTime"` // 入库时间
	Contact *string `form:"contact"` // 联系人
	ContactNumber *string `form:"contactNumber"` // 联系电话
	EntryPerson *int64 `form:"entryPerson"` // 录入人
	ContactInformation *string `form:"contactInformation"` // 联系方式
	Mailbox *string `form:"mailbox"` // 邮箱
	Duties *string `form:"duties"` // 职务
	ActivitiesUndertaken *string `form:"activitiesUndertaken"` // 拟开展活动
	LogoPicture *string `form:"logoPicture"` // logo图片
	UpdatedBy *int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *FinanceListRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// Handle 获取查询条件
func (r *FinanceListRequest) Handle() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
        if r.Id != nil {
            // 默认等于查询
            db = db.Where("id = ?", *r.Id)
        }
        if r.FinanceNo != nil {
            // 默认等于查询
            db = db.Where("finance_no = ?", *r.FinanceNo)
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
        if r.Location != nil {
            // 默认等于查询
            db = db.Where("location = ?", *r.Location)
        }
        if r.FinType != nil {
            // 默认等于查询
            db = db.Where("fin_type = ?", *r.FinType)
        }
        if r.FinNature != nil {
            // 默认等于查询
            db = db.Where("fin_nature = ?", *r.FinNature)
        }
        if r.IsConfidential != nil {
            // 默认等于查询
            db = db.Where("is_confidential = ?", *r.IsConfidential)
        }
        if r.CooperationTime != nil {
            // 默认等于查询
            db = db.Where("cooperation_time = ?", *r.CooperationTime)
        }
        if r.Cooperativeness != nil {
            // 默认等于查询
            db = db.Where("cooperativeness = ?", *r.Cooperativeness)
        }
        if r.CooperationScope != nil {
            // 默认等于查询
            db = db.Where("cooperation_scope = ?", *r.CooperationScope)
        }
        if r.Remark != nil {
            // 默认等于查询
            db = db.Where("remark = ?", *r.Remark)
        }
        if r.InvestRoundName != nil {
            // 默认等于查询
            db = db.Where("invest_round_name = ?", *r.InvestRoundName)
        }
        if r.OperatingStatusName != nil {
            // 默认等于查询
            db = db.Where("operating_status_name = ?", *r.OperatingStatusName)
        }
        if r.RegCapitalCurrencyName != nil {
            // 默认等于查询
            db = db.Where("reg_capital_currency_name = ?", *r.RegCapitalCurrencyName)
        }
        if r.PaidCapitalCurrencyName != nil {
            // 默认等于查询
            db = db.Where("paid_capital_currency_name = ?", *r.PaidCapitalCurrencyName)
        }
        if r.RegTypeName != nil {
            // 默认等于查询
            db = db.Where("reg_type_name = ?", *r.RegTypeName)
        }
        if r.DataSource != nil {
            // 默认等于查询
            db = db.Where("data_source = ?", *r.DataSource)
        }
        if r.IsBlacklist != nil {
            // 默认等于查询
            db = db.Where("is_blacklist = ?", *r.IsBlacklist)
        }
        if r.PositionNum != nil {
            // 默认等于查询
            db = db.Where("position_num = ?", *r.PositionNum)
        }
        if r.PositionAdvancedNum != nil {
            // 默认等于查询
            db = db.Where("position_advanced_num = ?", *r.PositionAdvancedNum)
        }
        if r.Tag != nil {
            // 默认等于查询
            db = db.Where("tag = ?", *r.Tag)
        }
        if r.DockingPerson != nil {
            // 默认等于查询
            db = db.Where("docking_person = ?", *r.DockingPerson)
        }
        if r.CorrespondingSourceId != nil {
            // 默认等于查询
            db = db.Where("corresponding_source_id = ?", *r.CorrespondingSourceId)
        }
        if r.DockingPersonnel != nil {
            // 默认等于查询
            db = db.Where("docking_personnel = ?", *r.DockingPersonnel)
        }
        if r.Level != nil {
            // 默认等于查询
            db = db.Where("level = ?", *r.Level)
        }
        if r.Partner != nil {
            // 默认等于查询
            db = db.Where("partner = ?", *r.Partner)
        }
        if r.WarehousingTime != nil {
            // 默认等于查询
            db = db.Where("warehousing_time = ?", *r.WarehousingTime)
        }
        if r.Contact != nil {
            // 默认等于查询
            db = db.Where("contact = ?", *r.Contact)
        }
        if r.ContactNumber != nil {
            // 默认等于查询
            db = db.Where("contact_number = ?", *r.ContactNumber)
        }
        if r.EntryPerson != nil {
            // 默认等于查询
            db = db.Where("entry_person = ?", *r.EntryPerson)
        }
        if r.ContactInformation != nil {
            // 默认等于查询
            db = db.Where("contact_information = ?", *r.ContactInformation)
        }
        if r.Mailbox != nil {
            // 默认等于查询
            db = db.Where("mailbox = ?", *r.Mailbox)
        }
        if r.Duties != nil {
            // 默认等于查询
            db = db.Where("duties = ?", *r.Duties)
        }
        if r.ActivitiesUndertaken != nil {
            // 默认等于查询
            db = db.Where("activities_undertaken = ?", *r.ActivitiesUndertaken)
        }
        if r.LogoPicture != nil {
            // 默认等于查询
            db = db.Where("logo_picture = ?", *r.LogoPicture)
        }
        if r.UpdatedBy != nil {
            // 默认等于查询
            db = db.Where("updated_by = ?", *r.UpdatedBy)
        }
		return db
	}
}

// FinanceCreateRequest 创建res_finance请求参数
type FinanceCreateRequest struct {
	models.Validator
	FinanceNo string `form:"financeNo" validate:"required" message:"金融编码不能为空"` // 金融编码
	BriefName string `form:"briefName" validate:"required" message:"机构简称不能为空"` // 机构简称
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
	Location string `form:"location" validate:"required" message:"所在地不能为空"` // 所在地
	FinType string `form:"finType" validate:"required" message:"金融机构类型不能为空"` // 金融机构类型
	FinNature string `form:"finNature" validate:"required" message:"机构性质不能为空"` // 机构性质
	IsConfidential string `form:"isConfidential" validate:"required" message:"是否签订保密协议不能为空"` // 是否签订保密协议
	CooperationTime *time.Time `form:"cooperationTime" validate:"required" message:"合作建立时间不能为空"` // 合作建立时间
	Cooperativeness string `form:"cooperativeness" validate:"required" message:"合作程度不能为空"` // 合作程度
	CooperationScope string `form:"cooperationScope" validate:"required" message:"合作范围不能为空"` // 合作范围
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	InvestRoundName string `form:"investRoundName" validate:"required" message:"当前轮次名称不能为空"` // 当前轮次名称
	OperatingStatusName string `form:"operatingStatusName" validate:"required" message:"运营状态名称不能为空"` // 运营状态名称
	RegCapitalCurrencyName string `form:"regCapitalCurrencyName" validate:"required" message:"注册资本币种名称不能为空"` // 注册资本币种名称
	PaidCapitalCurrencyName string `form:"paidCapitalCurrencyName" validate:"required" message:"实缴资本币种名称不能为空"` // 实缴资本币种名称
	RegTypeName string `form:"regTypeName" validate:"required" message:"企业注册类型名称不能为空"` // 企业注册类型名称
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	PositionNum float64 `form:"positionNum"` // 挂职人数
	PositionAdvancedNum float64 `form:"positionAdvancedNum"` // 高级挂职人数
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	DockingPerson string `form:"dockingPerson" validate:"required" message:"对接人不能为空"` // 对接人
	CorrespondingSourceId float64 `form:"correspondingSourceId"` // 对应数据源id
	DockingPersonnel string `form:"dockingPersonnel" validate:"required" message:"对接人员不能为空"` // 对接人员
	Level string `form:"level" validate:"required" message:"级别不能为空"` // 级别
	Partner int `form:"partner"` // 是否合作伙伴: 0否 1是
	WarehousingTime *time.Time `form:"warehousingTime" validate:"required" message:"入库时间不能为空"` // 入库时间
	Contact string `form:"contact" validate:"required" message:"联系人不能为空"` // 联系人
	ContactNumber string `form:"contactNumber" validate:"required" message:"联系电话不能为空"` // 联系电话
	EntryPerson int64 `form:"entryPerson"` // 录入人
	ContactInformation string `form:"contactInformation" validate:"required" message:"联系方式不能为空"` // 联系方式
	Mailbox string `form:"mailbox" validate:"required" message:"邮箱不能为空"` // 邮箱
	Duties string `form:"duties" validate:"required" message:"职务不能为空"` // 职务
	ActivitiesUndertaken string `form:"activitiesUndertaken" validate:"required" message:"拟开展活动不能为空"` // 拟开展活动
	LogoPicture string `form:"logoPicture" validate:"required" message:"logo图片不能为空"` // logo图片
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *FinanceCreateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// FinanceUpdateRequest 更新res_finance请求参数
type FinanceUpdateRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
	FinanceNo string `form:"financeNo" validate:"required" message:"金融编码不能为空"` // 金融编码
	BriefName string `form:"briefName" validate:"required" message:"机构简称不能为空"` // 机构简称
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
	Location string `form:"location" validate:"required" message:"所在地不能为空"` // 所在地
	FinType string `form:"finType" validate:"required" message:"金融机构类型不能为空"` // 金融机构类型
	FinNature string `form:"finNature" validate:"required" message:"机构性质不能为空"` // 机构性质
	IsConfidential string `form:"isConfidential" validate:"required" message:"是否签订保密协议不能为空"` // 是否签订保密协议
	CooperationTime *time.Time `form:"cooperationTime" validate:"required" message:"合作建立时间不能为空"` // 合作建立时间
	Cooperativeness string `form:"cooperativeness" validate:"required" message:"合作程度不能为空"` // 合作程度
	CooperationScope string `form:"cooperationScope" validate:"required" message:"合作范围不能为空"` // 合作范围
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	InvestRoundName string `form:"investRoundName" validate:"required" message:"当前轮次名称不能为空"` // 当前轮次名称
	OperatingStatusName string `form:"operatingStatusName" validate:"required" message:"运营状态名称不能为空"` // 运营状态名称
	RegCapitalCurrencyName string `form:"regCapitalCurrencyName" validate:"required" message:"注册资本币种名称不能为空"` // 注册资本币种名称
	PaidCapitalCurrencyName string `form:"paidCapitalCurrencyName" validate:"required" message:"实缴资本币种名称不能为空"` // 实缴资本币种名称
	RegTypeName string `form:"regTypeName" validate:"required" message:"企业注册类型名称不能为空"` // 企业注册类型名称
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	PositionNum float64 `form:"positionNum"` // 挂职人数
	PositionAdvancedNum float64 `form:"positionAdvancedNum"` // 高级挂职人数
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	DockingPerson string `form:"dockingPerson" validate:"required" message:"对接人不能为空"` // 对接人
	CorrespondingSourceId float64 `form:"correspondingSourceId"` // 对应数据源id
	DockingPersonnel string `form:"dockingPersonnel" validate:"required" message:"对接人员不能为空"` // 对接人员
	Level string `form:"level" validate:"required" message:"级别不能为空"` // 级别
	Partner int `form:"partner"` // 是否合作伙伴: 0否 1是
	WarehousingTime *time.Time `form:"warehousingTime" validate:"required" message:"入库时间不能为空"` // 入库时间
	Contact string `form:"contact" validate:"required" message:"联系人不能为空"` // 联系人
	ContactNumber string `form:"contactNumber" validate:"required" message:"联系电话不能为空"` // 联系电话
	EntryPerson int64 `form:"entryPerson"` // 录入人
	ContactInformation string `form:"contactInformation" validate:"required" message:"联系方式不能为空"` // 联系方式
	Mailbox string `form:"mailbox" validate:"required" message:"邮箱不能为空"` // 邮箱
	Duties string `form:"duties" validate:"required" message:"职务不能为空"` // 职务
	ActivitiesUndertaken string `form:"activitiesUndertaken" validate:"required" message:"拟开展活动不能为空"` // 拟开展活动
	LogoPicture string `form:"logoPicture" validate:"required" message:"logo图片不能为空"` // logo图片
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *FinanceUpdateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// FinanceDeleteRequest 删除res_finance请求参数
type FinanceDeleteRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *FinanceDeleteRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// FinanceGetByIDRequest 根据ID获取res_finance请求参数
type FinanceGetByIDRequest struct {
	models.Validator
	Id int64 `uri:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *FinanceGetByIDRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}