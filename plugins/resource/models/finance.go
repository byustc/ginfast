package models

import (
	"context"
	"time"
	"gin-fast/app/global/app"
	"gorm.io/gorm"
)

// Finance res_finance 模型结构体
type Finance struct {
	Id int64 `gorm:"column:id;primaryKey;not null;autoIncrement" json:"id"` // 主键
	TenantId uint `gorm:"column:tenant_id;default:0;index" json:"tenantId"` // 租户ID字段
	FinanceNo string `gorm:"column:finance_no;index" json:"financeNo"` // 金融编码
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
	FinType string `gorm:"column:fin_type" json:"finType"` // 金融机构类型
	FinNature string `gorm:"column:fin_nature" json:"finNature"` // 机构性质
	IsConfidential string `gorm:"column:is_confidential" json:"isConfidential"` // 是否签订保密协议
	CooperationTime *time.Time `gorm:"column:cooperation_time" json:"cooperationTime"` // 合作建立时间
	Cooperativeness string `gorm:"column:cooperativeness" json:"cooperativeness"` // 合作程度
	CooperationScope string `gorm:"column:cooperation_scope" json:"cooperationScope"` // 合作范围
	Remark string `gorm:"column:remark" json:"remark"` // 备注
	InvestRoundName string `gorm:"column:invest_round_name" json:"investRoundName"` // 当前轮次名称
	OperatingStatusName string `gorm:"column:operating_status_name" json:"operatingStatusName"` // 运营状态名称
	RegCapitalCurrencyName string `gorm:"column:reg_capital_currency_name" json:"regCapitalCurrencyName"` // 注册资本币种名称
	PaidCapitalCurrencyName string `gorm:"column:paid_capital_currency_name" json:"paidCapitalCurrencyName"` // 实缴资本币种名称
	RegTypeName string `gorm:"column:reg_type_name" json:"regTypeName"` // 企业注册类型名称
	DataSource string `gorm:"column:data_source" json:"dataSource"` // 数据来源
	IsBlacklist int `gorm:"column:is_blacklist;default:0" json:"isBlacklist"` // 是否负面清单: 0否 1是
	PositionNum float64 `gorm:"column:position_num" json:"positionNum"` // 挂职人数
	PositionAdvancedNum float64 `gorm:"column:position_advanced_num" json:"positionAdvancedNum"` // 高级挂职人数
	Tag string `gorm:"column:tag" json:"tag"` // 行业数据标签
	DockingPerson string `gorm:"column:docking_person" json:"dockingPerson"` // 对接人
	CorrespondingSourceId float64 `gorm:"column:corresponding_source_id" json:"correspondingSourceId"` // 对应数据源id
	DockingPersonnel string `gorm:"column:docking_personnel" json:"dockingPersonnel"` // 对接人员
	Level string `gorm:"column:level" json:"level"` // 级别
	Partner int `gorm:"column:partner" json:"partner"` // 是否合作伙伴: 0否 1是
	WarehousingTime *time.Time `gorm:"column:warehousing_time" json:"warehousingTime"` // 入库时间
	Contact string `gorm:"column:contact" json:"contact"` // 联系人
	ContactNumber string `gorm:"column:contact_number" json:"contactNumber"` // 联系电话
	EntryPerson int64 `gorm:"column:entry_person" json:"entryPerson"` // 录入人
	ContactInformation string `gorm:"column:contact_information" json:"contactInformation"` // 联系方式
	Mailbox string `gorm:"column:mailbox" json:"mailbox"` // 邮箱
	Duties string `gorm:"column:duties" json:"duties"` // 职务
	ActivitiesUndertaken string `gorm:"column:activities_undertaken" json:"activitiesUndertaken"` // 拟开展活动
	LogoPicture string `gorm:"column:logo_picture" json:"logoPicture"` // logo图片
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64 `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"` // 修改时间
	UpdatedBy int64 `gorm:"column:updated_by" json:"updatedBy"` // 修改人
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// FinanceList res_finance 列表
type FinanceList []Finance

// NewFinance 创建res_finance实例
func NewFinance() *Finance {
	return &Finance{}
}

// NewFinanceList 创建res_finance列表实例
func NewFinanceList() *FinanceList {
	return &FinanceList{}
}

// TableName 指定表名
func (Finance) TableName() string {
	return "res_finance"
}

// GetByID 根据ID获取res_finance
func (m *Finance) GetByID(c context.Context, id int64) error {
	return app.DB().WithContext(c).First(m, id).Error
}

// Create 创建res_finance记录
func (m *Finance) Create(c context.Context) error {
	return app.DB().WithContext(c).Create(m).Error
}

// Update 更新res_finance记录
func (m *Finance) Update(c context.Context) error {
	return app.DB().WithContext(c).Save(m).Error
}

// Delete 软删除res_finance记录
func (m *Finance) Delete(c context.Context) error {
	return app.DB().WithContext(c).Delete(m).Error
}

// IsEmpty 检查模型是否为空
func (m *Finance) IsEmpty() bool {
	return m == nil || m.Id == 0
}

// Find 查询res_finance列表
func (l *FinanceList) Find(c context.Context, funcs ...func(*gorm.DB) *gorm.DB) error {
	return app.DB().WithContext(c).Model(&Finance{}).Scopes(funcs...).Find(l).Error
}

// GetTotal 获取res_finance总数
func (l *FinanceList) GetTotal(c context.Context, query ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := app.DB().WithContext(c).Model(&Finance{}).Scopes(query...).Count(&count).Error
	return count, err
}