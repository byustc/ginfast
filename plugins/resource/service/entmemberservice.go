package service

import (
	"gin-fast/plugins/resource/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// EntMemberService res_enterprise_member服务
type EntMemberService struct{}

// NewEntMemberService 创建res_enterprise_member服务
func NewEntMemberService() *EntMemberService {
	return &EntMemberService{}
}

// Create 创建res_enterprise_member
func (s *EntMemberService) Create(c *gin.Context, req models.EntMemberCreateRequest) (*models.EntMember, error) {
	// 创建res_enterprise_member记录
	entMember := models.NewEntMember()
    entMember.EnterpriseId = req.EnterpriseId
    entMember.PersonName = req.PersonName
    entMember.Position = req.Position
    entMember.PersonIntro = req.PersonIntro
    entMember.RankNum = req.RankNum
    entMember.Phone = req.Phone
    entMember.EduBackground = req.EduBackground
    entMember.DataSource = req.DataSource
    entMember.IsBlacklist = req.IsBlacklist
    entMember.Tag = req.Tag
    entMember.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := entMember.Create(c); err != nil {
		return nil, err
	}

	return entMember, nil
}

// Update 更新res_enterprise_member
func (s *EntMemberService) Update(c *gin.Context, req models.EntMemberUpdateRequest) error {
	// 查找res_enterprise_member记录
	entMember := models.NewEntMember()
	if err := entMember.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新res_enterprise_member信息
    entMember.EnterpriseId = req.EnterpriseId
    entMember.PersonName = req.PersonName
    entMember.Position = req.Position
    entMember.PersonIntro = req.PersonIntro
    entMember.RankNum = req.RankNum
    entMember.Phone = req.Phone
    entMember.EduBackground = req.EduBackground
    entMember.DataSource = req.DataSource
    entMember.IsBlacklist = req.IsBlacklist
    entMember.Tag = req.Tag
    entMember.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := entMember.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除res_enterprise_member
func (s *EntMemberService) Delete(c *gin.Context, id int64) error {
	// 查找res_enterprise_member记录
	entMember := models.NewEntMember()
	if err := entMember.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := entMember.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取res_enterprise_member
func (s *EntMemberService) GetByID(c *gin.Context, id int64) (*models.EntMember, error) {
	// 查找res_enterprise_member记录
	entMember := models.NewEntMember()
	if err := entMember.GetByID(c, id); err != nil {
		return nil, err
	}

	return entMember, nil
}

// List res_enterprise_member列表（分页查询）
func (s *EntMemberService) List(c *gin.Context, req models.EntMemberListRequest) (*models.EntMemberList, int64, error) {
	// 获取总数
	entMemberList := models.NewEntMemberList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := entMemberList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = entMemberList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return entMemberList, total, nil
}