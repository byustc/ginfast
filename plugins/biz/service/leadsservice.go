package service

import (
	"gin-fast/plugins/biz/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// LeadsService biz_leads服务
type LeadsService struct{}

// NewLeadsService 创建biz_leads服务
func NewLeadsService() *LeadsService {
	return &LeadsService{}
}

// Create 创建biz_leads
func (s *LeadsService) Create(c *gin.Context, req models.LeadsCreateRequest) (*models.Leads, error) {
	// 创建biz_leads记录
	leads := models.NewLeads()
    leads.LeadsNo = req.LeadsNo
    leads.CompanyName = req.CompanyName
    leads.Region = req.Region
    leads.Mobile = req.Mobile
    leads.Telephone = req.Telephone
    leads.Email = req.Email
    leads.LeadsSource = req.LeadsSource
    leads.FlowStatus = req.FlowStatus
    leads.DistributionStatus = req.DistributionStatus
    leads.TransformationStatus = req.TransformationStatus
    leads.LeadsPool = req.LeadsPool
    leads.MarketingActivity = req.MarketingActivity
    leads.PromotionChannel = req.PromotionChannel
    leads.ReturnReason = req.ReturnReason
    leads.ReturnReasonDescription = req.ReturnReasonDescription
    leads.ReturnTime = req.ReturnTime
    leads.DistributionTime = req.DistributionTime
    leads.LeadsDescription = req.LeadsDescription
    leads.CustomerId = req.CustomerId
    leads.ContactId = req.ContactId
    leads.BusinessId = req.BusinessId
    leads.TransformationTime = req.TransformationTime
    leads.Sex = req.Sex
    leads.Department = req.Department
    leads.JobTitle = req.JobTitle
    leads.Age = req.Age
    leads.Birthday = req.Birthday
    leads.Wechat = req.Wechat
    leads.Qq = req.Qq
    leads.Fax = req.Fax
    leads.CompanyAddress = req.CompanyAddress
    leads.CompanyPhone = req.CompanyPhone
    leads.Industry = req.Industry
    leads.StaffSize = req.StaffSize
    leads.CompanyEmail = req.CompanyEmail
    leads.CompanyFax = req.CompanyFax
    leads.LastFlowTime = req.LastFlowTime
    leads.HandoverDescription = req.HandoverDescription
    leads.DistributionDescription = req.DistributionDescription
    leads.GetTime = req.GetTime
    leads.CountryRegion = req.CountryRegion
    leads.ReturnOriginPool = req.ReturnOriginPool
    leads.BusinessLicense = req.BusinessLicense
    leads.GmName = req.GmName
    leads.SalesExperience = req.SalesExperience
    leads.SalesDescription = req.SalesDescription
    leads.AnnualSales = req.AnnualSales
    leads.SameIndustry = req.SameIndustry
    leads.CompanyBrand = req.CompanyBrand
    leads.OtherInformation = req.OtherInformation
    leads.CompanyWeb = req.CompanyWeb
    leads.LicenseValid = req.LicenseValid
    leads.SalesYears = req.SalesYears
    leads.SalesDeptUserNum = req.SalesDeptUserNum
    leads.MarketingDeptUserNum = req.MarketingDeptUserNum
    leads.ServiceDeptUserNum = req.ServiceDeptUserNum
    leads.OtherDeptUserNum = req.OtherDeptUserNum
    leads.Role = req.Role
    leads.EnterpriseType = req.EnterpriseType
    leads.StartDate = req.StartDate
    leads.EndDate = req.EndDate
    leads.ActivityCity = req.ActivityCity
    leads.Sponsor = req.Sponsor
    leads.ActivityName = req.ActivityName
    leads.OwnBu = req.OwnBu
    leads.CooperationDirection = req.CooperationDirection
    leads.Solution = req.Solution
    leads.ContactStatus = req.ContactStatus
    leads.SignedProject = req.SignedProject
    leads.EstimatedGrossProfit = req.EstimatedGrossProfit
    leads.BaseCurrency = req.BaseCurrency
    leads.ExchangeRate = req.ExchangeRate
    leads.ExpectedSigningTime = req.ExpectedSigningTime
    leads.FollowUpPerson = req.FollowUpPerson
    leads.FeedbackResults = req.FeedbackResults
    leads.NewOldCustomer = req.NewOldCustomer
    leads.CustomerPartner = req.CustomerPartner
    leads.IndustryQs = req.IndustryQs
    leads.Stage = req.Stage
    leads.OriginDataOwnerId = req.OriginDataOwnerId
    leads.FirstGetFromPool = req.FirstGetFromPool
    leads.AuditStartUser = req.AuditStartUser
    leads.AuditCode = req.AuditCode
    leads.AuditStatus = req.AuditStatus
    leads.AuditStartTime = req.AuditStartTime
    leads.AuditEndTime = req.AuditEndTime
    leads.ActivityCityText = req.ActivityCityText
    leads.HistoryDataOwnerId = req.HistoryDataOwnerId
    leads.LeaderAssigned = req.LeaderAssigned
    leads.BureauLeaderAssigned = req.BureauLeaderAssigned
    leads.OriginNotes = req.OriginNotes
    leads.District = req.District
    leads.PlatformCompany = req.PlatformCompany
    leads.FromCityPlan = req.FromCityPlan
    leads.ServiceAdvisor = req.ServiceAdvisor
    leads.ServiceAdvisorContactWay = req.ServiceAdvisorContactWay
    leads.IndustryAffiliation = req.IndustryAffiliation
    leads.InvestorOriginPlace = req.InvestorOriginPlace
    leads.ForeignInvestment = req.ForeignInvestment
    leads.BpAttachment = req.BpAttachment
    leads.ProposedProjectName = req.ProposedProjectName
    leads.RegisterMoney = req.RegisterMoney
    leads.TotalInvestmentAmount = req.TotalInvestmentAmount
    leads.TaxEstimation = req.TaxEstimation
    leads.RecruitmentPlanContent = req.RecruitmentPlanContent
    leads.ShareParallelClasses = req.ShareParallelClasses
    leads.ResidencePlace = req.ResidencePlace
    leads.LeadsRecommender = req.LeadsRecommender
    leads.LeadsProvideOrg = req.LeadsProvideOrg
    leads.LeadsOwner = req.LeadsOwner
    leads.InputTime = req.InputTime
    leads.InvestProjectName = req.InvestProjectName
    leads.TechAdvanced = req.TechAdvanced
    leads.LandingContent = req.LandingContent
    leads.ToolCarrier = req.ToolCarrier
    leads.ProductionFactors = req.ProductionFactors
    leads.Financing = req.Financing
    leads.Market = req.Market
    leads.Policy = req.Policy
    leads.IsShareCyzb = req.IsShareCyzb
    leads.IsDistributeTask = req.IsDistributeTask
    leads.IsShareLeads = req.IsShareLeads
    leads.InvestCompany = req.InvestCompany
    leads.PhaseStatus = req.PhaseStatus
    leads.Importance = req.Importance
    leads.EnterpriseDemand = req.EnterpriseDemand
    leads.LeadDepartmentForClues = req.LeadDepartmentForClues
    leads.AnalysisContent = req.AnalysisContent
    leads.JudgmentOpinion = req.JudgmentOpinion
    leads.ReasonTermination = req.ReasonTermination
    leads.Destermiantion = req.Destermiantion
    leads.TechnicalAdvantages = req.TechnicalAdvantages
    leads.BpAttachments = req.BpAttachments
    leads.DockingType = req.DockingType
    leads.DockingInfo = req.DockingInfo
    leads.IndustryStatus = req.IndustryStatus
    leads.PecuniaryCondition = req.PecuniaryCondition
    leads.CoreAdvantages = req.CoreAdvantages
    leads.LeadsOrigin = req.LeadsOrigin
    leads.InvestRound = req.InvestRound
    leads.NeedIndustry = req.NeedIndustry
    leads.NeedFinance = req.NeedFinance
    leads.IndustryQsLabel = req.IndustryQsLabel
    leads.ImportanceNew = req.ImportanceNew
    leads.BureauLeadsId = req.BureauLeadsId
    leads.InvestorOriginPlaceCode = req.InvestorOriginPlaceCode
    leads.Tag = req.Tag
    leads.MaximumFactoryRent = req.MaximumFactoryRent
    leads.LeasedArea = req.LeasedArea
    leads.Period = req.Period
    leads.FactoryRentType = req.FactoryRentType
    leads.FloorHeightStarts = req.FloorHeightStarts
    leads.ParkLevel = req.ParkLevel
    leads.LandAreaRequired = req.LandAreaRequired
    leads.MinimumRent = req.MinimumRent
    leads.WhetherThereIsANeedForDe = req.WhetherThereIsANeedForDe
    leads.MinimumFactoryRent = req.MinimumFactoryRent
    leads.EndOfFloorHeight = req.EndOfFloorHeight
    leads.InvestmentArea = req.InvestmentArea
    leads.ParkType = req.ParkType
    leads.MaximumRent = req.MaximumRent
    leads.InterestRate = req.InterestRate
    leads.LandUse = req.LandUse
    leads.Quota = req.Quota
    leads.GuaranteeMethods = req.GuaranteeMethods
    leads.RentType = req.RentType
    leads.FinancingAmountMillionY = req.FinancingAmountMillionY
    leads.ExternalClueId = req.ExternalClueId
    leads.ClueResearcher = req.ClueResearcher
    leads.AssessmentTeam = req.AssessmentTeam
    leads.EnterpriseintermediaryContact = req.EnterpriseintermediaryContact
    leads.ProportionOfEquitySold = req.ProportionOfEquitySold
    leads.BureauClueProgress = req.BureauClueProgress
    leads.CurrentTransferMarketValue = req.CurrentTransferMarketValue
    leads.OperatingIncomeLastYear = req.OperatingIncomeLastYear
    leads.DetailedDescription = req.DetailedDescription
    leads.TimeRequirementsBefore = req.TimeRequirementsBefore
    leads.WhetherToGambleOnPerforman = req.WhetherToGambleOnPerforman
    leads.WhetherItIsAListedCompany = req.WhetherItIsAListedCompany
    leads.CityCode = req.CityCode
    leads.FinancialSituationInThePas = req.FinancialSituationInThePas
    leads.ActualControllerSituation = req.ActualControllerSituation
    leads.RequirementsForAcquirers = req.RequirementsForAcquirers
    leads.TransferConsiderationTransac = req.TransferConsiderationTransac
    leads.Seller = req.Seller
    leads.NeedFurtherUnderstandingOf = req.NeedFurtherUnderstandingOf
    leads.ProjectMatchmaker = req.ProjectMatchmaker
    leads.Country = req.Country
    leads.ListingPlate = req.ListingPlate
    leads.LatestValuation = req.LatestValuation
    leads.ReasonForSelling = req.ReasonForSelling
    leads.NetProfitLastYear = req.NetProfitLastYear
    leads.CanWePromiseToImplementPr = req.CanWePromiseToImplementPr
    leads.PerformanceCommitment = req.PerformanceCommitment
    leads.MainBusinessMainProductsAn = req.MainBusinessMainProductsAn
    leads.ProjectIntroduction = req.ProjectIntroduction
    leads.CityDockUserPhone = req.CityDockUserPhone
    leads.ProjectClueContactPerson = req.ProjectClueContactPerson
    leads.CityDockUserName = req.CityDockUserName
    leads.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := leads.Create(c); err != nil {
		return nil, err
	}

	return leads, nil
}

// Update 更新biz_leads
func (s *LeadsService) Update(c *gin.Context, req models.LeadsUpdateRequest) error {
	// 查找biz_leads记录
	leads := models.NewLeads()
	if err := leads.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新biz_leads信息
    leads.LeadsNo = req.LeadsNo
    leads.CompanyName = req.CompanyName
    leads.Region = req.Region
    leads.Mobile = req.Mobile
    leads.Telephone = req.Telephone
    leads.Email = req.Email
    leads.LeadsSource = req.LeadsSource
    leads.FlowStatus = req.FlowStatus
    leads.DistributionStatus = req.DistributionStatus
    leads.TransformationStatus = req.TransformationStatus
    leads.LeadsPool = req.LeadsPool
    leads.MarketingActivity = req.MarketingActivity
    leads.PromotionChannel = req.PromotionChannel
    leads.ReturnReason = req.ReturnReason
    leads.ReturnReasonDescription = req.ReturnReasonDescription
    leads.ReturnTime = req.ReturnTime
    leads.DistributionTime = req.DistributionTime
    leads.LeadsDescription = req.LeadsDescription
    leads.CustomerId = req.CustomerId
    leads.ContactId = req.ContactId
    leads.BusinessId = req.BusinessId
    leads.TransformationTime = req.TransformationTime
    leads.Sex = req.Sex
    leads.Department = req.Department
    leads.JobTitle = req.JobTitle
    leads.Age = req.Age
    leads.Birthday = req.Birthday
    leads.Wechat = req.Wechat
    leads.Qq = req.Qq
    leads.Fax = req.Fax
    leads.CompanyAddress = req.CompanyAddress
    leads.CompanyPhone = req.CompanyPhone
    leads.Industry = req.Industry
    leads.StaffSize = req.StaffSize
    leads.CompanyEmail = req.CompanyEmail
    leads.CompanyFax = req.CompanyFax
    leads.LastFlowTime = req.LastFlowTime
    leads.HandoverDescription = req.HandoverDescription
    leads.DistributionDescription = req.DistributionDescription
    leads.GetTime = req.GetTime
    leads.CountryRegion = req.CountryRegion
    leads.ReturnOriginPool = req.ReturnOriginPool
    leads.BusinessLicense = req.BusinessLicense
    leads.GmName = req.GmName
    leads.SalesExperience = req.SalesExperience
    leads.SalesDescription = req.SalesDescription
    leads.AnnualSales = req.AnnualSales
    leads.SameIndustry = req.SameIndustry
    leads.CompanyBrand = req.CompanyBrand
    leads.OtherInformation = req.OtherInformation
    leads.CompanyWeb = req.CompanyWeb
    leads.LicenseValid = req.LicenseValid
    leads.SalesYears = req.SalesYears
    leads.SalesDeptUserNum = req.SalesDeptUserNum
    leads.MarketingDeptUserNum = req.MarketingDeptUserNum
    leads.ServiceDeptUserNum = req.ServiceDeptUserNum
    leads.OtherDeptUserNum = req.OtherDeptUserNum
    leads.Role = req.Role
    leads.EnterpriseType = req.EnterpriseType
    leads.StartDate = req.StartDate
    leads.EndDate = req.EndDate
    leads.ActivityCity = req.ActivityCity
    leads.Sponsor = req.Sponsor
    leads.ActivityName = req.ActivityName
    leads.OwnBu = req.OwnBu
    leads.CooperationDirection = req.CooperationDirection
    leads.Solution = req.Solution
    leads.ContactStatus = req.ContactStatus
    leads.SignedProject = req.SignedProject
    leads.EstimatedGrossProfit = req.EstimatedGrossProfit
    leads.BaseCurrency = req.BaseCurrency
    leads.ExchangeRate = req.ExchangeRate
    leads.ExpectedSigningTime = req.ExpectedSigningTime
    leads.FollowUpPerson = req.FollowUpPerson
    leads.FeedbackResults = req.FeedbackResults
    leads.NewOldCustomer = req.NewOldCustomer
    leads.CustomerPartner = req.CustomerPartner
    leads.IndustryQs = req.IndustryQs
    leads.Stage = req.Stage
    leads.OriginDataOwnerId = req.OriginDataOwnerId
    leads.FirstGetFromPool = req.FirstGetFromPool
    leads.AuditStartUser = req.AuditStartUser
    leads.AuditCode = req.AuditCode
    leads.AuditStatus = req.AuditStatus
    leads.AuditStartTime = req.AuditStartTime
    leads.AuditEndTime = req.AuditEndTime
    leads.ActivityCityText = req.ActivityCityText
    leads.HistoryDataOwnerId = req.HistoryDataOwnerId
    leads.LeaderAssigned = req.LeaderAssigned
    leads.BureauLeaderAssigned = req.BureauLeaderAssigned
    leads.OriginNotes = req.OriginNotes
    leads.District = req.District
    leads.PlatformCompany = req.PlatformCompany
    leads.FromCityPlan = req.FromCityPlan
    leads.ServiceAdvisor = req.ServiceAdvisor
    leads.ServiceAdvisorContactWay = req.ServiceAdvisorContactWay
    leads.IndustryAffiliation = req.IndustryAffiliation
    leads.InvestorOriginPlace = req.InvestorOriginPlace
    leads.ForeignInvestment = req.ForeignInvestment
    leads.BpAttachment = req.BpAttachment
    leads.ProposedProjectName = req.ProposedProjectName
    leads.RegisterMoney = req.RegisterMoney
    leads.TotalInvestmentAmount = req.TotalInvestmentAmount
    leads.TaxEstimation = req.TaxEstimation
    leads.RecruitmentPlanContent = req.RecruitmentPlanContent
    leads.ShareParallelClasses = req.ShareParallelClasses
    leads.ResidencePlace = req.ResidencePlace
    leads.LeadsRecommender = req.LeadsRecommender
    leads.LeadsProvideOrg = req.LeadsProvideOrg
    leads.LeadsOwner = req.LeadsOwner
    leads.InputTime = req.InputTime
    leads.InvestProjectName = req.InvestProjectName
    leads.TechAdvanced = req.TechAdvanced
    leads.LandingContent = req.LandingContent
    leads.ToolCarrier = req.ToolCarrier
    leads.ProductionFactors = req.ProductionFactors
    leads.Financing = req.Financing
    leads.Market = req.Market
    leads.Policy = req.Policy
    leads.IsShareCyzb = req.IsShareCyzb
    leads.IsDistributeTask = req.IsDistributeTask
    leads.IsShareLeads = req.IsShareLeads
    leads.InvestCompany = req.InvestCompany
    leads.PhaseStatus = req.PhaseStatus
    leads.Importance = req.Importance
    leads.EnterpriseDemand = req.EnterpriseDemand
    leads.LeadDepartmentForClues = req.LeadDepartmentForClues
    leads.AnalysisContent = req.AnalysisContent
    leads.JudgmentOpinion = req.JudgmentOpinion
    leads.ReasonTermination = req.ReasonTermination
    leads.Destermiantion = req.Destermiantion
    leads.TechnicalAdvantages = req.TechnicalAdvantages
    leads.BpAttachments = req.BpAttachments
    leads.DockingType = req.DockingType
    leads.DockingInfo = req.DockingInfo
    leads.IndustryStatus = req.IndustryStatus
    leads.PecuniaryCondition = req.PecuniaryCondition
    leads.CoreAdvantages = req.CoreAdvantages
    leads.LeadsOrigin = req.LeadsOrigin
    leads.InvestRound = req.InvestRound
    leads.NeedIndustry = req.NeedIndustry
    leads.NeedFinance = req.NeedFinance
    leads.IndustryQsLabel = req.IndustryQsLabel
    leads.ImportanceNew = req.ImportanceNew
    leads.BureauLeadsId = req.BureauLeadsId
    leads.InvestorOriginPlaceCode = req.InvestorOriginPlaceCode
    leads.Tag = req.Tag
    leads.MaximumFactoryRent = req.MaximumFactoryRent
    leads.LeasedArea = req.LeasedArea
    leads.Period = req.Period
    leads.FactoryRentType = req.FactoryRentType
    leads.FloorHeightStarts = req.FloorHeightStarts
    leads.ParkLevel = req.ParkLevel
    leads.LandAreaRequired = req.LandAreaRequired
    leads.MinimumRent = req.MinimumRent
    leads.WhetherThereIsANeedForDe = req.WhetherThereIsANeedForDe
    leads.MinimumFactoryRent = req.MinimumFactoryRent
    leads.EndOfFloorHeight = req.EndOfFloorHeight
    leads.InvestmentArea = req.InvestmentArea
    leads.ParkType = req.ParkType
    leads.MaximumRent = req.MaximumRent
    leads.InterestRate = req.InterestRate
    leads.LandUse = req.LandUse
    leads.Quota = req.Quota
    leads.GuaranteeMethods = req.GuaranteeMethods
    leads.RentType = req.RentType
    leads.FinancingAmountMillionY = req.FinancingAmountMillionY
    leads.ExternalClueId = req.ExternalClueId
    leads.ClueResearcher = req.ClueResearcher
    leads.AssessmentTeam = req.AssessmentTeam
    leads.EnterpriseintermediaryContact = req.EnterpriseintermediaryContact
    leads.ProportionOfEquitySold = req.ProportionOfEquitySold
    leads.BureauClueProgress = req.BureauClueProgress
    leads.CurrentTransferMarketValue = req.CurrentTransferMarketValue
    leads.OperatingIncomeLastYear = req.OperatingIncomeLastYear
    leads.DetailedDescription = req.DetailedDescription
    leads.TimeRequirementsBefore = req.TimeRequirementsBefore
    leads.WhetherToGambleOnPerforman = req.WhetherToGambleOnPerforman
    leads.WhetherItIsAListedCompany = req.WhetherItIsAListedCompany
    leads.CityCode = req.CityCode
    leads.FinancialSituationInThePas = req.FinancialSituationInThePas
    leads.ActualControllerSituation = req.ActualControllerSituation
    leads.RequirementsForAcquirers = req.RequirementsForAcquirers
    leads.TransferConsiderationTransac = req.TransferConsiderationTransac
    leads.Seller = req.Seller
    leads.NeedFurtherUnderstandingOf = req.NeedFurtherUnderstandingOf
    leads.ProjectMatchmaker = req.ProjectMatchmaker
    leads.Country = req.Country
    leads.ListingPlate = req.ListingPlate
    leads.LatestValuation = req.LatestValuation
    leads.ReasonForSelling = req.ReasonForSelling
    leads.NetProfitLastYear = req.NetProfitLastYear
    leads.CanWePromiseToImplementPr = req.CanWePromiseToImplementPr
    leads.PerformanceCommitment = req.PerformanceCommitment
    leads.MainBusinessMainProductsAn = req.MainBusinessMainProductsAn
    leads.ProjectIntroduction = req.ProjectIntroduction
    leads.CityDockUserPhone = req.CityDockUserPhone
    leads.ProjectClueContactPerson = req.ProjectClueContactPerson
    leads.CityDockUserName = req.CityDockUserName
    leads.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := leads.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除biz_leads
func (s *LeadsService) Delete(c *gin.Context, id int64) error {
	// 查找biz_leads记录
	leads := models.NewLeads()
	if err := leads.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := leads.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取biz_leads
func (s *LeadsService) GetByID(c *gin.Context, id int64) (*models.Leads, error) {
	// 查找biz_leads记录
	leads := models.NewLeads()
	if err := leads.GetByID(c, id); err != nil {
		return nil, err
	}

	return leads, nil
}

// List biz_leads列表（分页查询）
func (s *LeadsService) List(c *gin.Context, req models.LeadsListRequest) (*models.LeadsList, int64, error) {
	// 获取总数
	leadsList := models.NewLeadsList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := leadsList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = leadsList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return leadsList, total, nil
}