package service

import (
	"gin-fast/plugins/biz/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// ProjectResService biz_project_res服务
type ProjectResService struct{}

// NewProjectResService 创建biz_project_res服务
func NewProjectResService() *ProjectResService {
	return &ProjectResService{}
}

// Create 创建biz_project_res
func (s *ProjectResService) Create(c *gin.Context, req models.ProjectResCreateRequest) (*models.ProjectRes, error) {
	// 创建biz_project_res记录
	projectRes := models.NewProjectRes()
    projectRes.ProjectId = req.ProjectId
    projectRes.AgencyType = req.AgencyType
    projectRes.FinanceId = req.FinanceId
    projectRes.InvestId = req.InvestId
    projectRes.ThinktankId = req.ThinktankId
    projectRes.SceneId = req.SceneId
    projectRes.DistrictId = req.DistrictId
    projectRes.CarrierId = req.CarrierId
    projectRes.CarrierInfoId = req.CarrierInfoId
    projectRes.Tag = req.Tag
    projectRes.PolicyId = req.PolicyId
    projectRes.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := projectRes.Create(c); err != nil {
		return nil, err
	}

	return projectRes, nil
}

// Update 更新biz_project_res
func (s *ProjectResService) Update(c *gin.Context, req models.ProjectResUpdateRequest) error {
	// 查找biz_project_res记录
	projectRes := models.NewProjectRes()
	if err := projectRes.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新biz_project_res信息
    projectRes.ProjectId = req.ProjectId
    projectRes.AgencyType = req.AgencyType
    projectRes.FinanceId = req.FinanceId
    projectRes.InvestId = req.InvestId
    projectRes.ThinktankId = req.ThinktankId
    projectRes.SceneId = req.SceneId
    projectRes.DistrictId = req.DistrictId
    projectRes.CarrierId = req.CarrierId
    projectRes.CarrierInfoId = req.CarrierInfoId
    projectRes.Tag = req.Tag
    projectRes.PolicyId = req.PolicyId
    projectRes.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := projectRes.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除biz_project_res
func (s *ProjectResService) Delete(c *gin.Context, id int64) error {
	// 查找biz_project_res记录
	projectRes := models.NewProjectRes()
	if err := projectRes.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := projectRes.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取biz_project_res
func (s *ProjectResService) GetByID(c *gin.Context, id int64) (*models.ProjectRes, error) {
	// 查找biz_project_res记录
	projectRes := models.NewProjectRes()
	if err := projectRes.GetByID(c, id); err != nil {
		return nil, err
	}

	return projectRes, nil
}

// List biz_project_res列表（分页查询）
func (s *ProjectResService) List(c *gin.Context, req models.ProjectResListRequest) (*models.ProjectResList, int64, error) {
	// 获取总数
	projectResList := models.NewProjectResList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := projectResList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = projectResList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return projectResList, total, nil
}