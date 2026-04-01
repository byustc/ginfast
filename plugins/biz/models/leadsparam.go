package models

import (
	"time"
	"gin-fast/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LeadsListRequest biz_leads列表请求参数
type LeadsListRequest struct {
	models.BasePaging
	models.Validator
	Id *int64 `form:"id"` // 主键
	LeadsNo *string `form:"leadsNo"` // 线索编码
	CompanyName *string `form:"companyName"` // 公司名称
	Region *string `form:"region"` // 所在区域
	Mobile *string `form:"mobile"` // 手机
	Telephone *string `form:"telephone"` // 电话
	Email *string `form:"email"` // 邮箱
	LeadsSource *string `form:"leadsSource"` // 线索来源
	FlowStatus *string `form:"flowStatus"` // 跟进状态
	DistributionStatus *string `form:"distributionStatus"` // 分配状态
	TransformationStatus *string `form:"transformationStatus"` // 转化状态
	LeadsPool *int64 `form:"leadsPool"` // 线索池
	MarketingActivity *int64 `form:"marketingActivity"` // 营销活动
	PromotionChannel *string `form:"promotionChannel"` // 推广渠道
	ReturnReason *string `form:"returnReason"` // 退回/回收原因
	ReturnReasonDescription *string `form:"returnReasonDescription"` // 退回/回收原因说明
	ReturnTime *time.Time `form:"returnTime"` // 退回/回收时间
	DistributionTime *time.Time `form:"distributionTime"` // 分配/领取时间
	LeadsDescription *string `form:"leadsDescription"` // 线索说明
	CustomerId *int64 `form:"customerId"` // 客户id
	ContactId *int64 `form:"contactId"` // 联系人id
	BusinessId *int64 `form:"businessId"` // 业务id
	TransformationTime *time.Time `form:"transformationTime"` // 转化时间
	Sex *string `form:"sex"` // 性别
	Department *string `form:"department"` // 部门
	JobTitle *string `form:"jobTitle"` // 职位
	Age *int `form:"age"` // 年龄
	Birthday *time.Time `form:"birthday"` // 生日
	Wechat *string `form:"wechat"` // 微信
	Qq *string `form:"qq"` // QQ
	Fax *string `form:"fax"` // 传真
	CompanyAddress *string `form:"companyAddress"` // 公司地址
	CompanyPhone *string `form:"companyPhone"` // 公司电话
	Industry *string `form:"industry"` // 所属行业
	StaffSize *string `form:"staffSize"` // 人员规模
	CompanyEmail *string `form:"companyEmail"` // 公司邮箱
	CompanyFax *string `form:"companyFax"` // 公司传真
	LastFlowTime *time.Time `form:"lastFlowTime"` // 最近跟进时间
	HandoverDescription *string `form:"handoverDescription"` // 移交说明
	DistributionDescription *string `form:"distributionDescription"` // 分配说明
	GetTime *time.Time `form:"getTime"` // 获取时间
	CountryRegion *string `form:"countryRegion"` // 国家地区
	ReturnOriginPool *int `form:"returnOriginPool"` // 退回原池
	BusinessLicense *string `form:"businessLicense"` // 营业执照
	GmName *string `form:"gmName"` // 总经理姓名
	SalesExperience *int `form:"salesExperience"` // 是否有销售经验
	SalesDescription *string `form:"salesDescription"` // 销售经验简述
	AnnualSales *string `form:"annualSales"` // 年销售额
	SameIndustry *int `form:"sameIndustry"` // 是否是同一领域的其他公司/品牌的合作伙伴
	CompanyBrand *string `form:"companyBrand"` // 具体公司和品牌
	OtherInformation *string `form:"otherInformation"` // 其他信息
	CompanyWeb *string `form:"companyWeb"` // 公司网址
	LicenseValid *time.Time `form:"licenseValid"` // 执照有效期
	SalesYears *float64 `form:"salesYears"` // 累计销售年数
	SalesDeptUserNum *float64 `form:"salesDeptUserNum"` // 销售部门人数
	MarketingDeptUserNum *float64 `form:"marketingDeptUserNum"` // 营销部门人数
	ServiceDeptUserNum *float64 `form:"serviceDeptUserNum"` // 服务部门人数
	OtherDeptUserNum *float64 `form:"otherDeptUserNum"` // 其他部门人数
	Role *string `form:"role"` // 角色
	EnterpriseType *string `form:"enterpriseType"` // 企业类型
	StartDate *time.Time `form:"startDate"` // 活动开始时间
	EndDate *time.Time `form:"endDate"` // 活动结束时间
	ActivityCity *string `form:"activityCity"` // 活动城市
	Sponsor *string `form:"sponsor"` // 主办方
	ActivityName *string `form:"activityName"` // 活动名称
	OwnBu *string `form:"ownBu"` // 客户于合作方向所属BU
	CooperationDirection *string `form:"cooperationDirection"` // 合作方向
	Solution *string `form:"solution"` // 额外的产品/解决方案需求
	ContactStatus *string `form:"contactStatus"` // 联系状态
	SignedProject *string `form:"signedProject"` // 签约项目
	EstimatedGrossProfit *float64 `form:"estimatedGrossProfit"` // 预计毛利
	BaseCurrency *string `form:"baseCurrency"` // 基准货币
	ExchangeRate *string `form:"exchangeRate"` // 汇率
	ExpectedSigningTime *time.Time `form:"expectedSigningTime"` // 预计签约时间
	FollowUpPerson *int64 `form:"followUpPerson"` // 跟进人
	FeedbackResults *string `form:"feedbackResults"` // 反馈结果
	NewOldCustomer *string `form:"newOldCustomer"` // 新客/老客
	CustomerPartner *string `form:"customerPartner"` // 客户/合作伙伴
	IndustryQs *string `form:"industryQs"` // 所属行业
	Stage *string `form:"stage"` // 阶段
	OriginDataOwnerId *int64 `form:"originDataOwnerId"` // 原归属人
	FirstGetFromPool *int `form:"firstGetFromPool"` // 是否为从池内的第一次获取
	AuditStartUser *int64 `form:"auditStartUser"` // 审批发起人
	AuditCode *string `form:"auditCode"` // 审批编号
	AuditStatus *string `form:"auditStatus"` // 审批状态
	AuditStartTime *time.Time `form:"auditStartTime"` // 审批发起时间
	AuditEndTime *time.Time `form:"auditEndTime"` // 审批结束时间
	ActivityCityText *string `form:"activityCityText"` // 活动城市（文本）
	HistoryDataOwnerId *string `form:"historyDataOwnerId"` // 历史归属人信息
	LeaderAssigned *string `form:"leaderAssigned"` // 市领导交办
	BureauLeaderAssigned *string `form:"bureauLeaderAssigned"` // 局领导交办
	OriginNotes *string `form:"originNotes"` // 来源备注
	District *string `form:"district"` // 区县
	PlatformCompany *string `form:"platformCompany"` // 平台公司
	FromCityPlan *int `form:"fromCityPlan"` // 是否来自创投城市计划生态合伙人
	ServiceAdvisor *int64 `form:"serviceAdvisor"` // 市级服务专员
	ServiceAdvisorContactWay *string `form:"serviceAdvisorContactWay"` // 市级服务专员联系方式
	IndustryAffiliation *string `form:"industryAffiliation"` // 所属产业
	InvestorOriginPlace *string `form:"investorOriginPlace"` // 投资方来源地
	ForeignInvestment *int `form:"foreignInvestment"` // 是否外资
	BpAttachment *string `form:"bpAttachment"` // BP附件
	ProposedProjectName *string `form:"proposedProjectName"` // 拟投项目名称
	RegisterMoney *float64 `form:"registerMoney"` // 注册资金（万元）
	TotalInvestmentAmount *float64 `form:"totalInvestmentAmount"` // 总投资额（亿元）
	TaxEstimation *float64 `form:"taxEstimation"` // 纳税预估（万元）
	RecruitmentPlanContent *string `form:"recruitmentPlanContent"` // 招引谋划内容
	ShareParallelClasses *int `form:"shareParallelClasses"` // 是否共享平行专班
	ResidencePlace *string `form:"residencePlace"` // 拟落户地
	LeadsRecommender *string `form:"leadsRecommender"` // 线索推荐方
	LeadsProvideOrg *int64 `form:"leadsProvideOrg"` // 线索提供部门/机构
	LeadsOwner *int64 `form:"leadsOwner"` // 线索归属人员
	InputTime *time.Time `form:"inputTime"` // 线索录入时间
	InvestProjectName *string `form:"investProjectName"` // 拟投资项目名称
	TechAdvanced *string `form:"techAdvanced"` // 技术先进性
	LandingContent *string `form:"landingContent"` // 落地内容
	ToolCarrier *string `form:"toolCarrier"` // 载体
	ProductionFactors *string `form:"productionFactors"` // 生产要素
	Financing *string `form:"financing"` // 融资
	Market *string `form:"market"` // 场景/市场
	Policy *string `form:"policy"` // 政策
	IsShareCyzb *string `form:"isShareCyzb"` // 是否共享产业专班
	IsDistributeTask *int `form:"isDistributeTask"` // 同时下发行业研判任务
	IsShareLeads *int `form:"isShareLeads"` // 同时与行业牵头人共享线索
	InvestCompany *int64 `form:"investCompany"` // 投资方名称
	PhaseStatus *string `form:"phaseStatus"` // 阶段状态
	Importance *string `form:"importance"` // 重要性
	EnterpriseDemand *string `form:"enterpriseDemand"` // 企业诉求
	LeadDepartmentForClues *int64 `form:"leadDepartmentForClues"` // 线索牵头部门
	AnalysisContent *string `form:"analysisContent"` // 研判内容
	JudgmentOpinion *string `form:"judgmentOpinion"` // 研判意见
	ReasonTermination *string `form:"reasonTermination"` // 线索终止原因
	Destermiantion *string `form:"destermiantion"` // 终止说明
	TechnicalAdvantages *string `form:"technicalAdvantages"` // 技术优势
	BpAttachments *string `form:"bpAttachments"` // BP附件库
	DockingType *string `form:"dockingType"` // 对接形式
	DockingInfo *string `form:"dockingInfo"` // 对接情况
	IndustryStatus *string `form:"industryStatus"` // 行业地位
	PecuniaryCondition *string `form:"pecuniaryCondition"` // 财务情况
	CoreAdvantages *string `form:"coreAdvantages"` // 核心优势
	LeadsOrigin *string `form:"leadsOrigin"` // 线索来源
	InvestRound *string `form:"investRound"` // 拟投资轮次
	NeedIndustry *int `form:"needIndustry"` // 是否需要产业投资人
	NeedFinance *int `form:"needFinance"` // 是否需要财务投资人
	IndustryQsLabel *string `form:"industryQsLabel"` // 所属行业-中文名
	ImportanceNew *string `form:"importanceNew"` // 重要性
	BureauLeadsId *string `form:"bureauLeadsId"` // 市局系统线索主键id
	InvestorOriginPlaceCode *string `form:"investorOriginPlaceCode"` // 投资方来源地(省市区)
	Tag *string `form:"tag"` // 数据标签
	MaximumFactoryRent *float64 `form:"maximumFactoryRent"` // 厂房租金最大值
	LeasedArea *float64 `form:"leasedArea"` // 厂房租用面积
	Period *string `form:"period"` // 期限
	FactoryRentType *string `form:"factoryRentType"` // 厂房租金类型
	FloorHeightStarts *float64 `form:"floorHeightStarts"` // 厂房层高开始
	ParkLevel *string `form:"parkLevel"` // 园区等级
	LandAreaRequired *float64 `form:"landAreaRequired"` // 需地块面积
	MinimumRent *float64 `form:"minimumRent"` // 楼宇租金最小值
	WhetherThereIsANeedForDe *int `form:"whetherThereIsANeedForDe"` // 是否有债权融资需求
	MinimumFactoryRent *float64 `form:"minimumFactoryRent"` // 厂房租金最小值
	EndOfFloorHeight *float64 `form:"endOfFloorHeight"` // 厂房层高结束
	InvestmentArea *float64 `form:"investmentArea"` // 楼宇招商面积
	ParkType *string `form:"parkType"` // 园区类型
	MaximumRent *float64 `form:"maximumRent"` // 楼宇租金最大值
	InterestRate *float64 `form:"interestRate"` // 利率
	LandUse *string `form:"landUse"` // 土地用途
	Quota *string `form:"quota"` // 额度
	GuaranteeMethods *string `form:"guaranteeMethods"` // 担保方式
	RentType *string `form:"rentType"` // 楼宇租金类型
	FinancingAmountMillionY *float64 `form:"financingAmountMillionY"` // 融资金额(亿元)
	ExternalClueId *string `form:"externalClueId"` // 外部推荐线索id
	ClueResearcher *int64 `form:"clueResearcher"` // 线索研判人
	AssessmentTeam *string `form:"assessmentTeam"` // 研判团队
	EnterpriseintermediaryContact *int64 `form:"enterpriseintermediaryContact"` // 企业/中介机构联系人
	ProportionOfEquitySold *float64 `form:"proportionOfEquitySold"` // 出售股权比例
	BureauClueProgress *string `form:"bureauClueProgress"` // 市局线索进度
	CurrentTransferMarketValue *float64 `form:"currentTransferMarketValue"` // 当前转让市值
	OperatingIncomeLastYear *float64 `form:"operatingIncomeLastYear"` // 去年营业收入
	DetailedDescription *string `form:"detailedDescription"` // 详细描述
	TimeRequirementsBefore *time.Time `form:"timeRequirementsBefore"` // 时间要求（之前）
	WhetherToGambleOnPerforman *string `form:"whetherToGambleOnPerforman"` // 是否业绩对赌
	WhetherItIsAListedCompany *string `form:"whetherItIsAListedCompany"` // 是否上市公司
	CityCode *string `form:"cityCode"` // 所属城市(省市区)
	FinancialSituationInThePas *string `form:"financialSituationInThePas"` // 近三年财务情况
	ActualControllerSituation *string `form:"actualControllerSituation"` // 实际控制人情况
	RequirementsForAcquirers *string `form:"requirementsForAcquirers"` // 对收购方要求
	TransferConsiderationTransac *float64 `form:"transferConsiderationTransac"` // 转让对价（交易金额）
	Seller *string `form:"seller"` // 出售方
	NeedFurtherUnderstandingOf *string `form:"needFurtherUnderstandingOf"` // 需要进一步了解问题
	ProjectMatchmaker *int64 `form:"projectMatchmaker"` // 项目对接人
	Country *string `form:"country"` // 国家
	ListingPlate *string `form:"listingPlate"` // 上市板块
	LatestValuation *float64 `form:"latestValuation"` // 最新估值
	ReasonForSelling *string `form:"reasonForSelling"` // 出售原因
	NetProfitLastYear *float64 `form:"netProfitLastYear"` // 去年净利润
	CanWePromiseToImplementPr *string `form:"canWePromiseToImplementPr"` // 能否承诺落地产能或总部搬迁
	PerformanceCommitment *string `form:"performanceCommitment"` // 业绩承诺
	MainBusinessMainProductsAn *string `form:"mainBusinessMainProductsAn"` // 主营业务、主要产品及客户情况
	ProjectIntroduction *string `form:"projectIntroduction"` // 项目介绍
	CityDockUserPhone *string `form:"cityDockUserPhone"` // 市级服务专员联系方式
	ProjectClueContactPerson *string `form:"projectClueContactPerson"` // 项目线索联系人
	CityDockUserName *string `form:"cityDockUserName"` // 市级服务专员姓名
	UpdatedBy *int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *LeadsListRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// Handle 获取查询条件
func (r *LeadsListRequest) Handle() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
        if r.Id != nil {
            // 默认等于查询
            db = db.Where("id = ?", *r.Id)
        }
        if r.LeadsNo != nil {
            // 默认等于查询
            db = db.Where("leads_no = ?", *r.LeadsNo)
        }
        if r.CompanyName != nil {
            // 默认等于查询
            db = db.Where("company_name = ?", *r.CompanyName)
        }
        if r.Region != nil {
            // 默认等于查询
            db = db.Where("region = ?", *r.Region)
        }
        if r.Mobile != nil {
            // 默认等于查询
            db = db.Where("mobile = ?", *r.Mobile)
        }
        if r.Telephone != nil {
            // 默认等于查询
            db = db.Where("telephone = ?", *r.Telephone)
        }
        if r.Email != nil {
            // 默认等于查询
            db = db.Where("email = ?", *r.Email)
        }
        if r.LeadsSource != nil {
            // 默认等于查询
            db = db.Where("leads_source = ?", *r.LeadsSource)
        }
        if r.FlowStatus != nil {
            // 默认等于查询
            db = db.Where("flow_status = ?", *r.FlowStatus)
        }
        if r.DistributionStatus != nil {
            // 默认等于查询
            db = db.Where("distribution_status = ?", *r.DistributionStatus)
        }
        if r.TransformationStatus != nil {
            // 默认等于查询
            db = db.Where("transformation_status = ?", *r.TransformationStatus)
        }
        if r.LeadsPool != nil {
            // 默认等于查询
            db = db.Where("leads_pool = ?", *r.LeadsPool)
        }
        if r.MarketingActivity != nil {
            // 默认等于查询
            db = db.Where("marketing_activity = ?", *r.MarketingActivity)
        }
        if r.PromotionChannel != nil {
            // 默认等于查询
            db = db.Where("promotion_channel = ?", *r.PromotionChannel)
        }
        if r.ReturnReason != nil {
            // 默认等于查询
            db = db.Where("return_reason = ?", *r.ReturnReason)
        }
        if r.ReturnReasonDescription != nil {
            // 默认等于查询
            db = db.Where("return_reason_description = ?", *r.ReturnReasonDescription)
        }
        if r.ReturnTime != nil {
            // 默认等于查询
            db = db.Where("return_time = ?", *r.ReturnTime)
        }
        if r.DistributionTime != nil {
            // 默认等于查询
            db = db.Where("distribution_time = ?", *r.DistributionTime)
        }
        if r.LeadsDescription != nil {
            // 默认等于查询
            db = db.Where("leads_description = ?", *r.LeadsDescription)
        }
        if r.CustomerId != nil {
            // 默认等于查询
            db = db.Where("customer_id = ?", *r.CustomerId)
        }
        if r.ContactId != nil {
            // 默认等于查询
            db = db.Where("contact_id = ?", *r.ContactId)
        }
        if r.BusinessId != nil {
            // 默认等于查询
            db = db.Where("business_id = ?", *r.BusinessId)
        }
        if r.TransformationTime != nil {
            // 默认等于查询
            db = db.Where("transformation_time = ?", *r.TransformationTime)
        }
        if r.Sex != nil {
            // 默认等于查询
            db = db.Where("sex = ?", *r.Sex)
        }
        if r.Department != nil {
            // 默认等于查询
            db = db.Where("department = ?", *r.Department)
        }
        if r.JobTitle != nil {
            // 默认等于查询
            db = db.Where("job_title = ?", *r.JobTitle)
        }
        if r.Age != nil {
            // 默认等于查询
            db = db.Where("age = ?", *r.Age)
        }
        if r.Birthday != nil {
            // 默认等于查询
            db = db.Where("birthday = ?", *r.Birthday)
        }
        if r.Wechat != nil {
            // 默认等于查询
            db = db.Where("wechat = ?", *r.Wechat)
        }
        if r.Qq != nil {
            // 默认等于查询
            db = db.Where("qq = ?", *r.Qq)
        }
        if r.Fax != nil {
            // 默认等于查询
            db = db.Where("fax = ?", *r.Fax)
        }
        if r.CompanyAddress != nil {
            // 默认等于查询
            db = db.Where("company_address = ?", *r.CompanyAddress)
        }
        if r.CompanyPhone != nil {
            // 默认等于查询
            db = db.Where("company_phone = ?", *r.CompanyPhone)
        }
        if r.Industry != nil {
            // 默认等于查询
            db = db.Where("industry = ?", *r.Industry)
        }
        if r.StaffSize != nil {
            // 默认等于查询
            db = db.Where("staff_size = ?", *r.StaffSize)
        }
        if r.CompanyEmail != nil {
            // 默认等于查询
            db = db.Where("company_email = ?", *r.CompanyEmail)
        }
        if r.CompanyFax != nil {
            // 默认等于查询
            db = db.Where("company_fax = ?", *r.CompanyFax)
        }
        if r.LastFlowTime != nil {
            // 默认等于查询
            db = db.Where("last_flow_time = ?", *r.LastFlowTime)
        }
        if r.HandoverDescription != nil {
            // 默认等于查询
            db = db.Where("handover_description = ?", *r.HandoverDescription)
        }
        if r.DistributionDescription != nil {
            // 默认等于查询
            db = db.Where("distribution_description = ?", *r.DistributionDescription)
        }
        if r.GetTime != nil {
            // 默认等于查询
            db = db.Where("get_time = ?", *r.GetTime)
        }
        if r.CountryRegion != nil {
            // 默认等于查询
            db = db.Where("country_region = ?", *r.CountryRegion)
        }
        if r.ReturnOriginPool != nil {
            // 默认等于查询
            db = db.Where("return_origin_pool = ?", *r.ReturnOriginPool)
        }
        if r.BusinessLicense != nil {
            // 默认等于查询
            db = db.Where("business_license = ?", *r.BusinessLicense)
        }
        if r.GmName != nil {
            // 默认等于查询
            db = db.Where("gm_name = ?", *r.GmName)
        }
        if r.SalesExperience != nil {
            // 默认等于查询
            db = db.Where("sales_experience = ?", *r.SalesExperience)
        }
        if r.SalesDescription != nil {
            // 默认等于查询
            db = db.Where("sales_description = ?", *r.SalesDescription)
        }
        if r.AnnualSales != nil {
            // 默认等于查询
            db = db.Where("annual_sales = ?", *r.AnnualSales)
        }
        if r.SameIndustry != nil {
            // 默认等于查询
            db = db.Where("same_industry = ?", *r.SameIndustry)
        }
        if r.CompanyBrand != nil {
            // 默认等于查询
            db = db.Where("company_brand = ?", *r.CompanyBrand)
        }
        if r.OtherInformation != nil {
            // 默认等于查询
            db = db.Where("other_information = ?", *r.OtherInformation)
        }
        if r.CompanyWeb != nil {
            // 默认等于查询
            db = db.Where("company_web = ?", *r.CompanyWeb)
        }
        if r.LicenseValid != nil {
            // 默认等于查询
            db = db.Where("license_valid = ?", *r.LicenseValid)
        }
        if r.SalesYears != nil {
            // 默认等于查询
            db = db.Where("sales_years = ?", *r.SalesYears)
        }
        if r.SalesDeptUserNum != nil {
            // 默认等于查询
            db = db.Where("sales_dept_user_num = ?", *r.SalesDeptUserNum)
        }
        if r.MarketingDeptUserNum != nil {
            // 默认等于查询
            db = db.Where("marketing_dept_user_num = ?", *r.MarketingDeptUserNum)
        }
        if r.ServiceDeptUserNum != nil {
            // 默认等于查询
            db = db.Where("service_dept_user_num = ?", *r.ServiceDeptUserNum)
        }
        if r.OtherDeptUserNum != nil {
            // 默认等于查询
            db = db.Where("other_dept_user_num = ?", *r.OtherDeptUserNum)
        }
        if r.Role != nil {
            // 默认等于查询
            db = db.Where("role = ?", *r.Role)
        }
        if r.EnterpriseType != nil {
            // 默认等于查询
            db = db.Where("enterprise_type = ?", *r.EnterpriseType)
        }
        if r.StartDate != nil {
            // 默认等于查询
            db = db.Where("start_date = ?", *r.StartDate)
        }
        if r.EndDate != nil {
            // 默认等于查询
            db = db.Where("end_date = ?", *r.EndDate)
        }
        if r.ActivityCity != nil {
            // 默认等于查询
            db = db.Where("activity_city = ?", *r.ActivityCity)
        }
        if r.Sponsor != nil {
            // 默认等于查询
            db = db.Where("sponsor = ?", *r.Sponsor)
        }
        if r.ActivityName != nil {
            // 默认等于查询
            db = db.Where("activity_name = ?", *r.ActivityName)
        }
        if r.OwnBu != nil {
            // 默认等于查询
            db = db.Where("own_bu = ?", *r.OwnBu)
        }
        if r.CooperationDirection != nil {
            // 默认等于查询
            db = db.Where("cooperation_direction = ?", *r.CooperationDirection)
        }
        if r.Solution != nil {
            // 默认等于查询
            db = db.Where("solution = ?", *r.Solution)
        }
        if r.ContactStatus != nil {
            // 默认等于查询
            db = db.Where("contact_status = ?", *r.ContactStatus)
        }
        if r.SignedProject != nil {
            // 默认等于查询
            db = db.Where("signed_project = ?", *r.SignedProject)
        }
        if r.EstimatedGrossProfit != nil {
            // 默认等于查询
            db = db.Where("estimated_gross_profit = ?", *r.EstimatedGrossProfit)
        }
        if r.BaseCurrency != nil {
            // 默认等于查询
            db = db.Where("base_currency = ?", *r.BaseCurrency)
        }
        if r.ExchangeRate != nil {
            // 默认等于查询
            db = db.Where("exchange_rate = ?", *r.ExchangeRate)
        }
        if r.ExpectedSigningTime != nil {
            // 默认等于查询
            db = db.Where("expected_signing_time = ?", *r.ExpectedSigningTime)
        }
        if r.FollowUpPerson != nil {
            // 默认等于查询
            db = db.Where("follow_up_person = ?", *r.FollowUpPerson)
        }
        if r.FeedbackResults != nil {
            // 默认等于查询
            db = db.Where("feedback_results = ?", *r.FeedbackResults)
        }
        if r.NewOldCustomer != nil {
            // 默认等于查询
            db = db.Where("new_old_customer = ?", *r.NewOldCustomer)
        }
        if r.CustomerPartner != nil {
            // 默认等于查询
            db = db.Where("customer_partner = ?", *r.CustomerPartner)
        }
        if r.IndustryQs != nil {
            // 默认等于查询
            db = db.Where("industry_qs = ?", *r.IndustryQs)
        }
        if r.Stage != nil {
            // 默认等于查询
            db = db.Where("stage = ?", *r.Stage)
        }
        if r.OriginDataOwnerId != nil {
            // 默认等于查询
            db = db.Where("origin_data_owner_id = ?", *r.OriginDataOwnerId)
        }
        if r.FirstGetFromPool != nil {
            // 默认等于查询
            db = db.Where("first_get_from_pool = ?", *r.FirstGetFromPool)
        }
        if r.AuditStartUser != nil {
            // 默认等于查询
            db = db.Where("audit_start_user = ?", *r.AuditStartUser)
        }
        if r.AuditCode != nil {
            // 默认等于查询
            db = db.Where("audit_code = ?", *r.AuditCode)
        }
        if r.AuditStatus != nil {
            // 默认等于查询
            db = db.Where("audit_status = ?", *r.AuditStatus)
        }
        if r.AuditStartTime != nil {
            // 默认等于查询
            db = db.Where("audit_start_time = ?", *r.AuditStartTime)
        }
        if r.AuditEndTime != nil {
            // 默认等于查询
            db = db.Where("audit_end_time = ?", *r.AuditEndTime)
        }
        if r.ActivityCityText != nil {
            // 默认等于查询
            db = db.Where("activity_city_text = ?", *r.ActivityCityText)
        }
        if r.HistoryDataOwnerId != nil {
            // 默认等于查询
            db = db.Where("history_data_owner_id = ?", *r.HistoryDataOwnerId)
        }
        if r.LeaderAssigned != nil {
            // 默认等于查询
            db = db.Where("leader_assigned = ?", *r.LeaderAssigned)
        }
        if r.BureauLeaderAssigned != nil {
            // 默认等于查询
            db = db.Where("bureau_leader_assigned = ?", *r.BureauLeaderAssigned)
        }
        if r.OriginNotes != nil {
            // 默认等于查询
            db = db.Where("origin_notes = ?", *r.OriginNotes)
        }
        if r.District != nil {
            // 默认等于查询
            db = db.Where("district = ?", *r.District)
        }
        if r.PlatformCompany != nil {
            // 默认等于查询
            db = db.Where("platform_company = ?", *r.PlatformCompany)
        }
        if r.FromCityPlan != nil {
            // 默认等于查询
            db = db.Where("from_city_plan = ?", *r.FromCityPlan)
        }
        if r.ServiceAdvisor != nil {
            // 默认等于查询
            db = db.Where("service_advisor = ?", *r.ServiceAdvisor)
        }
        if r.ServiceAdvisorContactWay != nil {
            // 默认等于查询
            db = db.Where("service_advisor_contact_way = ?", *r.ServiceAdvisorContactWay)
        }
        if r.IndustryAffiliation != nil {
            // 默认等于查询
            db = db.Where("industry_affiliation = ?", *r.IndustryAffiliation)
        }
        if r.InvestorOriginPlace != nil {
            // 默认等于查询
            db = db.Where("investor_origin_place = ?", *r.InvestorOriginPlace)
        }
        if r.ForeignInvestment != nil {
            // 默认等于查询
            db = db.Where("foreign_investment = ?", *r.ForeignInvestment)
        }
        if r.BpAttachment != nil {
            // 默认等于查询
            db = db.Where("bp_attachment = ?", *r.BpAttachment)
        }
        if r.ProposedProjectName != nil {
            // 默认等于查询
            db = db.Where("proposed_project_name = ?", *r.ProposedProjectName)
        }
        if r.RegisterMoney != nil {
            // 默认等于查询
            db = db.Where("register_money = ?", *r.RegisterMoney)
        }
        if r.TotalInvestmentAmount != nil {
            // 默认等于查询
            db = db.Where("total_investment_amount = ?", *r.TotalInvestmentAmount)
        }
        if r.TaxEstimation != nil {
            // 默认等于查询
            db = db.Where("tax_estimation = ?", *r.TaxEstimation)
        }
        if r.RecruitmentPlanContent != nil {
            // 默认等于查询
            db = db.Where("recruitment_plan_content = ?", *r.RecruitmentPlanContent)
        }
        if r.ShareParallelClasses != nil {
            // 默认等于查询
            db = db.Where("share_parallel_classes = ?", *r.ShareParallelClasses)
        }
        if r.ResidencePlace != nil {
            // 默认等于查询
            db = db.Where("residence_place = ?", *r.ResidencePlace)
        }
        if r.LeadsRecommender != nil {
            // 默认等于查询
            db = db.Where("leads_recommender = ?", *r.LeadsRecommender)
        }
        if r.LeadsProvideOrg != nil {
            // 默认等于查询
            db = db.Where("leads_provide_org = ?", *r.LeadsProvideOrg)
        }
        if r.LeadsOwner != nil {
            // 默认等于查询
            db = db.Where("leads_owner = ?", *r.LeadsOwner)
        }
        if r.InputTime != nil {
            // 默认等于查询
            db = db.Where("input_time = ?", *r.InputTime)
        }
        if r.InvestProjectName != nil {
            // 默认等于查询
            db = db.Where("invest_project_name = ?", *r.InvestProjectName)
        }
        if r.TechAdvanced != nil {
            // 默认等于查询
            db = db.Where("tech_advanced = ?", *r.TechAdvanced)
        }
        if r.LandingContent != nil {
            // 默认等于查询
            db = db.Where("landing_content = ?", *r.LandingContent)
        }
        if r.ToolCarrier != nil {
            // 默认等于查询
            db = db.Where("tool_carrier = ?", *r.ToolCarrier)
        }
        if r.ProductionFactors != nil {
            // 默认等于查询
            db = db.Where("production_factors = ?", *r.ProductionFactors)
        }
        if r.Financing != nil {
            // 默认等于查询
            db = db.Where("financing = ?", *r.Financing)
        }
        if r.Market != nil {
            // 默认等于查询
            db = db.Where("market = ?", *r.Market)
        }
        if r.Policy != nil {
            // 默认等于查询
            db = db.Where("policy = ?", *r.Policy)
        }
        if r.IsShareCyzb != nil {
            // 默认等于查询
            db = db.Where("is_share_cyzb = ?", *r.IsShareCyzb)
        }
        if r.IsDistributeTask != nil {
            // 默认等于查询
            db = db.Where("is_distribute_task = ?", *r.IsDistributeTask)
        }
        if r.IsShareLeads != nil {
            // 默认等于查询
            db = db.Where("is_share_leads = ?", *r.IsShareLeads)
        }
        if r.InvestCompany != nil {
            // 默认等于查询
            db = db.Where("invest_company = ?", *r.InvestCompany)
        }
        if r.PhaseStatus != nil {
            // 默认等于查询
            db = db.Where("phase_status = ?", *r.PhaseStatus)
        }
        if r.Importance != nil {
            // 默认等于查询
            db = db.Where("importance = ?", *r.Importance)
        }
        if r.EnterpriseDemand != nil {
            // 默认等于查询
            db = db.Where("enterprise_demand = ?", *r.EnterpriseDemand)
        }
        if r.LeadDepartmentForClues != nil {
            // 默认等于查询
            db = db.Where("lead_department_for_clues = ?", *r.LeadDepartmentForClues)
        }
        if r.AnalysisContent != nil {
            // 默认等于查询
            db = db.Where("analysis_content = ?", *r.AnalysisContent)
        }
        if r.JudgmentOpinion != nil {
            // 默认等于查询
            db = db.Where("judgment_opinion = ?", *r.JudgmentOpinion)
        }
        if r.ReasonTermination != nil {
            // 默认等于查询
            db = db.Where("reason_termination = ?", *r.ReasonTermination)
        }
        if r.Destermiantion != nil {
            // 默认等于查询
            db = db.Where("destermiantion = ?", *r.Destermiantion)
        }
        if r.TechnicalAdvantages != nil {
            // 默认等于查询
            db = db.Where("technical_advantages = ?", *r.TechnicalAdvantages)
        }
        if r.BpAttachments != nil {
            // 默认等于查询
            db = db.Where("bp_attachments = ?", *r.BpAttachments)
        }
        if r.DockingType != nil {
            // 默认等于查询
            db = db.Where("docking_type = ?", *r.DockingType)
        }
        if r.DockingInfo != nil {
            // 默认等于查询
            db = db.Where("docking_info = ?", *r.DockingInfo)
        }
        if r.IndustryStatus != nil {
            // 默认等于查询
            db = db.Where("industry_status = ?", *r.IndustryStatus)
        }
        if r.PecuniaryCondition != nil {
            // 默认等于查询
            db = db.Where("pecuniary_condition = ?", *r.PecuniaryCondition)
        }
        if r.CoreAdvantages != nil {
            // 默认等于查询
            db = db.Where("core_advantages = ?", *r.CoreAdvantages)
        }
        if r.LeadsOrigin != nil {
            // 默认等于查询
            db = db.Where("leads_origin = ?", *r.LeadsOrigin)
        }
        if r.InvestRound != nil {
            // 默认等于查询
            db = db.Where("invest_round = ?", *r.InvestRound)
        }
        if r.NeedIndustry != nil {
            // 默认等于查询
            db = db.Where("need_industry = ?", *r.NeedIndustry)
        }
        if r.NeedFinance != nil {
            // 默认等于查询
            db = db.Where("need_finance = ?", *r.NeedFinance)
        }
        if r.IndustryQsLabel != nil {
            // 默认等于查询
            db = db.Where("industry_qs_label = ?", *r.IndustryQsLabel)
        }
        if r.ImportanceNew != nil {
            // 默认等于查询
            db = db.Where("importance_new = ?", *r.ImportanceNew)
        }
        if r.BureauLeadsId != nil {
            // 默认等于查询
            db = db.Where("bureau_leads_id = ?", *r.BureauLeadsId)
        }
        if r.InvestorOriginPlaceCode != nil {
            // 默认等于查询
            db = db.Where("investor_origin_place_code = ?", *r.InvestorOriginPlaceCode)
        }
        if r.Tag != nil {
            // 默认等于查询
            db = db.Where("tag = ?", *r.Tag)
        }
        if r.MaximumFactoryRent != nil {
            // 默认等于查询
            db = db.Where("maximum_factory_rent = ?", *r.MaximumFactoryRent)
        }
        if r.LeasedArea != nil {
            // 默认等于查询
            db = db.Where("leased_area = ?", *r.LeasedArea)
        }
        if r.Period != nil {
            // 默认等于查询
            db = db.Where("period = ?", *r.Period)
        }
        if r.FactoryRentType != nil {
            // 默认等于查询
            db = db.Where("factory_rent_type = ?", *r.FactoryRentType)
        }
        if r.FloorHeightStarts != nil {
            // 默认等于查询
            db = db.Where("floor_height_starts = ?", *r.FloorHeightStarts)
        }
        if r.ParkLevel != nil {
            // 默认等于查询
            db = db.Where("park_level = ?", *r.ParkLevel)
        }
        if r.LandAreaRequired != nil {
            // 默认等于查询
            db = db.Where("land_area_required = ?", *r.LandAreaRequired)
        }
        if r.MinimumRent != nil {
            // 默认等于查询
            db = db.Where("minimum_rent = ?", *r.MinimumRent)
        }
        if r.WhetherThereIsANeedForDe != nil {
            // 默认等于查询
            db = db.Where("whether_there_is_a_need_for_de = ?", *r.WhetherThereIsANeedForDe)
        }
        if r.MinimumFactoryRent != nil {
            // 默认等于查询
            db = db.Where("minimum_factory_rent = ?", *r.MinimumFactoryRent)
        }
        if r.EndOfFloorHeight != nil {
            // 默认等于查询
            db = db.Where("end_of_floor_height = ?", *r.EndOfFloorHeight)
        }
        if r.InvestmentArea != nil {
            // 默认等于查询
            db = db.Where("investment_area = ?", *r.InvestmentArea)
        }
        if r.ParkType != nil {
            // 默认等于查询
            db = db.Where("park_type = ?", *r.ParkType)
        }
        if r.MaximumRent != nil {
            // 默认等于查询
            db = db.Where("maximum_rent = ?", *r.MaximumRent)
        }
        if r.InterestRate != nil {
            // 默认等于查询
            db = db.Where("interest_rate = ?", *r.InterestRate)
        }
        if r.LandUse != nil {
            // 默认等于查询
            db = db.Where("land_use = ?", *r.LandUse)
        }
        if r.Quota != nil {
            // 默认等于查询
            db = db.Where("quota = ?", *r.Quota)
        }
        if r.GuaranteeMethods != nil {
            // 默认等于查询
            db = db.Where("guarantee_methods = ?", *r.GuaranteeMethods)
        }
        if r.RentType != nil {
            // 默认等于查询
            db = db.Where("rent_type = ?", *r.RentType)
        }
        if r.FinancingAmountMillionY != nil {
            // 默认等于查询
            db = db.Where("financing_amount_100_million_y = ?", *r.FinancingAmountMillionY)
        }
        if r.ExternalClueId != nil {
            // 默认等于查询
            db = db.Where("external_clue_id = ?", *r.ExternalClueId)
        }
        if r.ClueResearcher != nil {
            // 默认等于查询
            db = db.Where("clue_researcher = ?", *r.ClueResearcher)
        }
        if r.AssessmentTeam != nil {
            // 默认等于查询
            db = db.Where("assessment_team = ?", *r.AssessmentTeam)
        }
        if r.EnterpriseintermediaryContact != nil {
            // 默认等于查询
            db = db.Where("enterpriseintermediary_contact = ?", *r.EnterpriseintermediaryContact)
        }
        if r.ProportionOfEquitySold != nil {
            // 默认等于查询
            db = db.Where("proportion_of_equity_sold = ?", *r.ProportionOfEquitySold)
        }
        if r.BureauClueProgress != nil {
            // 默认等于查询
            db = db.Where("bureau_clue_progress = ?", *r.BureauClueProgress)
        }
        if r.CurrentTransferMarketValue != nil {
            // 默认等于查询
            db = db.Where("current_transfer_market_value = ?", *r.CurrentTransferMarketValue)
        }
        if r.OperatingIncomeLastYear != nil {
            // 默认等于查询
            db = db.Where("operating_income_last_year = ?", *r.OperatingIncomeLastYear)
        }
        if r.DetailedDescription != nil {
            // 默认等于查询
            db = db.Where("detailed_description = ?", *r.DetailedDescription)
        }
        if r.TimeRequirementsBefore != nil {
            // 默认等于查询
            db = db.Where("time_requirements_before = ?", *r.TimeRequirementsBefore)
        }
        if r.WhetherToGambleOnPerforman != nil {
            // 默认等于查询
            db = db.Where("whether_to_gamble_on_performan = ?", *r.WhetherToGambleOnPerforman)
        }
        if r.WhetherItIsAListedCompany != nil {
            // 默认等于查询
            db = db.Where("whether_it_is_a_listed_company = ?", *r.WhetherItIsAListedCompany)
        }
        if r.CityCode != nil {
            // 默认等于查询
            db = db.Where("city_code = ?", *r.CityCode)
        }
        if r.FinancialSituationInThePas != nil {
            // 默认等于查询
            db = db.Where("financial_situation_in_the_pas = ?", *r.FinancialSituationInThePas)
        }
        if r.ActualControllerSituation != nil {
            // 默认等于查询
            db = db.Where("actual_controller_situation = ?", *r.ActualControllerSituation)
        }
        if r.RequirementsForAcquirers != nil {
            // 默认等于查询
            db = db.Where("requirements_for_acquirers = ?", *r.RequirementsForAcquirers)
        }
        if r.TransferConsiderationTransac != nil {
            // 默认等于查询
            db = db.Where("transfer_consideration_transac = ?", *r.TransferConsiderationTransac)
        }
        if r.Seller != nil {
            // 默认等于查询
            db = db.Where("seller = ?", *r.Seller)
        }
        if r.NeedFurtherUnderstandingOf != nil {
            // 默认等于查询
            db = db.Where("need_further_understanding_of = ?", *r.NeedFurtherUnderstandingOf)
        }
        if r.ProjectMatchmaker != nil {
            // 默认等于查询
            db = db.Where("project_matchmaker = ?", *r.ProjectMatchmaker)
        }
        if r.Country != nil {
            // 默认等于查询
            db = db.Where("country = ?", *r.Country)
        }
        if r.ListingPlate != nil {
            // 默认等于查询
            db = db.Where("listing_plate = ?", *r.ListingPlate)
        }
        if r.LatestValuation != nil {
            // 默认等于查询
            db = db.Where("latest_valuation = ?", *r.LatestValuation)
        }
        if r.ReasonForSelling != nil {
            // 默认等于查询
            db = db.Where("reason_for_selling = ?", *r.ReasonForSelling)
        }
        if r.NetProfitLastYear != nil {
            // 默认等于查询
            db = db.Where("net_profit_last_year = ?", *r.NetProfitLastYear)
        }
        if r.CanWePromiseToImplementPr != nil {
            // 默认等于查询
            db = db.Where("can_we_promise_to_implement_pr = ?", *r.CanWePromiseToImplementPr)
        }
        if r.PerformanceCommitment != nil {
            // 默认等于查询
            db = db.Where("performance_commitment = ?", *r.PerformanceCommitment)
        }
        if r.MainBusinessMainProductsAn != nil {
            // 默认等于查询
            db = db.Where("main_business_main_products_an = ?", *r.MainBusinessMainProductsAn)
        }
        if r.ProjectIntroduction != nil {
            // 默认等于查询
            db = db.Where("project_introduction = ?", *r.ProjectIntroduction)
        }
        if r.CityDockUserPhone != nil {
            // 默认等于查询
            db = db.Where("city_dock_user_phone = ?", *r.CityDockUserPhone)
        }
        if r.ProjectClueContactPerson != nil {
            // 默认等于查询
            db = db.Where("project_clue_contact_person = ?", *r.ProjectClueContactPerson)
        }
        if r.CityDockUserName != nil {
            // 默认等于查询
            db = db.Where("city_dock_user_name = ?", *r.CityDockUserName)
        }
        if r.UpdatedBy != nil {
            // 默认等于查询
            db = db.Where("updated_by = ?", *r.UpdatedBy)
        }
		return db
	}
}

