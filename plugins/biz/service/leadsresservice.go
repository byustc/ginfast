package service

import (
	"gin-fast/plugins/biz/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// LeadsResService biz_leads_res服务
type LeadsResService struct{}

// NewLeadsResService 创建biz_leads_res服务
func NewLeadsResService() *LeadsResService {
	return &LeadsResService{}
}

// Create 创建biz_leads_res
func (s *LeadsResService) Create(c *gin.Context, req models.LeadsResCreateRequest) (*models.LeadsRes, error) {
	// 创建biz_leads_res记录
	leadsRes := models.NewLeadsRes()
    leadsRes.LeadsId = req.LeadsId
    leadsRes.Code = req.Code
    leadsRes.AgencyType = req.AgencyType
    leadsRes.FinanceId = req.FinanceId
    leadsRes.InvestId = req.InvestId
    leadsRes.ThinktankId = req.ThinktankId
    leadsRes.SceneId = req.SceneId
    leadsRes.DistrictId = req.DistrictId
    leadsRes.CarrierId = req.CarrierId
    leadsRes.Tag = req.Tag
    leadsRes.PolicyId = req.PolicyId
    leadsRes.BuyerId = req.BuyerId
    leadsRes.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := leadsRes.Create(c); err != nil {
		return nil, err
	}

	return leadsRes, nil
}

// Update 更新biz_leads_res
func (s *LeadsResService) Update(c *gin.Context, req models.LeadsResUpdateRequest) error {
	// 查找biz_leads_res记录
	leadsRes := models.NewLeadsRes()
	if err := leadsRes.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新biz_leads_res信息
    leadsRes.LeadsId = req.LeadsId
    leadsRes.Code = req.Code
    leadsRes.AgencyType = req.AgencyType
    leadsRes.FinanceId = req.FinanceId
    leadsRes.InvestId = req.InvestId
    leadsRes.ThinktankId = req.ThinktankId
    leadsRes.SceneId = req.SceneId
    leadsRes.DistrictId = req.DistrictId
    leadsRes.CarrierId = req.CarrierId
    leadsRes.Tag = req.Tag
    leadsRes.PolicyId = req.PolicyId
    leadsRes.BuyerId = req.BuyerId
    leadsRes.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := leadsRes.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除biz_leads_res
func (s *LeadsResService) Delete(c *gin.Context, id int64) error {
	// 查找biz_leads_res记录
	leadsRes := models.NewLeadsRes()
	if err := leadsRes.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := leadsRes.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取biz_leads_res
func (s *LeadsResService) GetByID(c *gin.Context, id int64) (*models.LeadsRes, error) {
	// 查找biz_leads_res记录
	leadsRes := models.NewLeadsRes()
	if err := leadsRes.GetByID(c, id); err != nil {
		return nil, err
	}

	return leadsRes, nil
}

// List biz_leads_res列表（分页查询）
func (s *LeadsResService) List(c *gin.Context, req models.LeadsResListRequest) (*models.LeadsResList, int64, error) {
	// 获取总数
	leadsResList := models.NewLeadsResList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := leadsResList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = leadsResList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return leadsResList, total, nil
}