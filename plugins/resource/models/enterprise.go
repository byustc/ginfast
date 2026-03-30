package models

import (
	"context"
	"time"
	"gin-fast/app/global/app"
	"gorm.io/gorm"
)

// Enterprise res_enterprise 模型结构体
type Enterprise struct {
	Id int64 `gorm:"column:id;primaryKey;not null;autoIncrement" json:"id"` // 主键
	TenantId uint `gorm:"column:tenant_id;default:0;index" json:"tenantId"` // 租户ID字段
	EnterpriseNo string `gorm:"column:enterprise_no;index" json:"enterpriseNo"` // 企业编码
	BriefName string `gorm:"column:brief_name" json:"briefName"` // 企业简称
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
	LatestInvestmentId string `gorm:"column:latest_investment_id" json:"latestInvestmentId"` // 最新融资id
	LatestInvestTime *time.Time `gorm:"column:latest_invest_time" json:"latestInvestTime"` // 最新融资时间
	LatestExposureTime *time.Time `gorm:"column:latest_exposure_time" json:"latestExposureTime"` // 最新融资曝光时间
	LatestInvestRound float64 `gorm:"column:latest_invest_round" json:"latestInvestRound"` // 最新融资轮次
	LatestInvestAmountType float64 `gorm:"column:latest_invest_amount_type" json:"latestInvestAmountType"` // 最新融资金额类型
	LatestInvestAmount string `gorm:"column:latest_invest_amount" json:"latestInvestAmount"` // 最新融资金额
	UnifiedInvestAmount float64 `gorm:"column:unified_invest_amount" json:"unifiedInvestAmount"` // 最新融资金额转化（单位：万）
	LatestInvestCurrency float64 `gorm:"column:latest_invest_currency" json:"latestInvestCurrency"` // 最新融资金融币种
	LatestInvestor string `gorm:"column:latest_investor" json:"latestInvestor"` // 最新融资投资方
	InvestmentNum float64 `gorm:"column:investment_num" json:"investmentNum"` // 历史融资次数
	TotalInvestors string `gorm:"column:total_investors" json:"totalInvestors"` // 历史投资方
	TotalInvestAmount float64 `gorm:"column:total_invest_amount" json:"totalInvestAmount"` // 历史融资总额（统一化：万元）
	ExternalInvestmentCount float64 `gorm:"column:external_investment_count" json:"externalInvestmentCount"` // 对外投资总数
	CompanyContact string `gorm:"column:company_contact" json:"companyContact"` // 联系信息
	OrganizationalStructure string `gorm:"column:organizational_structure" json:"organizationalStructure"` // 企业组织架构
	SecuritiesInfo string `gorm:"column:securities_info" json:"securitiesInfo"` // 证券信息
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
	InvestRoundName string `gorm:"column:invest_round_name" json:"investRoundName"` // 当前轮次名称
	OperatingStatusName string `gorm:"column:operating_status_name" json:"operatingStatusName"` // 运营状态名称
	LatestInvestRoundName string `gorm:"column:latest_invest_round_name" json:"latestInvestRoundName"` // 最新融资轮次名称
	RegTypeName string `gorm:"column:reg_type_name" json:"regTypeName"` // 企业注册类型名称
	RegCapitalCurrencyName string `gorm:"column:reg_capital_currency_name" json:"regCapitalCurrencyName"` // 注册资本币种名称
	PaidCapitalCurrencyName string `gorm:"column:paid_capital_currency_name" json:"paidCapitalCurrencyName"` // 实缴资本币种名称
	LatestInvestAmountTypeName string `gorm:"column:latest_invest_amount_type_name" json:"latestInvestAmountTypeName"` // 最新融资金额类型名称
	LatestInvestCurrencyName string `gorm:"column:latest_invest_currency_name" json:"latestInvestCurrencyName"` // 最新融资金额币种名称
	Remark string `gorm:"column:remark" json:"remark"` // 备注
	IndustryStatus string `gorm:"column:industry_status" json:"industryStatus"` // 行业地位
	PecuniaryCondition string `gorm:"column:pecuniary_condition" json:"pecuniaryCondition"` // 财务情况
	TechAdvanced string `gorm:"column:tech_advanced" json:"techAdvanced"` // 技术先进性
	DataSource string `gorm:"column:data_source" json:"dataSource"` // 数据来源
	IsBlacklist int `gorm:"column:is_blacklist;default:0" json:"isBlacklist"` // 是否负面清单: 0否 1是
	Tag string `gorm:"column:tag" json:"tag"` // 行业数据标签
	LogoPicture string `gorm:"column:logo_picture" json:"logoPicture"` // logo图片
	OrganizationalStructureNew string `gorm:"column:organizational_structure_new" json:"organizationalStructureNew"` // 组织架构（新）
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64 `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"` // 修改时间
	UpdatedBy int64 `gorm:"column:updated_by" json:"updatedBy"` // 修改人
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// EnterpriseList res_enterprise 列表
type EnterpriseList []Enterprise

// NewEnterprise 创建res_enterprise实例
func NewEnterprise() *Enterprise {
	return &Enterprise{}
}

// NewEnterpriseList 创建res_enterprise列表实例
func NewEnterpriseList() *EnterpriseList {
	return &EnterpriseList{}
}

// TableName 指定表名
func (Enterprise) TableName() string {
	return "res_enterprise"
}

// GetByID 根据ID获取res_enterprise
func (m *Enterprise) GetByID(c context.Context, id int64) error {
	return app.DB().WithContext(c).First(m, id).Error
}

// Create 创建res_enterprise记录
func (m *Enterprise) Create(c context.Context) error {
	return app.DB().WithContext(c).Create(m).Error
}

// Update 更新res_enterprise记录
func (m *Enterprise) Update(c context.Context) error {
	return app.DB().WithContext(c).Save(m).Error
}

// Delete 软删除res_enterprise记录
func (m *Enterprise) Delete(c context.Context) error {
	return app.DB().WithContext(c).Delete(m).Error
}

// IsEmpty 检查模型是否为空
func (m *Enterprise) IsEmpty() bool {
	return m == nil || m.Id == 0
}

// Find 查询res_enterprise列表
func (l *EnterpriseList) Find(c context.Context, funcs ...func(*gorm.DB) *gorm.DB) error {
	return app.DB().WithContext(c).Model(&Enterprise{}).Scopes(funcs...).Find(l).Error
}

// GetTotal 获取res_enterprise总数
func (l *EnterpriseList) GetTotal(c context.Context, query ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := app.DB().WithContext(c).Model(&Enterprise{}).Scopes(query...).Count(&count).Error
	return count, err
}