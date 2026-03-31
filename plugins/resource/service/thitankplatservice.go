package service

import (
	"gin-fast/plugins/resource/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// ThitankPlatService res_thinktank_plat服务
type ThitankPlatService struct{}

// NewThitankPlatService 创建res_thinktank_plat服务
func NewThitankPlatService() *ThitankPlatService {
	return &ThitankPlatService{}
}

// Create 创建res_thinktank_plat
func (s *ThitankPlatService) Create(c *gin.Context, req models.ThitankPlatCreateRequest) (*models.ThitankPlat, error) {
	// 创建res_thinktank_plat记录
	thitankPlat := models.NewThitankPlat()
    thitankPlat.PlatformName = req.PlatformName
    thitankPlat.ThinktankId = req.ThinktankId
    thitankPlat.SupportFacilities = req.SupportFacilities
    thitankPlat.PlatformType = req.PlatformType
    thitankPlat.ServiceIndustry = req.ServiceIndustry
    thitankPlat.ServiceObject = req.ServiceObject
    thitankPlat.ServiceCapability = req.ServiceCapability
    thitankPlat.ServiceContent = req.ServiceContent
    thitankPlat.Qualification = req.Qualification
    thitankPlat.Remark = req.Remark
    thitankPlat.DataSource = req.DataSource
    thitankPlat.IsBlacklist = req.IsBlacklist
    thitankPlat.ServiceField = req.ServiceField
    thitankPlat.Tag = req.Tag
    thitankPlat.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := thitankPlat.Create(c); err != nil {
		return nil, err
	}

	return thitankPlat, nil
}

// Update 更新res_thinktank_plat
func (s *ThitankPlatService) Update(c *gin.Context, req models.ThitankPlatUpdateRequest) error {
	// 查找res_thinktank_plat记录
	thitankPlat := models.NewThitankPlat()
	if err := thitankPlat.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新res_thinktank_plat信息
    thitankPlat.PlatformName = req.PlatformName
    thitankPlat.ThinktankId = req.ThinktankId
    thitankPlat.SupportFacilities = req.SupportFacilities
    thitankPlat.PlatformType = req.PlatformType
    thitankPlat.ServiceIndustry = req.ServiceIndustry
    thitankPlat.ServiceObject = req.ServiceObject
    thitankPlat.ServiceCapability = req.ServiceCapability
    thitankPlat.ServiceContent = req.ServiceContent
    thitankPlat.Qualification = req.Qualification
    thitankPlat.Remark = req.Remark
    thitankPlat.DataSource = req.DataSource
    thitankPlat.IsBlacklist = req.IsBlacklist
    thitankPlat.ServiceField = req.ServiceField
    thitankPlat.Tag = req.Tag
    thitankPlat.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := thitankPlat.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除res_thinktank_plat
func (s *ThitankPlatService) Delete(c *gin.Context, id int64) error {
	// 查找res_thinktank_plat记录
	thitankPlat := models.NewThitankPlat()
	if err := thitankPlat.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := thitankPlat.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取res_thinktank_plat
func (s *ThitankPlatService) GetByID(c *gin.Context, id int64) (*models.ThitankPlat, error) {
	// 查找res_thinktank_plat记录
	thitankPlat := models.NewThitankPlat()
	if err := thitankPlat.GetByID(c, id); err != nil {
		return nil, err
	}

	return thitankPlat, nil
}

// List res_thinktank_plat列表（分页查询）
func (s *ThitankPlatService) List(c *gin.Context, req models.ThitankPlatListRequest) (*models.ThitankPlatList, int64, error) {
	// 获取总数
	thitankPlatList := models.NewThitankPlatList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := thitankPlatList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = thitankPlatList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return thitankPlatList, total, nil
}