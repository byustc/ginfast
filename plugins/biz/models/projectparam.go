package models

import (
	"time"
	"gin-fast/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProjectListRequest biz_project列表请求参数
type ProjectListRequest struct {
	models.BasePaging
	models.Validator
	Id *int64 `form:"id"` // 主键
	ProjectNo *string `form:"projectNo"` // 项目编号
	Name *string `form:"name"` // 客户名称
	UpdatedBy *int64 `form:"updatedBy"` // 修改人
	Source *string `form:"source"` // 客户来源
	LastFlowTime *time.Time `form:"lastFlowTime"` // 最近跟进时间
	TransactionStatus *string `form:"transactionStatus"` // 成交状态
	RecycleTime *time.Time `form:"recycleTime"` // 退回时间
	DistributionStatus *string `form:"distributionStatus"` // 分配状态
	BaseCurrency *float64 `form:"baseCurrency"` // BaseCurrency
	ExchangeRate *string `form:"exchangeRate"` // 汇率
	RecycleReason *string `form:"recycleReason"` // 退回/回收原因
	PoolId *int64 `form:"poolId"` // PoolId
	Code *string `form:"code"` // 客户编号
	GetTime *time.Time `form:"getTime"` // 获取时间
	DistributionTime *time.Time `form:"distributionTime"` // 领取/分配时间
	HandoverObjects *string `form:"handoverObjects"` // 同时移交客户关联的业务对象
	HandoverRetainOwner *int `form:"handoverRetainOwner"` // 移交保留归属人在团队
	OtherHandoverRetainOwner *int `form:"otherHandoverRetainOwner"` // 其他对象移交保留归属人在团队
	RecycleStartTime *time.Time `form:"recycleStartTime"` // 回收计算起始时间
	AuditCode *string `form:"auditCode"` // 审批编号
	AuditStartTime *time.Time `form:"auditStartTime"` // 审批发起时间
	AuditEndTime *time.Time `form:"auditEndTime"` // 审批结束时间
	AuditStatus *string `form:"auditStatus"` // 审批状态
	AuditMsg *string `form:"auditMsg"` // 审批留言
	ApplyTime *time.Time `form:"applyTime"` // 申请时间
	OriginDataOwnerId *int64 `form:"originDataOwnerId"` // 原归属人
	AuditStartUser *int64 `form:"auditStartUser"` // 审批发起人
	FirstGetFromPool *int `form:"firstGetFromPool"` // 是否为第一次从池里获取
	ContactUser *int64 `form:"contactUser"` // 联系人
	Remark *string `form:"remark"` // 备注
	GenerationDate *time.Time `form:"generationDate"` // 项目生成时间
	FixedInvestmentRemark *string `form:"fixedInvestmentRemark"` // 固投备注
	SourceOfInvestment *string `form:"sourceOfInvestment"` // 投资来源地
	TotalInvestmentAmount *float64 `form:"totalInvestmentAmount"` // 总投资额（亿）
	FixedAssetsInvestment *float64 `form:"fixedAssetsInvestment"` // 固定资产投资额（亿）
	InvestmentArea *string `form:"investmentArea"` // 投资区域
	MunicipalServiceTel *string `form:"municipalServiceTel"` // 市级服务专员联系方式
	ProjectLabels *string `form:"projectLabels"` // 项目标签
	ExpectedSigningTime *time.Time `form:"expectedSigningTime"` // 预计签约时间
	InvestmentType *string `form:"investmentType"` // 投资类型
	BelongIndustryChain *string `form:"belongIndustryChain"` // 所属产业链
	Industry *string `form:"industry"` // 行业
	IndustrySector *string `form:"industrySector"` // 行业领域
	IsForeignInvestment *int `form:"isForeignInvestment"` // 是否外资
	IsAcquireLand *int `form:"isAcquireLand"` // 是否拿地
	PromotionTypes *string `form:"promotionTypes"` // 推进形式
	ProjectContent *string `form:"projectContent"` // 项目内容
	AdvancePlan *string `form:"advancePlan"` // 推进计划
	JudgingMaterials *string `form:"judgingMaterials"` // 研判材料
	ListOfClueSource *string `form:"listOfClueSource"` // 清单线索来源
	LeaderOfMunicipalTrans *string `form:"leaderOfMunicipalTrans"` // 市交办领导
	LeaderOfBureauTrans *string `form:"leaderOfBureauTrans"` // 局交办领导
	SourceRemark *string `form:"sourceRemark"` // 来源备注
	Dept *string `form:"dept"` // 部门
	District *string `form:"district"` // 区县
	PlatformCompany *string `form:"platformCompany"` // 平台公司
	IsCityPlanPartner *int `form:"isCityPlanPartner"` // 是否来自创投城市计划生态合伙人
	EnterpriseType *string `form:"enterpriseType"` // 企业类型
	InvestorProfile *string `form:"investorProfile"` // 投资方简介
	ProjectContact *string `form:"projectContact"` // 项目联系人
	EnterpriseContactInfo *string `form:"enterpriseContactInfo"` // 企业联系方式
	Referrer *string `form:"referrer"` // 引荐人
	ReferrerContactInfo *string `form:"referrerContactInfo"` // 引荐人联系方式
	ProjectPhase *string `form:"projectPhase"` // 在谈节点
	PhaseId *string `form:"phaseId"` // 项目阶段
	BpAttachments *string `form:"bpAttachments"` // BP附件库
	Stress *string `form:"stress"` // 重要性
	IndustryType *string `form:"industryType"` // 行业分类
	City *string `form:"city"` // 所在城市
	InvestCompany *int64 `form:"investCompany"` // 投资方名称
	ItemProvideDepartment *int64 `form:"itemProvideDepartment"` // 项目提供部门/机构
	ItemRecommender *string `form:"itemRecommender"` // 项目推荐方
	ItemRecommendTime *time.Time `form:"itemRecommendTime"` // 项目推荐时间
	LeadDepartment *int64 `form:"leadDepartment"` // 牵头部门
	DistributeTask *int `form:"distributeTask"` // 同时下发行业研判任务
	ShareItem *int `form:"shareItem"` // 同时与行业牵头人或产业专班共享项目
	LandContent *string `form:"landContent"` // 落地内容
	EnterpriseAppeal *string `form:"enterpriseAppeal"` // 企业诉求
	Carrier *string `form:"carrier"` // 载体
	ProductionFactors *string `form:"productionFactors"` // 生产要素
	Capital *string `form:"capital"` // 融资/资本
	Scene *string `form:"scene"` // 场景/市场
	Policy *string `form:"policy"` // 政策
	InvestmentPromoteUser *int64 `form:"investmentPromoteUser"` // 投促推进人
	DockDistrict *string `form:"dockDistrict"` // 对接区县
	ItemProgress *string `form:"itemProgress"` // 项目进展
	CapitalFactor *int `form:"capitalFactor"` // 资本要素
	CapitalAssessmentContent *string `form:"capitalAssessmentContent"` // 研判内容
	CapitalAssessmentView *string `form:"capitalAssessmentView"` // 资本要素-研判意见
	SceneFactor *int `form:"sceneFactor"` // 场景要素
	SceneAssessmentContent *string `form:"sceneAssessmentContent"` // 场景要素-研判内容
	SceneAssessmentView *string `form:"sceneAssessmentView"` // 场景要素-研判意见
	AreaFactor *int `form:"areaFactor"` // 场地要素
	AreaAssessmentContent *string `form:"areaAssessmentContent"` // 场地要素-研判内容
	AreaAssessmentView *string `form:"areaAssessmentView"` // 场地要素-研判意见
	PolicyFactor *int `form:"policyFactor"` // 政策要素
	PolicyAssessmentContent *string `form:"policyAssessmentContent"` // 政策要素-研判内容
	PolicyAssessmentView *string `form:"policyAssessmentView"` // 政策要素-研判意见
	ItemOwnerName *int64 `form:"itemOwnerName"` // 项目归属人员
	RelatedRask *int64 `form:"relatedRask"` // 关联任务
	RelatedLead *int64 `form:"relatedLead"` // 关联线索
	WinRate *float64 `form:"winRate"` // 赢率
	ProjectContactInfo *string `form:"projectContactInfo"` // 项目联系方式
	ProjectType *string `form:"projectType"` // 项目类型
	AssessmentContent *string `form:"assessmentContent"` // 研判内容
	AssessmentView *string `form:"assessmentView"` // 研判意见
	CoreAdvantages *string `form:"coreAdvantages"` // 核心优势
	IndustryStatus *string `form:"industryStatus"` // 行业地位
	FinancialSituation *string `form:"financialSituation"` // 财务情况
	Speed *string `form:"speed"` // 进度
	InvestRound *string `form:"investRound"` // 拟投资轮次
	NeedIndustry *int `form:"needIndustry"` // 是否需要产业投资人
	NeedFinance *int `form:"needFinance"` // 是否需要财务投资人
	IndustryTypeCn *string `form:"industryTypeCn"` // 行业分类-中文
	BureauProjectId *string `form:"bureauProjectId"` // 市局项目id
	InvestCompanyName *string `form:"investCompanyName"` // 投资方名称（与市局对接预留字段）
	CityCode *string `form:"cityCode"` // 所属城市(省市区)
	CapitalInvestInfo *string `form:"capitalInvestInfo"` // 资本要素的投资信息（json）
	Tag *string `form:"tag"` // 数据标签
	FinancingAmount *float64 `form:"financingAmount"` // 融资金额(亿元)
	BuildingInvestmentArea *float64 `form:"buildingInvestmentArea"` // 楼宇招商面积
	GuaranteeMethods *string `form:"guaranteeMethods"` // 担保方式
	Period *string `form:"period"` // 期限
	FactoryRentalArea *float64 `form:"factoryRentalArea"` // 厂房租用面积
	FactoryFloorHeightBegins *float64 `form:"factoryFloorHeightBegins"` // 厂房层高开始
	Quota *string `form:"quota"` // 额度
	ParkType *string `form:"parkType"` // 园区类型
	MinimumRentalValue *float64 `form:"minimumRentalValue"` // 楼宇租金最小值
	MaximumFactoryRent *float64 `form:"maximumFactoryRent"` // 厂房租金最大值
	WhetherThereIsANeedForDe *int `form:"whetherThereIsANeedForDe"` // 是否有债权融资需求
	EndOfFactoryFloorHeight *float64 `form:"endOfFactoryFloorHeight"` // 厂房层高结束
	MaximumRentalValue *float64 `form:"maximumRentalValue"` // 楼宇租金最大值
	MinimumFactoryRent *float64 `form:"minimumFactoryRent"` // 厂房租金最小值
	LandUse *string `form:"landUse"` // 土地用途
	ParkLevel *string `form:"parkLevel"` // 园区等级
	TypeOfBuildingRent *string `form:"typeOfBuildingRent"` // 楼宇租金类型
	LandAreaRequired *float64 `form:"landAreaRequired"` // 需地块面积
	FactoryRentType *string `form:"factoryRentType"` // 厂房租金类型
	InterestRate *float64 `form:"interestRate"` // 利率
	ClueResearcher *int64 `form:"clueResearcher"` // 线索研判人
	AssessmentTeam *string `form:"assessmentTeam"` // 研判团队
	EquityFinancingAmount *float64 `form:"equityFinancingAmount"` // 股权融资金额
	BondFinancingAmount *float64 `form:"bondFinancingAmount"` // 债券融资金额
	SimilarFinancialAmount *float64 `form:"similarFinancialAmount"` // 类金融融资金额
	BureauProjectProgress *string `form:"bureauProjectProgress"` // 市局项目进度
	Country *string `form:"country"` // 国家
	IssueId *string `form:"issueId"` // 市局项目轮次id
	ProjectIntroduction *string `form:"projectIntroduction"` // 项目介绍
	IssueType *string `form:"issueType"` // 问题类型
	AreaDockContent *string `form:"areaDockContent"` // 区级对接情况
	MoneyValue *string `form:"moneyValue"` // 问题金额
	IssueContent *string `form:"issueContent"` // 问题描述
	AreaAdvice *string `form:"areaAdvice"` // 区级建议
}

// Validate 验证请求参数
func (r *ProjectListRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// Handle 获取查询条件
func (r *ProjectListRequest) Handle() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
        if r.Id != nil {
            // 默认等于查询
            db = db.Where("id = ?", *r.Id)
        }
        if r.ProjectNo != nil {
            // 默认等于查询
            db = db.Where("project_no = ?", *r.ProjectNo)
        }
        if r.Name != nil {
            // 默认等于查询
            db = db.Where("name = ?", *r.Name)
        }
        if r.UpdatedBy != nil {
            // 默认等于查询
            db = db.Where("updated_by = ?", *r.UpdatedBy)
        }
        if r.Source != nil {
            // 默认等于查询
            db = db.Where("source = ?", *r.Source)
        }
        if r.LastFlowTime != nil {
            // 默认等于查询
            db = db.Where("last_flow_time = ?", *r.LastFlowTime)
        }
        if r.TransactionStatus != nil {
            // 默认等于查询
            db = db.Where("transaction_status = ?", *r.TransactionStatus)
        }
        if r.RecycleTime != nil {
            // 默认等于查询
            db = db.Where("recycle_time = ?", *r.RecycleTime)
        }
        if r.DistributionStatus != nil {
            // 默认等于查询
            db = db.Where("distribution_status = ?", *r.DistributionStatus)
        }
        if r.BaseCurrency != nil {
            // 默认等于查询
            db = db.Where("base_currency = ?", *r.BaseCurrency)
        }
        if r.ExchangeRate != nil {
            // 默认等于查询
            db = db.Where("exchange_rate = ?", *r.ExchangeRate)
        }
        if r.RecycleReason != nil {
            // 默认等于查询
            db = db.Where("recycle_reason = ?", *r.RecycleReason)
        }
        if r.PoolId != nil {
            // 默认等于查询
            db = db.Where("pool_id = ?", *r.PoolId)
        }
        if r.Code != nil {
            // 默认等于查询
            db = db.Where("code = ?", *r.Code)
        }
        if r.GetTime != nil {
            // 默认等于查询
            db = db.Where("get_time = ?", *r.GetTime)
        }
        if r.DistributionTime != nil {
            // 默认等于查询
            db = db.Where("distribution_time = ?", *r.DistributionTime)
        }
        if r.HandoverObjects != nil {
            // 默认等于查询
            db = db.Where("handover_objects = ?", *r.HandoverObjects)
        }
        if r.HandoverRetainOwner != nil {
            // 默认等于查询
            db = db.Where("handover_retain_owner = ?", *r.HandoverRetainOwner)
        }
        if r.OtherHandoverRetainOwner != nil {
            // 默认等于查询
            db = db.Where("other_handover_retain_owner = ?", *r.OtherHandoverRetainOwner)
        }
        if r.RecycleStartTime != nil {
            // 默认等于查询
            db = db.Where("recycle_start_time = ?", *r.RecycleStartTime)
        }
        if r.AuditCode != nil {
            // 默认等于查询
            db = db.Where("audit_code = ?", *r.AuditCode)
        }
        if r.AuditStartTime != nil {
            // 默认等于查询
            db = db.Where("audit_start_time = ?", *r.AuditStartTime)
        }
        if r.AuditEndTime != nil {
            // 默认等于查询
            db = db.Where("audit_end_time = ?", *r.AuditEndTime)
        }
        if r.AuditStatus != nil {
            // 默认等于查询
            db = db.Where("audit_status = ?", *r.AuditStatus)
        }
        if r.AuditMsg != nil {
            // 默认等于查询
            db = db.Where("audit_msg = ?", *r.AuditMsg)
        }
        if r.ApplyTime != nil {
            // 默认等于查询
            db = db.Where("apply_time = ?", *r.ApplyTime)
        }
        if r.OriginDataOwnerId != nil {
            // 默认等于查询
            db = db.Where("origin_data_owner_id = ?", *r.OriginDataOwnerId)
        }
        if r.AuditStartUser != nil {
            // 默认等于查询
            db = db.Where("audit_start_user = ?", *r.AuditStartUser)
        }
        if r.FirstGetFromPool != nil {
            // 默认等于查询
            db = db.Where("first_get_from_pool = ?", *r.FirstGetFromPool)
        }
        if r.ContactUser != nil {
            // 默认等于查询
            db = db.Where("contact_user = ?", *r.ContactUser)
        }
        if r.Remark != nil {
            // 默认等于查询
            db = db.Where("remark = ?", *r.Remark)
        }
        if r.GenerationDate != nil {
            // 默认等于查询
            db = db.Where("generation_date = ?", *r.GenerationDate)
        }
        if r.FixedInvestmentRemark != nil {
            // 默认等于查询
            db = db.Where("fixed_investment_remark = ?", *r.FixedInvestmentRemark)
        }
        if r.SourceOfInvestment != nil {
            // 默认等于查询
            db = db.Where("source_of_investment = ?", *r.SourceOfInvestment)
        }
        if r.TotalInvestmentAmount != nil {
            // 默认等于查询
            db = db.Where("total_investment_amount = ?", *r.TotalInvestmentAmount)
        }
        if r.FixedAssetsInvestment != nil {
            // 默认等于查询
            db = db.Where("fixed_assets_investment = ?", *r.FixedAssetsInvestment)
        }
        if r.InvestmentArea != nil {
            // 默认等于查询
            db = db.Where("investment_area = ?", *r.InvestmentArea)
        }
        if r.MunicipalServiceTel != nil {
            // 默认等于查询
            db = db.Where("municipal_service_tel = ?", *r.MunicipalServiceTel)
        }
        if r.ProjectLabels != nil {
            // 默认等于查询
            db = db.Where("project_labels = ?", *r.ProjectLabels)
        }
        if r.ExpectedSigningTime != nil {
            // 默认等于查询
            db = db.Where("expected_signing_time = ?", *r.ExpectedSigningTime)
        }
        if r.InvestmentType != nil {
            // 默认等于查询
            db = db.Where("investment_type = ?", *r.InvestmentType)
        }
        if r.BelongIndustryChain != nil {
            // 默认等于查询
            db = db.Where("belong_industry_chain = ?", *r.BelongIndustryChain)
        }
        if r.Industry != nil {
            // 默认等于查询
            db = db.Where("industry = ?", *r.Industry)
        }
        if r.IndustrySector != nil {
            // 默认等于查询
            db = db.Where("industry_sector = ?", *r.IndustrySector)
        }
        if r.IsForeignInvestment != nil {
            // 默认等于查询
            db = db.Where("is_foreign_investment = ?", *r.IsForeignInvestment)
        }
        if r.IsAcquireLand != nil {
            // 默认等于查询
            db = db.Where("is_acquire_land = ?", *r.IsAcquireLand)
        }
        if r.PromotionTypes != nil {
            // 默认等于查询
            db = db.Where("promotion_types = ?", *r.PromotionTypes)
        }
        if r.ProjectContent != nil {
            // 默认等于查询
            db = db.Where("project_content = ?", *r.ProjectContent)
        }
        if r.AdvancePlan != nil {
            // 默认等于查询
            db = db.Where("advance_plan = ?", *r.AdvancePlan)
        }
        if r.JudgingMaterials != nil {
            // 默认等于查询
            db = db.Where("judging_materials = ?", *r.JudgingMaterials)
        }
        if r.ListOfClueSource != nil {
            // 默认等于查询
            db = db.Where("list_of_clue_source = ?", *r.ListOfClueSource)
        }
        if r.LeaderOfMunicipalTrans != nil {
            // 默认等于查询
            db = db.Where("leader_of_municipal_trans = ?", *r.LeaderOfMunicipalTrans)
        }
        if r.LeaderOfBureauTrans != nil {
            // 默认等于查询
            db = db.Where("leader_of_bureau_trans = ?", *r.LeaderOfBureauTrans)
        }
        if r.SourceRemark != nil {
            // 默认等于查询
            db = db.Where("source_remark = ?", *r.SourceRemark)
        }
        if r.Dept != nil {
            // 默认等于查询
            db = db.Where("dept = ?", *r.Dept)
        }
        if r.District != nil {
            // 默认等于查询
            db = db.Where("district = ?", *r.District)
        }
        if r.PlatformCompany != nil {
            // 默认等于查询
            db = db.Where("platform_company = ?", *r.PlatformCompany)
        }
        if r.IsCityPlanPartner != nil {
            // 默认等于查询
            db = db.Where("is_city_plan_partner = ?", *r.IsCityPlanPartner)
        }
        if r.EnterpriseType != nil {
            // 默认等于查询
            db = db.Where("enterprise_type = ?", *r.EnterpriseType)
        }
        if r.InvestorProfile != nil {
            // 默认等于查询
            db = db.Where("investor_profile = ?", *r.InvestorProfile)
        }
        if r.ProjectContact != nil {
            // 默认等于查询
            db = db.Where("project_contact = ?", *r.ProjectContact)
        }
        if r.EnterpriseContactInfo != nil {
            // 默认等于查询
            db = db.Where("enterprise_contact_info = ?", *r.EnterpriseContactInfo)
        }
        if r.Referrer != nil {
            // 默认等于查询
            db = db.Where("referrer = ?", *r.Referrer)
        }
        if r.ReferrerContactInfo != nil {
            // 默认等于查询
            db = db.Where("referrer_contact_info = ?", *r.ReferrerContactInfo)
        }
        if r.ProjectPhase != nil {
            // 默认等于查询
            db = db.Where("project_phase = ?", *r.ProjectPhase)
        }
        if r.PhaseId != nil {
            // 默认等于查询
            db = db.Where("phase_id = ?", *r.PhaseId)
        }
        if r.BpAttachments != nil {
            // 默认等于查询
            db = db.Where("bp_attachments = ?", *r.BpAttachments)
        }
        if r.Stress != nil {
            // 默认等于查询
            db = db.Where("stress = ?", *r.Stress)
        }
        if r.IndustryType != nil {
            // 默认等于查询
            db = db.Where("industry_type = ?", *r.IndustryType)
        }
        if r.City != nil {
            // 默认等于查询
            db = db.Where("city = ?", *r.City)
        }
        if r.InvestCompany != nil {
            // 默认等于查询
            db = db.Where("invest_company = ?", *r.InvestCompany)
        }
        if r.ItemProvideDepartment != nil {
            // 默认等于查询
            db = db.Where("item_provide_department = ?", *r.ItemProvideDepartment)
        }
        if r.ItemRecommender != nil {
            // 默认等于查询
            db = db.Where("item_recommender = ?", *r.ItemRecommender)
        }
        if r.ItemRecommendTime != nil {
            // 默认等于查询
            db = db.Where("item_recommend_time = ?", *r.ItemRecommendTime)
        }
        if r.LeadDepartment != nil {
            // 默认等于查询
            db = db.Where("lead_department = ?", *r.LeadDepartment)
        }
        if r.DistributeTask != nil {
            // 默认等于查询
            db = db.Where("distribute_task = ?", *r.DistributeTask)
        }
        if r.ShareItem != nil {
            // 默认等于查询
            db = db.Where("share_item = ?", *r.ShareItem)
        }
        if r.LandContent != nil {
            // 默认等于查询
            db = db.Where("land_content = ?", *r.LandContent)
        }
        if r.EnterpriseAppeal != nil {
            // 默认等于查询
            db = db.Where("enterprise_appeal = ?", *r.EnterpriseAppeal)
        }
        if r.Carrier != nil {
            // 默认等于查询
            db = db.Where("carrier = ?", *r.Carrier)
        }
        if r.ProductionFactors != nil {
            // 默认等于查询
            db = db.Where("production_factors = ?", *r.ProductionFactors)
        }
        if r.Capital != nil {
            // 默认等于查询
            db = db.Where("capital = ?", *r.Capital)
        }
        if r.Scene != nil {
            // 默认等于查询
            db = db.Where("scene = ?", *r.Scene)
        }
        if r.Policy != nil {
            // 默认等于查询
            db = db.Where("policy = ?", *r.Policy)
        }
        if r.InvestmentPromoteUser != nil {
            // 默认等于查询
            db = db.Where("investment_promote_user = ?", *r.InvestmentPromoteUser)
        }
        if r.DockDistrict != nil {
            // 默认等于查询
            db = db.Where("dock_district = ?", *r.DockDistrict)
        }
        if r.ItemProgress != nil {
            // 默认等于查询
            db = db.Where("item_progress = ?", *r.ItemProgress)
        }
        if r.CapitalFactor != nil {
            // 默认等于查询
            db = db.Where("capital_factor = ?", *r.CapitalFactor)
        }
        if r.CapitalAssessmentContent != nil {
            // 默认等于查询
            db = db.Where("capital_assessment_content = ?", *r.CapitalAssessmentContent)
        }
        if r.CapitalAssessmentView != nil {
            // 默认等于查询
            db = db.Where("capital_assessment_view = ?", *r.CapitalAssessmentView)
        }
        if r.SceneFactor != nil {
            // 默认等于查询
            db = db.Where("scene_factor = ?", *r.SceneFactor)
        }
        if r.SceneAssessmentContent != nil {
            // 默认等于查询
            db = db.Where("scene_assessment_content = ?", *r.SceneAssessmentContent)
        }
        if r.SceneAssessmentView != nil {
            // 默认等于查询
            db = db.Where("scene_assessment_view = ?", *r.SceneAssessmentView)
        }
        if r.AreaFactor != nil {
            // 默认等于查询
            db = db.Where("area_factor = ?", *r.AreaFactor)
        }
        if r.AreaAssessmentContent != nil {
            // 默认等于查询
            db = db.Where("area_assessment_content = ?", *r.AreaAssessmentContent)
        }
        if r.AreaAssessmentView != nil {
            // 默认等于查询
            db = db.Where("area_assessment_view = ?", *r.AreaAssessmentView)
        }
        if r.PolicyFactor != nil {
            // 默认等于查询
            db = db.Where("policy_factor = ?", *r.PolicyFactor)
        }
        if r.PolicyAssessmentContent != nil {
            // 默认等于查询
            db = db.Where("policy_assessment_content = ?", *r.PolicyAssessmentContent)
        }
        if r.PolicyAssessmentView != nil {
            // 默认等于查询
            db = db.Where("policy_assessment_view = ?", *r.PolicyAssessmentView)
        }
        if r.ItemOwnerName != nil {
            // 默认等于查询
            db = db.Where("item_owner_name = ?", *r.ItemOwnerName)
        }
        if r.RelatedRask != nil {
            // 默认等于查询
            db = db.Where("related_rask = ?", *r.RelatedRask)
        }
        if r.RelatedLead != nil {
            // 默认等于查询
            db = db.Where("related_lead = ?", *r.RelatedLead)
        }
        if r.WinRate != nil {
            // 默认等于查询
            db = db.Where("win_rate = ?", *r.WinRate)
        }
        if r.ProjectContactInfo != nil {
            // 默认等于查询
            db = db.Where("project_contact_info = ?", *r.ProjectContactInfo)
        }
        if r.ProjectType != nil {
            // 默认等于查询
            db = db.Where("project_type = ?", *r.ProjectType)
        }
        if r.AssessmentContent != nil {
            // 默认等于查询
            db = db.Where("assessment_content = ?", *r.AssessmentContent)
        }
        if r.AssessmentView != nil {
            // 默认等于查询
            db = db.Where("assessment_view = ?", *r.AssessmentView)
        }
        if r.CoreAdvantages != nil {
            // 默认等于查询
            db = db.Where("core_advantages = ?", *r.CoreAdvantages)
        }
        if r.IndustryStatus != nil {
            // 默认等于查询
            db = db.Where("industry_status = ?", *r.IndustryStatus)
        }
        if r.FinancialSituation != nil {
            // 默认等于查询
            db = db.Where("financial_situation = ?", *r.FinancialSituation)
        }
        if r.Speed != nil {
            // 默认等于查询
            db = db.Where("speed = ?", *r.Speed)
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
        if r.IndustryTypeCn != nil {
            // 默认等于查询
            db = db.Where("industry_type_cn = ?", *r.IndustryTypeCn)
        }
        if r.BureauProjectId != nil {
            // 默认等于查询
            db = db.Where("bureau_project_id = ?", *r.BureauProjectId)
        }
        if r.InvestCompanyName != nil {
            // 默认等于查询
            db = db.Where("invest_company_name = ?", *r.InvestCompanyName)
        }
        if r.CityCode != nil {
            // 默认等于查询
            db = db.Where("city_code = ?", *r.CityCode)
        }
        if r.CapitalInvestInfo != nil {
            // 默认等于查询
            db = db.Where("capital_invest_info = ?", *r.CapitalInvestInfo)
        }
        if r.Tag != nil {
            // 默认等于查询
            db = db.Where("tag = ?", *r.Tag)
        }
        if r.FinancingAmount != nil {
            // 默认等于查询
            db = db.Where("financing_amount = ?", *r.FinancingAmount)
        }
        if r.BuildingInvestmentArea != nil {
            // 默认等于查询
            db = db.Where("building_investment_area = ?", *r.BuildingInvestmentArea)
        }
        if r.GuaranteeMethods != nil {
            // 默认等于查询
            db = db.Where("guarantee_methods = ?", *r.GuaranteeMethods)
        }
        if r.Period != nil {
            // 默认等于查询
            db = db.Where("period = ?", *r.Period)
        }
        if r.FactoryRentalArea != nil {
            // 默认等于查询
            db = db.Where("factory_rental_area = ?", *r.FactoryRentalArea)
        }
        if r.FactoryFloorHeightBegins != nil {
            // 默认等于查询
            db = db.Where("factory_floor_height_begins = ?", *r.FactoryFloorHeightBegins)
        }
        if r.Quota != nil {
            // 默认等于查询
            db = db.Where("quota = ?", *r.Quota)
        }
        if r.ParkType != nil {
            // 默认等于查询
            db = db.Where("park_type = ?", *r.ParkType)
        }
        if r.MinimumRentalValue != nil {
            // 默认等于查询
            db = db.Where("minimum_rental_value = ?", *r.MinimumRentalValue)
        }
        if r.MaximumFactoryRent != nil {
            // 默认等于查询
            db = db.Where("maximum_factory_rent = ?", *r.MaximumFactoryRent)
        }
        if r.WhetherThereIsANeedForDe != nil {
            // 默认等于查询
            db = db.Where("whether_there_is_a_need_for_de = ?", *r.WhetherThereIsANeedForDe)
        }
        if r.EndOfFactoryFloorHeight != nil {
            // 默认等于查询
            db = db.Where("end_of_factory_floor_height = ?", *r.EndOfFactoryFloorHeight)
        }
        if r.MaximumRentalValue != nil {
            // 默认等于查询
            db = db.Where("maximum_rental_value = ?", *r.MaximumRentalValue)
        }
        if r.MinimumFactoryRent != nil {
            // 默认等于查询
            db = db.Where("minimum_factory_rent = ?", *r.MinimumFactoryRent)
        }
        if r.LandUse != nil {
            // 默认等于查询
            db = db.Where("land_use = ?", *r.LandUse)
        }
        if r.ParkLevel != nil {
            // 默认等于查询
            db = db.Where("park_level = ?", *r.ParkLevel)
        }
        if r.TypeOfBuildingRent != nil {
            // 默认等于查询
            db = db.Where("type_of_building_rent = ?", *r.TypeOfBuildingRent)
        }
        if r.LandAreaRequired != nil {
            // 默认等于查询
            db = db.Where("land_area_required = ?", *r.LandAreaRequired)
        }
        if r.FactoryRentType != nil {
            // 默认等于查询
            db = db.Where("factory_rent_type = ?", *r.FactoryRentType)
        }
        if r.InterestRate != nil {
            // 默认等于查询
            db = db.Where("interest_rate = ?", *r.InterestRate)
        }
        if r.ClueResearcher != nil {
            // 默认等于查询
            db = db.Where("clue_researcher = ?", *r.ClueResearcher)
        }
        if r.AssessmentTeam != nil {
            // 默认等于查询
            db = db.Where("assessment_team = ?", *r.AssessmentTeam)
        }
        if r.EquityFinancingAmount != nil {
            // 默认等于查询
            db = db.Where("equity_financing_amount = ?", *r.EquityFinancingAmount)
        }
        if r.BondFinancingAmount != nil {
            // 默认等于查询
            db = db.Where("bond_financing_amount = ?", *r.BondFinancingAmount)
        }
        if r.SimilarFinancialAmount != nil {
            // 默认等于查询
            db = db.Where("similar_financial_amount = ?", *r.SimilarFinancialAmount)
        }
        if r.BureauProjectProgress != nil {
            // 默认等于查询
            db = db.Where("bureau_project_progress = ?", *r.BureauProjectProgress)
        }
        if r.Country != nil {
            // 默认等于查询
            db = db.Where("country = ?", *r.Country)
        }
        if r.IssueId != nil {
            // 默认等于查询
            db = db.Where("issue_id = ?", *r.IssueId)
        }
        if r.ProjectIntroduction != nil {
            // 默认等于查询
            db = db.Where("project_introduction = ?", *r.ProjectIntroduction)
        }
        if r.IssueType != nil {
            // 默认等于查询
            db = db.Where("issue_type = ?", *r.IssueType)
        }
        if r.AreaDockContent != nil {
            // 默认等于查询
            db = db.Where("area_dock_content = ?", *r.AreaDockContent)
        }
        if r.MoneyValue != nil {
            // 默认等于查询
            db = db.Where("money_value = ?", *r.MoneyValue)
        }
        if r.IssueContent != nil {
            // 默认等于查询
            db = db.Where("issue_content = ?", *r.IssueContent)
        }
        if r.AreaAdvice != nil {
            // 默认等于查询
            db = db.Where("area_advice = ?", *r.AreaAdvice)
        }
		return db
	}
}

// ProjectCreateRequest 创建biz_project请求参数
type ProjectCreateRequest struct {
	models.Validator
	ProjectNo string `form:"projectNo" validate:"required" message:"项目编号不能为空"` // 项目编号
	Name string `form:"name" validate:"required" message:"客户名称不能为空"` // 客户名称
	UpdatedBy int64 `form:"updatedBy"` // 修改人
	Source string `form:"source" validate:"required" message:"客户来源不能为空"` // 客户来源
	LastFlowTime *time.Time `form:"lastFlowTime" validate:"required" message:"最近跟进时间不能为空"` // 最近跟进时间
	TransactionStatus string `form:"transactionStatus" validate:"required" message:"成交状态不能为空"` // 成交状态
	RecycleTime *time.Time `form:"recycleTime" validate:"required" message:"退回时间不能为空"` // 退回时间
	DistributionStatus string `form:"distributionStatus" validate:"required" message:"分配状态不能为空"` // 分配状态
	BaseCurrency float64 `form:"baseCurrency"` // BaseCurrency
	ExchangeRate string `form:"exchangeRate" validate:"required" message:"汇率不能为空"` // 汇率
	RecycleReason string `form:"recycleReason" validate:"required" message:"退回/回收原因不能为空"` // 退回/回收原因
	PoolId int64 `form:"poolId"` // PoolId
	Code string `form:"code" validate:"required" message:"客户编号不能为空"` // 客户编号
	GetTime *time.Time `form:"getTime" validate:"required" message:"获取时间不能为空"` // 获取时间
	DistributionTime *time.Time `form:"distributionTime" validate:"required" message:"领取/分配时间不能为空"` // 领取/分配时间
	HandoverObjects string `form:"handoverObjects" validate:"required" message:"同时移交客户关联的业务对象不能为空"` // 同时移交客户关联的业务对象
	HandoverRetainOwner int `form:"handoverRetainOwner"` // 移交保留归属人在团队
	OtherHandoverRetainOwner int `form:"otherHandoverRetainOwner"` // 其他对象移交保留归属人在团队
	RecycleStartTime *time.Time `form:"recycleStartTime" validate:"required" message:"回收计算起始时间不能为空"` // 回收计算起始时间
	AuditCode string `form:"auditCode" validate:"required" message:"审批编号不能为空"` // 审批编号
	AuditStartTime *time.Time `form:"auditStartTime" validate:"required" message:"审批发起时间不能为空"` // 审批发起时间
	AuditEndTime *time.Time `form:"auditEndTime" validate:"required" message:"审批结束时间不能为空"` // 审批结束时间
	AuditStatus string `form:"auditStatus" validate:"required" message:"审批状态不能为空"` // 审批状态
	AuditMsg string `form:"auditMsg" validate:"required" message:"审批留言不能为空"` // 审批留言
	ApplyTime *time.Time `form:"applyTime" validate:"required" message:"申请时间不能为空"` // 申请时间
	OriginDataOwnerId int64 `form:"originDataOwnerId"` // 原归属人
	AuditStartUser int64 `form:"auditStartUser"` // 审批发起人
	FirstGetFromPool int `form:"firstGetFromPool"` // 是否为第一次从池里获取
	ContactUser int64 `form:"contactUser"` // 联系人
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	GenerationDate *time.Time `form:"generationDate" validate:"required" message:"项目生成时间不能为空"` // 项目生成时间
	FixedInvestmentRemark string `form:"fixedInvestmentRemark" validate:"required" message:"固投备注不能为空"` // 固投备注
	SourceOfInvestment string `form:"sourceOfInvestment" validate:"required" message:"投资来源地不能为空"` // 投资来源地
	TotalInvestmentAmount float64 `form:"totalInvestmentAmount"` // 总投资额（亿）
	FixedAssetsInvestment float64 `form:"fixedAssetsInvestment"` // 固定资产投资额（亿）
	InvestmentArea string `form:"investmentArea" validate:"required" message:"投资区域不能为空"` // 投资区域
	MunicipalServiceTel string `form:"municipalServiceTel" validate:"required" message:"市级服务专员联系方式不能为空"` // 市级服务专员联系方式
	ProjectLabels string `form:"projectLabels" validate:"required" message:"项目标签不能为空"` // 项目标签
	ExpectedSigningTime *time.Time `form:"expectedSigningTime" validate:"required" message:"预计签约时间不能为空"` // 预计签约时间
	InvestmentType string `form:"investmentType" validate:"required" message:"投资类型不能为空"` // 投资类型
	BelongIndustryChain string `form:"belongIndustryChain" validate:"required" message:"所属产业链不能为空"` // 所属产业链
	Industry string `form:"industry" validate:"required" message:"行业不能为空"` // 行业
	IndustrySector string `form:"industrySector" validate:"required" message:"行业领域不能为空"` // 行业领域
	IsForeignInvestment int `form:"isForeignInvestment"` // 是否外资
	IsAcquireLand int `form:"isAcquireLand"` // 是否拿地
	PromotionTypes string `form:"promotionTypes" validate:"required" message:"推进形式不能为空"` // 推进形式
	ProjectContent string `form:"projectContent" validate:"required" message:"项目内容不能为空"` // 项目内容
	AdvancePlan string `form:"advancePlan" validate:"required" message:"推进计划不能为空"` // 推进计划
	JudgingMaterials string `form:"judgingMaterials" validate:"required" message:"研判材料不能为空"` // 研判材料
	ListOfClueSource string `form:"listOfClueSource" validate:"required" message:"清单线索来源不能为空"` // 清单线索来源
	LeaderOfMunicipalTrans string `form:"leaderOfMunicipalTrans" validate:"required" message:"市交办领导不能为空"` // 市交办领导
	LeaderOfBureauTrans string `form:"leaderOfBureauTrans" validate:"required" message:"局交办领导不能为空"` // 局交办领导
	SourceRemark string `form:"sourceRemark" validate:"required" message:"来源备注不能为空"` // 来源备注
	Dept string `form:"dept" validate:"required" message:"部门不能为空"` // 部门
	District string `form:"district" validate:"required" message:"区县不能为空"` // 区县
	PlatformCompany string `form:"platformCompany" validate:"required" message:"平台公司不能为空"` // 平台公司
	IsCityPlanPartner int `form:"isCityPlanPartner"` // 是否来自创投城市计划生态合伙人
	EnterpriseType string `form:"enterpriseType" validate:"required" message:"企业类型不能为空"` // 企业类型
	InvestorProfile string `form:"investorProfile" validate:"required" message:"投资方简介不能为空"` // 投资方简介
	ProjectContact string `form:"projectContact" validate:"required" message:"项目联系人不能为空"` // 项目联系人
	EnterpriseContactInfo string `form:"enterpriseContactInfo" validate:"required" message:"企业联系方式不能为空"` // 企业联系方式
	Referrer string `form:"referrer" validate:"required" message:"引荐人不能为空"` // 引荐人
	ReferrerContactInfo string `form:"referrerContactInfo" validate:"required" message:"引荐人联系方式不能为空"` // 引荐人联系方式
	ProjectPhase string `form:"projectPhase" validate:"required" message:"在谈节点不能为空"` // 在谈节点
	PhaseId string `form:"phaseId" validate:"required" message:"项目阶段不能为空"` // 项目阶段
	BpAttachments string `form:"bpAttachments" validate:"required" message:"BP附件库不能为空"` // BP附件库
	Stress string `form:"stress" validate:"required" message:"重要性不能为空"` // 重要性
	IndustryType string `form:"industryType" validate:"required" message:"行业分类不能为空"` // 行业分类
	City string `form:"city" validate:"required" message:"所在城市不能为空"` // 所在城市
	InvestCompany int64 `form:"investCompany"` // 投资方名称
	ItemProvideDepartment int64 `form:"itemProvideDepartment"` // 项目提供部门/机构
	ItemRecommender string `form:"itemRecommender" validate:"required" message:"项目推荐方不能为空"` // 项目推荐方
	ItemRecommendTime *time.Time `form:"itemRecommendTime" validate:"required" message:"项目推荐时间不能为空"` // 项目推荐时间
	LeadDepartment int64 `form:"leadDepartment"` // 牵头部门
	DistributeTask int `form:"distributeTask"` // 同时下发行业研判任务
	ShareItem int `form:"shareItem"` // 同时与行业牵头人或产业专班共享项目
	LandContent string `form:"landContent" validate:"required" message:"落地内容不能为空"` // 落地内容
	EnterpriseAppeal string `form:"enterpriseAppeal" validate:"required" message:"企业诉求不能为空"` // 企业诉求
	Carrier string `form:"carrier" validate:"required" message:"载体不能为空"` // 载体
	ProductionFactors string `form:"productionFactors" validate:"required" message:"生产要素不能为空"` // 生产要素
	Capital string `form:"capital" validate:"required" message:"融资/资本不能为空"` // 融资/资本
	Scene string `form:"scene" validate:"required" message:"场景/市场不能为空"` // 场景/市场
	Policy string `form:"policy" validate:"required" message:"政策不能为空"` // 政策
	InvestmentPromoteUser int64 `form:"investmentPromoteUser"` // 投促推进人
	DockDistrict string `form:"dockDistrict" validate:"required" message:"对接区县不能为空"` // 对接区县
	ItemProgress string `form:"itemProgress" validate:"required" message:"项目进展不能为空"` // 项目进展
	CapitalFactor int `form:"capitalFactor"` // 资本要素
	CapitalAssessmentContent string `form:"capitalAssessmentContent" validate:"required" message:"研判内容不能为空"` // 研判内容
	CapitalAssessmentView string `form:"capitalAssessmentView" validate:"required" message:"资本要素-研判意见不能为空"` // 资本要素-研判意见
	SceneFactor int `form:"sceneFactor"` // 场景要素
	SceneAssessmentContent string `form:"sceneAssessmentContent" validate:"required" message:"场景要素-研判内容不能为空"` // 场景要素-研判内容
	SceneAssessmentView string `form:"sceneAssessmentView" validate:"required" message:"场景要素-研判意见不能为空"` // 场景要素-研判意见
	AreaFactor int `form:"areaFactor"` // 场地要素
	AreaAssessmentContent string `form:"areaAssessmentContent" validate:"required" message:"场地要素-研判内容不能为空"` // 场地要素-研判内容
	AreaAssessmentView string `form:"areaAssessmentView" validate:"required" message:"场地要素-研判意见不能为空"` // 场地要素-研判意见
	PolicyFactor int `form:"policyFactor"` // 政策要素
	PolicyAssessmentContent string `form:"policyAssessmentContent" validate:"required" message:"政策要素-研判内容不能为空"` // 政策要素-研判内容
	PolicyAssessmentView string `form:"policyAssessmentView" validate:"required" message:"政策要素-研判意见不能为空"` // 政策要素-研判意见
	ItemOwnerName int64 `form:"itemOwnerName"` // 项目归属人员
	RelatedRask int64 `form:"relatedRask"` // 关联任务
	RelatedLead int64 `form:"relatedLead"` // 关联线索
	WinRate float64 `form:"winRate"` // 赢率
	ProjectContactInfo string `form:"projectContactInfo" validate:"required" message:"项目联系方式不能为空"` // 项目联系方式
	ProjectType string `form:"projectType" validate:"required" message:"项目类型不能为空"` // 项目类型
	AssessmentContent string `form:"assessmentContent" validate:"required" message:"研判内容不能为空"` // 研判内容
	AssessmentView string `form:"assessmentView" validate:"required" message:"研判意见不能为空"` // 研判意见
	CoreAdvantages string `form:"coreAdvantages" validate:"required" message:"核心优势不能为空"` // 核心优势
	IndustryStatus string `form:"industryStatus" validate:"required" message:"行业地位不能为空"` // 行业地位
	FinancialSituation string `form:"financialSituation" validate:"required" message:"财务情况不能为空"` // 财务情况
	Speed string `form:"speed" validate:"required" message:"进度不能为空"` // 进度
	InvestRound string `form:"investRound" validate:"required" message:"拟投资轮次不能为空"` // 拟投资轮次
	NeedIndustry int `form:"needIndustry"` // 是否需要产业投资人
	NeedFinance int `form:"needFinance"` // 是否需要财务投资人
	IndustryTypeCn string `form:"industryTypeCn" validate:"required" message:"行业分类-中文不能为空"` // 行业分类-中文
	BureauProjectId string `form:"bureauProjectId" validate:"required" message:"市局项目id不能为空"` // 市局项目id
	InvestCompanyName string `form:"investCompanyName" validate:"required" message:"投资方名称（与市局对接预留字段）不能为空"` // 投资方名称（与市局对接预留字段）
	CityCode string `form:"cityCode" validate:"required" message:"所属城市(省市区)不能为空"` // 所属城市(省市区)
	CapitalInvestInfo string `form:"capitalInvestInfo" validate:"required" message:"资本要素的投资信息（json）不能为空"` // 资本要素的投资信息（json）
	Tag string `form:"tag" validate:"required" message:"数据标签不能为空"` // 数据标签
	FinancingAmount float64 `form:"financingAmount"` // 融资金额(亿元)
	BuildingInvestmentArea float64 `form:"buildingInvestmentArea"` // 楼宇招商面积
	GuaranteeMethods string `form:"guaranteeMethods" validate:"required" message:"担保方式不能为空"` // 担保方式
	Period string `form:"period" validate:"required" message:"期限不能为空"` // 期限
	FactoryRentalArea float64 `form:"factoryRentalArea"` // 厂房租用面积
	FactoryFloorHeightBegins float64 `form:"factoryFloorHeightBegins"` // 厂房层高开始
	Quota string `form:"quota" validate:"required" message:"额度不能为空"` // 额度
	ParkType string `form:"parkType" validate:"required" message:"园区类型不能为空"` // 园区类型
	MinimumRentalValue float64 `form:"minimumRentalValue"` // 楼宇租金最小值
	MaximumFactoryRent float64 `form:"maximumFactoryRent"` // 厂房租金最大值
	WhetherThereIsANeedForDe int `form:"whetherThereIsANeedForDe"` // 是否有债权融资需求
	EndOfFactoryFloorHeight float64 `form:"endOfFactoryFloorHeight"` // 厂房层高结束
	MaximumRentalValue float64 `form:"maximumRentalValue"` // 楼宇租金最大值
	MinimumFactoryRent float64 `form:"minimumFactoryRent"` // 厂房租金最小值
	LandUse string `form:"landUse" validate:"required" message:"土地用途不能为空"` // 土地用途
	ParkLevel string `form:"parkLevel" validate:"required" message:"园区等级不能为空"` // 园区等级
	TypeOfBuildingRent string `form:"typeOfBuildingRent" validate:"required" message:"楼宇租金类型不能为空"` // 楼宇租金类型
	LandAreaRequired float64 `form:"landAreaRequired"` // 需地块面积
	FactoryRentType string `form:"factoryRentType" validate:"required" message:"厂房租金类型不能为空"` // 厂房租金类型
	InterestRate float64 `form:"interestRate"` // 利率
	ClueResearcher int64 `form:"clueResearcher"` // 线索研判人
	AssessmentTeam string `form:"assessmentTeam" validate:"required" message:"研判团队不能为空"` // 研判团队
	EquityFinancingAmount float64 `form:"equityFinancingAmount"` // 股权融资金额
	BondFinancingAmount float64 `form:"bondFinancingAmount"` // 债券融资金额
	SimilarFinancialAmount float64 `form:"similarFinancialAmount"` // 类金融融资金额
	BureauProjectProgress string `form:"bureauProjectProgress" validate:"required" message:"市局项目进度不能为空"` // 市局项目进度
	Country string `form:"country" validate:"required" message:"国家不能为空"` // 国家
	IssueId string `form:"issueId" validate:"required" message:"市局项目轮次id不能为空"` // 市局项目轮次id
	ProjectIntroduction string `form:"projectIntroduction" validate:"required" message:"项目介绍不能为空"` // 项目介绍
	IssueType string `form:"issueType" validate:"required" message:"问题类型不能为空"` // 问题类型
	AreaDockContent string `form:"areaDockContent" validate:"required" message:"区级对接情况不能为空"` // 区级对接情况
	MoneyValue string `form:"moneyValue" validate:"required" message:"问题金额不能为空"` // 问题金额
	IssueContent string `form:"issueContent" validate:"required" message:"问题描述不能为空"` // 问题描述
	AreaAdvice string `form:"areaAdvice" validate:"required" message:"区级建议不能为空"` // 区级建议
}

// Validate 验证请求参数
func (r *ProjectCreateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ProjectUpdateRequest 更新biz_project请求参数
type ProjectUpdateRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
	ProjectNo string `form:"projectNo" validate:"required" message:"项目编号不能为空"` // 项目编号
	Name string `form:"name" validate:"required" message:"客户名称不能为空"` // 客户名称
	UpdatedBy int64 `form:"updatedBy"` // 修改人
	Source string `form:"source" validate:"required" message:"客户来源不能为空"` // 客户来源
	LastFlowTime *time.Time `form:"lastFlowTime" validate:"required" message:"最近跟进时间不能为空"` // 最近跟进时间
	TransactionStatus string `form:"transactionStatus" validate:"required" message:"成交状态不能为空"` // 成交状态
	RecycleTime *time.Time `form:"recycleTime" validate:"required" message:"退回时间不能为空"` // 退回时间
	DistributionStatus string `form:"distributionStatus" validate:"required" message:"分配状态不能为空"` // 分配状态
	BaseCurrency float64 `form:"baseCurrency"` // BaseCurrency
	ExchangeRate string `form:"exchangeRate" validate:"required" message:"汇率不能为空"` // 汇率
	RecycleReason string `form:"recycleReason" validate:"required" message:"退回/回收原因不能为空"` // 退回/回收原因
	PoolId int64 `form:"poolId"` // PoolId
	Code string `form:"code" validate:"required" message:"客户编号不能为空"` // 客户编号
	GetTime *time.Time `form:"getTime" validate:"required" message:"获取时间不能为空"` // 获取时间
	DistributionTime *time.Time `form:"distributionTime" validate:"required" message:"领取/分配时间不能为空"` // 领取/分配时间
	HandoverObjects string `form:"handoverObjects" validate:"required" message:"同时移交客户关联的业务对象不能为空"` // 同时移交客户关联的业务对象
	HandoverRetainOwner int `form:"handoverRetainOwner"` // 移交保留归属人在团队
	OtherHandoverRetainOwner int `form:"otherHandoverRetainOwner"` // 其他对象移交保留归属人在团队
	RecycleStartTime *time.Time `form:"recycleStartTime" validate:"required" message:"回收计算起始时间不能为空"` // 回收计算起始时间
	AuditCode string `form:"auditCode" validate:"required" message:"审批编号不能为空"` // 审批编号
	AuditStartTime *time.Time `form:"auditStartTime" validate:"required" message:"审批发起时间不能为空"` // 审批发起时间
	AuditEndTime *time.Time `form:"auditEndTime" validate:"required" message:"审批结束时间不能为空"` // 审批结束时间
	AuditStatus string `form:"auditStatus" validate:"required" message:"审批状态不能为空"` // 审批状态
	AuditMsg string `form:"auditMsg" validate:"required" message:"审批留言不能为空"` // 审批留言
	ApplyTime *time.Time `form:"applyTime" validate:"required" message:"申请时间不能为空"` // 申请时间
	OriginDataOwnerId int64 `form:"originDataOwnerId"` // 原归属人
	AuditStartUser int64 `form:"auditStartUser"` // 审批发起人
	FirstGetFromPool int `form:"firstGetFromPool"` // 是否为第一次从池里获取
	ContactUser int64 `form:"contactUser"` // 联系人
	Remark string `form:"remark" validate:"required" message:"备注不能为空"` // 备注
	GenerationDate *time.Time `form:"generationDate" validate:"required" message:"项目生成时间不能为空"` // 项目生成时间
	FixedInvestmentRemark string `form:"fixedInvestmentRemark" validate:"required" message:"固投备注不能为空"` // 固投备注
	SourceOfInvestment string `form:"sourceOfInvestment" validate:"required" message:"投资来源地不能为空"` // 投资来源地
	TotalInvestmentAmount float64 `form:"totalInvestmentAmount"` // 总投资额（亿）
	FixedAssetsInvestment float64 `form:"fixedAssetsInvestment"` // 固定资产投资额（亿）
	InvestmentArea string `form:"investmentArea" validate:"required" message:"投资区域不能为空"` // 投资区域
	MunicipalServiceTel string `form:"municipalServiceTel" validate:"required" message:"市级服务专员联系方式不能为空"` // 市级服务专员联系方式
	ProjectLabels string `form:"projectLabels" validate:"required" message:"项目标签不能为空"` // 项目标签
	ExpectedSigningTime *time.Time `form:"expectedSigningTime" validate:"required" message:"预计签约时间不能为空"` // 预计签约时间
	InvestmentType string `form:"investmentType" validate:"required" message:"投资类型不能为空"` // 投资类型
	BelongIndustryChain string `form:"belongIndustryChain" validate:"required" message:"所属产业链不能为空"` // 所属产业链
	Industry string `form:"industry" validate:"required" message:"行业不能为空"` // 行业
	IndustrySector string `form:"industrySector" validate:"required" message:"行业领域不能为空"` // 行业领域
	IsForeignInvestment int `form:"isForeignInvestment"` // 是否外资
	IsAcquireLand int `form:"isAcquireLand"` // 是否拿地
	PromotionTypes string `form:"promotionTypes" validate:"required" message:"推进形式不能为空"` // 推进形式
	ProjectContent string `form:"projectContent" validate:"required" message:"项目内容不能为空"` // 项目内容
	AdvancePlan string `form:"advancePlan" validate:"required" message:"推进计划不能为空"` // 推进计划
	JudgingMaterials string `form:"judgingMaterials" validate:"required" message:"研判材料不能为空"` // 研判材料
	ListOfClueSource string `form:"listOfClueSource" validate:"required" message:"清单线索来源不能为空"` // 清单线索来源
	LeaderOfMunicipalTrans string `form:"leaderOfMunicipalTrans" validate:"required" message:"市交办领导不能为空"` // 市交办领导
	LeaderOfBureauTrans string `form:"leaderOfBureauTrans" validate:"required" message:"局交办领导不能为空"` // 局交办领导
	SourceRemark string `form:"sourceRemark" validate:"required" message:"来源备注不能为空"` // 来源备注
	Dept string `form:"dept" validate:"required" message:"部门不能为空"` // 部门
	District string `form:"district" validate:"required" message:"区县不能为空"` // 区县
	PlatformCompany string `form:"platformCompany" validate:"required" message:"平台公司不能为空"` // 平台公司
	IsCityPlanPartner int `form:"isCityPlanPartner"` // 是否来自创投城市计划生态合伙人
	EnterpriseType string `form:"enterpriseType" validate:"required" message:"企业类型不能为空"` // 企业类型
	InvestorProfile string `form:"investorProfile" validate:"required" message:"投资方简介不能为空"` // 投资方简介
	ProjectContact string `form:"projectContact" validate:"required" message:"项目联系人不能为空"` // 项目联系人
	EnterpriseContactInfo string `form:"enterpriseContactInfo" validate:"required" message:"企业联系方式不能为空"` // 企业联系方式
	Referrer string `form:"referrer" validate:"required" message:"引荐人不能为空"` // 引荐人
	ReferrerContactInfo string `form:"referrerContactInfo" validate:"required" message:"引荐人联系方式不能为空"` // 引荐人联系方式
	ProjectPhase string `form:"projectPhase" validate:"required" message:"在谈节点不能为空"` // 在谈节点
	PhaseId string `form:"phaseId" validate:"required" message:"项目阶段不能为空"` // 项目阶段
	BpAttachments string `form:"bpAttachments" validate:"required" message:"BP附件库不能为空"` // BP附件库
	Stress string `form:"stress" validate:"required" message:"重要性不能为空"` // 重要性
	IndustryType string `form:"industryType" validate:"required" message:"行业分类不能为空"` // 行业分类
	City string `form:"city" validate:"required" message:"所在城市不能为空"` // 所在城市
	InvestCompany int64 `form:"investCompany"` // 投资方名称
	ItemProvideDepartment int64 `form:"itemProvideDepartment"` // 项目提供部门/机构
	ItemRecommender string `form:"itemRecommender" validate:"required" message:"项目推荐方不能为空"` // 项目推荐方
	ItemRecommendTime *time.Time `form:"itemRecommendTime" validate:"required" message:"项目推荐时间不能为空"` // 项目推荐时间
	LeadDepartment int64 `form:"leadDepartment"` // 牵头部门
	DistributeTask int `form:"distributeTask"` // 同时下发行业研判任务
	ShareItem int `form:"shareItem"` // 同时与行业牵头人或产业专班共享项目
	LandContent string `form:"landContent" validate:"required" message:"落地内容不能为空"` // 落地内容
	EnterpriseAppeal string `form:"enterpriseAppeal" validate:"required" message:"企业诉求不能为空"` // 企业诉求
	Carrier string `form:"carrier" validate:"required" message:"载体不能为空"` // 载体
	ProductionFactors string `form:"productionFactors" validate:"required" message:"生产要素不能为空"` // 生产要素
	Capital string `form:"capital" validate:"required" message:"融资/资本不能为空"` // 融资/资本
	Scene string `form:"scene" validate:"required" message:"场景/市场不能为空"` // 场景/市场
	Policy string `form:"policy" validate:"required" message:"政策不能为空"` // 政策
	InvestmentPromoteUser int64 `form:"investmentPromoteUser"` // 投促推进人
	DockDistrict string `form:"dockDistrict" validate:"required" message:"对接区县不能为空"` // 对接区县
	ItemProgress string `form:"itemProgress" validate:"required" message:"项目进展不能为空"` // 项目进展
	CapitalFactor int `form:"capitalFactor"` // 资本要素
	CapitalAssessmentContent string `form:"capitalAssessmentContent" validate:"required" message:"研判内容不能为空"` // 研判内容
	CapitalAssessmentView string `form:"capitalAssessmentView" validate:"required" message:"资本要素-研判意见不能为空"` // 资本要素-研判意见
	SceneFactor int `form:"sceneFactor"` // 场景要素
	SceneAssessmentContent string `form:"sceneAssessmentContent" validate:"required" message:"场景要素-研判内容不能为空"` // 场景要素-研判内容
	SceneAssessmentView string `form:"sceneAssessmentView" validate:"required" message:"场景要素-研判意见不能为空"` // 场景要素-研判意见
	AreaFactor int `form:"areaFactor"` // 场地要素
	AreaAssessmentContent string `form:"areaAssessmentContent" validate:"required" message:"场地要素-研判内容不能为空"` // 场地要素-研判内容
	AreaAssessmentView string `form:"areaAssessmentView" validate:"required" message:"场地要素-研判意见不能为空"` // 场地要素-研判意见
	PolicyFactor int `form:"policyFactor"` // 政策要素
	PolicyAssessmentContent string `form:"policyAssessmentContent" validate:"required" message:"政策要素-研判内容不能为空"` // 政策要素-研判内容
	PolicyAssessmentView string `form:"policyAssessmentView" validate:"required" message:"政策要素-研判意见不能为空"` // 政策要素-研判意见
	ItemOwnerName int64 `form:"itemOwnerName"` // 项目归属人员
	RelatedRask int64 `form:"relatedRask"` // 关联任务
	RelatedLead int64 `form:"relatedLead"` // 关联线索
	WinRate float64 `form:"winRate"` // 赢率
	ProjectContactInfo string `form:"projectContactInfo" validate:"required" message:"项目联系方式不能为空"` // 项目联系方式
	ProjectType string `form:"projectType" validate:"required" message:"项目类型不能为空"` // 项目类型
	AssessmentContent string `form:"assessmentContent" validate:"required" message:"研判内容不能为空"` // 研判内容
	AssessmentView string `form:"assessmentView" validate:"required" message:"研判意见不能为空"` // 研判意见
	CoreAdvantages string `form:"coreAdvantages" validate:"required" message:"核心优势不能为空"` // 核心优势
	IndustryStatus string `form:"industryStatus" validate:"required" message:"行业地位不能为空"` // 行业地位
	FinancialSituation string `form:"financialSituation" validate:"required" message:"财务情况不能为空"` // 财务情况
	Speed string `form:"speed" validate:"required" message:"进度不能为空"` // 进度
	InvestRound string `form:"investRound" validate:"required" message:"拟投资轮次不能为空"` // 拟投资轮次
	NeedIndustry int `form:"needIndustry"` // 是否需要产业投资人
	NeedFinance int `form:"needFinance"` // 是否需要财务投资人
	IndustryTypeCn string `form:"industryTypeCn" validate:"required" message:"行业分类-中文不能为空"` // 行业分类-中文
	BureauProjectId string `form:"bureauProjectId" validate:"required" message:"市局项目id不能为空"` // 市局项目id
	InvestCompanyName string `form:"investCompanyName" validate:"required" message:"投资方名称（与市局对接预留字段）不能为空"` // 投资方名称（与市局对接预留字段）
	CityCode string `form:"cityCode" validate:"required" message:"所属城市(省市区)不能为空"` // 所属城市(省市区)
	CapitalInvestInfo string `form:"capitalInvestInfo" validate:"required" message:"资本要素的投资信息（json）不能为空"` // 资本要素的投资信息（json）
	Tag string `form:"tag" validate:"required" message:"数据标签不能为空"` // 数据标签
	FinancingAmount float64 `form:"financingAmount"` // 融资金额(亿元)
	BuildingInvestmentArea float64 `form:"buildingInvestmentArea"` // 楼宇招商面积
	GuaranteeMethods string `form:"guaranteeMethods" validate:"required" message:"担保方式不能为空"` // 担保方式
	Period string `form:"period" validate:"required" message:"期限不能为空"` // 期限
	FactoryRentalArea float64 `form:"factoryRentalArea"` // 厂房租用面积
	FactoryFloorHeightBegins float64 `form:"factoryFloorHeightBegins"` // 厂房层高开始
	Quota string `form:"quota" validate:"required" message:"额度不能为空"` // 额度
	ParkType string `form:"parkType" validate:"required" message:"园区类型不能为空"` // 园区类型
	MinimumRentalValue float64 `form:"minimumRentalValue"` // 楼宇租金最小值
	MaximumFactoryRent float64 `form:"maximumFactoryRent"` // 厂房租金最大值
	WhetherThereIsANeedForDe int `form:"whetherThereIsANeedForDe"` // 是否有债权融资需求
	EndOfFactoryFloorHeight float64 `form:"endOfFactoryFloorHeight"` // 厂房层高结束
	MaximumRentalValue float64 `form:"maximumRentalValue"` // 楼宇租金最大值
	MinimumFactoryRent float64 `form:"minimumFactoryRent"` // 厂房租金最小值
	LandUse string `form:"landUse" validate:"required" message:"土地用途不能为空"` // 土地用途
	ParkLevel string `form:"parkLevel" validate:"required" message:"园区等级不能为空"` // 园区等级
	TypeOfBuildingRent string `form:"typeOfBuildingRent" validate:"required" message:"楼宇租金类型不能为空"` // 楼宇租金类型
	LandAreaRequired float64 `form:"landAreaRequired"` // 需地块面积
	FactoryRentType string `form:"factoryRentType" validate:"required" message:"厂房租金类型不能为空"` // 厂房租金类型
	InterestRate float64 `form:"interestRate"` // 利率
	ClueResearcher int64 `form:"clueResearcher"` // 线索研判人
	AssessmentTeam string `form:"assessmentTeam" validate:"required" message:"研判团队不能为空"` // 研判团队
	EquityFinancingAmount float64 `form:"equityFinancingAmount"` // 股权融资金额
	BondFinancingAmount float64 `form:"bondFinancingAmount"` // 债券融资金额
	SimilarFinancialAmount float64 `form:"similarFinancialAmount"` // 类金融融资金额
	BureauProjectProgress string `form:"bureauProjectProgress" validate:"required" message:"市局项目进度不能为空"` // 市局项目进度
	Country string `form:"country" validate:"required" message:"国家不能为空"` // 国家
	IssueId string `form:"issueId" validate:"required" message:"市局项目轮次id不能为空"` // 市局项目轮次id
	ProjectIntroduction string `form:"projectIntroduction" validate:"required" message:"项目介绍不能为空"` // 项目介绍
	IssueType string `form:"issueType" validate:"required" message:"问题类型不能为空"` // 问题类型
	AreaDockContent string `form:"areaDockContent" validate:"required" message:"区级对接情况不能为空"` // 区级对接情况
	MoneyValue string `form:"moneyValue" validate:"required" message:"问题金额不能为空"` // 问题金额
	IssueContent string `form:"issueContent" validate:"required" message:"问题描述不能为空"` // 问题描述
	AreaAdvice string `form:"areaAdvice" validate:"required" message:"区级建议不能为空"` // 区级建议
}

// Validate 验证请求参数
func (r *ProjectUpdateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ProjectDeleteRequest 删除biz_project请求参数
type ProjectDeleteRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *ProjectDeleteRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// ProjectGetByIDRequest 根据ID获取biz_project请求参数
type ProjectGetByIDRequest struct {
	models.Validator
	Id int64 `uri:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *ProjectGetByIDRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}