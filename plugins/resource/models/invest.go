package models

import (
	"context"
	"time"
	"gin-fast/app/global/app"
	"gorm.io/gorm"
)

// Invest res_invest 模型结构体
type Invest struct {
	Id int64 `gorm:"column:id;primaryKey;not null;autoIncrement" json:"id"` // 主键
	TenantId uint `gorm:"column:tenant_id;default:0;index" json:"tenantId"` // 租户ID字段
	InvestNo string `gorm:"column:invest_no;index" json:"investNo"` // 投资编码
	BriefName string `gorm:"column:brief_name" json:"briefName"` // 机构简称
	FullName string `gorm:"column:full_name" json:"fullName"` // 企业全称
	ComOtherName string `gorm:"column:com_other_name" json:"comOtherName"` // 其它简称
	Logo string `gorm:"column:logo" json:"logo"` // 企业Logo
	ComType string `gorm:"column:com_type" json:"comType"` // 企业类型
	Website string `gorm:"column:website" json:"website"` // 企业网址
	BriefIntro string `gorm:"column:brief_intro" json:"briefIntro"` // 一句话简介
	DetailIntro string `gorm:"column:detail_intro" json:"detailIntro"` // 企业介绍
	EstablishTime *time.Time `gorm:"column:establish_time" json:"establishTime"` // 成立时间
	InvestRound float64 `gorm:"column:invest_round" json:"investRound"` // 当前轮次
	OperatingStatus float64 `gorm:"column:operating_status" json:"operatingStatus"` // 运营状态
	Staff string `gorm:"column:staff" json:"staff"` // 人员规模
	BaikeLink string `gorm:"column:baike_link" json:"baikeLink"` // 百度百科链接
	Industry string `gorm:"column:industry" json:"industry"` // 一级行业标签
	IndustryName string `gorm:"column:industry_name" json:"industryName"` // 一级行业标签名称
	MarketValuationRmb float64 `gorm:"column:market_valuation_rmb" json:"marketValuationRmb"` // 公司最新市值/估值-亿元
	Tags string `gorm:"column:tags" json:"tags"` // 企业全部标签
	RegFullName string `gorm:"column:reg_full_name" json:"regFullName"` // 注册名称
	RegTime *time.Time `gorm:"column:reg_time" json:"regTime"` // 注册时间
	RegLocation string `gorm:"column:reg_location" json:"regLocation"` // 注册地址
	RegCapital float64 `gorm:"column:reg_capital" json:"regCapital"` // 注册资本（万）
	RegCapitalCurrency float64 `gorm:"column:reg_capital_currency" json:"regCapitalCurrency"` // 注册资本币种
	PaidCapital float64 `gorm:"column:paid_capital" json:"paidCapital"` // 实缴资本（万）
	PaidCapitalCurrency float64 `gorm:"column:paid_capital_currency" json:"paidCapitalCurrency"` // 实缴资本币种
	RegType float64 `gorm:"column:reg_type" json:"regType"` // 企业注册类型
	LegalRepresent string `gorm:"column:legal_represent" json:"legalRepresent"` // 法人代表
	RegNumber string `gorm:"column:reg_number" json:"regNumber"` // 工商注册号
	CreditNumber string `gorm:"column:credit_number" json:"creditNumber"` // 统一社会信用代码
	BusinessPeriod string `gorm:"column:business_period" json:"businessPeriod"` // 营业期限
	BusinessScope string `gorm:"column:business_scope" json:"businessScope"` // 经营范围
	RegIndustry float64 `gorm:"column:reg_industry" json:"regIndustry"` // 行业门类
	RegIndustryName string `gorm:"column:reg_industry_name" json:"regIndustryName"` // 行业门类名称
	RegSubIndustry float64 `gorm:"column:reg_sub_industry" json:"regSubIndustry"` // 行业大类
	RegSubIndustryName string `gorm:"column:reg_sub_industry_name" json:"regSubIndustryName"` // 行业大类名称
	RegMiddleIndustry float64 `gorm:"column:reg_middle_industry" json:"regMiddleIndustry"` // 行业中类
	RegMiddleIndustryName string `gorm:"column:reg_middle_industry_name" json:"regMiddleIndustryName"` // 行业中类名称
	RegSmallIndustry float64 `gorm:"column:reg_small_industry" json:"regSmallIndustry"` // 行业小类
	RegSmallIndustryName string `gorm:"column:reg_small_industry_name" json:"regSmallIndustryName"` // 行业小类名称
	RegInstitute string `gorm:"column:reg_institute" json:"regInstitute"` // 登记机关
	ApproveTime *time.Time `gorm:"column:approve_time" json:"approveTime"` // 核准日期
	Location string `gorm:"column:location" json:"location"` // 所在地
	OrgNature string `gorm:"column:org_nature" json:"orgNature"` // 机构性质
	IsCvc string `gorm:"column:is_cvc" json:"isCvc"` // 是否产业CVC
	CvcDesc string `gorm:"column:cvc_desc" json:"cvcDesc"` // 产业CVC描述
	MgtScale string `gorm:"column:mgt_scale" json:"mgtScale"` // 管理规模
	InvestIndustry string `gorm:"column:invest_industry" json:"investIndustry"` // 投资行业
	SingleInvestAmount string `gorm:"column:single_invest_amount" json:"singleInvestAmount"` // 单笔投资金额
	RoundPrefer string `gorm:"column:round_prefer" json:"roundPrefer"` // 投资阶段偏好
	InvestMethod string `gorm:"column:invest_method" json:"investMethod"` // 投资方式
	InvestCase string `gorm:"column:invest_case" json:"investCase"` // 投资案例
	StarProj string `gorm:"column:star_proj" json:"starProj"` // 明星项目
	MainDemands string `gorm:"column:main_demands" json:"mainDemands"` // 主要诉求/痛点
	Remark string `gorm:"column:remark" json:"remark"` // 备注
	InvestRoundName string `gorm:"column:invest_round_name" json:"investRoundName"` // 当前轮次名称
	OperatingStatusName string `gorm:"column:operating_status_name" json:"operatingStatusName"` // 运营状态名称
	RegCapitalCurrencyName string `gorm:"column:reg_capital_currency_name" json:"regCapitalCurrencyName"` // 注册资本币种名称
	PaidCapitalCurrencyName string `gorm:"column:paid_capital_currency_name" json:"paidCapitalCurrencyName"` // 实缴资本币种名称
	RegTypeName string `gorm:"column:reg_type_name" json:"regTypeName"` // 企业注册类型名称
	InvestType string `gorm:"column:invest_type" json:"investType"` // 投资类型
	IsBlacklist int `gorm:"column:is_blacklist;default:0" json:"isBlacklist"` // 是否负面清单: 0否 1是
	NeedInvestorType string `gorm:"column:need_investor_type" json:"needInvestorType"` // 是否需要产业投资人、财务投资人
	InvestableBalance float64 `gorm:"column:investable_balance" json:"investableBalance"` // 可投余额
	DataSource string `gorm:"column:data_source" json:"dataSource"` // 数据来源
	InvestmentField string `gorm:"column:investment_field" json:"investmentField"` // 投资领域
	Tag string `gorm:"column:tag" json:"tag"` // 行业数据标签
	DockingPerson string `gorm:"column:docking_person" json:"dockingPerson"` // 对接人
	DockingPersonnel string `gorm:"column:docking_personnel" json:"dockingPersonnel"` // 对接人员
	Level string `gorm:"column:level" json:"level"` // 级别
	Partner int `gorm:"column:partner" json:"partner"` // 是否合作伙伴: 0否 1是
	DockingContent string `gorm:"column:docking_content" json:"dockingContent"` // 对接内容
	WhetherConfidentialityAgre int `gorm:"column:whether_confidentiality_agre" json:"whetherConfidentialityAgre"` // 是否签订保密协议
	Duties string `gorm:"column:duties" json:"duties"` // 职务
	Mailbox string `gorm:"column:mailbox" json:"mailbox"` // 邮箱
	CompletionReverseInvest string `gorm:"column:completion_reverse_invest" json:"completionReverseInvest"` // 反投完成情况
	AvailableBalance string `gorm:"column:available_balance" json:"availableBalance"` // 可投余额
	Dockee string `gorm:"column:dockee" json:"dockee"` // 对接方
	EntryPerson int64 `gorm:"column:entry_person" json:"entryPerson"` // 录入人
	FundName string `gorm:"column:fund_name" json:"fundName"` // 基金名称
	Contact string `gorm:"column:contact" json:"contact"` // 联系人
	CounterinvestmentRequirements string `gorm:"column:counterinvestment_requirements" json:"counterinvestmentRequirements"` // 反投要求
	WarehousingTime *time.Time `gorm:"column:warehousing_time" json:"warehousingTime"` // 入库时间
	ContactInformation string `gorm:"column:contact_information" json:"contactInformation"` // 联系方式
	InvestmentPreferences string `gorm:"column:investment_preferences" json:"investmentPreferences"` // 投资偏好
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64 `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"` // 修改时间
	UpdatedBy int64 `gorm:"column:updated_by" json:"updatedBy"` // 修改人
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// InvestList res_invest 列表
type InvestList []Invest

// NewInvest 创建res_invest实例
func NewInvest() *Invest {
	return &Invest{}
}

// NewInvestList 创建res_invest列表实例
func NewInvestList() *InvestList {
	return &InvestList{}
}

// TableName 指定表名
func (Invest) TableName() string {
	return "res_invest"
}

// GetByID 根据ID获取res_invest
func (m *Invest) GetByID(c context.Context, id int64) error {
	return app.DB().WithContext(c).First(m, id).Error
}

// Create 创建res_invest记录
func (m *Invest) Create(c context.Context) error {
	return app.DB().WithContext(c).Create(m).Error
}

// Update 更新res_invest记录
func (m *Invest) Update(c context.Context) error {
	return app.DB().WithContext(c).Save(m).Error
}

// Delete 软删除res_invest记录
func (m *Invest) Delete(c context.Context) error {
	return app.DB().WithContext(c).Delete(m).Error
}

// IsEmpty 检查模型是否为空
func (m *Invest) IsEmpty() bool {
	return m == nil || m.Id == 0
}

// Find 查询res_invest列表
func (l *InvestList) Find(c context.Context, funcs ...func(*gorm.DB) *gorm.DB) error {
	return app.DB().WithContext(c).Model(&Invest{}).Scopes(funcs...).Find(l).Error
}

// GetTotal 获取res_invest总数
func (l *InvestList) GetTotal(c context.Context, query ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := app.DB().WithContext(c).Model(&Invest{}).Scopes(query...).Count(&count).Error
	return count, err
}