package service

import (
	"gin-fast/plugins/resource/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// ThitankMemberService res_thinktank_member服务
type ThitankMemberService struct{}

// NewThitankMemberService 创建res_thinktank_member服务
func NewThitankMemberService() *ThitankMemberService {
	return &ThitankMemberService{}
}

// Create 创建res_thinktank_member
func (s *ThitankMemberService) Create(c *gin.Context, req models.ThitankMemberCreateRequest) (*models.ThitankMember, error) {
	// 创建res_thinktank_member记录
	thitankMember := models.NewThitankMember()
    thitankMember.ThinktankId = req.ThinktankId
    thitankMember.PersonName = req.PersonName
    thitankMember.Position = req.Position
    thitankMember.PersonIntro = req.PersonIntro
    thitankMember.Phone = req.Phone
    thitankMember.EduBackground = req.EduBackground
    thitankMember.EnterpriseId = req.EnterpriseId
    thitankMember.RankNum = req.RankNum
    thitankMember.DataSource = req.DataSource
    thitankMember.IsBlacklist = req.IsBlacklist
    thitankMember.Tag = req.Tag
    thitankMember.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := thitankMember.Create(c); err != nil {
		return nil, err
	}

	return thitankMember, nil
}

// Update 更新res_thinktank_member
func (s *ThitankMemberService) Update(c *gin.Context, req models.ThitankMemberUpdateRequest) error {
	// 查找res_thinktank_member记录
	thitankMember := models.NewThitankMember()
	if err := thitankMember.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新res_thinktank_member信息
    thitankMember.ThinktankId = req.ThinktankId
    thitankMember.PersonName = req.PersonName
    thitankMember.Position = req.Position
    thitankMember.PersonIntro = req.PersonIntro
    thitankMember.Phone = req.Phone
    thitankMember.EduBackground = req.EduBackground
    thitankMember.EnterpriseId = req.EnterpriseId
    thitankMember.RankNum = req.RankNum
    thitankMember.DataSource = req.DataSource
    thitankMember.IsBlacklist = req.IsBlacklist
    thitankMember.Tag = req.Tag
    thitankMember.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := thitankMember.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除res_thinktank_member
func (s *ThitankMemberService) Delete(c *gin.Context, id int64) error {
	// 查找res_thinktank_member记录
	thitankMember := models.NewThitankMember()
	if err := thitankMember.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := thitankMember.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取res_thinktank_member
func (s *ThitankMemberService) GetByID(c *gin.Context, id int64) (*models.ThitankMember, error) {
	// 查找res_thinktank_member记录
	thitankMember := models.NewThitankMember()
	if err := thitankMember.GetByID(c, id); err != nil {
		return nil, err
	}

	return thitankMember, nil
}

// List res_thinktank_member列表（分页查询）
func (s *ThitankMemberService) List(c *gin.Context, req models.ThitankMemberListRequest) (*models.ThitankMemberList, int64, error) {
	// 获取总数
	thitankMemberList := models.NewThitankMemberList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := thitankMemberList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = thitankMemberList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return thitankMemberList, total, nil
}