package service

import (
	"gin-fast/plugins/resource/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// EnterpriseService res_enterprise服务
type EnterpriseService struct{}

// NewEnterpriseService 创建res_enterprise服务
func NewEnterpriseService() *EnterpriseService {
	return &EnterpriseService{}
}

// Create 创建res_enterprise
func (s *EnterpriseService) Create(c *gin.Context, req models.EnterpriseCreateRequest) (*models.Enterprise, error) {
	// 创建res_enterprise记录
	enterprise := models.NewEnterprise()
    enterprise.EnterpriseNo = req.EnterpriseNo
    enterprise.BriefName = req.BriefName
    enterprise.FullName = req.FullName
    enterprise.ComOtherName = req.ComOtherName
    enterprise.Logo = req.Logo
    enterprise.ComType = req.ComType
    enterprise.Website = req.Website
    enterprise.BriefIntro = req.BriefIntro
    enterprise.DetailIntro = req.DetailIntro
    enterprise.EstablishTime = req.EstablishTime
    enterprise.InvestRound = req.InvestRound
    enterprise.OperatingStatus = req.OperatingStatus
    enterprise.Staff = req.Staff
    enterprise.BaikeLink = req.BaikeLink
    enterprise.Industry = req.Industry
    enterprise.IndustryName = req.IndustryName
    enterprise.MarketValuationRmb = req.MarketValuationRmb
    enterprise.Tags = req.Tags
    enterprise.LatestInvestmentId = req.LatestInvestmentId
    enterprise.LatestInvestTime = req.LatestInvestTime
    enterprise.LatestExposureTime = req.LatestExposureTime
    enterprise.LatestInvestRound = req.LatestInvestRound
    enterprise.LatestInvestAmountType = req.LatestInvestAmountType
    enterprise.LatestInvestAmount = req.LatestInvestAmount
    enterprise.UnifiedInvestAmount = req.UnifiedInvestAmount
    enterprise.LatestInvestCurrency = req.LatestInvestCurrency
    enterprise.LatestInvestor = req.LatestInvestor
    enterprise.InvestmentNum = req.InvestmentNum
    enterprise.TotalInvestors = req.TotalInvestors
    enterprise.TotalInvestAmount = req.TotalInvestAmount
    enterprise.ExternalInvestmentCount = req.ExternalInvestmentCount
    enterprise.CompanyContact = req.CompanyContact
    enterprise.OrganizationalStructure = req.OrganizationalStructure
    enterprise.SecuritiesInfo = req.SecuritiesInfo
    enterprise.RegFullName = req.RegFullName
    enterprise.RegTime = req.RegTime
    enterprise.RegLocation = req.RegLocation
    enterprise.RegCapital = req.RegCapital
    enterprise.RegCapitalCurrency = req.RegCapitalCurrency
    enterprise.PaidCapital = req.PaidCapital
    enterprise.PaidCapitalCurrency = req.PaidCapitalCurrency
    enterprise.RegType = req.RegType
    enterprise.LegalRepresent = req.LegalRepresent
    enterprise.RegNumber = req.RegNumber
    enterprise.CreditNumber = req.CreditNumber
    enterprise.BusinessPeriod = req.BusinessPeriod
    enterprise.BusinessScope = req.BusinessScope
    enterprise.RegIndustry = req.RegIndustry
    enterprise.RegIndustryName = req.RegIndustryName
    enterprise.RegSubIndustry = req.RegSubIndustry
    enterprise.RegSubIndustryName = req.RegSubIndustryName
    enterprise.RegMiddleIndustry = req.RegMiddleIndustry
    enterprise.RegMiddleIndustryName = req.RegMiddleIndustryName
    enterprise.RegSmallIndustry = req.RegSmallIndustry
    enterprise.RegSmallIndustryName = req.RegSmallIndustryName
    enterprise.RegInstitute = req.RegInstitute
    enterprise.ApproveTime = req.ApproveTime
    enterprise.InvestRoundName = req.InvestRoundName
    enterprise.OperatingStatusName = req.OperatingStatusName
    enterprise.LatestInvestRoundName = req.LatestInvestRoundName
    enterprise.RegTypeName = req.RegTypeName
    enterprise.RegCapitalCurrencyName = req.RegCapitalCurrencyName
    enterprise.PaidCapitalCurrencyName = req.PaidCapitalCurrencyName
    enterprise.LatestInvestAmountTypeName = req.LatestInvestAmountTypeName
    enterprise.LatestInvestCurrencyName = req.LatestInvestCurrencyName
    enterprise.Remark = req.Remark
    enterprise.IndustryStatus = req.IndustryStatus
    enterprise.PecuniaryCondition = req.PecuniaryCondition
    enterprise.TechAdvanced = req.TechAdvanced
    enterprise.DataSource = req.DataSource
    enterprise.IsBlacklist = req.IsBlacklist
    enterprise.Tag = req.Tag
    enterprise.LogoPicture = req.LogoPicture
    enterprise.OrganizationalStructureNew = req.OrganizationalStructureNew
    enterprise.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := enterprise.Create(c); err != nil {
		return nil, err
	}

	return enterprise, nil
}

// Update 更新res_enterprise
func (s *EnterpriseService) Update(c *gin.Context, req models.EnterpriseUpdateRequest) error {
	// 查找res_enterprise记录
	enterprise := models.NewEnterprise()
	if err := enterprise.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新res_enterprise信息
    enterprise.EnterpriseNo = req.EnterpriseNo
    enterprise.BriefName = req.BriefName
    enterprise.FullName = req.FullName
    enterprise.ComOtherName = req.ComOtherName
    enterprise.Logo = req.Logo
    enterprise.ComType = req.ComType
    enterprise.Website = req.Website
    enterprise.BriefIntro = req.BriefIntro
    enterprise.DetailIntro = req.DetailIntro
    enterprise.EstablishTime = req.EstablishTime
    enterprise.InvestRound = req.InvestRound
    enterprise.OperatingStatus = req.OperatingStatus
    enterprise.Staff = req.Staff
    enterprise.BaikeLink = req.BaikeLink
    enterprise.Industry = req.Industry
    enterprise.IndustryName = req.IndustryName
    enterprise.MarketValuationRmb = req.MarketValuationRmb
    enterprise.Tags = req.Tags
    enterprise.LatestInvestmentId = req.LatestInvestmentId
    enterprise.LatestInvestTime = req.LatestInvestTime
    enterprise.LatestExposureTime = req.LatestExposureTime
    enterprise.LatestInvestRound = req.LatestInvestRound
    enterprise.LatestInvestAmountType = req.LatestInvestAmountType
    enterprise.LatestInvestAmount = req.LatestInvestAmount
    enterprise.UnifiedInvestAmount = req.UnifiedInvestAmount
    enterprise.LatestInvestCurrency = req.LatestInvestCurrency
    enterprise.LatestInvestor = req.LatestInvestor
    enterprise.InvestmentNum = req.InvestmentNum
    enterprise.TotalInvestors = req.TotalInvestors
    enterprise.TotalInvestAmount = req.TotalInvestAmount
    enterprise.ExternalInvestmentCount = req.ExternalInvestmentCount
    enterprise.CompanyContact = req.CompanyContact
    enterprise.OrganizationalStructure = req.OrganizationalStructure
    enterprise.SecuritiesInfo = req.SecuritiesInfo
    enterprise.RegFullName = req.RegFullName
    enterprise.RegTime = req.RegTime
    enterprise.RegLocation = req.RegLocation
    enterprise.RegCapital = req.RegCapital
    enterprise.RegCapitalCurrency = req.RegCapitalCurrency
    enterprise.PaidCapital = req.PaidCapital
    enterprise.PaidCapitalCurrency = req.PaidCapitalCurrency
    enterprise.RegType = req.RegType
    enterprise.LegalRepresent = req.LegalRepresent
    enterprise.RegNumber = req.RegNumber
    enterprise.CreditNumber = req.CreditNumber
    enterprise.BusinessPeriod = req.BusinessPeriod
    enterprise.BusinessScope = req.BusinessScope
    enterprise.RegIndustry = req.RegIndustry
    enterprise.RegIndustryName = req.RegIndustryName
    enterprise.RegSubIndustry = req.RegSubIndustry
    enterprise.RegSubIndustryName = req.RegSubIndustryName
    enterprise.RegMiddleIndustry = req.RegMiddleIndustry
    enterprise.RegMiddleIndustryName = req.RegMiddleIndustryName
    enterprise.RegSmallIndustry = req.RegSmallIndustry
    enterprise.RegSmallIndustryName = req.RegSmallIndustryName
    enterprise.RegInstitute = req.RegInstitute
    enterprise.ApproveTime = req.ApproveTime
    enterprise.InvestRoundName = req.InvestRoundName
    enterprise.OperatingStatusName = req.OperatingStatusName
    enterprise.LatestInvestRoundName = req.LatestInvestRoundName
    enterprise.RegTypeName = req.RegTypeName
    enterprise.RegCapitalCurrencyName = req.RegCapitalCurrencyName
    enterprise.PaidCapitalCurrencyName = req.PaidCapitalCurrencyName
    enterprise.LatestInvestAmountTypeName = req.LatestInvestAmountTypeName
    enterprise.LatestInvestCurrencyName = req.LatestInvestCurrencyName
    enterprise.Remark = req.Remark
    enterprise.IndustryStatus = req.IndustryStatus
    enterprise.PecuniaryCondition = req.PecuniaryCondition
    enterprise.TechAdvanced = req.TechAdvanced
    enterprise.DataSource = req.DataSource
    enterprise.IsBlacklist = req.IsBlacklist
    enterprise.Tag = req.Tag
    enterprise.LogoPicture = req.LogoPicture
    enterprise.OrganizationalStructureNew = req.OrganizationalStructureNew
    enterprise.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := enterprise.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除res_enterprise
func (s *EnterpriseService) Delete(c *gin.Context, id int64) error {
	// 查找res_enterprise记录
	enterprise := models.NewEnterprise()
	if err := enterprise.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := enterprise.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取res_enterprise
func (s *EnterpriseService) GetByID(c *gin.Context, id int64) (*models.Enterprise, error) {
	// 查找res_enterprise记录
	enterprise := models.NewEnterprise()
	if err := enterprise.GetByID(c, id); err != nil {
		return nil, err
	}

	return enterprise, nil
}

// List res_enterprise列表（分页查询）
func (s *EnterpriseService) List(c *gin.Context, req models.EnterpriseListRequest) (*models.EnterpriseList, int64, error) {
	// 获取总数
	enterpriseList := models.NewEnterpriseList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := enterpriseList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = enterpriseList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return enterpriseList, total, nil
}