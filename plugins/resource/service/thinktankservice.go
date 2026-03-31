package service

import (
	"gin-fast/plugins/resource/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// ThinktankService res_thinktank服务
type ThinktankService struct{}

// NewThinktankService 创建res_thinktank服务
func NewThinktankService() *ThinktankService {
	return &ThinktankService{}
}

// Create 创建res_thinktank
func (s *ThinktankService) Create(c *gin.Context, req models.ThinktankCreateRequest) (*models.Thinktank, error) {
	// 创建res_thinktank记录
	thinktank := models.NewThinktank()
    thinktank.ThinktankNo = req.ThinktankNo
    thinktank.BriefName = req.BriefName
    thinktank.FullName = req.FullName
    thinktank.ComOtherName = req.ComOtherName
    thinktank.Logo = req.Logo
    thinktank.ComType = req.ComType
    thinktank.Website = req.Website
    thinktank.BriefIntro = req.BriefIntro
    thinktank.DetailIntro = req.DetailIntro
    thinktank.EstablishTime = req.EstablishTime
    thinktank.InvestRound = req.InvestRound
    thinktank.InvestRoundName = req.InvestRoundName
    thinktank.OperatingStatus = req.OperatingStatus
    thinktank.OperatingStatusName = req.OperatingStatusName
    thinktank.Staff = req.Staff
    thinktank.BaikeLink = req.BaikeLink
    thinktank.Industry = req.Industry
    thinktank.IndustryName = req.IndustryName
    thinktank.MarketValuationRmb = req.MarketValuationRmb
    thinktank.Tags = req.Tags
    thinktank.Location = req.Location
    thinktank.Address = req.Address
    thinktank.CooperationType = req.CooperationType
    thinktank.CooperationForm = req.CooperationForm
    thinktank.ResourceAnalysis = req.ResourceAnalysis
    thinktank.ConstructionContent = req.ConstructionContent
    thinktank.IndustryFocus = req.IndustryFocus
    thinktank.CooperationUnits = req.CooperationUnits
    thinktank.BusinessFocus = req.BusinessFocus
    thinktank.OperationMode = req.OperationMode
    thinktank.ResponsibleDepartment = req.ResponsibleDepartment
    thinktank.ResearchTopic = req.ResearchTopic
    thinktank.Remark = req.Remark
    thinktank.RegFullName = req.RegFullName
    thinktank.RegTime = req.RegTime
    thinktank.RegLocation = req.RegLocation
    thinktank.RegCapital = req.RegCapital
    thinktank.RegCapitalCurrency = req.RegCapitalCurrency
    thinktank.RegCapitalCurrencyName = req.RegCapitalCurrencyName
    thinktank.PaidCapital = req.PaidCapital
    thinktank.PaidCapitalCurrency = req.PaidCapitalCurrency
    thinktank.PaidCapitalCurrencyName = req.PaidCapitalCurrencyName
    thinktank.RegType = req.RegType
    thinktank.RegTypeName = req.RegTypeName
    thinktank.LegalRepresent = req.LegalRepresent
    thinktank.RegNumber = req.RegNumber
    thinktank.CreditNumber = req.CreditNumber
    thinktank.BusinessPeriod = req.BusinessPeriod
    thinktank.BusinessScope = req.BusinessScope
    thinktank.RegIndustry = req.RegIndustry
    thinktank.RegIndustryName = req.RegIndustryName
    thinktank.RegSubIndustry = req.RegSubIndustry
    thinktank.RegSubIndustryName = req.RegSubIndustryName
    thinktank.RegMiddleIndustry = req.RegMiddleIndustry
    thinktank.RegMiddleIndustryName = req.RegMiddleIndustryName
    thinktank.RegSmallIndustry = req.RegSmallIndustry
    thinktank.RegSmallIndustryName = req.RegSmallIndustryName
    thinktank.RegInstitute = req.RegInstitute
    thinktank.ApproveTime = req.ApproveTime
    thinktank.DataSource = req.DataSource
    thinktank.IsBlacklist = req.IsBlacklist
    thinktank.Tag = req.Tag
    thinktank.Partner = req.Partner
    thinktank.PlatformResources = req.PlatformResources
    thinktank.Contact = req.Contact
    thinktank.ContactNumber = req.ContactNumber
    thinktank.Duties = req.Duties
    thinktank.LocationChina = req.LocationChina
    thinktank.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := thinktank.Create(c); err != nil {
		return nil, err
	}

	return thinktank, nil
}

// Update 更新res_thinktank
func (s *ThinktankService) Update(c *gin.Context, req models.ThinktankUpdateRequest) error {
	// 查找res_thinktank记录
	thinktank := models.NewThinktank()
	if err := thinktank.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新res_thinktank信息
    thinktank.ThinktankNo = req.ThinktankNo
    thinktank.BriefName = req.BriefName
    thinktank.FullName = req.FullName
    thinktank.ComOtherName = req.ComOtherName
    thinktank.Logo = req.Logo
    thinktank.ComType = req.ComType
    thinktank.Website = req.Website
    thinktank.BriefIntro = req.BriefIntro
    thinktank.DetailIntro = req.DetailIntro
    thinktank.EstablishTime = req.EstablishTime
    thinktank.InvestRound = req.InvestRound
    thinktank.InvestRoundName = req.InvestRoundName
    thinktank.OperatingStatus = req.OperatingStatus
    thinktank.OperatingStatusName = req.OperatingStatusName
    thinktank.Staff = req.Staff
    thinktank.BaikeLink = req.BaikeLink
    thinktank.Industry = req.Industry
    thinktank.IndustryName = req.IndustryName
    thinktank.MarketValuationRmb = req.MarketValuationRmb
    thinktank.Tags = req.Tags
    thinktank.Location = req.Location
    thinktank.Address = req.Address
    thinktank.CooperationType = req.CooperationType
    thinktank.CooperationForm = req.CooperationForm
    thinktank.ResourceAnalysis = req.ResourceAnalysis
    thinktank.ConstructionContent = req.ConstructionContent
    thinktank.IndustryFocus = req.IndustryFocus
    thinktank.CooperationUnits = req.CooperationUnits
    thinktank.BusinessFocus = req.BusinessFocus
    thinktank.OperationMode = req.OperationMode
    thinktank.ResponsibleDepartment = req.ResponsibleDepartment
    thinktank.ResearchTopic = req.ResearchTopic
    thinktank.Remark = req.Remark
    thinktank.RegFullName = req.RegFullName
    thinktank.RegTime = req.RegTime
    thinktank.RegLocation = req.RegLocation
    thinktank.RegCapital = req.RegCapital
    thinktank.RegCapitalCurrency = req.RegCapitalCurrency
    thinktank.RegCapitalCurrencyName = req.RegCapitalCurrencyName
    thinktank.PaidCapital = req.PaidCapital
    thinktank.PaidCapitalCurrency = req.PaidCapitalCurrency
    thinktank.PaidCapitalCurrencyName = req.PaidCapitalCurrencyName
    thinktank.RegType = req.RegType
    thinktank.RegTypeName = req.RegTypeName
    thinktank.LegalRepresent = req.LegalRepresent
    thinktank.RegNumber = req.RegNumber
    thinktank.CreditNumber = req.CreditNumber
    thinktank.BusinessPeriod = req.BusinessPeriod
    thinktank.BusinessScope = req.BusinessScope
    thinktank.RegIndustry = req.RegIndustry
    thinktank.RegIndustryName = req.RegIndustryName
    thinktank.RegSubIndustry = req.RegSubIndustry
    thinktank.RegSubIndustryName = req.RegSubIndustryName
    thinktank.RegMiddleIndustry = req.RegMiddleIndustry
    thinktank.RegMiddleIndustryName = req.RegMiddleIndustryName
    thinktank.RegSmallIndustry = req.RegSmallIndustry
    thinktank.RegSmallIndustryName = req.RegSmallIndustryName
    thinktank.RegInstitute = req.RegInstitute
    thinktank.ApproveTime = req.ApproveTime
    thinktank.DataSource = req.DataSource
    thinktank.IsBlacklist = req.IsBlacklist
    thinktank.Tag = req.Tag
    thinktank.Partner = req.Partner
    thinktank.PlatformResources = req.PlatformResources
    thinktank.Contact = req.Contact
    thinktank.ContactNumber = req.ContactNumber
    thinktank.Duties = req.Duties
    thinktank.LocationChina = req.LocationChina
    thinktank.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := thinktank.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除res_thinktank
func (s *ThinktankService) Delete(c *gin.Context, id int64) error {
	// 查找res_thinktank记录
	thinktank := models.NewThinktank()
	if err := thinktank.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := thinktank.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取res_thinktank
func (s *ThinktankService) GetByID(c *gin.Context, id int64) (*models.Thinktank, error) {
	// 查找res_thinktank记录
	thinktank := models.NewThinktank()
	if err := thinktank.GetByID(c, id); err != nil {
		return nil, err
	}

	return thinktank, nil
}

// List res_thinktank列表（分页查询）
func (s *ThinktankService) List(c *gin.Context, req models.ThinktankListRequest) (*models.ThinktankList, int64, error) {
	// 获取总数
	thinktankList := models.NewThinktankList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := thinktankList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = thinktankList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return thinktankList, total, nil
}