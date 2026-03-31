package service

import (
	"gin-fast/plugins/resource/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// FinanceService res_finance服务
type FinanceService struct{}

// NewFinanceService 创建res_finance服务
func NewFinanceService() *FinanceService {
	return &FinanceService{}
}

// Create 创建res_finance
func (s *FinanceService) Create(c *gin.Context, req models.FinanceCreateRequest) (*models.Finance, error) {
	// 创建res_finance记录
	finance := models.NewFinance()
    finance.FinanceNo = req.FinanceNo
    finance.BriefName = req.BriefName
    finance.FullName = req.FullName
    finance.ComOtherName = req.ComOtherName
    finance.Logo = req.Logo
    finance.ComType = req.ComType
    finance.Website = req.Website
    finance.BriefIntro = req.BriefIntro
    finance.DetailIntro = req.DetailIntro
    finance.EstablishTime = req.EstablishTime
    finance.InvestRound = req.InvestRound
    finance.OperatingStatus = req.OperatingStatus
    finance.Staff = req.Staff
    finance.BaikeLink = req.BaikeLink
    finance.Industry = req.Industry
    finance.IndustryName = req.IndustryName
    finance.MarketValuationRmb = req.MarketValuationRmb
    finance.Tags = req.Tags
    finance.RegFullName = req.RegFullName
    finance.RegTime = req.RegTime
    finance.RegLocation = req.RegLocation
    finance.RegCapital = req.RegCapital
    finance.RegCapitalCurrency = req.RegCapitalCurrency
    finance.PaidCapital = req.PaidCapital
    finance.PaidCapitalCurrency = req.PaidCapitalCurrency
    finance.RegType = req.RegType
    finance.LegalRepresent = req.LegalRepresent
    finance.RegNumber = req.RegNumber
    finance.CreditNumber = req.CreditNumber
    finance.BusinessPeriod = req.BusinessPeriod
    finance.BusinessScope = req.BusinessScope
    finance.RegIndustry = req.RegIndustry
    finance.RegIndustryName = req.RegIndustryName
    finance.RegSubIndustry = req.RegSubIndustry
    finance.RegSubIndustryName = req.RegSubIndustryName
    finance.RegMiddleIndustry = req.RegMiddleIndustry
    finance.RegMiddleIndustryName = req.RegMiddleIndustryName
    finance.RegSmallIndustry = req.RegSmallIndustry
    finance.RegSmallIndustryName = req.RegSmallIndustryName
    finance.RegInstitute = req.RegInstitute
    finance.ApproveTime = req.ApproveTime
    finance.Location = req.Location
    finance.FinType = req.FinType
    finance.FinNature = req.FinNature
    finance.IsConfidential = req.IsConfidential
    finance.CooperationTime = req.CooperationTime
    finance.Cooperativeness = req.Cooperativeness
    finance.CooperationScope = req.CooperationScope
    finance.Remark = req.Remark
    finance.InvestRoundName = req.InvestRoundName
    finance.OperatingStatusName = req.OperatingStatusName
    finance.RegCapitalCurrencyName = req.RegCapitalCurrencyName
    finance.PaidCapitalCurrencyName = req.PaidCapitalCurrencyName
    finance.RegTypeName = req.RegTypeName
    finance.DataSource = req.DataSource
    finance.IsBlacklist = req.IsBlacklist
    finance.PositionNum = req.PositionNum
    finance.PositionAdvancedNum = req.PositionAdvancedNum
    finance.Tag = req.Tag
    finance.DockingPerson = req.DockingPerson
    finance.CorrespondingSourceId = req.CorrespondingSourceId
    finance.DockingPersonnel = req.DockingPersonnel
    finance.Level = req.Level
    finance.Partner = req.Partner
    finance.WarehousingTime = req.WarehousingTime
    finance.Contact = req.Contact
    finance.ContactNumber = req.ContactNumber
    finance.EntryPerson = req.EntryPerson
    finance.ContactInformation = req.ContactInformation
    finance.Mailbox = req.Mailbox
    finance.Duties = req.Duties
    finance.ActivitiesUndertaken = req.ActivitiesUndertaken
    finance.LogoPicture = req.LogoPicture
    finance.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := finance.Create(c); err != nil {
		return nil, err
	}

	return finance, nil
}

// Update 更新res_finance
func (s *FinanceService) Update(c *gin.Context, req models.FinanceUpdateRequest) error {
	// 查找res_finance记录
	finance := models.NewFinance()
	if err := finance.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新res_finance信息
    finance.FinanceNo = req.FinanceNo
    finance.BriefName = req.BriefName
    finance.FullName = req.FullName
    finance.ComOtherName = req.ComOtherName
    finance.Logo = req.Logo
    finance.ComType = req.ComType
    finance.Website = req.Website
    finance.BriefIntro = req.BriefIntro
    finance.DetailIntro = req.DetailIntro
    finance.EstablishTime = req.EstablishTime
    finance.InvestRound = req.InvestRound
    finance.OperatingStatus = req.OperatingStatus
    finance.Staff = req.Staff
    finance.BaikeLink = req.BaikeLink
    finance.Industry = req.Industry
    finance.IndustryName = req.IndustryName
    finance.MarketValuationRmb = req.MarketValuationRmb
    finance.Tags = req.Tags
    finance.RegFullName = req.RegFullName
    finance.RegTime = req.RegTime
    finance.RegLocation = req.RegLocation
    finance.RegCapital = req.RegCapital
    finance.RegCapitalCurrency = req.RegCapitalCurrency
    finance.PaidCapital = req.PaidCapital
    finance.PaidCapitalCurrency = req.PaidCapitalCurrency
    finance.RegType = req.RegType
    finance.LegalRepresent = req.LegalRepresent
    finance.RegNumber = req.RegNumber
    finance.CreditNumber = req.CreditNumber
    finance.BusinessPeriod = req.BusinessPeriod
    finance.BusinessScope = req.BusinessScope
    finance.RegIndustry = req.RegIndustry
    finance.RegIndustryName = req.RegIndustryName
    finance.RegSubIndustry = req.RegSubIndustry
    finance.RegSubIndustryName = req.RegSubIndustryName
    finance.RegMiddleIndustry = req.RegMiddleIndustry
    finance.RegMiddleIndustryName = req.RegMiddleIndustryName
    finance.RegSmallIndustry = req.RegSmallIndustry
    finance.RegSmallIndustryName = req.RegSmallIndustryName
    finance.RegInstitute = req.RegInstitute
    finance.ApproveTime = req.ApproveTime
    finance.Location = req.Location
    finance.FinType = req.FinType
    finance.FinNature = req.FinNature
    finance.IsConfidential = req.IsConfidential
    finance.CooperationTime = req.CooperationTime
    finance.Cooperativeness = req.Cooperativeness
    finance.CooperationScope = req.CooperationScope
    finance.Remark = req.Remark
    finance.InvestRoundName = req.InvestRoundName
    finance.OperatingStatusName = req.OperatingStatusName
    finance.RegCapitalCurrencyName = req.RegCapitalCurrencyName
    finance.PaidCapitalCurrencyName = req.PaidCapitalCurrencyName
    finance.RegTypeName = req.RegTypeName
    finance.DataSource = req.DataSource
    finance.IsBlacklist = req.IsBlacklist
    finance.PositionNum = req.PositionNum
    finance.PositionAdvancedNum = req.PositionAdvancedNum
    finance.Tag = req.Tag
    finance.DockingPerson = req.DockingPerson
    finance.CorrespondingSourceId = req.CorrespondingSourceId
    finance.DockingPersonnel = req.DockingPersonnel
    finance.Level = req.Level
    finance.Partner = req.Partner
    finance.WarehousingTime = req.WarehousingTime
    finance.Contact = req.Contact
    finance.ContactNumber = req.ContactNumber
    finance.EntryPerson = req.EntryPerson
    finance.ContactInformation = req.ContactInformation
    finance.Mailbox = req.Mailbox
    finance.Duties = req.Duties
    finance.ActivitiesUndertaken = req.ActivitiesUndertaken
    finance.LogoPicture = req.LogoPicture
    finance.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := finance.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除res_finance
func (s *FinanceService) Delete(c *gin.Context, id int64) error {
	// 查找res_finance记录
	finance := models.NewFinance()
	if err := finance.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := finance.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取res_finance
func (s *FinanceService) GetByID(c *gin.Context, id int64) (*models.Finance, error) {
	// 查找res_finance记录
	finance := models.NewFinance()
	if err := finance.GetByID(c, id); err != nil {
		return nil, err
	}

	return finance, nil
}

// List res_finance列表（分页查询）
func (s *FinanceService) List(c *gin.Context, req models.FinanceListRequest) (*models.FinanceList, int64, error) {
	// 获取总数
	financeList := models.NewFinanceList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := financeList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = financeList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return financeList, total, nil
}