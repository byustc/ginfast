package service

import (
	"gin-fast/plugins/resource/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// DistrictService res_district服务
type DistrictService struct{}

// NewDistrictService 创建res_district服务
func NewDistrictService() *DistrictService {
	return &DistrictService{}
}

// Create 创建res_district
func (s *DistrictService) Create(c *gin.Context, req models.DistrictCreateRequest) (*models.District, error) {
	// 创建res_district记录
	district := models.NewDistrict()
    district.DistrictNo = req.DistrictNo
    district.Name = req.Name
    district.Introduction = req.Introduction
    district.Tags = req.Tags
    district.ContactPerson = req.ContactPerson
    district.Phone = req.Phone
    district.Address = req.Address
    district.Remark = req.Remark
    district.DataSource = req.DataSource
    district.IsBlacklist = req.IsBlacklist
    district.DataSourceId = req.DataSourceId
    district.Tag = req.Tag
    district.CujuChanye = req.CujuChanye
    district.Industry = req.Industry
    district.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := district.Create(c); err != nil {
		return nil, err
	}

	return district, nil
}

// Update 更新res_district
func (s *DistrictService) Update(c *gin.Context, req models.DistrictUpdateRequest) error {
	// 查找res_district记录
	district := models.NewDistrict()
	if err := district.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新res_district信息
    district.DistrictNo = req.DistrictNo
    district.Name = req.Name
    district.Introduction = req.Introduction
    district.Tags = req.Tags
    district.ContactPerson = req.ContactPerson
    district.Phone = req.Phone
    district.Address = req.Address
    district.Remark = req.Remark
    district.DataSource = req.DataSource
    district.IsBlacklist = req.IsBlacklist
    district.DataSourceId = req.DataSourceId
    district.Tag = req.Tag
    district.CujuChanye = req.CujuChanye
    district.Industry = req.Industry
    district.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := district.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除res_district
func (s *DistrictService) Delete(c *gin.Context, id int64) error {
	// 查找res_district记录
	district := models.NewDistrict()
	if err := district.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := district.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取res_district
func (s *DistrictService) GetByID(c *gin.Context, id int64) (*models.District, error) {
	// 查找res_district记录
	district := models.NewDistrict()
	if err := district.GetByID(c, id); err != nil {
		return nil, err
	}

	return district, nil
}

// List res_district列表（分页查询）
func (s *DistrictService) List(c *gin.Context, req models.DistrictListRequest) (*models.DistrictList, int64, error) {
	// 获取总数
	districtList := models.NewDistrictList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := districtList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = districtList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return districtList, total, nil
}