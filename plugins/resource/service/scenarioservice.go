package service

import (
	"gin-fast/plugins/resource/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// ScenarioService res_scenario服务
type ScenarioService struct{}

// NewScenarioService 创建res_scenario服务
func NewScenarioService() *ScenarioService {
	return &ScenarioService{}
}

// Create 创建res_scenario
func (s *ScenarioService) Create(c *gin.Context, req models.ScenarioCreateRequest) (*models.Scenario, error) {
	// 创建res_scenario记录
	scenario := models.NewScenario()
    scenario.ScenarioNo = req.ScenarioNo
    scenario.PublishTime = req.PublishTime
    scenario.OwnerUnit = req.OwnerUnit
    scenario.SceneName = req.SceneName
    scenario.Industry = req.Industry
    scenario.OpportunityZone = req.OpportunityZone
    scenario.ValidityPeriod = req.ValidityPeriod
    scenario.BudgetGuarantee = req.BudgetGuarantee
    scenario.OtherSituations = req.OtherSituations
    scenario.OpportunityDesc = req.OpportunityDesc
    scenario.BuildBasement = req.BuildBasement
    scenario.CooperationRequirements = req.CooperationRequirements
    scenario.CooperationForm = req.CooperationForm
    scenario.ContactPerson = req.ContactPerson
    scenario.Position = req.Position
    scenario.Phone = req.Phone
    scenario.Remark = req.Remark
    scenario.DataSource = req.DataSource
    scenario.IsBlacklist = req.IsBlacklist
    scenario.Tag = req.Tag
    scenario.IndustrialFields = req.IndustrialFields
    scenario.XcujuChanye = req.XcujuChanye
    scenario.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := scenario.Create(c); err != nil {
		return nil, err
	}

	return scenario, nil
}

// Update 更新res_scenario
func (s *ScenarioService) Update(c *gin.Context, req models.ScenarioUpdateRequest) error {
	// 查找res_scenario记录
	scenario := models.NewScenario()
	if err := scenario.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新res_scenario信息
    scenario.ScenarioNo = req.ScenarioNo
    scenario.PublishTime = req.PublishTime
    scenario.OwnerUnit = req.OwnerUnit
    scenario.SceneName = req.SceneName
    scenario.Industry = req.Industry
    scenario.OpportunityZone = req.OpportunityZone
    scenario.ValidityPeriod = req.ValidityPeriod
    scenario.BudgetGuarantee = req.BudgetGuarantee
    scenario.OtherSituations = req.OtherSituations
    scenario.OpportunityDesc = req.OpportunityDesc
    scenario.BuildBasement = req.BuildBasement
    scenario.CooperationRequirements = req.CooperationRequirements
    scenario.CooperationForm = req.CooperationForm
    scenario.ContactPerson = req.ContactPerson
    scenario.Position = req.Position
    scenario.Phone = req.Phone
    scenario.Remark = req.Remark
    scenario.DataSource = req.DataSource
    scenario.IsBlacklist = req.IsBlacklist
    scenario.Tag = req.Tag
    scenario.IndustrialFields = req.IndustrialFields
    scenario.XcujuChanye = req.XcujuChanye
    scenario.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := scenario.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除res_scenario
func (s *ScenarioService) Delete(c *gin.Context, id int64) error {
	// 查找res_scenario记录
	scenario := models.NewScenario()
	if err := scenario.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := scenario.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取res_scenario
func (s *ScenarioService) GetByID(c *gin.Context, id int64) (*models.Scenario, error) {
	// 查找res_scenario记录
	scenario := models.NewScenario()
	if err := scenario.GetByID(c, id); err != nil {
		return nil, err
	}

	return scenario, nil
}

// List res_scenario列表（分页查询）
func (s *ScenarioService) List(c *gin.Context, req models.ScenarioListRequest) (*models.ScenarioList, int64, error) {
	// 获取总数
	scenarioList := models.NewScenarioList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := scenarioList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = scenarioList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return scenarioList, total, nil
}