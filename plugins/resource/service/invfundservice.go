package service

import (
	"gin-fast/plugins/resource/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// InvFundService res_invest_fund服务
type InvFundService struct{}

// NewInvFundService 创建res_invest_fund服务
func NewInvFundService() *InvFundService {
	return &InvFundService{}
}

// Create 创建res_invest_fund
func (s *InvFundService) Create(c *gin.Context, req models.InvFundCreateRequest) (*models.InvFund, error) {
	// 创建res_invest_fund记录
	invFund := models.NewInvFund()
    invFund.InvestId = req.InvestId
    invFund.FundName = req.FundName
    invFund.FundScale = req.FundScale
    invFund.InvestIndustry = req.InvestIndustry
    invFund.RoundPrefer = req.RoundPrefer
    invFund.SingleInvestAmount = req.SingleInvestAmount
    invFund.InvestMethod = req.InvestMethod
    invFund.InvestableBalance = req.InvestableBalance
    invFund.ReturnRequirements = req.ReturnRequirements
    invFund.CompletionStatus = req.CompletionStatus
    invFund.FundType = req.FundType
    invFund.EstablishTime = req.EstablishTime
    invFund.DataSource = req.DataSource
    invFund.IsBlacklist = req.IsBlacklist
    invFund.InvestmentField = req.InvestmentField
    invFund.AvailableBalanceFund = req.AvailableBalanceFund
    invFund.Tag = req.Tag
    invFund.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := invFund.Create(c); err != nil {
		return nil, err
	}

	return invFund, nil
}

// Update 更新res_invest_fund
func (s *InvFundService) Update(c *gin.Context, req models.InvFundUpdateRequest) error {
	// 查找res_invest_fund记录
	invFund := models.NewInvFund()
	if err := invFund.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新res_invest_fund信息
    invFund.InvestId = req.InvestId
    invFund.FundName = req.FundName
    invFund.FundScale = req.FundScale
    invFund.InvestIndustry = req.InvestIndustry
    invFund.RoundPrefer = req.RoundPrefer
    invFund.SingleInvestAmount = req.SingleInvestAmount
    invFund.InvestMethod = req.InvestMethod
    invFund.InvestableBalance = req.InvestableBalance
    invFund.ReturnRequirements = req.ReturnRequirements
    invFund.CompletionStatus = req.CompletionStatus
    invFund.FundType = req.FundType
    invFund.EstablishTime = req.EstablishTime
    invFund.DataSource = req.DataSource
    invFund.IsBlacklist = req.IsBlacklist
    invFund.InvestmentField = req.InvestmentField
    invFund.AvailableBalanceFund = req.AvailableBalanceFund
    invFund.Tag = req.Tag
    invFund.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := invFund.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除res_invest_fund
func (s *InvFundService) Delete(c *gin.Context, id int64) error {
	// 查找res_invest_fund记录
	invFund := models.NewInvFund()
	if err := invFund.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := invFund.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取res_invest_fund
func (s *InvFundService) GetByID(c *gin.Context, id int64) (*models.InvFund, error) {
	// 查找res_invest_fund记录
	invFund := models.NewInvFund()
	if err := invFund.GetByID(c, id); err != nil {
		return nil, err
	}

	return invFund, nil
}

// List res_invest_fund列表（分页查询）
func (s *InvFundService) List(c *gin.Context, req models.InvFundListRequest) (*models.InvFundList, int64, error) {
	// 获取总数
	invFundList := models.NewInvFundList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := invFundList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = invFundList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return invFundList, total, nil
}