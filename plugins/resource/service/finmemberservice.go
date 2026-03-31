package service

import (
	"gin-fast/plugins/resource/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// FinMemberService res_finance_member服务
type FinMemberService struct{}

// NewFinMemberService 创建res_finance_member服务
func NewFinMemberService() *FinMemberService {
	return &FinMemberService{}
}

// Create 创建res_finance_member
func (s *FinMemberService) Create(c *gin.Context, req models.FinMemberCreateRequest) (*models.FinMember, error) {
	// 创建res_finance_member记录
	finMember := models.NewFinMember()
    finMember.FinanceId = req.FinanceId
    finMember.PersonName = req.PersonName
    finMember.Position = req.Position
    finMember.PersonIntro = req.PersonIntro
    finMember.RankNum = req.RankNum
    finMember.Phone = req.Phone
    finMember.EduBackground = req.EduBackground
    finMember.DataSource = req.DataSource
    finMember.IsBlacklist = req.IsBlacklist
    finMember.Tag = req.Tag
    finMember.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := finMember.Create(c); err != nil {
		return nil, err
	}

	return finMember, nil
}

// Update 更新res_finance_member
func (s *FinMemberService) Update(c *gin.Context, req models.FinMemberUpdateRequest) error {
	// 查找res_finance_member记录
	finMember := models.NewFinMember()
	if err := finMember.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新res_finance_member信息
    finMember.FinanceId = req.FinanceId
    finMember.PersonName = req.PersonName
    finMember.Position = req.Position
    finMember.PersonIntro = req.PersonIntro
    finMember.RankNum = req.RankNum
    finMember.Phone = req.Phone
    finMember.EduBackground = req.EduBackground
    finMember.DataSource = req.DataSource
    finMember.IsBlacklist = req.IsBlacklist
    finMember.Tag = req.Tag
    finMember.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := finMember.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除res_finance_member
func (s *FinMemberService) Delete(c *gin.Context, id int64) error {
	// 查找res_finance_member记录
	finMember := models.NewFinMember()
	if err := finMember.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := finMember.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取res_finance_member
func (s *FinMemberService) GetByID(c *gin.Context, id int64) (*models.FinMember, error) {
	// 查找res_finance_member记录
	finMember := models.NewFinMember()
	if err := finMember.GetByID(c, id); err != nil {
		return nil, err
	}

	return finMember, nil
}

// List res_finance_member列表（分页查询）
func (s *FinMemberService) List(c *gin.Context, req models.FinMemberListRequest) (*models.FinMemberList, int64, error) {
	// 获取总数
	finMemberList := models.NewFinMemberList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := finMemberList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = finMemberList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return finMemberList, total, nil
}