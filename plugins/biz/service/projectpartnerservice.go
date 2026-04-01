package service

import (
	"gin-fast/plugins/biz/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// ProjectPartnerService biz_project_partner服务
type ProjectPartnerService struct{}

// NewProjectPartnerService 创建biz_project_partner服务
func NewProjectPartnerService() *ProjectPartnerService {
	return &ProjectPartnerService{}
}

// Create 创建biz_project_partner
func (s *ProjectPartnerService) Create(c *gin.Context, req models.ProjectPartnerCreateRequest) (*models.ProjectPartner, error) {
	// 创建biz_project_partner记录
	projectPartner := models.NewProjectPartner()
    projectPartner.ProjectId = req.ProjectId
    projectPartner.RelationProject = req.RelationProject
    projectPartner.RelationPartner = req.RelationPartner
    projectPartner.Organization = req.Organization
    projectPartner.InvestmentAmount = req.InvestmentAmount
    projectPartner.Unit = req.Unit
    projectPartner.Tag = req.Tag
    projectPartner.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := projectPartner.Create(c); err != nil {
		return nil, err
	}

	return projectPartner, nil
}

// Update 更新biz_project_partner
func (s *ProjectPartnerService) Update(c *gin.Context, req models.ProjectPartnerUpdateRequest) error {
	// 查找biz_project_partner记录
	projectPartner := models.NewProjectPartner()
	if err := projectPartner.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新biz_project_partner信息
    projectPartner.ProjectId = req.ProjectId
    projectPartner.RelationProject = req.RelationProject
    projectPartner.RelationPartner = req.RelationPartner
    projectPartner.Organization = req.Organization
    projectPartner.InvestmentAmount = req.InvestmentAmount
    projectPartner.Unit = req.Unit
    projectPartner.Tag = req.Tag
    projectPartner.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := projectPartner.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除biz_project_partner
func (s *ProjectPartnerService) Delete(c *gin.Context, id int64) error {
	// 查找biz_project_partner记录
	projectPartner := models.NewProjectPartner()
	if err := projectPartner.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := projectPartner.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取biz_project_partner
func (s *ProjectPartnerService) GetByID(c *gin.Context, id int64) (*models.ProjectPartner, error) {
	// 查找biz_project_partner记录
	projectPartner := models.NewProjectPartner()
	if err := projectPartner.GetByID(c, id); err != nil {
		return nil, err
	}

	return projectPartner, nil
}

// List biz_project_partner列表（分页查询）
func (s *ProjectPartnerService) List(c *gin.Context, req models.ProjectPartnerListRequest) (*models.ProjectPartnerList, int64, error) {
	// 获取总数
	projectPartnerList := models.NewProjectPartnerList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := projectPartnerList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = projectPartnerList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return projectPartnerList, total, nil
}