// LeadsCreateRequest 创建biz_leads请求参数
type LeadsCreateRequest struct {
	models.Validator
	LeadsNo string `form:"leadsNo" validate:"required" message:"线索编码不能为空"` // 线索编码
	CompanyName string `form:"companyName" validate:"required" message:"公司名称不能为空"` // 公司名称
	Region string `form:"region" validate:"required" message:"所在区域不能为空"` // 所在区域
	Mobile string `form:"mobile" validate:"required" message:"手机不能为空"` // 手机
	Telephone string `form:"telephone" validate:"required" message:"电话不能为空"` // 电话
	Email string `form:"email" validate:"required" message:"邮箱不能为空"` // 邮箱
	LeadsSource string `form:"leadsSource" validate:"required" message:"线索来源不能为空"` // 线索来源
	FlowStatus string `form:"flowStatus" validate:"required" message:"跟进状态不能为空"` // 跟进状态
	DistributionStatus string `form:"distributionStatus" validate:"required" message:"分配状态不能为空"` // 分配状态
	TransformationStatus string `form:"transformationStatus" validate:"required" message:"转化状态不能为空"` // 转化状态
	LeadsPool int64 `form:"leadsPool"` // 线索池
	MarketingActivity int64 `form:"marketingActivity"` // 营销活动
	PromotionChannel string `form:"promotionChannel" validate:"required" message:"推广渠道不能为空"` // 推广渠道
	ReturnReason string `form:"returnReason" validate:"required" message:"退回/回收原因不能为空"` // 退回/回收原因
	ReturnReasonDescription string `form:"returnReasonDescription" validate:"required" message:"退回/回收原因说明不能为空"` // 退回/回收原因说明
	ReturnTime *time.Time `form:"returnTime" validate:"required" message:"退回/回收时间不能为空"` // 退回/回收时间
	DistributionTime *time.Time `form:"distributionTime" validate:"required" message:"分配/领取时间不能为空"` // 分配/领取时间
	LeadsDescription string `form:"leadsDescription" validate:"required" message:"线索说明不能为空"` // 线索说明
	CustomerId int64 `form:"customerId"` // 客户id
	ContactId int64 `form:"contactId"` // 联系人id
	BusinessId int64 `form:"businessId"` // 业务id
	TransformationTime *time.Time `form:"transformationTime" validate:"required" message:"转化时间不能为空"` // 转化时间
	Sex string `form:"sex" validate:"required" message:"性别不能为空"` // 性别
	Department string `form:"department" validate:"required" message:"部门不能为空"` // 部门
	JobTitle string `form:"jobTitle" validate:"required" message:"职位不能为空"` // 职位
	Age int `form:"age"` // 年龄
	Birthday *time.Time `form:"birthday" validate:"required" message:"生日不能为空"` // 生日
	Wechat string `form:"wechat" validate:"required" message:"微信不能为空"` // 微信
	Qq string `form:"qq" validate:"required" message:"QQ不能为空"` // QQ
	Fax string `form:"fax" validate:"required" message:"传真不能为空"` // 传真
	CompanyAddress string `form:"companyAddress" validate:"required" message:"公司地址不能为空"` // 公司地址
	CompanyPhone string `form:"companyPhone" validate:"required" message:"公司电话不能为空"` // 公司电话
	Industry string `form:"industry" validate:"required" message:"所属行业不能为空"` // 所属行业
	StaffSize string `form:"staffSize" validate:"required" message:"人员规模不能为空"` // 人员规模
	CompanyEmail string `form:"companyEmail" validate:"required" message:"公司邮箱不能为空"` // 公司邮箱
	CompanyFax string `form:"companyFax" validate:"required" message:"公司传真不能为空"` // 公司传真
	LastFlowTime *time.Time `form:"lastFlowTime" validate:"required" message:"最近跟进时间不能为空"` // 最近跟进时间
	HandoverDescription string `form:"handoverDescription" validate:"required" message:"移交说明不能为空"` // 移交说明
	DistributionDescription string `form:"distributionDescription" validate:"required" message:"分配说明不能为空"` // 分配说明
	GetTime *time.Time `form:"getTime" validate:"required" message:"获取时间不能为空"` // 获取时间
	CountryRegion string `form:"countryRegion" validate:"required" message:"国家地区不能为空"` // 国家地区
	ReturnOriginPool int `form:"returnOriginPool"` // 退回原池
	BusinessLicense string `form:"businessLicense" validate:"required" message:"营业执照不能为空"` // 营业执照
	GmName string `form:"gmName" validate:"required" message:"总经理姓名不能为空"` // 总经理姓名
	SalesExperience int `form:"salesExperience"` // 是否有销售经验
	SalesDescription string `form:"salesDescription" validate:"required" message:"销售经验简述不能为空"` // 销售经验简述
	AnnualSales string `form:"annualSales" validate:"required" message:"年销售额不能为空"` // 年销售额
	SameIndustry int `form:"sameIndustry"` // 是否是同一领域的其他公司/品牌的合作伙伴
	CompanyBrand string `form:"companyBrand" validate:"required" message:"具体公司和品牌不能为空"` // 具体公司和品牌
	OtherInformation string `form:"otherInformation" validate:"required" message:"其他信息不能为空"` // 其他信息
	CompanyWeb string `form:"companyWeb" validate:"required" message:"公司网址不能为空"` // 公司网址
	LicenseValid *time.Time `form:"licenseValid" validate:"required" message:"执照有效期不能为空"` // 执照有效期
	SalesYears float64 `form:"salesYears"` // 累计销售年数
	SalesDeptUserNum float64 `form:"salesDeptUserNum"` // 销售部门人数
	MarketingDeptUserNum float64 `form:"marketingDeptUserNum"` // 营销部门人数
	ServiceDeptUserNum float64 `form:"serviceDeptUserNum"` // 服务部门人数
	OtherDeptUserNum float64 `form:"otherDeptUserNum"` // 其他部门人数
	Role string `form:"role" validate:"required" message:"角色不能为空"` // 角色
	EnterpriseType string `form:"enterpriseType" validate:"required" message:"企业类型不能为空"` // 企业类型
	StartDate *time.Time `form:"startDate" validate:"required" message:"活动开始时间不能为空"` // 活动开始时间
	EndDate *time.Time `form:"endDate" validate:"required" message:"活动结束时间不能为空"` // 活动结束时间
	ActivityCity string `form:"activityCity" validate:"required" message:"活动城市不能为空"` // 活动城市
	Sponsor string `form:"sponsor" validate:"required" message:"主办方不能为空"` // 主办方
	ActivityName string `form:"activityName" validate:"required" message:"活动名称不能为空"` // 活动名称
	OwnBu string `form:"ownBu" validate:"required" message:"客户于合作方向所属BU不能为空"` // 客户于合作方向所属BU
	CooperationDirection string `form:"cooperationDirection" validate:"required" message:"合作方向不能为空"` // 合作方向
	Solution string `form:"solution" validate:"required" message:"额外的产品/解决方案需求不能为空"` // 额外的产品/解决方案需求
	ContactStatus string `form:"contactStatus" validate:"required" message:"联系状态不能为空"` // 联系状态
	SignedProject string `form:"signedProject" validate:"required" message:"签约项目不能为空"` // 签约项目
	EstimatedGrossProfit float64 `form:"estimatedGrossProfit"` // 预计毛利
	BaseCurrency string `form:"baseCurrency" validate:"required" message:"基准货币不能为空"` // 基准货币
	ExchangeRate string `form:"exchangeRate" validate:"required" message:"汇率不能为空"` // 汇率
	ExpectedSigningTime *time.Time `form:"expectedSigningTime" validate:"required" message:"预计签约时间不能为空"` // 预计签约时间
	FollowUpPerson int64 `form:"followUpPerson"` // 跟进人
	FeedbackResults string `form:"feedbackResults" validate:"required" message:"反馈结果不能为空"` // 反馈结果
	NewOldCustomer string `form:"newOldCustomer" validate:"required" message:"新客/老客不能为空"` // 新客/老客
	CustomerPartner string `form:"customerPartner" validate:"required" message:"客户/合作伙伴不能为空"` // 客户/合作伙伴
	IndustryQs string `form:"industryQs" validate:"required" message:"所属行业不能为空"` // 所属行业
	Stage string `form:"stage" validate:"required" message:"阶段不能为空"` // 阶段
	OriginDataOwnerId int64 `form:"originDataOwnerId"` // 原归属人
	FirstGetFromPool int `form:"firstGetFromPool"` // 是否为从池内的第一次获取
	AuditStartUser int64 `form:"auditStartUser"` // 审批发起人
	AuditCode string `form:"auditCode" validate:"required" message:"审批编号不能为空"` // 审批编号
	AuditStatus string `form:"auditStatus" validate:"required" message:"审批状态不能为空"` // 审批状态
	AuditStartTime *time.Time `form:"auditStartTime" validate:"required" message:"审批发起时间不能为空"` // 审批发起时间
	AuditEndTime *time.Time `form:"auditEndTime" validate:"required" message:"审批结束时间不能为空"` // 审批结束时间
	ActivityCityText string `form:"activityCityText" validate:"required" message:"活动城市（文本）不能为空"` // 活动城市（文本）
	HistoryDataOwnerId string `form:"historyDataOwnerId" validate:"required" message:"历史归属人信息不能为空"` // 历史归属人信息
	LeaderAssigned string `form:"leaderAssigned" validate:"required" message:"市领导交办不能为空"` // 市领导交办
	BureauLeaderAssigned string `form:"bureauLeaderAssigned" validate:"required" message:"局领导交办不能为空"` // 局领导交办
	OriginNotes string `form:"originNotes" validate:"required" message:"来源备注不能为空"` // 来源备注
	District string `form:"district" validate:"required" message:"区县不能为空"` // 区县
	PlatformCompany string `form:"platformCompany" validate:"required" message:"平台公司不能为空"` // 平台公司
	FromCityPlan int `form:"fromCityPlan"` // 是否来自创投城市计划生态合伙人
	ServiceAdvisor int64 `form:"serviceAdvisor"` // 市级服务专员
	ServiceAdvisorContactWay string `form:"serviceAdvisorContactWay" validate:"required" message:"市级服务专员联系方式不能为空"` // 市级服务专员联系方式
	IndustryAffiliation string `form:"industryAffiliation" validate:"required" message:"所属产业不能为空"` // 所属产业
	InvestorOriginPlace string `form:"investorOriginPlace" validate:"required" message:"投资方来源地不能为空"` // 投资方来源地
	ForeignInvestment int `form:"foreignInvestment"` // 是否外资
	BpAttachment string `form:"bpAttachment" validate:"required" message:"BP附件不能为空"` // BP附件
	ProposedProjectName string `form:"proposedProjectName" validate:"required" message:"拟投项目名称不能为空"` // 拟投项目名称
	RegisterMoney float64 `form:"registerMoney"` // 注册资金（万元）
	TotalInvestmentAmount float64 `form:"totalInvestmentAmount"` // 总投资额（亿元）
	TaxEstimation float64 `form:"taxEstimation"` // 纳税预估（万元）
	RecruitmentPlanContent string `form:"recruitmentPlanContent" validate:"required" message:"招引谋划内容不能为空"` // 招引谋划内容
	ShareParallelClasses int `form:"shareParallelClasses"` // 是否共享平行专班
	ResidencePlace string `form:"residencePlace" validate:"required" message:"拟落户地不能为空"` // 拟落户地
	LeadsRecommender string `form:"leadsRecommender" validate:"required" message:"线索推荐方不能为空"` // 线索推荐方
	LeadsProvideOrg int64 `form:"leadsProvideOrg"` // 线索提供部门/机构
	LeadsOwner int64 `form:"leadsOwner"` // 线索归属人员
	InputTime *time.Time `form:"inputTime" validate:"required" message:"线索录入时间不能为空"` // 线索录入时间
	InvestProjectName string `form:"investProjectName" validate:"required" message:"拟投资项目名称不能为空"` // 拟投资项目名称
	TechAdvanced string `form:"techAdvanced" validate:"required" message:"技术先进性不能为空"` // 技术先进性
	LandingContent string `form:"landingContent" validate:"required" message:"落地内容不能为空"` // 落地内容
	ToolCarrier string `form:"toolCarrier" validate:"required" message:"载体不能为空"` // 载体
	ProductionFactors string `form:"productionFactors" validate:"required" message:"生产要素不能为空"` // 生产要素
	Financing string `form:"financing" validate:"required" message:"融资不能为空"` // 融资
	Market string `form:"market" validate:"required" message:"场景/市场不能为空"` // 场景/市场
	Policy string `form:"policy" validate:"required" message:"政策不能为空"` // 政策
	IsShareCyzb string `form:"isShareCyzb" validate:"required" message:"是否共享产业专班不能为空"` // 是否共享产业专班
	IsDistributeTask int `form:"isDistributeTask"` // 同时下发行业研判任务
	IsShareLeads int `form:"isShareLeads"` // 同时与行业牵头人共享线索
	InvestCompany int64 `form:"investCompany"` // 投资方名称
	PhaseStatus string `form:"phaseStatus" validate:"required" message:"阶段状态不能为空"` // 阶段状态
	Importance string `form:"importance" validate:"required" message:"重要性不能为空"` // 重要性
	EnterpriseDemand string `form:"enterpriseDemand" validate:"required" message:"企业诉求不能为空"` // 企业诉求
	LeadDepartmentForClues int64 `form:"leadDepartmentForClues"` // 线索牵头部门
	AnalysisContent string `form:"analysisContent" validate:"required" message:"研判内容不能为空"` // 研判内容
	JudgmentOpinion string `form:"judgmentOpinion" validate:"required" message:"研判意见不能为空"` // 研判意见
	ReasonTermination string `form:"reasonTermination" validate:"required" message:"线索终止原因不能为空"` // 线索终止原因
	Destermiantion string `form:"destermiantion" validate:"required" message:"终止说明不能为空"` // 终止说明
	TechnicalAdvantages string `form:"technicalAdvantages" validate:"required" message:"技术优势不能为空"` // 技术优势
	BpAttachments string `form:"bpAttachments" validate:"required" message:"BP附件库不能为空"` // BP附件库
	DockingType string `form:"dockingType" validate:"required" message:"对接形式不能为空"` // 对接形式
	DockingInfo string `form:"dockingInfo" validate:"required" message:"对接情况不能为空"` // 对接情况
	IndustryStatus string `form:"industryStatus" validate:"required" message:"行业地位不能为空"` // 行业地位
	PecuniaryCondition string `form:"pecuniaryCondition" validate:"required" message:"财务情况不能为空"` // 财务情况
	CoreAdvantages string `form:"coreAdvantages" validate:"required" message:"核心优势不能为空"` // 核心优势
	LeadsOrigin string `form:"leadsOrigin" validate:"required" message:"线索来源不能为空"` // 线索来源
	InvestRound string `form:"investRound" validate:"required" message:"拟投资轮次不能为空"` // 拟投资轮次
	NeedIndustry int `form:"needIndustry"` // 是否需要产业投资人
	NeedFinance int `form:"needFinance"` // 是否需要财务投资人
	IndustryQsLabel string `form:"industryQsLabel" validate:"required" message:"所属行业-中文名不能为空"` // 所属行业-中文名
	ImportanceNew string `form:"importanceNew" validate:"required" message:"重要性不能为空"` // 重要性
	BureauLeadsId string `form:"bureauLeadsId" validate:"required" message:"市局系统线索主键id不能为空"` // 市局系统线索主键id
	InvestorOriginPlaceCode string `form:"investorOriginPlaceCode" validate:"required" message:"投资方来源地(省市区)不能为空"` // 投资方来源地(省市区)
	Tag string `form:"tag" validate:"required" message:"数据标签不能为空"` // 数据标签
	MaximumFactoryRent float64 `form:"maximumFactoryRent"` // 厂房租金最大值
	LeasedArea float64 `form:"leasedArea"` // 厂房租用面积
	Period string `form:"period" validate:"required" message:"期限不能为空"` // 期限
	FactoryRentType string `form:"factoryRentType" validate:"required" message:"厂房租金类型不能为空"` // 厂房租金类型
	FloorHeightStarts float64 `form:"floorHeightStarts"` // 厂房层高开始
	ParkLevel string `form:"parkLevel" validate:"required" message:"园区等级不能为空"` // 园区等级
	LandAreaRequired float64 `form:"landAreaRequired"` // 需地块面积
	MinimumRent float64 `form:"minimumRent"` // 楼宇租金最小值
	WhetherThereIsANeedForDe int `form:"whetherThereIsANeedForDe"` // 是否有债权融资需求
	MinimumFactoryRent float64 `form:"minimumFactoryRent"` // 厂房租金最小值
	EndOfFloorHeight float64 `form:"endOfFloorHeight"` // 厂房层高结束
	InvestmentArea float64 `form:"investmentArea"` // 楼宇招商面积
	ParkType string `form:"parkType" validate:"required" message:"园区类型不能为空"` // 园区类型
	MaximumRent float64 `form:"maximumRent"` // 楼宇租金最大值
	InterestRate float64 `form:"interestRate"` // 利率
	LandUse string `form:"landUse" validate:"required" message:"土地用途不能为空"` // 土地用途
	Quota string `form:"quota" validate:"required" message:"额度不能为空"` // 额度
	GuaranteeMethods string `form:"guaranteeMethods" validate:"required" message:"担保方式不能为空"` // 担保方式
	RentType string `form:"rentType" validate:"required" message:"楼宇租金类型不能为空"` // 楼宇租金类型
	FinancingAmountMillionY float64 `form:"financingAmountMillionY"` // 融资金额(亿元)
	ExternalClueId string `form:"externalClueId" validate:"required" message:"外部推荐线索id不能为空"` // 外部推荐线索id
	ClueResearcher int64 `form:"clueResearcher"` // 线索研判人
	AssessmentTeam string `form:"assessmentTeam" validate:"required" message:"研判团队不能为空"` // 研判团队
	EnterpriseintermediaryContact int64 `form:"enterpriseintermediaryContact"` // 企业/中介机构联系人
	ProportionOfEquitySold float64 `form:"proportionOfEquitySold"` // 出售股权比例
	BureauClueProgress string `form:"bureauClueProgress" validate:"required" message:"市局线索进度不能为空"` // 市局线索进度
	CurrentTransferMarketValue float64 `form:"currentTransferMarketValue"` // 当前转让市值
	OperatingIncomeLastYear float64 `form:"operatingIncomeLastYear"` // 去年营业收入
	DetailedDescription string `form:"detailedDescription" validate:"required" message:"详细描述不能为空"` // 详细描述
	TimeRequirementsBefore *time.Time `form:"timeRequirementsBefore" validate:"required" message:"时间要求（之前）不能为空"` // 时间要求（之前）
	WhetherToGambleOnPerforman string `form:"whetherToGambleOnPerforman" validate:"required" message:"是否业绩对赌不能为空"` // 是否业绩对赌
	WhetherItIsAListedCompany string `form:"whetherItIsAListedCompany" validate:"required" message:"是否上市公司不能为空"` // 是否上市公司
	CityCode string `form:"cityCode" validate:"required" message:"所属城市(省市区)不能为空"` // 所属城市(省市区)
	FinancialSituationInThePas string `form:"financialSituationInThePas" validate:"required" message:"近三年财务情况不能为空"` // 近三年财务情况
	ActualControllerSituation string `form:"actualControllerSituation" validate:"required" message:"实际控制人情况不能为空"` // 实际控制人情况
	RequirementsForAcquirers string `form:"requirementsForAcquirers" validate:"required" message:"对收购方要求不能为空"` // 对收购方要求
	TransferConsiderationTransac float64 `form:"transferConsiderationTransac"` // 转让对价（交易金额）
	Seller string `form:"seller" validate:"required" message:"出售方不能为空"` // 出售方
	NeedFurtherUnderstandingOf string `form:"needFurtherUnderstandingOf" validate:"required" message:"需要进一步了解问题不能为空"` // 需要进一步了解问题
	ProjectMatchmaker int64 `form:"projectMatchmaker"` // 项目对接人
	Country string `form:"country" validate:"required" message:"国家不能为空"` // 国家
	ListingPlate string `form:"listingPlate" validate:"required" message:"上市板块不能为空"` // 上市板块
	LatestValuation float64 `form:"latestValuation"` // 最新估值
	ReasonForSelling string `form:"reasonForSelling" validate:"required" message:"出售原因不能为空"` // 出售原因
	NetProfitLastYear float64 `form:"netProfitLastYear"` // 去年净利润
	CanWePromiseToImplementPr string `form:"canWePromiseToImplementPr" validate:"required" message:"能否承诺落地产能或总部搬迁不能为空"` // 能否承诺落地产能或总部搬迁
	PerformanceCommitment string `form:"performanceCommitment" validate:"required" message:"业绩承诺不能为空"` // 业绩承诺
	MainBusinessMainProductsAn string `form:"mainBusinessMainProductsAn" validate:"required" message:"主营业务、主要产品及客户情况不能为空"` // 主营业务、主要产品及客户情况
	ProjectIntroduction string `form:"projectIntroduction" validate:"required" message:"项目介绍不能为空"` // 项目介绍
	CityDockUserPhone string `form:"cityDockUserPhone" validate:"required" message:"市级服务专员联系方式不能为空"` // 市级服务专员联系方式
	ProjectClueContactPerson string `form:"projectClueContactPerson" validate:"required" message:"项目线索联系人不能为空"` // 项目线索联系人
	CityDockUserName string `form:"cityDockUserName" validate:"required" message:"市级服务专员姓名不能为空"` // 市级服务专员姓名
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *LeadsCreateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// LeadsUpdateRequest 更新biz_leads请求参数
type LeadsUpdateRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
	LeadsNo string `form:"leadsNo" validate:"required" message:"线索编码不能为空"` // 线索编码
	CompanyName string `form:"companyName" validate:"required" message:"公司名称不能为空"` // 公司名称
	Region string `form:"region" validate:"required" message:"所在区域不能为空"` // 所在区域
	Mobile string `form:"mobile" validate:"required" message:"手机不能为空"` // 手机
	Telephone string `form:"telephone" validate:"required" message:"电话不能为空"` // 电话
	Email string `form:"email" validate:"required" message:"邮箱不能为空"` // 邮箱
	LeadsSource string `form:"leadsSource" validate:"required" message:"线索来源不能为空"` // 线索来源
	FlowStatus string `form:"flowStatus" validate:"required" message:"跟进状态不能为空"` // 跟进状态
	DistributionStatus string `form:"distributionStatus" validate:"required" message:"分配状态不能为空"` // 分配状态
	TransformationStatus string `form:"transformationStatus" validate:"required" message:"转化状态不能为空"` // 转化状态
	LeadsPool int64 `form:"leadsPool"` // 线索池
	MarketingActivity int64 `form:"marketingActivity"` // 营销活动
	PromotionChannel string `form:"promotionChannel" validate:"required" message:"推广渠道不能为空"` // 推广渠道
	ReturnReason string `form:"returnReason" validate:"required" message:"退回/回收原因不能为空"` // 退回/回收原因
	ReturnReasonDescription string `form:"returnReasonDescription" validate:"required" message:"退回/回收原因说明不能为空"` // 退回/回收原因说明
	ReturnTime *time.Time `form:"returnTime" validate:"required" message:"退回/回收时间不能为空"` // 退回/回收时间
	DistributionTime *time.Time `form:"distributionTime" validate:"required" message:"分配/领取时间不能为空"` // 分配/领取时间
	LeadsDescription string `form:"leadsDescription" validate:"required" message:"线索说明不能为空"` // 线索说明
	CustomerId int64 `form:"customerId"` // 客户id
	ContactId int64 `form:"contactId"` // 联系人id
	BusinessId int64 `form:"businessId"` // 业务id
	TransformationTime *time.Time `form:"transformationTime" validate:"required" message:"转化时间不能为空"` // 转化时间
	Sex string `form:"sex" validate:"required" message:"性别不能为空"` // 性别
	Department string `form:"department" validate:"required" message:"部门不能为空"` // 部门
	JobTitle string `form:"jobTitle" validate:"required" message:"职位不能为空"` // 职位
	Age int `form:"age"` // 年龄
	Birthday *time.Time `form:"birthday" validate:"required" message:"生日不能为空"` // 生日
	Wechat string `form:"wechat" validate:"required" message:"微信不能为空"` // 微信
	Qq string `form:"qq" validate:"required" message:"QQ不能为空"` // QQ
	Fax string `form:"fax" validate:"required" message:"传真不能为空"` // 传真
	CompanyAddress string `form:"companyAddress" validate:"required" message:"公司地址不能为空"` // 公司地址
	CompanyPhone string `form:"companyPhone" validate:"required" message:"公司电话不能为空"` // 公司电话
	Industry string `form:"industry" validate:"required" message:"所属行业不能为空"` // 所属行业
	StaffSize string `form:"staffSize" validate:"required" message:"人员规模不能为空"` // 人员规模
	CompanyEmail string `form:"companyEmail" validate:"required" message:"公司邮箱不能为空"` // 公司邮箱
	CompanyFax string `form:"companyFax" validate:"required" message:"公司传真不能为空"` // 公司传真
	LastFlowTime *time.Time `form:"lastFlowTime" validate:"required" message:"最近跟进时间不能为空"` // 最近跟进时间
	HandoverDescription string `form:"handoverDescription" validate:"required" message:"移交说明不能为空"` // 移交说明
	DistributionDescription string `form:"distributionDescription" validate:"required" message:"分配说明不能为空"` // 分配说明
	GetTime *time.Time `form:"getTime" validate:"required" message:"获取时间不能为空"` // 获取时间
	CountryRegion string `form:"countryRegion" validate:"required" message:"国家地区不能为空"` // 国家地区
	ReturnOriginPool int `form:"returnOriginPool"` // 退回原池
	BusinessLicense string `form:"businessLicense" validate:"required" message:"营业执照不能为空"` // 营业执照
	GmName string `form:"gmName" validate:"required" message:"总经理姓名不能为空"` // 总经理姓名
	SalesExperience int `form:"salesExperience"` // 是否有销售经验
	SalesDescription string `form:"salesDescription" validate:"required" message:"销售经验简述不能为空"` // 销售经验简述
	AnnualSales string `form:"annualSales" validate:"required" message:"年销售额不能为空"` // 年销售额
	SameIndustry int `form:"sameIndustry"` // 是否是同一领域的其他公司/品牌的合作伙伴
	CompanyBrand string `form:"companyBrand" validate:"required" message:"具体公司和品牌不能为空"` // 具体公司和品牌
	OtherInformation string `form:"otherInformation" validate:"required" message:"其他信息不能为空"` // 其他信息
	CompanyWeb string `form:"companyWeb" validate:"required" message:"公司网址不能为空"` // 公司网址
	LicenseValid *time.Time `form:"licenseValid" validate:"required" message:"执照有效期不能为空"` // 执照有效期
	SalesYears float64 `form:"salesYears"` // 累计销售年数
	SalesDeptUserNum float64 `form:"salesDeptUserNum"` // 销售部门人数
	MarketingDeptUserNum float64 `form:"marketingDeptUserNum"` // 营销部门人数
	ServiceDeptUserNum float64 `form:"serviceDeptUserNum"` // 服务部门人数
	OtherDeptUserNum float64 `form:"otherDeptUserNum"` // 其他部门人数
	Role string `form:"role" validate:"required" message:"角色不能为空"` // 角色
	EnterpriseType string `form:"enterpriseType" validate:"required" message:"企业类型不能为空"` // 企业类型
	StartDate *time.Time `form:"startDate" validate:"required" message:"活动开始时间不能为空"` // 活动开始时间
	EndDate *time.Time `form:"endDate" validate:"required" message:"活动结束时间不能为空"` // 活动结束时间
	ActivityCity string `form:"activityCity" validate:"required" message:"活动城市不能为空"` // 活动城市
	Sponsor string `form:"sponsor" validate:"required" message:"主办方不能为空"` // 主办方
	ActivityName string `form:"activityName" validate:"required" message:"活动名称不能为空"` // 活动名称
	OwnBu string `form:"ownBu" validate:"required" message:"客户于合作方向所属BU不能为空"` // 客户于合作方向所属BU
	CooperationDirection string `form:"cooperationDirection" validate:"required" message:"合作方向不能为空"` // 合作方向
	Solution string `form:"solution" validate:"required" message:"额外的产品/解决方案需求不能为空"` // 额外的产品/解决方案需求
	ContactStatus string `form:"contactStatus" validate:"required" message:"联系状态不能为空"` // 联系状态
	SignedProject string `form:"signedProject" validate:"required" message:"签约项目不能为空"` // 签约项目
	EstimatedGrossProfit float64 `form:"estimatedGrossProfit"` // 预计毛利
	BaseCurrency string `form:"baseCurrency" validate:"required" message:"基准货币不能为空"` // 基准货币
	ExchangeRate string `form:"exchangeRate" validate:"required" message:"汇率不能为空"` // 汇率
	ExpectedSigningTime *time.Time `form:"expectedSigningTime" validate:"required" message:"预计签约时间不能为空"` // 预计签约时间
	FollowUpPerson int64 `form:"followUpPerson"` // 跟进人
	FeedbackResults string `form:"feedbackResults" validate:"required" message:"反馈结果不能为空"` // 反馈结果
	NewOldCustomer string `form:"newOldCustomer" validate:"required" message:"新客/老客不能为空"` // 新客/老客
	CustomerPartner string `form:"customerPartner" validate:"required" message:"客户/合作伙伴不能为空"` // 客户/合作伙伴
	IndustryQs string `form:"industryQs" validate:"required" message:"所属行业不能为空"` // 所属行业
	Stage string `form:"stage" validate:"required" message:"阶段不能为空"` // 阶段
	OriginDataOwnerId int64 `form:"originDataOwnerId"` // 原归属人
	FirstGetFromPool int `form:"firstGetFromPool"` // 是否为从池内的第一次获取
	AuditStartUser int64 `form:"auditStartUser"` // 审批发起人
	AuditCode string `form:"auditCode" validate:"required" message:"审批编号不能为空"` // 审批编号
	AuditStatus string `form:"auditStatus" validate:"required" message:"审批状态不能为空"` // 审批状态
	AuditStartTime *time.Time `form:"auditStartTime" validate:"required" message:"审批发起时间不能为空"` // 审批发起时间
	AuditEndTime *time.Time `form:"auditEndTime" validate:"required" message:"审批结束时间不能为空"` // 审批结束时间
	ActivityCityText string `form:"activityCityText" validate:"required" message:"活动城市（文本）不能为空"` // 活动城市（文本）
	HistoryDataOwnerId string `form:"historyDataOwnerId" validate:"required" message:"历史归属人信息不能为空"` // 历史归属人信息
	LeaderAssigned string `form:"leaderAssigned" validate:"required" message:"市领导交办不能为空"` // 市领导交办
	BureauLeaderAssigned string `form:"bureauLeaderAssigned" validate:"required" message:"局领导交办不能为空"` // 局领导交办
	OriginNotes string `form:"originNotes" validate:"required" message:"来源备注不能为空"` // 来源备注
	District string `form:"district" validate:"required" message:"区县不能为空"` // 区县
	PlatformCompany string `form:"platformCompany" validate:"required" message:"平台公司不能为空"` // 平台公司
	FromCityPlan int `form:"fromCityPlan"` // 是否来自创投城市计划生态合伙人
	ServiceAdvisor int64 `form:"serviceAdvisor"` // 市级服务专员
	ServiceAdvisorContactWay string `form:"serviceAdvisorContactWay" validate:"required" message:"市级服务专员联系方式不能为空"` // 市级服务专员联系方式
	IndustryAffiliation string `form:"industryAffiliation" validate:"required" message:"所属产业不能为空"` // 所属产业
	InvestorOriginPlace string `form:"investorOriginPlace" validate:"required" message:"投资方来源地不能为空"` // 投资方来源地
	ForeignInvestment int `form:"foreignInvestment"` // 是否外资
	BpAttachment string `form:"bpAttachment" validate:"required" message:"BP附件不能为空"` // BP附件
	ProposedProjectName string `form:"proposedProjectName" validate:"required" message:"拟投项目名称不能为空"` // 拟投项目名称
	RegisterMoney float64 `form:"registerMoney"` // 注册资金（万元）
	TotalInvestmentAmount float64 `form:"totalInvestmentAmount"` // 总投资额（亿元）
	TaxEstimation float64 `form:"taxEstimation"` // 纳税预估（万元）
	RecruitmentPlanContent string `form:"recruitmentPlanContent" validate:"required" message:"招引谋划内容不能为空"` // 招引谋划内容
	ShareParallelClasses int `form:"shareParallelClasses"` // 是否共享平行专班
	ResidencePlace string `form:"residencePlace" validate:"required" message:"拟落户地不能为空"` // 拟落户地
	LeadsRecommender string `form:"leadsRecommender" validate:"required" message:"线索推荐方不能为空"` // 线索推荐方
	LeadsProvideOrg int64 `form:"leadsProvideOrg"` // 线索提供部门/机构
	LeadsOwner int64 `form:"leadsOwner"` // 线索归属人员
	InputTime *time.Time `form:"inputTime" validate:"required" message:"线索录入时间不能为空"` // 线索录入时间
	InvestProjectName string `form:"investProjectName" validate:"required" message:"拟投资项目名称不能为空"` // 拟投资项目名称
	TechAdvanced string `form:"techAdvanced" validate:"required" message:"技术先进性不能为空"` // 技术先进性
	LandingContent string `form:"landingContent" validate:"required" message:"落地内容不能为空"` // 落地内容
	ToolCarrier string `form:"toolCarrier" validate:"required" message:"载体不能为空"` // 载体
	ProductionFactors string `form:"productionFactors" validate:"required" message:"生产要素不能为空"` // 生产要素
	Financing string `form:"financing" validate:"required" message:"融资不能为空"` // 融资
	Market string `form:"market" validate:"required" message:"场景/市场不能为空"` // 场景/市场
	Policy string `form:"policy" validate:"required" message:"政策不能为空"` // 政策
	IsShareCyzb string `form:"isShareCyzb" validate:"required" message:"是否共享产业专班不能为空"` // 是否共享产业专班
	IsDistributeTask int `form:"isDistributeTask"` // 同时下发行业研判任务
	IsShareLeads int `form:"isShareLeads"` // 同时与行业牵头人共享线索
	InvestCompany int64 `form:"investCompany"` // 投资方名称
	PhaseStatus string `form:"phaseStatus" validate:"required" message:"阶段状态不能为空"` // 阶段状态
	Importance string `form:"importance" validate:"required" message:"重要性不能为空"` // 重要性
	EnterpriseDemand string `form:"enterpriseDemand" validate:"required" message:"企业诉求不能为空"` // 企业诉求
	LeadDepartmentForClues int64 `form:"leadDepartmentForClues"` // 线索牵头部门
	AnalysisContent string `form:"analysisContent" validate:"required" message:"研判内容不能为空"` // 研判内容
	JudgmentOpinion string `form:"judgmentOpinion" validate:"required" message:"研判意见不能为空"` // 研判意见
	ReasonTermination string `form:"reasonTermination" validate:"required" message:"线索终止原因不能为空"` // 线索终止原因
	Destermiantion string `form:"destermiantion" validate:"required" message:"终止说明不能为空"` // 终止说明
	TechnicalAdvantages string `form:"technicalAdvantages" validate:"required" message:"技术优势不能为空"` // 技术优势
	BpAttachments string `form:"bpAttachments" validate:"required" message:"BP附件库不能为空"` // BP附件库
	DockingType string `form:"dockingType" validate:"required" message:"对接形式不能为空"` // 对接形式
	DockingInfo string `form:"dockingInfo" validate:"required" message:"对接情况不能为空"` // 对接情况
	IndustryStatus string `form:"industryStatus" validate:"required" message:"行业地位不能为空"` // 行业地位
	PecuniaryCondition string `form:"pecuniaryCondition" validate:"required" message:"财务情况不能为空"` // 财务情况
	CoreAdvantages string `form:"coreAdvantages" validate:"required" message:"核心优势不能为空"` // 核心优势
	LeadsOrigin string `form:"leadsOrigin" validate:"required" message:"线索来源不能为空"` // 线索来源
	InvestRound string `form:"investRound" validate:"required" message:"拟投资轮次不能为空"` // 拟投资轮次
	NeedIndustry int `form:"needIndustry"` // 是否需要产业投资人
	NeedFinance int `form:"needFinance"` // 是否需要财务投资人
	IndustryQsLabel string `form:"industryQsLabel" validate:"required" message:"所属行业-中文名不能为空"` // 所属行业-中文名
	ImportanceNew string `form:"importanceNew" validate:"required" message:"重要性不能为空"` // 重要性
	BureauLeadsId string `form:"bureauLeadsId" validate:"required" message:"市局系统线索主键id不能为空"` // 市局系统线索主键id
	InvestorOriginPlaceCode string `form:"investorOriginPlaceCode" validate:"required" message:"投资方来源地(省市区)不能为空"` // 投资方来源地(省市区)
	Tag string `form:"tag" validate:"required" message:"数据标签不能为空"` // 数据标签
	MaximumFactoryRent float64 `form:"maximumFactoryRent"` // 厂房租金最大值
	LeasedArea float64 `form:"leasedArea"` // 厂房租用面积
	Period string `form:"period" validate:"required" message:"期限不能为空"` // 期限
	FactoryRentType string `form:"factoryRentType" validate:"required" message:"厂房租金类型不能为空"` // 厂房租金类型
	FloorHeightStarts float64 `form:"floorHeightStarts"` // 厂房层高开始
	ParkLevel string `form:"parkLevel" validate:"required" message:"园区等级不能为空"` // 园区等级
	LandAreaRequired float64 `form:"landAreaRequired"` // 需地块面积
	MinimumRent float64 `form:"minimumRent"` // 楼宇租金最小值
	WhetherThereIsANeedForDe int `form:"whetherThereIsANeedForDe"` // 是否有债权融资需求
	MinimumFactoryRent float64 `form:"minimumFactoryRent"` // 厂房租金最小值
	EndOfFloorHeight float64 `form:"endOfFloorHeight"` // 厂房层高结束
	InvestmentArea float64 `form:"investmentArea"` // 楼宇招商面积
	ParkType string `form:"parkType" validate:"required" message:"园区类型不能为空"` // 园区类型
	MaximumRent float64 `form:"maximumRent"` // 楼宇租金最大值
	InterestRate float64 `form:"interestRate"` // 利率
	LandUse string `form:"landUse" validate:"required" message:"土地用途不能为空"` // 土地用途
	Quota string `form:"quota" validate:"required" message:"额度不能为空"` // 额度
	GuaranteeMethods string `form:"guaranteeMethods" validate:"required" message:"担保方式不能为空"` // 担保方式
	RentType string `form:"rentType" validate:"required" message:"楼宇租金类型不能为空"` // 楼宇租金类型
	FinancingAmountMillionY float64 `form:"financingAmountMillionY"` // 融资金额(亿元)
	ExternalClueId string `form:"externalClueId" validate:"required" message:"外部推荐线索id不能为空"` // 外部推荐线索id
	ClueResearcher int64 `form:"clueResearcher"` // 线索研判人
	AssessmentTeam string `form:"assessmentTeam" validate:"required" message:"研判团队不能为空"` // 研判团队
	EnterpriseintermediaryContact int64 `form:"enterpriseintermediaryContact"` // 企业/中介机构联系人
	ProportionOfEquitySold float64 `form:"proportionOfEquitySold"` // 出售股权比例
	BureauClueProgress string `form:"bureauClueProgress" validate:"required" message:"市局线索进度不能为空"` // 市局线索进度
	CurrentTransferMarketValue float64 `form:"currentTransferMarketValue"` // 当前转让市值
	OperatingIncomeLastYear float64 `form:"operatingIncomeLastYear"` // 去年营业收入
	DetailedDescription string `form:"detailedDescription" validate:"required" message:"详细描述不能为空"` // 详细描述
	TimeRequirementsBefore *time.Time `form:"timeRequirementsBefore" validate:"required" message:"时间要求（之前）不能为空"` // 时间要求（之前）
	WhetherToGambleOnPerforman string `form:"whetherToGambleOnPerforman" validate:"required" message:"是否业绩对赌不能为空"` // 是否业绩对赌
	WhetherItIsAListedCompany string `form:"whetherItIsAListedCompany" validate:"required" message:"是否上市公司不能为空"` // 是否上市公司
	CityCode string `form:"cityCode" validate:"required" message:"所属城市(省市区)不能为空"` // 所属城市(省市区)
	FinancialSituationInThePas string `form:"financialSituationInThePas" validate:"required" message:"近三年财务情况不能为空"` // 近三年财务情况
	ActualControllerSituation string `form:"actualControllerSituation" validate:"required" message:"实际控制人情况不能为空"` // 实际控制人情况
	RequirementsForAcquirers string `form:"requirementsForAcquirers" validate:"required" message:"对收购方要求不能为空"` // 对收购方要求
	TransferConsiderationTransac float64 `form:"transferConsiderationTransac"` // 转让对价（交易金额）
	Seller string `form:"seller" validate:"required" message:"出售方不能为空"` // 出售方
	NeedFurtherUnderstandingOf string `form:"needFurtherUnderstandingOf" validate:"required" message:"需要进一步了解问题不能为空"` // 需要进一步了解问题
	ProjectMatchmaker int64 `form:"projectMatchmaker"` // 项目对接人
	Country string `form:"country" validate:"required" message:"国家不能为空"` // 国家
	ListingPlate string `form:"listingPlate" validate:"required" message:"上市板块不能为空"` // 上市板块
	LatestValuation float64 `form:"latestValuation"` // 最新估值
	ReasonForSelling string `form:"reasonForSelling" validate:"required" message:"出售原因不能为空"` // 出售原因
	NetProfitLastYear float64 `form:"netProfitLastYear"` // 去年净利润
	CanWePromiseToImplementPr string `form:"canWePromiseToImplementPr" validate:"required" message:"能否承诺落地产能或总部搬迁不能为空"` // 能否承诺落地产能或总部搬迁
	PerformanceCommitment string `form:"performanceCommitment" validate:"required" message:"业绩承诺不能为空"` // 业绩承诺
	MainBusinessMainProductsAn string `form:"mainBusinessMainProductsAn" validate:"required" message:"主营业务、主要产品及客户情况不能为空"` // 主营业务、主要产品及客户情况
	ProjectIntroduction string `form:"projectIntroduction" validate:"required" message:"项目介绍不能为空"` // 项目介绍
	CityDockUserPhone string `form:"cityDockUserPhone" validate:"required" message:"市级服务专员联系方式不能为空"` // 市级服务专员联系方式
	ProjectClueContactPerson string `form:"projectClueContactPerson" validate:"required" message:"项目线索联系人不能为空"` // 项目线索联系人
	CityDockUserName string `form:"cityDockUserName" validate:"required" message:"市级服务专员姓名不能为空"` // 市级服务专员姓名
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *LeadsUpdateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// LeadsDeleteRequest 删除biz_leads请求参数
type LeadsDeleteRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *LeadsDeleteRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// LeadsGetByIDRequest 根据ID获取biz_leads请求参数
type LeadsGetByIDRequest struct {
	models.Validator
	Id int64 `uri:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *LeadsGetByIDRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}