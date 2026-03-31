package service

import (
	"gin-fast/plugins/resource/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// ThitankFundService res_thinktank_fund服务
type ThitankFundService struct{}

// NewThitankFundService 创建res_thinktank_fund服务
func NewThitankFundService() *ThitankFundService {
	return &ThitankFundService{}
}

// Create 创建res_thinktank_fund
func (s *ThitankFundService) Create(c *gin.Context, req models.ThitankFundCreateRequest) (*models.ThitankFund, error) {
	// 创建res_thinktank_fund记录
	thitankFund := models.NewThitankFund()
    thitankFund.FundName = req.FundName
    thitankFund.FundScale = req.FundScale
    thitankFund.ThinktankId = req.ThinktankId
    thitankFund.InvestIndustry = req.InvestIndustry
    thitankFund.RoundPrefer = req.RoundPrefer
    thitankFund.SingleInvestAmount = req.SingleInvestAmount
    thitankFund.InvestMethod = req.InvestMethod
    thitankFund.InvestableBalance = req.InvestableBalance
    thitankFund.ReturnRequirements = req.ReturnRequirements
    thitankFund.CompletionStatus = req.CompletionStatus
    thitankFund.FundType = req.FundType
    thitankFund.EstablishTime = req.EstablishTime
    thitankFund.DataSource = req.DataSource
    thitankFund.IsBlacklist = req.IsBlacklist
    thitankFund.InvestmentIndustries = req.InvestmentIndustries
    thitankFund.AvailableBalanceOfFund = req.AvailableBalanceOfFund
    thitankFund.Tag = req.Tag
    thitankFund.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := thitankFund.Create(c); err != nil {
		return nil, err
	}

	return thitankFund, nil
}

// Update 更新res_thinktank_fund
func (s *ThitankFundService) Update(c *gin.Context, req models.ThitankFundUpdateRequest) error {
	// 查找res_thinktank_fund记录
	thitankFund := models.NewThitankFund()
	if err := thitankFund.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新res_thinktank_fund信息
    thitankFund.FundName = req.FundName
    thitankFund.FundScale = req.FundScale
    thitankFund.ThinktankId = req.ThinktankId
    thitankFund.InvestIndustry = req.InvestIndustry
    thitankFund.RoundPrefer = req.RoundPrefer
    thitankFund.SingleInvestAmount = req.SingleInvestAmount
    thitankFund.InvestMethod = req.InvestMethod
    thitankFund.InvestableBalance = req.InvestableBalance
    thitankFund.ReturnRequirements = req.ReturnRequirements
    thitankFund.CompletionStatus = req.CompletionStatus
    thitankFund.FundType = req.FundType
    thitankFund.EstablishTime = req.EstablishTime
    thitankFund.DataSource = req.DataSource
    thitankFund.IsBlacklist = req.IsBlacklist
    thitankFund.InvestmentIndustries = req.InvestmentIndustries
    thitankFund.AvailableBalanceOfFund = req.AvailableBalanceOfFund
    thitankFund.Tag = req.Tag
    thitankFund.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := thitankFund.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除res_thinktank_fund
func (s *ThitankFundService) Delete(c *gin.Context, id int64) error {
	// 查找res_thinktank_fund记录
	thitankFund := models.NewThitankFund()
	if err := thitankFund.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := thitankFund.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取res_thinktank_fund
func (s *ThitankFundService) GetByID(c *gin.Context, id int64) (*models.ThitankFund, error) {
	// 查找res_thinktank_fund记录
	thitankFund := models.NewThitankFund()
	if err := thitankFund.GetByID(c, id); err != nil {
		return nil, err
	}

	return thitankFund, nil
}

// List res_thinktank_fund列表（分页查询）
func (s *ThitankFundService) List(c *gin.Context, req models.ThitankFundListRequest) (*models.ThitankFundList, int64, error) {
	// 获取总数
	thitankFundList := models.NewThitankFundList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := thitankFundList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = thitankFundList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return thitankFundList, total, nil
}