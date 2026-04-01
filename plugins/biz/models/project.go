package models

import (
	"context"
	"time"
	"gin-fast/app/global/app"
	"gorm.io/gorm"
)

// Project biz_project 模型结构体
type Project struct {
	Id int64 `gorm:"column:id;primaryKey;not null;autoIncrement" json:"id"` // 主键
	TenantId uint `gorm:"column:tenant_id;default:0;index" json:"tenantId"` // 租户ID字段
	ProjectNo string `gorm:"column:project_no;index" json:"projectNo"` // 项目编号
	Name string `gorm:"column:name" json:"name"` // 客户名称
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	CreatedBy int64 `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"` // 修改时间
	UpdatedBy int64 `gorm:"column:updated_by" json:"updatedBy"` // 修改人
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
	Source string `gorm:"column:source" json:"source"` // 客户来源
	LastFlowTime *time.Time `gorm:"column:last_flow_time" json:"lastFlowTime"` // 最近跟进时间
	TransactionStatus string `gorm:"column:transaction_status" json:"transactionStatus"` // 成交状态
	RecycleTime *time.Time `gorm:"column:recycle_time" json:"recycleTime"` // 退回时间
	DistributionStatus string `gorm:"column:distribution_status" json:"distributionStatus"` // 分配状态
	BaseCurrency float64 `gorm:"column:base_currency" json:"baseCurrency"` // BaseCurrency
	ExchangeRate string `gorm:"column:exchange_rate" json:"exchangeRate"` // 汇率
	RecycleReason string `gorm:"column:recycle_reason" json:"recycleReason"` // 退回/回收原因
	PoolId int64 `gorm:"column:pool_id" json:"poolId"` // PoolId
	Code string `gorm:"column:code" json:"code"` // 客户编号
	GetTime *time.Time `gorm:"column:get_time" json:"getTime"` // 获取时间
	DistributionTime *time.Time `gorm:"column:distribution_time" json:"distributionTime"` // 领取/分配时间
	HandoverObjects string `gorm:"column:handover_objects" json:"handoverObjects"` // 同时移交客户关联的业务对象
	HandoverRetainOwner int `gorm:"column:handover_retain_owner" json:"handoverRetainOwner"` // 移交保留归属人在团队
	OtherHandoverRetainOwner int `gorm:"column:other_handover_retain_owner" json:"otherHandoverRetainOwner"` // 其他对象移交保留归属人在团队
	RecycleStartTime *time.Time `gorm:"column:recycle_start_time" json:"recycleStartTime"` // 回收计算起始时间
	AuditCode string `gorm:"column:audit_code" json:"auditCode"` // 审批编号
	AuditStartTime *time.Time `gorm:"column:audit_start_time" json:"auditStartTime"` // 审批发起时间
	AuditEndTime *time.Time `gorm:"column:audit_end_time" json:"auditEndTime"` // 审批结束时间
	AuditStatus string `gorm:"column:audit_status;default:0" json:"auditStatus"` // 审批状态
	AuditMsg string `gorm:"column:audit_msg" json:"auditMsg"` // 审批留言
	ApplyTime *time.Time `gorm:"column:apply_time" json:"applyTime"` // 申请时间
	OriginDataOwnerId int64 `gorm:"column:origin_data_owner_id" json:"originDataOwnerId"` // 原归属人
	AuditStartUser int64 `gorm:"column:audit_start_user" json:"auditStartUser"` // 审批发起人
	FirstGetFromPool int `gorm:"column:first_get_from_pool;default:0" json:"firstGetFromPool"` // 是否为第一次从池里获取
	ContactUser int64 `gorm:"column:contact_user" json:"contactUser"` // 联系人
	Remark string `gorm:"column:remark" json:"remark"` // 备注
	GenerationDate *time.Time `gorm:"column:generation_date" json:"generationDate"` // 项目生成时间
	FixedInvestmentRemark string `gorm:"column:fixed_investment_remark" json:"fixedInvestmentRemark"` // 固投备注
	SourceOfInvestment string `gorm:"column:source_of_investment" json:"sourceOfInvestment"` // 投资来源地
	TotalInvestmentAmount float64 `gorm:"column:total_investment_amount" json:"totalInvestmentAmount"` // 总投资额（亿）
	FixedAssetsInvestment float64 `gorm:"column:fixed_assets_investment" json:"fixedAssetsInvestment"` // 固定资产投资额（亿）
	InvestmentArea string `gorm:"column:investment_area" json:"investmentArea"` // 投资区域
	MunicipalServiceTel string `gorm:"column:municipal_service_tel" json:"municipalServiceTel"` // 市级服务专员联系方式
	ProjectLabels string `gorm:"column:project_labels" json:"projectLabels"` // 项目标签
	ExpectedSigningTime *time.Time `gorm:"column:expected_signing_time" json:"expectedSigningTime"` // 预计签约时间
	InvestmentType string `gorm:"column:investment_type" json:"investmentType"` // 投资类型
	BelongIndustryChain string `gorm:"column:belong_industry_chain" json:"belongIndustryChain"` // 所属产业链
	Industry string `gorm:"column:industry" json:"industry"` // 行业
	IndustrySector string `gorm:"column:industry_sector" json:"industrySector"` // 行业领域
	IsForeignInvestment int `gorm:"column:is_foreign_investment;default:0" json:"isForeignInvestment"` // 是否外资
	IsAcquireLand int `gorm:"column:is_acquire_land;default:0" json:"isAcquireLand"` // 是否拿地
	PromotionTypes string `gorm:"column:promotion_types" json:"promotionTypes"` // 推进形式
	ProjectContent string `gorm:"column:project_content" json:"projectContent"` // 项目内容
	AdvancePlan string `gorm:"column:advance_plan" json:"advancePlan"` // 推进计划
	JudgingMaterials string `gorm:"column:judging_materials" json:"judgingMaterials"` // 研判材料
	ListOfClueSource string `gorm:"column:list_of_clue_source" json:"listOfClueSource"` // 清单线索来源
	LeaderOfMunicipalTrans string `gorm:"column:leader_of_municipal_trans" json:"leaderOfMunicipalTrans"` // 市交办领导
	LeaderOfBureauTrans string `gorm:"column:leader_of_bureau_trans" json:"leaderOfBureauTrans"` // 局交办领导
	SourceRemark string `gorm:"column:source_remark" json:"sourceRemark"` // 来源备注
	Dept string `gorm:"column:dept" json:"dept"` // 部门
	District string `gorm:"column:district" json:"district"` // 区县
	PlatformCompany string `gorm:"column:platform_company" json:"platformCompany"` // 平台公司
	IsCityPlanPartner int `gorm:"column:is_city_plan_partner;default:0" json:"isCityPlanPartner"` // 是否来自创投城市计划生态合伙人
	EnterpriseType string `gorm:"column:enterprise_type" json:"enterpriseType"` // 企业类型
	InvestorProfile string `gorm:"column:investor_profile" json:"investorProfile"` // 投资方简介
	ProjectContact string `gorm:"column:project_contact" json:"projectContact"` // 项目联系人
	EnterpriseContactInfo string `gorm:"column:enterprise_contact_info" json:"enterpriseContactInfo"` // 企业联系方式
	Referrer string `gorm:"column:referrer" json:"referrer"` // 引荐人
	ReferrerContactInfo string `gorm:"column:referrer_contact_info" json:"referrerContactInfo"` // 引荐人联系方式
	ProjectPhase string `gorm:"column:project_phase" json:"projectPhase"` // 在谈节点
	PhaseId string `gorm:"column:phase_id" json:"phaseId"` // 项目阶段
	BpAttachments string `gorm:"column:bp_attachments" json:"bpAttachments"` // BP附件库
	Stress string `gorm:"column:stress" json:"stress"` // 重要性
	IndustryType string `gorm:"column:industry_type" json:"industryType"` // 行业分类
	City string `gorm:"column:city" json:"city"` // 所在城市
	InvestCompany int64 `gorm:"column:invest_company" json:"investCompany"` // 投资方名称
	ItemProvideDepartment int64 `gorm:"column:item_provide_department" json:"itemProvideDepartment"` // 项目提供部门/机构
	ItemRecommender string `gorm:"column:item_recommender" json:"itemRecommender"` // 项目推荐方
	ItemRecommendTime *time.Time `gorm:"column:item_recommend_time" json:"itemRecommendTime"` // 项目推荐时间
	LeadDepartment int64 `gorm:"column:lead_department" json:"leadDepartment"` // 牵头部门
	DistributeTask int `gorm:"column:distribute_task;default:0" json:"distributeTask"` // 同时下发行业研判任务
	ShareItem int `gorm:"column:share_item;default:0" json:"shareItem"` // 同时与行业牵头人或产业专班共享项目
	LandContent string `gorm:"column:land_content" json:"landContent"` // 落地内容
	EnterpriseAppeal string `gorm:"column:enterprise_appeal" json:"enterpriseAppeal"` // 企业诉求
	Carrier string `gorm:"column:carrier" json:"carrier"` // 载体
	ProductionFactors string `gorm:"column:production_factors" json:"productionFactors"` // 生产要素
	Capital string `gorm:"column:capital" json:"capital"` // 融资/资本
	Scene string `gorm:"column:scene" json:"scene"` // 场景/市场
	Policy string `gorm:"column:policy" json:"policy"` // 政策
	InvestmentPromoteUser int64 `gorm:"column:investment_promote_user" json:"investmentPromoteUser"` // 投促推进人
	DockDistrict string `gorm:"column:dock_district" json:"dockDistrict"` // 对接区县
	ItemProgress string `gorm:"column:item_progress" json:"itemProgress"` // 项目进展
	CapitalFactor int `gorm:"column:capital_factor;default:0" json:"capitalFactor"` // 资本要素
	CapitalAssessmentContent string `gorm:"column:capital_assessment_content" json:"capitalAssessmentContent"` // 研判内容
	CapitalAssessmentView string `gorm:"column:capital_assessment_view" json:"capitalAssessmentView"` // 资本要素-研判意见
	SceneFactor int `gorm:"column:scene_factor;default:0" json:"sceneFactor"` // 场景要素
	SceneAssessmentContent string `gorm:"column:scene_assessment_content" json:"sceneAssessmentContent"` // 场景要素-研判内容
	SceneAssessmentView string `gorm:"column:scene_assessment_view" json:"sceneAssessmentView"` // 场景要素-研判意见
	AreaFactor int `gorm:"column:area_factor;default:0" json:"areaFactor"` // 场地要素
	AreaAssessmentContent string `gorm:"column:area_assessment_content" json:"areaAssessmentContent"` // 场地要素-研判内容
	AreaAssessmentView string `gorm:"column:area_assessment_view" json:"areaAssessmentView"` // 场地要素-研判意见
	PolicyFactor int `gorm:"column:policy_factor;default:0" json:"policyFactor"` // 政策要素
	PolicyAssessmentContent string `gorm:"column:policy_assessment_content" json:"policyAssessmentContent"` // 政策要素-研判内容
	PolicyAssessmentView string `gorm:"column:policy_assessment_view" json:"policyAssessmentView"` // 政策要素-研判意见
	ItemOwnerName int64 `gorm:"column:item_owner_name" json:"itemOwnerName"` // 项目归属人员
	RelatedRask int64 `gorm:"column:related_rask" json:"relatedRask"` // 关联任务
	RelatedLead int64 `gorm:"column:related_lead" json:"relatedLead"` // 关联线索
	WinRate float64 `gorm:"column:win_rate" json:"winRate"` // 赢率
	ProjectContactInfo string `gorm:"column:project_contact_info" json:"projectContactInfo"` // 项目联系方式
	ProjectType string `gorm:"column:project_type" json:"projectType"` // 项目类型
	AssessmentContent string `gorm:"column:assessment_content" json:"assessmentContent"` // 研判内容
	AssessmentView string `gorm:"column:assessment_view" json:"assessmentView"` // 研判意见
	CoreAdvantages string `gorm:"column:core_advantages" json:"coreAdvantages"` // 核心优势
	IndustryStatus string `gorm:"column:industry_status" json:"industryStatus"` // 行业地位
	FinancialSituation string `gorm:"column:financial_situation" json:"financialSituation"` // 财务情况
	Speed string `gorm:"column:speed" json:"speed"` // 进度
	InvestRound string `gorm:"column:invest_round" json:"investRound"` // 拟投资轮次
	NeedIndustry int `gorm:"column:need_industry;default:0" json:"needIndustry"` // 是否需要产业投资人
	NeedFinance int `gorm:"column:need_finance;default:0" json:"needFinance"` // 是否需要财务投资人
	IndustryTypeCn string `gorm:"column:industry_type_cn" json:"industryTypeCn"` // 行业分类-中文
	BureauProjectId string `gorm:"column:bureau_project_id" json:"bureauProjectId"` // 市局项目id
	InvestCompanyName string `gorm:"column:invest_company_name" json:"investCompanyName"` // 投资方名称（与市局对接预留字段）
	CityCode string `gorm:"column:city_code" json:"cityCode"` // 所属城市(省市区)
	CapitalInvestInfo string `gorm:"column:capital_invest_info" json:"capitalInvestInfo"` // 资本要素的投资信息（json）
	Tag string `gorm:"column:tag" json:"tag"` // 数据标签
	FinancingAmount float64 `gorm:"column:financing_amount" json:"financingAmount"` // 融资金额(亿元)
	BuildingInvestmentArea float64 `gorm:"column:building_investment_area" json:"buildingInvestmentArea"` // 楼宇招商面积
	GuaranteeMethods string `gorm:"column:guarantee_methods" json:"guaranteeMethods"` // 担保方式
	Period string `gorm:"column:period" json:"period"` // 期限
	FactoryRentalArea float64 `gorm:"column:factory_rental_area" json:"factoryRentalArea"` // 厂房租用面积
	FactoryFloorHeightBegins float64 `gorm:"column:factory_floor_height_begins" json:"factoryFloorHeightBegins"` // 厂房层高开始
	Quota string `gorm:"column:quota" json:"quota"` // 额度
	ParkType string `gorm:"column:park_type" json:"parkType"` // 园区类型
	MinimumRentalValue float64 `gorm:"column:minimum_rental_value" json:"minimumRentalValue"` // 楼宇租金最小值
	MaximumFactoryRent float64 `gorm:"column:maximum_factory_rent" json:"maximumFactoryRent"` // 厂房租金最大值
	WhetherThereIsANeedForDe int `gorm:"column:whether_there_is_a_need_for_de" json:"whetherThereIsANeedForDe"` // 是否有债权融资需求
	EndOfFactoryFloorHeight float64 `gorm:"column:end_of_factory_floor_height" json:"endOfFactoryFloorHeight"` // 厂房层高结束
	MaximumRentalValue float64 `gorm:"column:maximum_rental_value" json:"maximumRentalValue"` // 楼宇租金最大值
	MinimumFactoryRent float64 `gorm:"column:minimum_factory_rent" json:"minimumFactoryRent"` // 厂房租金最小值
	LandUse string `gorm:"column:land_use" json:"landUse"` // 土地用途
	ParkLevel string `gorm:"column:park_level" json:"parkLevel"` // 园区等级
	TypeOfBuildingRent string `gorm:"column:type_of_building_rent" json:"typeOfBuildingRent"` // 楼宇租金类型
	LandAreaRequired float64 `gorm:"column:land_area_required" json:"landAreaRequired"` // 需地块面积
	FactoryRentType string `gorm:"column:factory_rent_type" json:"factoryRentType"` // 厂房租金类型
	InterestRate float64 `gorm:"column:interest_rate" json:"interestRate"` // 利率
	ClueResearcher int64 `gorm:"column:clue_researcher" json:"clueResearcher"` // 线索研判人
	AssessmentTeam string `gorm:"column:assessment_team" json:"assessmentTeam"` // 研判团队
	EquityFinancingAmount float64 `gorm:"column:equity_financing_amount" json:"equityFinancingAmount"` // 股权融资金额
	BondFinancingAmount float64 `gorm:"column:bond_financing_amount" json:"bondFinancingAmount"` // 债券融资金额
	SimilarFinancialAmount float64 `gorm:"column:similar_financial_amount" json:"similarFinancialAmount"` // 类金融融资金额
	BureauProjectProgress string `gorm:"column:bureau_project_progress" json:"bureauProjectProgress"` // 市局项目进度
	Country string `gorm:"column:country" json:"country"` // 国家
	IssueId string `gorm:"column:issue_id" json:"issueId"` // 市局项目轮次id
	ProjectIntroduction string `gorm:"column:project_introduction" json:"projectIntroduction"` // 项目介绍
	IssueType string `gorm:"column:issue_type" json:"issueType"` // 问题类型
	AreaDockContent string `gorm:"column:area_dock_content" json:"areaDockContent"` // 区级对接情况
	MoneyValue string `gorm:"column:money_value" json:"moneyValue"` // 问题金额
	IssueContent string `gorm:"column:issue_content" json:"issueContent"` // 问题描述
	AreaAdvice string `gorm:"column:area_advice" json:"areaAdvice"` // 区级建议
}

