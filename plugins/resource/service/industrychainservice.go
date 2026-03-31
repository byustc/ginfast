package service

import (
	"gin-fast/plugins/resource/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// IndustryChainService res_industry_chain服务
type IndustryChainService struct{}

// NewIndustryChainService 创建res_industry_chain服务
func NewIndustryChainService() *IndustryChainService {
	return &IndustryChainService{}
}

// Create 创建res_industry_chain
func (s *IndustryChainService) Create(c *gin.Context, req models.IndustryChainCreateRequest) (*models.IndustryChain, error) {
	// 创建res_industry_chain记录
	industryChain := models.NewIndustryChain()
    industryChain.IndustryChainNo = req.IndustryChainNo
    industryChain.IndustryChainName = req.IndustryChainName
    industryChain.IndustryChainIntro = req.IndustryChainIntro
    industryChain.IndustryChainMap = req.IndustryChainMap
    industryChain.DataSource = req.DataSource
    industryChain.IsBlacklist = req.IsBlacklist
    industryChain.GraphLayer = req.GraphLayer
    industryChain.Tag = req.Tag
    industryChain.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := industryChain.Create(c); err != nil {
		return nil, err
	}

	return industryChain, nil
}

// Update 更新res_industry_chain
func (s *IndustryChainService) Update(c *gin.Context, req models.IndustryChainUpdateRequest) error {
	// 查找res_industry_chain记录
	industryChain := models.NewIndustryChain()
	if err := industryChain.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新res_industry_chain信息
    industryChain.IndustryChainNo = req.IndustryChainNo
    industryChain.IndustryChainName = req.IndustryChainName
    industryChain.IndustryChainIntro = req.IndustryChainIntro
    industryChain.IndustryChainMap = req.IndustryChainMap
    industryChain.DataSource = req.DataSource
    industryChain.IsBlacklist = req.IsBlacklist
    industryChain.GraphLayer = req.GraphLayer
    industryChain.Tag = req.Tag
    industryChain.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := industryChain.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除res_industry_chain
func (s *IndustryChainService) Delete(c *gin.Context, id int64) error {
	// 查找res_industry_chain记录
	industryChain := models.NewIndustryChain()
	if err := industryChain.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := industryChain.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取res_industry_chain
func (s *IndustryChainService) GetByID(c *gin.Context, id int64) (*models.IndustryChain, error) {
	// 查找res_industry_chain记录
	industryChain := models.NewIndustryChain()
	if err := industryChain.GetByID(c, id); err != nil {
		return nil, err
	}

	return industryChain, nil
}

// List res_industry_chain列表（分页查询）
func (s *IndustryChainService) List(c *gin.Context, req models.IndustryChainListRequest) (*models.IndustryChainList, int64, error) {
	// 获取总数
	industryChainList := models.NewIndustryChainList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := industryChainList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = industryChainList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return industryChainList, total, nil
}