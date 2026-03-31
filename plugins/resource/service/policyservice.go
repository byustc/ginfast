package service

import (
	"gin-fast/plugins/resource/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// PolicyService res_policy服务
type PolicyService struct{}

// NewPolicyService 创建res_policy服务
func NewPolicyService() *PolicyService {
	return &PolicyService{}
}

// Create 创建res_policy
func (s *PolicyService) Create(c *gin.Context, req models.PolicyCreateRequest) (*models.Policy, error) {
	// 创建res_policy记录
	policy := models.NewPolicy()
    policy.PolicyNo = req.PolicyNo
    policy.PolicyTitle = req.PolicyTitle
    policy.PolicyAbstract = req.PolicyAbstract
    policy.MainBody = req.MainBody
    policy.PublishTime = req.PublishTime
    policy.PolicyDepartment = req.PolicyDepartment
    policy.Source = req.Source
    policy.Link = req.Link
    policy.EffectDate = req.EffectDate
    policy.FailureDate = req.FailureDate
    policy.Status = req.Status
    policy.PolicyTimeliness = req.PolicyTimeliness
    policy.ContactName = req.ContactName
    policy.ContactPhone = req.ContactPhone
    policy.PolicyRegion = req.PolicyRegion
    policy.PolicyTag = req.PolicyTag
    policy.PolicyComType = req.PolicyComType
    policy.Level = req.Level
    policy.Attachment = req.Attachment
    policy.PolicyNumber = req.PolicyNumber
    policy.Remark = req.Remark
    policy.Title = req.Title
    policy.DataSource = req.DataSource
    policy.IsBlacklist = req.IsBlacklist
    policy.Tag = req.Tag
    policy.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := policy.Create(c); err != nil {
		return nil, err
	}

	return policy, nil
}

// Update 更新res_policy
func (s *PolicyService) Update(c *gin.Context, req models.PolicyUpdateRequest) error {
	// 查找res_policy记录
	policy := models.NewPolicy()
	if err := policy.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新res_policy信息
    policy.PolicyNo = req.PolicyNo
    policy.PolicyTitle = req.PolicyTitle
    policy.PolicyAbstract = req.PolicyAbstract
    policy.MainBody = req.MainBody
    policy.PublishTime = req.PublishTime
    policy.PolicyDepartment = req.PolicyDepartment
    policy.Source = req.Source
    policy.Link = req.Link
    policy.EffectDate = req.EffectDate
    policy.FailureDate = req.FailureDate
    policy.Status = req.Status
    policy.PolicyTimeliness = req.PolicyTimeliness
    policy.ContactName = req.ContactName
    policy.ContactPhone = req.ContactPhone
    policy.PolicyRegion = req.PolicyRegion
    policy.PolicyTag = req.PolicyTag
    policy.PolicyComType = req.PolicyComType
    policy.Level = req.Level
    policy.Attachment = req.Attachment
    policy.PolicyNumber = req.PolicyNumber
    policy.Remark = req.Remark
    policy.Title = req.Title
    policy.DataSource = req.DataSource
    policy.IsBlacklist = req.IsBlacklist
    policy.Tag = req.Tag
    policy.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := policy.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除res_policy
func (s *PolicyService) Delete(c *gin.Context, id int64) error {
	// 查找res_policy记录
	policy := models.NewPolicy()
	if err := policy.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := policy.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取res_policy
func (s *PolicyService) GetByID(c *gin.Context, id int64) (*models.Policy, error) {
	// 查找res_policy记录
	policy := models.NewPolicy()
	if err := policy.GetByID(c, id); err != nil {
		return nil, err
	}

	return policy, nil
}

// List res_policy列表（分页查询）
func (s *PolicyService) List(c *gin.Context, req models.PolicyListRequest) (*models.PolicyList, int64, error) {
	// 获取总数
	policyList := models.NewPolicyList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := policyList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = policyList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return policyList, total, nil
}