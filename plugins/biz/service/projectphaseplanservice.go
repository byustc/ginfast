package service

import (
	"gin-fast/plugins/biz/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// ProjectPhasePlanService biz_project_phase_plan服务
type ProjectPhasePlanService struct{}

// NewProjectPhasePlanService 创建biz_project_phase_plan服务
func NewProjectPhasePlanService() *ProjectPhasePlanService {
	return &ProjectPhasePlanService{}
}

// Create 创建biz_project_phase_plan
func (s *ProjectPhasePlanService) Create(c *gin.Context, req models.ProjectPhasePlanCreateRequest) (*models.ProjectPhasePlan, error) {
	// 创建biz_project_phase_plan记录
	projectPhasePlan := models.NewProjectPhasePlan()
    projectPhasePlan.ProjectId = req.ProjectId
    projectPhasePlan.ItemPhaseId = req.ItemPhaseId
    projectPhasePlan.PlanFinishTime = req.PlanFinishTime
    projectPhasePlan.ActualFinishTime = req.ActualFinishTime
    projectPhasePlan.Tag = req.Tag
    projectPhasePlan.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := projectPhasePlan.Create(c); err != nil {
		return nil, err
	}

	return projectPhasePlan, nil
}

// Update 更新biz_project_phase_plan
func (s *ProjectPhasePlanService) Update(c *gin.Context, req models.ProjectPhasePlanUpdateRequest) error {
	// 查找biz_project_phase_plan记录
	projectPhasePlan := models.NewProjectPhasePlan()
	if err := projectPhasePlan.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新biz_project_phase_plan信息
    projectPhasePlan.ProjectId = req.ProjectId
    projectPhasePlan.ItemPhaseId = req.ItemPhaseId
    projectPhasePlan.PlanFinishTime = req.PlanFinishTime
    projectPhasePlan.ActualFinishTime = req.ActualFinishTime
    projectPhasePlan.Tag = req.Tag
    projectPhasePlan.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := projectPhasePlan.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除biz_project_phase_plan
func (s *ProjectPhasePlanService) Delete(c *gin.Context, id int64) error {
	// 查找biz_project_phase_plan记录
	projectPhasePlan := models.NewProjectPhasePlan()
	if err := projectPhasePlan.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := projectPhasePlan.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取biz_project_phase_plan
func (s *ProjectPhasePlanService) GetByID(c *gin.Context, id int64) (*models.ProjectPhasePlan, error) {
	// 查找biz_project_phase_plan记录
	projectPhasePlan := models.NewProjectPhasePlan()
	if err := projectPhasePlan.GetByID(c, id); err != nil {
		return nil, err
	}

	return projectPhasePlan, nil
}

// List biz_project_phase_plan列表（分页查询）
func (s *ProjectPhasePlanService) List(c *gin.Context, req models.ProjectPhasePlanListRequest) (*models.ProjectPhasePlanList, int64, error) {
	// 获取总数
	projectPhasePlanList := models.NewProjectPhasePlanList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := projectPhasePlanList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = projectPhasePlanList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return projectPhasePlanList, total, nil
}