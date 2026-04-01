package service

import (
	"gin-fast/plugins/biz/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// ProjectService biz_project服务
type ProjectService struct{}

// NewProjectService 创建biz_project服务
func NewProjectService() *ProjectService {
	return &ProjectService{}
}

// Create 创建biz_project
func (s *ProjectService) Create(c *gin.Context, req models.ProjectCreateRequest) (*models.Project, error) {
	// 创建biz_project记录
	project := models.NewProject()
    project.ProjectNo = req.ProjectNo
    project.Name = req.Name
    project.UpdatedBy = req.UpdatedBy
    project.Source = req.Source
    project.LastFlowTime = req.LastFlowTime
    project.TransactionStatus = req.TransactionStatus
    project.RecycleTime = req.RecycleTime
    project.DistributionStatus = req.DistributionStatus
    project.BaseCurrency = req.BaseCurrency
    project.ExchangeRate = req.ExchangeRate
    project.RecycleReason = req.RecycleReason
    project.PoolId = req.PoolId
    project.Code = req.Code
    project.GetTime = req.GetTime
    project.DistributionTime = req.DistributionTime
    project.HandoverObjects = req.HandoverObjects
    project.HandoverRetainOwner = req.HandoverRetainOwner
    project.OtherHandoverRetainOwner = req.OtherHandoverRetainOwner
    project.RecycleStartTime = req.RecycleStartTime
    project.AuditCode = req.AuditCode
    project.AuditStartTime = req.AuditStartTime
    project.AuditEndTime = req.AuditEndTime
    project.AuditStatus = req.AuditStatus
    project.AuditMsg = req.AuditMsg
    project.ApplyTime = req.ApplyTime
    project.OriginDataOwnerId = req.OriginDataOwnerId
    project.AuditStartUser = req.AuditStartUser
    project.FirstGetFromPool = req.FirstGetFromPool
    project.ContactUser = req.ContactUser
    project.Remark = req.Remark
    project.GenerationDate = req.GenerationDate
    project.FixedInvestmentRemark = req.FixedInvestmentRemark
    project.SourceOfInvestment = req.SourceOfInvestment
    project.TotalInvestmentAmount = req.TotalInvestmentAmount
    project.FixedAssetsInvestment = req.FixedAssetsInvestment
    project.InvestmentArea = req.InvestmentArea
    project.MunicipalServiceTel = req.MunicipalServiceTel
    project.ProjectLabels = req.ProjectLabels
    project.ExpectedSigningTime = req.ExpectedSigningTime
    project.InvestmentType = req.InvestmentType
    project.BelongIndustryChain = req.BelongIndustryChain
    project.Industry = req.Industry
    project.IndustrySector = req.IndustrySector
    project.IsForeignInvestment = req.IsForeignInvestment
    project.IsAcquireLand = req.IsAcquireLand
    project.PromotionTypes = req.PromotionTypes
    project.ProjectContent = req.ProjectContent
    project.AdvancePlan = req.AdvancePlan
    project.JudgingMaterials = req.JudgingMaterials
    project.ListOfClueSource = req.ListOfClueSource
    project.LeaderOfMunicipalTrans = req.LeaderOfMunicipalTrans
    project.LeaderOfBureauTrans = req.LeaderOfBureauTrans
    project.SourceRemark = req.SourceRemark
    project.Dept = req.Dept
    project.District = req.District
    project.PlatformCompany = req.PlatformCompany
    project.IsCityPlanPartner = req.IsCityPlanPartner
    project.EnterpriseType = req.EnterpriseType
    project.InvestorProfile = req.InvestorProfile
    project.ProjectContact = req.ProjectContact
    project.EnterpriseContactInfo = req.EnterpriseContactInfo
    project.Referrer = req.Referrer
    project.ReferrerContactInfo = req.ReferrerContactInfo
    project.ProjectPhase = req.ProjectPhase
    project.PhaseId = req.PhaseId
    project.BpAttachments = req.BpAttachments
    project.Stress = req.Stress
    project.IndustryType = req.IndustryType
    project.City = req.City
    project.InvestCompany = req.InvestCompany
    project.ItemProvideDepartment = req.ItemProvideDepartment
    project.ItemRecommender = req.ItemRecommender
    project.ItemRecommendTime = req.ItemRecommendTime
    project.LeadDepartment = req.LeadDepartment
    project.DistributeTask = req.DistributeTask
    project.ShareItem = req.ShareItem
    project.LandContent = req.LandContent
    project.EnterpriseAppeal = req.EnterpriseAppeal
    project.Carrier = req.Carrier
    project.ProductionFactors = req.ProductionFactors
    project.Capital = req.Capital
    project.Scene = req.Scene
    project.Policy = req.Policy
    project.InvestmentPromoteUser = req.InvestmentPromoteUser
    project.DockDistrict = req.DockDistrict
    project.ItemProgress = req.ItemProgress
    project.CapitalFactor = req.CapitalFactor
    project.CapitalAssessmentContent = req.CapitalAssessmentContent
    project.CapitalAssessmentView = req.CapitalAssessmentView
    project.SceneFactor = req.SceneFactor
    project.SceneAssessmentContent = req.SceneAssessmentContent
    project.SceneAssessmentView = req.SceneAssessmentView
    project.AreaFactor = req.AreaFactor
    project.AreaAssessmentContent = req.AreaAssessmentContent
    project.AreaAssessmentView = req.AreaAssessmentView
    project.PolicyFactor = req.PolicyFactor
    project.PolicyAssessmentContent = req.PolicyAssessmentContent
    project.PolicyAssessmentView = req.PolicyAssessmentView
    project.ItemOwnerName = req.ItemOwnerName
    project.RelatedRask = req.RelatedRask
    project.RelatedLead = req.RelatedLead
    project.WinRate = req.WinRate
    project.ProjectContactInfo = req.ProjectContactInfo
    project.ProjectType = req.ProjectType
    project.AssessmentContent = req.AssessmentContent
    project.AssessmentView = req.AssessmentView
    project.CoreAdvantages = req.CoreAdvantages
    project.IndustryStatus = req.IndustryStatus
    project.FinancialSituation = req.FinancialSituation
    project.Speed = req.Speed
    project.InvestRound = req.InvestRound
    project.NeedIndustry = req.NeedIndustry
    project.NeedFinance = req.NeedFinance
    project.IndustryTypeCn = req.IndustryTypeCn
    project.BureauProjectId = req.BureauProjectId
    project.InvestCompanyName = req.InvestCompanyName
    project.CityCode = req.CityCode
    project.CapitalInvestInfo = req.CapitalInvestInfo
    project.Tag = req.Tag
    project.FinancingAmount = req.FinancingAmount
    project.BuildingInvestmentArea = req.BuildingInvestmentArea
    project.GuaranteeMethods = req.GuaranteeMethods
    project.Period = req.Period
    project.FactoryRentalArea = req.FactoryRentalArea
    project.FactoryFloorHeightBegins = req.FactoryFloorHeightBegins
    project.Quota = req.Quota
    project.ParkType = req.ParkType
    project.MinimumRentalValue = req.MinimumRentalValue
    project.MaximumFactoryRent = req.MaximumFactoryRent
    project.WhetherThereIsANeedForDe = req.WhetherThereIsANeedForDe
    project.EndOfFactoryFloorHeight = req.EndOfFactoryFloorHeight
    project.MaximumRentalValue = req.MaximumRentalValue
    project.MinimumFactoryRent = req.MinimumFactoryRent
    project.LandUse = req.LandUse
    project.ParkLevel = req.ParkLevel
    project.TypeOfBuildingRent = req.TypeOfBuildingRent
    project.LandAreaRequired = req.LandAreaRequired
    project.FactoryRentType = req.FactoryRentType
    project.InterestRate = req.InterestRate
    project.ClueResearcher = req.ClueResearcher
    project.AssessmentTeam = req.AssessmentTeam
    project.EquityFinancingAmount = req.EquityFinancingAmount
    project.BondFinancingAmount = req.BondFinancingAmount
    project.SimilarFinancialAmount = req.SimilarFinancialAmount
    project.BureauProjectProgress = req.BureauProjectProgress
    project.Country = req.Country
    project.IssueId = req.IssueId
    project.ProjectIntroduction = req.ProjectIntroduction
    project.IssueType = req.IssueType
    project.AreaDockContent = req.AreaDockContent
    project.MoneyValue = req.MoneyValue
    project.IssueContent = req.IssueContent
    project.AreaAdvice = req.AreaAdvice
	// 保存到数据库
	if err := project.Create(c); err != nil {
		return nil, err
	}

	return project, nil
}

// Update 更新biz_project
func (s *ProjectService) Update(c *gin.Context, req models.ProjectUpdateRequest) error {
	// 查找biz_project记录
	project := models.NewProject()
	if err := project.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新biz_project信息
    project.ProjectNo = req.ProjectNo
    project.Name = req.Name
    project.UpdatedBy = req.UpdatedBy
    project.Source = req.Source
    project.LastFlowTime = req.LastFlowTime
    project.TransactionStatus = req.TransactionStatus
    project.RecycleTime = req.RecycleTime
    project.DistributionStatus = req.DistributionStatus
    project.BaseCurrency = req.BaseCurrency
    project.ExchangeRate = req.ExchangeRate
    project.RecycleReason = req.RecycleReason
    project.PoolId = req.PoolId
    project.Code = req.Code
    project.GetTime = req.GetTime
    project.DistributionTime = req.DistributionTime
    project.HandoverObjects = req.HandoverObjects
    project.HandoverRetainOwner = req.HandoverRetainOwner
    project.OtherHandoverRetainOwner = req.OtherHandoverRetainOwner
    project.RecycleStartTime = req.RecycleStartTime
    project.AuditCode = req.AuditCode
    project.AuditStartTime = req.AuditStartTime
    project.AuditEndTime = req.AuditEndTime
    project.AuditStatus = req.AuditStatus
    project.AuditMsg = req.AuditMsg
    project.ApplyTime = req.ApplyTime
    project.OriginDataOwnerId = req.OriginDataOwnerId
    project.AuditStartUser = req.AuditStartUser
    project.FirstGetFromPool = req.FirstGetFromPool
    project.ContactUser = req.ContactUser
    project.Remark = req.Remark
    project.GenerationDate = req.GenerationDate
    project.FixedInvestmentRemark = req.FixedInvestmentRemark
    project.SourceOfInvestment = req.SourceOfInvestment
    project.TotalInvestmentAmount = req.TotalInvestmentAmount
    project.FixedAssetsInvestment = req.FixedAssetsInvestment
    project.InvestmentArea = req.InvestmentArea
    project.MunicipalServiceTel = req.MunicipalServiceTel
    project.ProjectLabels = req.ProjectLabels
    project.ExpectedSigningTime = req.ExpectedSigningTime
    project.InvestmentType = req.InvestmentType
    project.BelongIndustryChain = req.BelongIndustryChain
    project.Industry = req.Industry
    project.IndustrySector = req.IndustrySector
    project.IsForeignInvestment = req.IsForeignInvestment
    project.IsAcquireLand = req.IsAcquireLand
    project.PromotionTypes = req.PromotionTypes
    project.ProjectContent = req.ProjectContent
    project.AdvancePlan = req.AdvancePlan
    project.JudgingMaterials = req.JudgingMaterials
    project.ListOfClueSource = req.ListOfClueSource
    project.LeaderOfMunicipalTrans = req.LeaderOfMunicipalTrans
    project.LeaderOfBureauTrans = req.LeaderOfBureauTrans
    project.SourceRemark = req.SourceRemark
    project.Dept = req.Dept
    project.District = req.District
    project.PlatformCompany = req.PlatformCompany
    project.IsCityPlanPartner = req.IsCityPlanPartner
    project.EnterpriseType = req.EnterpriseType
    project.InvestorProfile = req.InvestorProfile
    project.ProjectContact = req.ProjectContact
    project.EnterpriseContactInfo = req.EnterpriseContactInfo
    project.Referrer = req.Referrer
    project.ReferrerContactInfo = req.ReferrerContactInfo
    project.ProjectPhase = req.ProjectPhase
    project.PhaseId = req.PhaseId
    project.BpAttachments = req.BpAttachments
    project.Stress = req.Stress
    project.IndustryType = req.IndustryType
    project.City = req.City
    project.InvestCompany = req.InvestCompany
    project.ItemProvideDepartment = req.ItemProvideDepartment
    project.ItemRecommender = req.ItemRecommender
    project.ItemRecommendTime = req.ItemRecommendTime
    project.LeadDepartment = req.LeadDepartment
    project.DistributeTask = req.DistributeTask
    project.ShareItem = req.ShareItem
    project.LandContent = req.LandContent
    project.EnterpriseAppeal = req.EnterpriseAppeal
    project.Carrier = req.Carrier
    project.ProductionFactors = req.ProductionFactors
    project.Capital = req.Capital
    project.Scene = req.Scene
    project.Policy = req.Policy
    project.InvestmentPromoteUser = req.InvestmentPromoteUser
    project.DockDistrict = req.DockDistrict
    project.ItemProgress = req.ItemProgress
    project.CapitalFactor = req.CapitalFactor
    project.CapitalAssessmentContent = req.CapitalAssessmentContent
    project.CapitalAssessmentView = req.CapitalAssessmentView
    project.SceneFactor = req.SceneFactor
    project.SceneAssessmentContent = req.SceneAssessmentContent
    project.SceneAssessmentView = req.SceneAssessmentView
    project.AreaFactor = req.AreaFactor
    project.AreaAssessmentContent = req.AreaAssessmentContent
    project.AreaAssessmentView = req.AreaAssessmentView
    project.PolicyFactor = req.PolicyFactor
    project.PolicyAssessmentContent = req.PolicyAssessmentContent
    project.PolicyAssessmentView = req.PolicyAssessmentView
    project.ItemOwnerName = req.ItemOwnerName
    project.RelatedRask = req.RelatedRask
    project.RelatedLead = req.RelatedLead
    project.WinRate = req.WinRate
    project.ProjectContactInfo = req.ProjectContactInfo
    project.ProjectType = req.ProjectType
    project.AssessmentContent = req.AssessmentContent
    project.AssessmentView = req.AssessmentView
    project.CoreAdvantages = req.CoreAdvantages
    project.IndustryStatus = req.IndustryStatus
    project.FinancialSituation = req.FinancialSituation
    project.Speed = req.Speed
    project.InvestRound = req.InvestRound
    project.NeedIndustry = req.NeedIndustry
    project.NeedFinance = req.NeedFinance
    project.IndustryTypeCn = req.IndustryTypeCn
    project.BureauProjectId = req.BureauProjectId
    project.InvestCompanyName = req.InvestCompanyName
    project.CityCode = req.CityCode
    project.CapitalInvestInfo = req.CapitalInvestInfo
    project.Tag = req.Tag
    project.FinancingAmount = req.FinancingAmount
    project.BuildingInvestmentArea = req.BuildingInvestmentArea
    project.GuaranteeMethods = req.GuaranteeMethods
    project.Period = req.Period
    project.FactoryRentalArea = req.FactoryRentalArea
    project.FactoryFloorHeightBegins = req.FactoryFloorHeightBegins
    project.Quota = req.Quota
    project.ParkType = req.ParkType
    project.MinimumRentalValue = req.MinimumRentalValue
    project.MaximumFactoryRent = req.MaximumFactoryRent
    project.WhetherThereIsANeedForDe = req.WhetherThereIsANeedForDe
    project.EndOfFactoryFloorHeight = req.EndOfFactoryFloorHeight
    project.MaximumRentalValue = req.MaximumRentalValue
    project.MinimumFactoryRent = req.MinimumFactoryRent
    project.LandUse = req.LandUse
    project.ParkLevel = req.ParkLevel
    project.TypeOfBuildingRent = req.TypeOfBuildingRent
    project.LandAreaRequired = req.LandAreaRequired
    project.FactoryRentType = req.FactoryRentType
    project.InterestRate = req.InterestRate
    project.ClueResearcher = req.ClueResearcher
    project.AssessmentTeam = req.AssessmentTeam
    project.EquityFinancingAmount = req.EquityFinancingAmount
    project.BondFinancingAmount = req.BondFinancingAmount
    project.SimilarFinancialAmount = req.SimilarFinancialAmount
    project.BureauProjectProgress = req.BureauProjectProgress
    project.Country = req.Country
    project.IssueId = req.IssueId
    project.ProjectIntroduction = req.ProjectIntroduction
    project.IssueType = req.IssueType
    project.AreaDockContent = req.AreaDockContent
    project.MoneyValue = req.MoneyValue
    project.IssueContent = req.IssueContent
    project.AreaAdvice = req.AreaAdvice
	// 保存到数据库
	if err := project.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除biz_project
func (s *ProjectService) Delete(c *gin.Context, id int64) error {
	// 查找biz_project记录
	project := models.NewProject()
	if err := project.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := project.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取biz_project
func (s *ProjectService) GetByID(c *gin.Context, id int64) (*models.Project, error) {
	// 查找biz_project记录
	project := models.NewProject()
	if err := project.GetByID(c, id); err != nil {
		return nil, err
	}

	return project, nil
}

// List biz_project列表（分页查询）
func (s *ProjectService) List(c *gin.Context, req models.ProjectListRequest) (*models.ProjectList, int64, error) {
	// 获取总数
	projectList := models.NewProjectList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := projectList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = projectList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return projectList, total, nil
}