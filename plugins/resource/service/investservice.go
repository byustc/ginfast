package service

import (
	"gin-fast/plugins/resource/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// InvestService res_invest服务
type InvestService struct{}

// NewInvestService 创建res_invest服务
func NewInvestService() *InvestService {
	return &InvestService{}
}

// Create 创建res_invest
func (s *InvestService) Create(c *gin.Context, req models.InvestCreateRequest) (*models.Invest, error) {
	// 创建res_invest记录
	invest := models.NewInvest()
    invest.InvestNo = req.InvestNo
    invest.BriefName = req.BriefName
    invest.FullName = req.FullName
    invest.ComOtherName = req.ComOtherName
    invest.Logo = req.Logo
    invest.ComType = req.ComType
    invest.Website = req.Website
    invest.BriefIntro = req.BriefIntro
    invest.DetailIntro = req.DetailIntro
    invest.EstablishTime = req.EstablishTime
    invest.InvestRound = req.InvestRound
    invest.OperatingStatus = req.OperatingStatus
    invest.Staff = req.Staff
    invest.BaikeLink = req.BaikeLink
    invest.Industry = req.Industry
    invest.IndustryName = req.IndustryName
    invest.MarketValuationRmb = req.MarketValuationRmb
    invest.Tags = req.Tags
    invest.RegFullName = req.RegFullName
    invest.RegTime = req.RegTime
    invest.RegLocation = req.RegLocation
    invest.RegCapital = req.RegCapital
    invest.RegCapitalCurrency = req.RegCapitalCurrency
    invest.PaidCapital = req.PaidCapital
    invest.PaidCapitalCurrency = req.PaidCapitalCurrency
    invest.RegType = req.RegType
    invest.LegalRepresent = req.LegalRepresent
    invest.RegNumber = req.RegNumber
    invest.CreditNumber = req.CreditNumber
    invest.BusinessPeriod = req.BusinessPeriod
    invest.BusinessScope = req.BusinessScope
    invest.RegIndustry = req.RegIndustry
    invest.RegIndustryName = req.RegIndustryName
    invest.RegSubIndustry = req.RegSubIndustry
    invest.RegSubIndustryName = req.RegSubIndustryName
    invest.RegMiddleIndustry = req.RegMiddleIndustry
    invest.RegMiddleIndustryName = req.RegMiddleIndustryName
    invest.RegSmallIndustry = req.RegSmallIndustry
    invest.RegSmallIndustryName = req.RegSmallIndustryName
    invest.RegInstitute = req.RegInstitute
    invest.ApproveTime = req.ApproveTime
    invest.Location = req.Location
    invest.OrgNature = req.OrgNature
    invest.IsCvc = req.IsCvc
    invest.CvcDesc = req.CvcDesc
    invest.MgtScale = req.MgtScale
    invest.InvestIndustry = req.InvestIndustry
    invest.SingleInvestAmount = req.SingleInvestAmount
    invest.RoundPrefer = req.RoundPrefer
    invest.InvestMethod = req.InvestMethod
    invest.InvestCase = req.InvestCase
    invest.StarProj = req.StarProj
    invest.MainDemands = req.MainDemands
    invest.Remark = req.Remark
    invest.InvestRoundName = req.InvestRoundName
    invest.OperatingStatusName = req.OperatingStatusName
    invest.RegCapitalCurrencyName = req.RegCapitalCurrencyName
    invest.PaidCapitalCurrencyName = req.PaidCapitalCurrencyName
    invest.RegTypeName = req.RegTypeName
    invest.InvestType = req.InvestType
    invest.IsBlacklist = req.IsBlacklist
    invest.NeedInvestorType = req.NeedInvestorType
    invest.InvestableBalance = req.InvestableBalance
    invest.DataSource = req.DataSource
    invest.InvestmentField = req.InvestmentField
    invest.Tag = req.Tag
    invest.DockingPerson = req.DockingPerson
    invest.DockingPersonnel = req.DockingPersonnel
    invest.Level = req.Level
    invest.Partner = req.Partner
    invest.DockingContent = req.DockingContent
    invest.WhetherConfidentialityAgre = req.WhetherConfidentialityAgre
    invest.Duties = req.Duties
    invest.Mailbox = req.Mailbox
    invest.CompletionReverseInvest = req.CompletionReverseInvest
    invest.AvailableBalance = req.AvailableBalance
    invest.Dockee = req.Dockee
    invest.EntryPerson = req.EntryPerson
    invest.FundName = req.FundName
    invest.Contact = req.Contact
    invest.CounterinvestmentRequirements = req.CounterinvestmentRequirements
    invest.WarehousingTime = req.WarehousingTime
    invest.ContactInformation = req.ContactInformation
    invest.InvestmentPreferences = req.InvestmentPreferences
    invest.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := invest.Create(c); err != nil {
		return nil, err
	}

	return invest, nil
}

// Update 更新res_invest
func (s *InvestService) Update(c *gin.Context, req models.InvestUpdateRequest) error {
	// 查找res_invest记录
	invest := models.NewInvest()
	if err := invest.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新res_invest信息
    invest.InvestNo = req.InvestNo
    invest.BriefName = req.BriefName
    invest.FullName = req.FullName
    invest.ComOtherName = req.ComOtherName
    invest.Logo = req.Logo
    invest.ComType = req.ComType
    invest.Website = req.Website
    invest.BriefIntro = req.BriefIntro
    invest.DetailIntro = req.DetailIntro
    invest.EstablishTime = req.EstablishTime
    invest.InvestRound = req.InvestRound
    invest.OperatingStatus = req.OperatingStatus
    invest.Staff = req.Staff
    invest.BaikeLink = req.BaikeLink
    invest.Industry = req.Industry
    invest.IndustryName = req.IndustryName
    invest.MarketValuationRmb = req.MarketValuationRmb
    invest.Tags = req.Tags
    invest.RegFullName = req.RegFullName
    invest.RegTime = req.RegTime
    invest.RegLocation = req.RegLocation
    invest.RegCapital = req.RegCapital
    invest.RegCapitalCurrency = req.RegCapitalCurrency
    invest.PaidCapital = req.PaidCapital
    invest.PaidCapitalCurrency = req.PaidCapitalCurrency
    invest.RegType = req.RegType
    invest.LegalRepresent = req.LegalRepresent
    invest.RegNumber = req.RegNumber
    invest.CreditNumber = req.CreditNumber
    invest.BusinessPeriod = req.BusinessPeriod
    invest.BusinessScope = req.BusinessScope
    invest.RegIndustry = req.RegIndustry
    invest.RegIndustryName = req.RegIndustryName
    invest.RegSubIndustry = req.RegSubIndustry
    invest.RegSubIndustryName = req.RegSubIndustryName
    invest.RegMiddleIndustry = req.RegMiddleIndustry
    invest.RegMiddleIndustryName = req.RegMiddleIndustryName
    invest.RegSmallIndustry = req.RegSmallIndustry
    invest.RegSmallIndustryName = req.RegSmallIndustryName
    invest.RegInstitute = req.RegInstitute
    invest.ApproveTime = req.ApproveTime
    invest.Location = req.Location
    invest.OrgNature = req.OrgNature
    invest.IsCvc = req.IsCvc
    invest.CvcDesc = req.CvcDesc
    invest.MgtScale = req.MgtScale
    invest.InvestIndustry = req.InvestIndustry
    invest.SingleInvestAmount = req.SingleInvestAmount
    invest.RoundPrefer = req.RoundPrefer
    invest.InvestMethod = req.InvestMethod
    invest.InvestCase = req.InvestCase
    invest.StarProj = req.StarProj
    invest.MainDemands = req.MainDemands
    invest.Remark = req.Remark
    invest.InvestRoundName = req.InvestRoundName
    invest.OperatingStatusName = req.OperatingStatusName
    invest.RegCapitalCurrencyName = req.RegCapitalCurrencyName
    invest.PaidCapitalCurrencyName = req.PaidCapitalCurrencyName
    invest.RegTypeName = req.RegTypeName
    invest.InvestType = req.InvestType
    invest.IsBlacklist = req.IsBlacklist
    invest.NeedInvestorType = req.NeedInvestorType
    invest.InvestableBalance = req.InvestableBalance
    invest.DataSource = req.DataSource
    invest.InvestmentField = req.InvestmentField
    invest.Tag = req.Tag
    invest.DockingPerson = req.DockingPerson
    invest.DockingPersonnel = req.DockingPersonnel
    invest.Level = req.Level
    invest.Partner = req.Partner
    invest.DockingContent = req.DockingContent
    invest.WhetherConfidentialityAgre = req.WhetherConfidentialityAgre
    invest.Duties = req.Duties
    invest.Mailbox = req.Mailbox
    invest.CompletionReverseInvest = req.CompletionReverseInvest
    invest.AvailableBalance = req.AvailableBalance
    invest.Dockee = req.Dockee
    invest.EntryPerson = req.EntryPerson
    invest.FundName = req.FundName
    invest.Contact = req.Contact
    invest.CounterinvestmentRequirements = req.CounterinvestmentRequirements
    invest.WarehousingTime = req.WarehousingTime
    invest.ContactInformation = req.ContactInformation
    invest.InvestmentPreferences = req.InvestmentPreferences
    invest.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := invest.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除res_invest
func (s *InvestService) Delete(c *gin.Context, id int64) error {
	// 查找res_invest记录
	invest := models.NewInvest()
	if err := invest.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := invest.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取res_invest
func (s *InvestService) GetByID(c *gin.Context, id int64) (*models.Invest, error) {
	// 查找res_invest记录
	invest := models.NewInvest()
	if err := invest.GetByID(c, id); err != nil {
		return nil, err
	}

	return invest, nil
}

// List res_invest列表（分页查询）
func (s *InvestService) List(c *gin.Context, req models.InvestListRequest) (*models.InvestList, int64, error) {
	// 获取总数
	investList := models.NewInvestList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := investList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = investList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return investList, total, nil
}