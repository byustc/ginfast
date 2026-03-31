package service

import (
	"gin-fast/plugins/resource/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// InvMemberService res_invest_member服务
type InvMemberService struct{}

// NewInvMemberService 创建res_invest_member服务
func NewInvMemberService() *InvMemberService {
	return &InvMemberService{}
}

// Create 创建res_invest_member
func (s *InvMemberService) Create(c *gin.Context, req models.InvMemberCreateRequest) (*models.InvMember, error) {
	// 创建res_invest_member记录
	invMember := models.NewInvMember()
    invMember.InvestId = req.InvestId
    invMember.PersonName = req.PersonName
    invMember.Position = req.Position
    invMember.PersonIntro = req.PersonIntro
    invMember.RankNum = req.RankNum
    invMember.Phone = req.Phone
    invMember.EduBackground = req.EduBackground
    invMember.DataSource = req.DataSource
    invMember.IsBlacklist = req.IsBlacklist
    invMember.Tag = req.Tag
    invMember.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := invMember.Create(c); err != nil {
		return nil, err
	}

	return invMember, nil
}

// Update 更新res_invest_member
func (s *InvMemberService) Update(c *gin.Context, req models.InvMemberUpdateRequest) error {
	// 查找res_invest_member记录
	invMember := models.NewInvMember()
	if err := invMember.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新res_invest_member信息
    invMember.InvestId = req.InvestId
    invMember.PersonName = req.PersonName
    invMember.Position = req.Position
    invMember.PersonIntro = req.PersonIntro
    invMember.RankNum = req.RankNum
    invMember.Phone = req.Phone
    invMember.EduBackground = req.EduBackground
    invMember.DataSource = req.DataSource
    invMember.IsBlacklist = req.IsBlacklist
    invMember.Tag = req.Tag
    invMember.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := invMember.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除res_invest_member
func (s *InvMemberService) Delete(c *gin.Context, id int64) error {
	// 查找res_invest_member记录
	invMember := models.NewInvMember()
	if err := invMember.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := invMember.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取res_invest_member
func (s *InvMemberService) GetByID(c *gin.Context, id int64) (*models.InvMember, error) {
	// 查找res_invest_member记录
	invMember := models.NewInvMember()
	if err := invMember.GetByID(c, id); err != nil {
		return nil, err
	}

	return invMember, nil
}

// List res_invest_member列表（分页查询）
func (s *InvMemberService) List(c *gin.Context, req models.InvMemberListRequest) (*models.InvMemberList, int64, error) {
	// 获取总数
	invMemberList := models.NewInvMemberList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := invMemberList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = invMemberList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return invMemberList, total, nil
}