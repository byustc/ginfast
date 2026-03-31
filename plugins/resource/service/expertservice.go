package service

import (
	"gin-fast/plugins/resource/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// ExpertService res_expert服务
type ExpertService struct{}

// NewExpertService 创建res_expert服务
func NewExpertService() *ExpertService {
	return &ExpertService{}
}

// Create 创建res_expert
func (s *ExpertService) Create(c *gin.Context, req models.ExpertCreateRequest) (*models.Expert, error) {
	// 创建res_expert记录
	expert := models.NewExpert()
    expert.ExpertNo = req.ExpertNo
    expert.Name = req.Name
    expert.Phone = req.Phone
    expert.Location = req.Location
    expert.Degree = req.Degree
    expert.Major = req.Major
    expert.MajorLevel = req.MajorLevel
    expert.MajorType = req.MajorType
    expert.ResearchIndustry = req.ResearchIndustry
    expert.WorkYears = req.WorkYears
    expert.WorkCom = req.WorkCom
    expert.Position = req.Position
    expert.Industry = req.Industry
    expert.IsLecturer = req.IsLecturer
    expert.ExpertType = req.ExpertType
    expert.ExpertLevel = req.ExpertLevel
    expert.Source = req.Source
    expert.Remark = req.Remark
    expert.DataSource = req.DataSource
    expert.IsBlacklist = req.IsBlacklist
    expert.Tag = req.Tag
    expert.IndustrySj = req.IndustrySj
    expert.IsPartner = req.IsPartner
    expert.Profile = req.Profile
    expert.InstitutionType = req.InstitutionType
    expert.ContactNumber = req.ContactNumber
    expert.AffiliatedInstitutions = req.AffiliatedInstitutions
    expert.EntryPerson = req.EntryPerson
    expert.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := expert.Create(c); err != nil {
		return nil, err
	}

	return expert, nil
}

// Update 更新res_expert
func (s *ExpertService) Update(c *gin.Context, req models.ExpertUpdateRequest) error {
	// 查找res_expert记录
	expert := models.NewExpert()
	if err := expert.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新res_expert信息
    expert.ExpertNo = req.ExpertNo
    expert.Name = req.Name
    expert.Phone = req.Phone
    expert.Location = req.Location
    expert.Degree = req.Degree
    expert.Major = req.Major
    expert.MajorLevel = req.MajorLevel
    expert.MajorType = req.MajorType
    expert.ResearchIndustry = req.ResearchIndustry
    expert.WorkYears = req.WorkYears
    expert.WorkCom = req.WorkCom
    expert.Position = req.Position
    expert.Industry = req.Industry
    expert.IsLecturer = req.IsLecturer
    expert.ExpertType = req.ExpertType
    expert.ExpertLevel = req.ExpertLevel
    expert.Source = req.Source
    expert.Remark = req.Remark
    expert.DataSource = req.DataSource
    expert.IsBlacklist = req.IsBlacklist
    expert.Tag = req.Tag
    expert.IndustrySj = req.IndustrySj
    expert.IsPartner = req.IsPartner
    expert.Profile = req.Profile
    expert.InstitutionType = req.InstitutionType
    expert.ContactNumber = req.ContactNumber
    expert.AffiliatedInstitutions = req.AffiliatedInstitutions
    expert.EntryPerson = req.EntryPerson
    expert.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := expert.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除res_expert
func (s *ExpertService) Delete(c *gin.Context, id int64) error {
	// 查找res_expert记录
	expert := models.NewExpert()
	if err := expert.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := expert.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取res_expert
func (s *ExpertService) GetByID(c *gin.Context, id int64) (*models.Expert, error) {
	// 查找res_expert记录
	expert := models.NewExpert()
	if err := expert.GetByID(c, id); err != nil {
		return nil, err
	}

	return expert, nil
}

// List res_expert列表（分页查询）
func (s *ExpertService) List(c *gin.Context, req models.ExpertListRequest) (*models.ExpertList, int64, error) {
	// 获取总数
	expertList := models.NewExpertList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := expertList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = expertList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return expertList, total, nil
}