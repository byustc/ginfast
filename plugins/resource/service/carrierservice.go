package service

import (
	"gin-fast/plugins/resource/models"
	"github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"gin-fast/app/utils/datascope"
	"gin-fast/app/utils/tenanthelper"
)

// CarrierService res_carrier服务
type CarrierService struct{}

// NewCarrierService 创建res_carrier服务
func NewCarrierService() *CarrierService {
	return &CarrierService{}
}

// Create 创建res_carrier
func (s *CarrierService) Create(c *gin.Context, req models.CarrierCreateRequest) (*models.Carrier, error) {
	// 创建res_carrier记录
	carrier := models.NewCarrier()
    carrier.CarrierNo = req.CarrierNo
    carrier.CarrierType = req.CarrierType
    carrier.SecType = req.SecType
    carrier.Location = req.Location
    carrier.Address = req.Address
    carrier.IndustryPark = req.IndustryPark
    carrier.CarrierName = req.CarrierName
    carrier.CarrierPicture = req.CarrierPicture
    carrier.CarrierIntro = req.CarrierIntro
    carrier.Contact = req.Contact
    carrier.Phone = req.Phone
    carrier.StarCom = req.StarCom
    carrier.Area = req.Area
    carrier.FloorHeight = req.FloorHeight
    carrier.Industry = req.Industry
    carrier.AvailableLand = req.AvailableLand
    carrier.InvestmentArea = req.InvestmentArea
    carrier.Rent = req.Rent
    carrier.Remark = req.Remark
    carrier.DataSource = req.DataSource
    carrier.IsBlacklist = req.IsBlacklist
    carrier.ParkType = req.ParkType
    carrier.EndOfFloor = req.EndOfFloor
    carrier.AvailabilityType = req.AvailabilityType
    carrier.LeadingIndustries = req.LeadingIndustries
    carrier.TotalArea = req.TotalArea
    carrier.DataSourceId = req.DataSourceId
    carrier.ParcelArea = req.ParcelArea
    carrier.Tag = req.Tag
    carrier.RentType = req.RentType
    carrier.MinimumRent = req.MinimumRent
    carrier.MaximumRent = req.MaximumRent
    carrier.LettableArea = req.LettableArea
    carrier.DistrictCounty = req.DistrictCounty
    carrier.LandUse = req.LandUse
    carrier.FloorHeightRange = req.FloorHeightRange
    carrier.ParkLevel = req.ParkLevel
    carrier.CujuChanye = req.CujuChanye
    carrier.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := carrier.Create(c); err != nil {
		return nil, err
	}

	return carrier, nil
}

// Update 更新res_carrier
func (s *CarrierService) Update(c *gin.Context, req models.CarrierUpdateRequest) error {
	// 查找res_carrier记录
	carrier := models.NewCarrier()
	if err := carrier.GetByID(c, req.Id); err != nil {
		return err
	}
	// 更新res_carrier信息
    carrier.CarrierNo = req.CarrierNo
    carrier.CarrierType = req.CarrierType
    carrier.SecType = req.SecType
    carrier.Location = req.Location
    carrier.Address = req.Address
    carrier.IndustryPark = req.IndustryPark
    carrier.CarrierName = req.CarrierName
    carrier.CarrierPicture = req.CarrierPicture
    carrier.CarrierIntro = req.CarrierIntro
    carrier.Contact = req.Contact
    carrier.Phone = req.Phone
    carrier.StarCom = req.StarCom
    carrier.Area = req.Area
    carrier.FloorHeight = req.FloorHeight
    carrier.Industry = req.Industry
    carrier.AvailableLand = req.AvailableLand
    carrier.InvestmentArea = req.InvestmentArea
    carrier.Rent = req.Rent
    carrier.Remark = req.Remark
    carrier.DataSource = req.DataSource
    carrier.IsBlacklist = req.IsBlacklist
    carrier.ParkType = req.ParkType
    carrier.EndOfFloor = req.EndOfFloor
    carrier.AvailabilityType = req.AvailabilityType
    carrier.LeadingIndustries = req.LeadingIndustries
    carrier.TotalArea = req.TotalArea
    carrier.DataSourceId = req.DataSourceId
    carrier.ParcelArea = req.ParcelArea
    carrier.Tag = req.Tag
    carrier.RentType = req.RentType
    carrier.MinimumRent = req.MinimumRent
    carrier.MaximumRent = req.MaximumRent
    carrier.LettableArea = req.LettableArea
    carrier.DistrictCounty = req.DistrictCounty
    carrier.LandUse = req.LandUse
    carrier.FloorHeightRange = req.FloorHeightRange
    carrier.ParkLevel = req.ParkLevel
    carrier.CujuChanye = req.CujuChanye
    carrier.UpdatedBy = req.UpdatedBy
	// 保存到数据库
	if err := carrier.Update(c); err != nil {
		return err
	}
	return nil
}

// Delete 删除res_carrier
func (s *CarrierService) Delete(c *gin.Context, id int64) error {
	// 查找res_carrier记录
	carrier := models.NewCarrier()
	if err := carrier.GetByID(c, id); err != nil {
		return err
	}

	// 删除数据库记录
	if err := carrier.Delete(c); err != nil {
		return err
	}

	return nil
}

// GetByID 根据ID获取res_carrier
func (s *CarrierService) GetByID(c *gin.Context, id int64) (*models.Carrier, error) {
	// 查找res_carrier记录
	carrier := models.NewCarrier()
	if err := carrier.GetByID(c, id); err != nil {
		return nil, err
	}

	return carrier, nil
}

// List res_carrier列表（分页查询）
func (s *CarrierService) List(c *gin.Context, req models.CarrierListRequest) (*models.CarrierList, int64, error) {
	// 获取总数
	carrierList := models.NewCarrierList()
	scopes := []func(*gorm.DB) *gorm.DB{req.Handle()}
	scopes = append(scopes, datascope.GetDataScope(c))
	scopes = append(scopes, tenanthelper.TenantScope(c))
	total, err := carrierList.GetTotal(c, scopes...)
	if err != nil {
		return nil, 0, err
	}
    scopes = append(scopes, req.Paginate())
	// 获取分页数据
	err = carrierList.Find(c, scopes...)
	if err != nil {
		return nil, 0, err
	}

	return carrierList, total, nil
}