// ProjectList biz_project 列表
type ProjectList []Project

// NewProject 创建biz_project实例
func NewProject() *Project {
	return &Project{}
}

// NewProjectList 创建biz_project列表实例
func NewProjectList() *ProjectList {
	return &ProjectList{}
}

// TableName 指定表名
func (Project) TableName() string {
	return "biz_project"
}

// GetByID 根据ID获取biz_project
func (m *Project) GetByID(c context.Context, id int64) error {
	return app.DB().WithContext(c).First(m, id).Error
}

// Create 创建biz_project记录
func (m *Project) Create(c context.Context) error {
	return app.DB().WithContext(c).Create(m).Error
}

// Update 更新biz_project记录
func (m *Project) Update(c context.Context) error {
	return app.DB().WithContext(c).Save(m).Error
}

// Delete 软删除biz_project记录
func (m *Project) Delete(c context.Context) error {
	return app.DB().WithContext(c).Delete(m).Error
}

// IsEmpty 检查模型是否为空
func (m *Project) IsEmpty() bool {
	return m == nil || m.Id == 0
}

// Find 查询biz_project列表
func (l *ProjectList) Find(c context.Context, funcs ...func(*gorm.DB) *gorm.DB) error {
	return app.DB().WithContext(c).Model(&Project{}).Scopes(funcs...).Find(l).Error
}

// GetTotal 获取biz_project总数
func (l *ProjectList) GetTotal(c context.Context, query ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := app.DB().WithContext(c).Model(&Project{}).Scopes(query...).Count(&count).Error
	return count, err
}