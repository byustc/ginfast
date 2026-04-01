package models

import (
	"context"
	"time"
	"gin-fast/app/global/app"
	"gorm.io/gorm"
)

// Leads biz_leads 模型结构体
type Leads struct {
	Id int64 `gorm:"column:id;primaryKey;not null;autoIncrement" json:"id"` // 主键
	TenantId uint `gorm:"column:tenant_id;default:0;index" json:"tenantId"` // 租户ID字段
	LeadsNo string `gorm:"column:leads_no;index" json:"leadsNo"` // 线索编码
	CompanyName string `gorm:"column:company_name" json:"companyName"` // 公司名称
	Region string `gorm:"column:region" json:"region"` // 所在区域
	Mobile string `gorm:"column:mobile" json:"mobile"` // 手机
	Telephone string `gorm:"column:telephone" json:"telephone"` // 电话
	Email string `gorm:"column:email" json:"email"` // 邮箱
	LeadsSource string `gorm:"column:leads_source" json:"leadsSource"` // 线索来源
	FlowStatus string `gorm:"column:flow_status" json:"flowStatus"` // 跟进状态
	DistributionStatus string `gorm:"column:distribution_status" json:"distributionStatus"` // 分配状态
	TransformationStatus string `gorm:"column:transformation_status" json:"transformationStatus"` // 转化状态
	LeadsPool int64 `gorm:"column:leads_pool" json:"leadsPool"` // 线索池
	MarketingActivity int64 `gorm:"column:marketing_activity" json:"marketingActivity"` // 营销活动
	PromotionChannel string `gorm:"column:promotion_channel" json:"promotionChannel"` // 推广渠道
	ReturnReason string `gorm:"column:return_reason" json:"returnReason"` // 退回/回收原因
	ReturnReasonDescription string `gorm:"column:return_reason_description" json:"returnReasonDescription"` // 退回/回收原因说明
	ReturnTime *time.Time `gorm:"column:return_time" json:"returnTime"` // 退回/回收时间
	DistributionTime *time.Time `gorm:"column:distribution_time" json:"distributionTime"` // 分配/领取时间
	LeadsDescription string `gorm:"column:leads_description" json:"leadsDescription"` // 线索说明
	CustomerId int64 `gorm:"column:customer_id" json:"customerId"` // 客户id
	ContactId int64 `gorm:"column:contact_id" json:"contactId"` // 联系人id
	BusinessId int64 `gorm:"column:business_id" json:"businessId"` // 业务id
	TransformationTime *time.Time `gorm:"column:transformation_time" json:"transformationTime"` // 转化时间
	Sex string `gorm:"column:sex" json:"sex"` // 性别
	Department string `gorm:"column:department" json:"department"` // 部门
	JobTitle string `gorm:"column:job_title" json:"jobTitle"` // 职位
	Age int `gorm:"column:age" json:"age"` // 年龄
	Birthday *time.Time `gorm:"column:birthday" json:"birthday"` // 生日
	Wechat string `gorm:"column:wechat" json:"wechat"` // 微信
	Qq string `gorm:"column:qq" json:"qq"` // QQ
	Fax string `gorm:"column:fax" json:"fax"` // 传真
	CompanyAddress string `gorm:"column:company_address" json:"companyAddress"` // 公司地址
	CompanyPhone string `gorm:"column:company_phone" json:"companyPhone"` // 公司电话
	Industry string `gorm:"column:industry" json:"industry"` // 所属行业
	StaffSize string `gorm:"column:staff_size" json:"staffSize"` // 人员规模
	CompanyEmail string `gorm:"column:company_email" json:"companyEmail"` // 公司邮箱
	CompanyFax string `gorm:"column:company_fax" json:"companyFax"` // 公司传真
	LastFlowTime *time.Time `gorm:"column:last_flow_time" json:"lastFlowTime"` // 最近跟进时间
	HandoverDescription string `gorm:"column:handover_description" json:"handoverDescription"` // 移交说明
	DistributionDescription string `gorm:"column:distribution_description" json:"distributionDescription"` // 分配说明
	GetTime *time.Time `gorm:"column:get_time" json:"getTime"` // 获取时间
	CountryRegion string `gorm:"column:country_region" json:"countryRegion"` // 国家地区
	ReturnOriginPool int `gorm:"column:return_origin_pool" json:"returnOriginPool"` // 退回原池
	BusinessLicense string `gorm:"column:business_license" json:"businessLicense"` // 营业执照
	GmName string `gorm:"column:gm_name" json:"gmName"` // 总经理姓名
	SalesExperience int `gorm:"column:sales_experience" json:"salesExperience"` // 是否有销售经验
	SalesDescription string `gorm:"column:sales_description" json:"salesDescription"` // 销售经验简述
	AnnualSales string `gorm:"column:annual_sales" json:"annualSales"` // 年销售额
	SameIndustry int `gorm:"column:same_industry" json:"sameIndustry"` // 是否是同一领域的其他公司/品牌的合作伙伴
	CompanyBrand string `gorm:"column:company_brand" json:"companyBrand"` // 具体公司和品牌
	OtherInformation string `gorm:"column:other_information" json:"otherInformation"` // 其他信息
	CompanyWeb string `gorm:"column:company_web" json:"companyWeb"` // 公司网址
	LicenseValid *time.Time `gorm:"column:license_valid" json:"licenseValid"` // 执照有效期
	SalesYears float64 `gorm:"column:sales_years" json:"salesYears"` // 累计销售年数
	SalesDeptUserNum float64 `gorm:"column:sales_dept_user_num" json:"salesDeptUserNum"` // 销售部门人数
	MarketingDeptUserNum float64 `gorm:"column:marketing_dept_user_num" json:"marketingDeptUserNum"` // 营销部门人数
	ServiceDeptUserNum float64 `gorm:"column:service_dept_user_num" json:"serviceDeptUserNum"` // 服务部门人数
	OtherDeptUserNum float64 `gorm:"column:other_dept_user_num" json:"otherDeptUserNum"` // 其他部门人数
	Role string `gorm:"column:role" json:"role"` // 角色
	EnterpriseType string `gorm:"column:enterprise_type" json:"enterpriseType"` // 企业类型
	StartDate *time.Time `gorm:"column:start_date" json:"startDate"` // 活动开始时间
	EndDate *time.Time `gorm:"column:end_date" json:"endDate"` // 活动结束时间
	ActivityCity string `gorm:"column:activity_city" json:"activityCity"` // 活动城市
	Sponsor string `gorm:"column:sponsor" json:"sponsor"` // 主办方
	ActivityName string `gorm:"column:activity_name" json:"activityName"` // 活动名称
	OwnBu string `gorm:"column:own_bu" json:"ownBu"` // 客户于合作方向所属BU
	CooperationDirection string `gorm:"column:cooperation_direction" json:"cooperationDirection"` // 合作方向
	Solution string `gorm:"column:solution" json:"solution"` // 额外的产品/解决方案需求
	ContactStatus string `gorm:"column:contact_status" json:"contactStatus"` // 联系状态
	SignedProject string `gorm:"column:signed_project" json:"signedProject"` // 签约项目
	EstimatedGrossProfit float64 `gorm:"column:estimated_gross_profit" json:"estimatedGrossProfit"` // 预计毛利
	BaseCurrency string `gorm:"column:base_currency" json:"baseCurrency"` // 基准货币
	ExchangeRate string `gorm:"column:exchange_rate" json:"exchangeRate"` // 汇率
	ExpectedSigningTime *time.Time `gorm:"column:expected_signing_time" json:"expectedSigningTime"` // 预计签约时间
	FollowUpPerson int64 `gorm:"column:follow_up_person" json:"followUpPerson"` // 跟进人
	FeedbackResults string `gorm:"column:feedback_results" json:"feedbackResults"` // 反馈结果
	NewOldCustomer string `gorm:"column:new_old_customer" json:"newOldCustomer"` // 新客/老客
	CustomerPartner string `gorm:"column:customer_partner" json:"customerPartner"` // 客户/合作伙伴
	IndustryQs string `gorm:"column:industry_qs" json:"industryQs"` // 所属行业
	Stage string `gorm:"column:stage" json:"stage"` // 阶段
	OriginDataOwnerId int64 `gorm:"column:origin_data_owner_id" json:"originDataOwnerId"` // 原归属人
	FirstGetFromPool int `gorm:"column:first_get_from_pool;default:0" json:"firstGetFromPool"` // 是否为从池内的第一次获取
	AuditStartUser int64 `gorm:"column:audit_start_user" json:"auditStartUser"` // 审批发起人
	AuditCode string `gorm:"column:audit_code" json:"auditCode"` // 审批编号
	AuditStatus string `gorm:"column:audit_status" json:"auditStatus"` // 审批状态
	AuditStartTime *time.Time `gorm:"column:audit_start_time" json:"auditStartTime"` // 审批发起时间
	AuditEndTime *time.Time `gorm:"column:audit_end_time" json:"auditEndTime"` // 审批结束时间
	ActivityCityText string `gorm:"column:activity_city_text" json:"activityCityText"` // 活动城市（文本）
	HistoryDataOwnerId string `gorm:"column:history_data_owner_id" json:"historyDataOwnerId"` // 历史归属人信息
	LeaderAssigned string `gorm:"column:leader_assigned" json:"leaderAssigned"` // 市领导交办
	BureauLeaderAssigned string `gorm:"column:bureau_leader_assigned" json:"bureauLeaderAssigned"` // 局领导交办
	OriginNotes string `gorm:"column:origin_notes" json:"originNotes"` // 来源备注
	District string `gorm:"column:district" json:"district"` // 区县
	PlatformCompany string `gorm:"column:platform_company" json:"platformCompany"` // 平台公司
	FromCityPlan int `gorm:"column:from_city_plan;default:0" json:"fromCityPlan"` // 是否来自创投城市计划生态合伙人
	ServiceAdvisor int64 `gorm:"column:service_advisor" json:"serviceAdvisor"` // 市级服务专员
	ServiceAdvisorContactWay string `gorm:"column:service_advisor_contact_way" json:"serviceAdvisorContactWay"` // 市级服务专员联系方式
	IndustryAffiliation string `gorm:"column:industry_affiliation" json:"industryAffiliation"` // 所属产业
	InvestorOriginPlace string `gorm:"column:investor_origin_place" json:"investorOriginPlace"` // 投资方来源地
	ForeignInvestment int `gorm:"column:foreign_investment;default:0" json:"foreignInvestment"` // 是否外资
	BpAttachment string `gorm:"column:bp_attachment" json:"bpAttachment"` // BP附件
	ProposedProjectName string `gorm:"column:proposed_project_name" json:"proposedProjectName"` // 拟投项目名称
	RegisterMoney float64 `gorm:"column:register_money" json:"registerMoney"` // 注册资金（万元）
	TotalInvestmentAmount float64 `gorm:"column:total_investment_amount" json:"totalInvestmentAmount"` // 总投资额（亿元）
	TaxEstimation float64 `gorm:"column:tax_estimation" json:"taxEstimation"` // 纳税预估（万元）
	RecruitmentPlanContent string `gorm:"column:recruitment_plan_content" json:"recruitmentPlanContent"` // 招引谋划内容
	ShareParallelClasses int `gorm:"column:share_parallel_classes;default:0" json:"shareParallelClasses"` // 是否共享平行专班
	ResidencePlace string `gorm:"column:residence_place" json:"residencePlace"` // 拟落户地
	LeadsRecommender string `gorm:"column:leads_recommender" json:"leadsRecommender"` // 线索推荐方
	LeadsProvideOrg int64 `gorm:"column:leads_provide_org" json:"leadsProvideOrg"` // 线索提供部门/机构
	LeadsOwner int64 `gorm:"column:leads_owner" json:"leadsOwner"` // 线索归属人员
	InputTime *time.Time `gorm:"column:input_time" json:"inputTime"` // 线索录入时间
	InvestProjectName string `gorm:"column:invest_project_name" json:"investProjectName"` // 拟投资项目名称
	TechAdvanced string `gorm:"column:tech_advanced" json:"techAdvanced"` // 技术先进性
	LandingContent string `gorm:"column:landing_content" json:"landingContent"` // 落地内容
	ToolCarrier string `gorm:"column:tool_carrier" json:"toolCarrier"` // 载体
	ProductionFactors string `gorm:"column:production_factors" json:"productionFactors"` // 生产要素
	Financing string `gorm:"column:financing" json:"financing"` // 融资
	Market string `gorm:"column:market" json:"market"` // 场景/市场
	Policy string `gorm:"column:policy" json:"policy"` // 政策
	IsShareCyzb string `gorm:"column:is_share_cyzb" json:"isShareCyzb"` // 是否共享产业专班
	IsDistributeTask int `gorm:"column:is_distribute_task;default:0" json:"isDistributeTask"` // 同时下发行业研判任务
	IsShareLeads int `gorm:"column:is_share_leads;default:0" json:"isShareLeads"` // 同时与行业牵头人共享线索
	InvestCompany int64 `gorm:"column:invest_company" json:"investCompany"` // 投资方名称
	PhaseStatus string `gorm:"column:phase_status" json:"phaseStatus"` // 阶段状态
	Importance string `gorm:"column:importance" json:"importance"` // 重要性
	EnterpriseDemand string `gorm:"column:enterprise_demand" json:"enterpriseDemand"` // 企业诉求
	LeadDepartmentForClues int64 `gorm:"column:lead_department_for_clues" json:"leadDepartmentForClues"` // 线索牵头部门
	AnalysisContent string `gorm:"column:analysis_content" json:"analysisContent"` // 研判内容
	JudgmentOpinion string `gorm:"column:judgment_opinion" json:"judgmentOpinion"` // 研判意见
	ReasonTermination string `gorm:"column:reason_termination" json:"reasonTermination"` // 线索终止原因
	Destermiantion string `gorm:"column:destermiantion" json:"destermiantion"` // 终止说明
	TechnicalAdvantages string `gorm:"column:technical_advantages" json:"technicalAdvantages"` // 技术优势
	BpAttachments string `gorm:"column:bp_attachments" json:"bpAttachments"` // BP附件库
	DockingType string `gorm:"column:docking_type" json:"dockingType"` // 对接形式
	DockingInfo string `gorm:"column:docking_info" json:"dockingInfo"` // 对接情况
	IndustryStatus string `gorm:"column:industry_status" json:"industryStatus"` // 行业地位
	PecuniaryCondition string `gorm:"column:pecuniary_condition" json:"pecuniaryCondition"` // 财务情况
	CoreAdvantages string `gorm:"column:core_advantages" json:"coreAdvantages"` // 核心优势
	LeadsOrigin string `gorm:"column:leads_origin" json:"leadsOrigin"` // 线索来源
	InvestRound string `gorm:"column:invest_round" json:"investRound"` // 拟投资轮次
	NeedIndustry int `gorm:"column:need_industry;default:0" json:"needIndustry"` // 是否需要产业投资人
	NeedFinance int `gorm:"column:need_finance;default:0" json:"needFinance"` // 是否需要财务投资人
	IndustryQsLabel string `gorm:"column:industry_qs_label" json:"industryQsLabel"` // 所属行业-中文名
	ImportanceNew string `gorm:"column:importance_new" json:"importanceNew"` // 重要性
	BureauLeadsId string `gorm:"column:bureau_leads_id" json:"bureauLeadsId"` // 市局系统线索主键id
	InvestorOriginPlaceCode string `gorm:"column:investor_origin_place_code" json:"investorOriginPlaceCode"` // 投资方来源地(省市区)
	Tag string `gorm:"column:tag" json:"tag"` // 数据标签
	MaximumFactoryRent float64 `gorm:"column:maximum_factory_rent" json:"maximumFactoryRent"` // 厂房租金最大值
	LeasedArea float64 `gorm:"column:leased_area" json:"leasedArea"` // 厂房租用面积
	Period string `gorm:"column:period" json:"period"` // 期限
	FactoryRentType string `gorm:"column:factory_rent_type" json:"factoryRentType"` // 厂房租金类型
	FloorHeightStarts float64 `gorm:"column:floor_height_starts" json:"floorHeightStarts"` // 厂房层高开始
	ParkLevel string `gorm:"column:park_level" json:"parkLevel"` // 园区等级
	LandAreaRequired float64 `gorm:"column:land_area_required" json:"landAreaRequired"` // 需地块面积
	MinimumRent float64 `gorm:"column:minimum_rent" json:"minimumRent"` // 楼宇租金最小值
	WhetherThereIsANeedForDe int `gorm:"column:whether_there_is_a_need_for_de" json:"whetherThereIsANeedForDe"` // 是否有债权融资需求
	MinimumFactoryRent float64 `gorm:"column:minimum_factory_rent" json:"minimumFactoryRent"` // 厂房租金最小值
	EndOfFloorHeight float64 `gorm:"column:end_of_floor_height" json:"endOfFloorHeight"` // 厂房层高结束
	InvestmentArea float64 `gorm:"column:investment_area" json:"investmentArea"` // 楼宇招商面积
	ParkType string `gorm:"column:park_type" json:"parkType"` // 园区类型
	MaximumRent float64 `gorm:"column:maximum_rent" json:"maximumRent"` // 楼宇租金最大值
	InterestRate float64 `gorm:"column:interest_rate" json:"interestRate"` // 利率
	LandUse string `gorm:"column:land_use" json:"landUse"` // 土地用途
	Quota string `gorm:"column:quota" json:"quota"` // 额度
	GuaranteeMethods string `gorm:"column:guarantee_methods" json:"guaranteeMethods"` // 担保方式
	RentType string `gorm:"column:rent_type" json:"rentType"` // 楼宇租金类型
	FinancingAmountMillionY float64 `gorm:"column:financing_amount_100_million_y" json:"financingAmountMillionY"` // 融资金额(亿元)
	ExternalClueId string `gorm:"column:external_clue_id" json:"externalClueId"` // 外部推荐线索id
	ClueResearcher int64 `gorm:"column:clue_researcher" json:"clueResearcher"` // 线索研判人
	AssessmentTeam string `gorm:"column:assessment_team" json:"assessmentTeam"` // 研判团队
	EnterpriseintermediaryContact int64 `gorm:"column:enterpriseintermediary_contact" json:"enterpriseintermediaryContact"` // 企业/中介机构联系人
	ProportionOfEquitySold float64 `gorm:"column:proportion_of_equity_sold" json:"proportionOfEquitySold"` // 出售股权比例
	BureauClueProgress string `gorm:"column:bureau_clue_progress" json:"bureauClueProgress"` // 市局线索进度
	CurrentTransferMarketValue float64 `gorm:"column:current_transfer_market_value" json:"currentTransferMarketValue"` // 当前转让市值
	OperatingIncomeLastYear float64 `gorm:"column:operating_income_last_year" json:"operatingIncomeLastYear"` // 去年营业收入
	DetailedDescription string `gorm:"column:detailed_description" json:"detailedDescription"` // 详细描述
	TimeRequirementsBefore *time.Time `gorm:"column:time_requirements_before" json:"timeRequirementsBefore"` // 时间要求（之前）
	WhetherToGambleOnPerforman string `gorm:"column:whether_to_gamble_on_performan" json:"whetherToGambleOnPerforman"` // 是否业绩对赌
	WhetherItIsAListedCompany string `gorm:"column:whether_it_is_a_listed_company" json:"whetherItIsAListedCompany"` // 是否上市公司
	CityCode string `gorm:"column:city_code" json:"cityCode"` // 所属城市(省市区)
	FinancialSituationInThePas string `gorm:"column:financial_situation_in_the_pas" json:"financialSituationInThePas"` // 近三年财务情况
	ActualControllerSituation string `gorm:"column:actual_controller_situation" json:"actualControllerSituation"` // 实际控制人情况
	RequirementsForAcquirers string `gorm:"column:requirements_for_acquirers" json:"requirementsForAcquirers"` // 对收购方要求
	TransferConsiderationTransac float64 `gorm:"column:transfer_consideration_transac" json:"transferConsiderationTransac"` // 转让对价（交易金额）
	Seller string `gorm:"column:seller" json:"seller"` // 出售方
	NeedFurtherUnderstandingOf string `gorm:"column:need_further_understanding_of" json:"needFurtherUnderstandingOf"` // 需要进一步了解问题
	ProjectMatchmaker int64 `gorm:"column:project_matchmaker" json:"projectMatchmaker"` // 项目对接人
	Country string `gorm:"column:country" json:"country"` // 国家
	ListingPlate string `gorm:"column:listing_plate" json:"listingPlate"` // 上市板块
	LatestValuation float64 `gorm:"column:latest_valuation" json:"latestValuation"` // 最新估值
	ReasonForSelling string `gorm:"column:reason_for_selling" json:"reasonForSelling"` // 出售原因
	NetProfitLastYear float64 `gorm:"column:net_profit_last_year" json:"netProfitLastYear"` // 去年净利润
	CanWePromiseToImplementPr string `gorm:"column:can_we_promise_to_implement_pr" json:"canWePromiseToImplementPr"` // 能否承诺落地产能或总部搬迁
	PerformanceCommitment string `gorm:"column:performance_commitment" json:"performanceCommitment"` // 业绩承诺
	MainBusinessMainProductsAn string `gorm:"column:main_business_main_products_an" json:"mainBusinessMainProductsAn"` // 主营业务、主要产品及客户情况
	ProjectIntroduction string `gorm:"column:project_introduction" json:"projectIntroduction"` // 项目介绍
	CityDockUserPhone string `gorm:"column:city_dock_user_phone" json:"cityDockUserPhone"` // 市级服务专员联系方式
	ProjectClueContactPerson string `gorm:"column:project_clue_contact_person" json:"projectClueContactPerson"` // 项目线索联系人
	CityDockUserName string `gorm:"column:city_dock_user_name" json:"cityDockUserName"` // 市级服务专员姓名
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64 `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"` // 修改时间
	UpdatedBy int64 `gorm:"column:updated_by" json:"updatedBy"` // 修改人
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// LeadsList biz_leads 列表
type LeadsList []Leads

// NewLeads 创建biz_leads实例
func NewLeads() *Leads {
	return &Leads{}
}

// NewLeadsList 创建biz_leads列表实例
func NewLeadsList() *LeadsList {
	return &LeadsList{}
}

// TableName 指定表名
func (Leads) TableName() string {
	return "biz_leads"
}

// GetByID 根据ID获取biz_leads
func (m *Leads) GetByID(c context.Context, id int64) error {
	return app.DB().WithContext(c).First(m, id).Error
}

// Create 创建biz_leads记录
func (m *Leads) Create(c context.Context) error {
	return app.DB().WithContext(c).Create(m).Error
}

// Update 更新biz_leads记录
func (m *Leads) Update(c context.Context) error {
	return app.DB().WithContext(c).Save(m).Error
}

// Delete 软删除biz_leads记录
func (m *Leads) Delete(c context.Context) error {
	return app.DB().WithContext(c).Delete(m).Error
}

// IsEmpty 检查模型是否为空
func (m *Leads) IsEmpty() bool {
	return m == nil || m.Id == 0
}

// Find 查询biz_leads列表
func (l *LeadsList) Find(c context.Context, funcs ...func(*gorm.DB) *gorm.DB) error {
	return app.DB().WithContext(c).Model(&Leads{}).Scopes(funcs...).Find(l).Error
}

// GetTotal 获取biz_leads总数
func (l *LeadsList) GetTotal(c context.Context, query ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := app.DB().WithContext(c).Model(&Leads{}).Scopes(query...).Count(&count).Error
	return count, err
}