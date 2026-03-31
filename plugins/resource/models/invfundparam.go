package models

import (
	"time"
	"gin-fast/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InvFundListRequest res_invest_fund列表请求参数
type InvFundListRequest struct {
	models.BasePaging
	models.Validator
	Id *int64 `form:"id"` // 主键
	InvestId *int64 `form:"investId"` // 投资机构-数据主键id
	FundName *string `form:"fundName"` // 基金名称
	FundScale *string `form:"fundScale"` // 基金规模
	InvestIndustry *string `form:"investIndustry"` // 投资行业
	RoundPrefer *string `form:"roundPrefer"` // 投资阶段偏好
	SingleInvestAmount *string `form:"singleInvestAmount"` // 单笔投资金额
	InvestMethod *string `form:"investMethod"` // 投资方式
	InvestableBalance *float64 `form:"investableBalance"` // 可投余额
	ReturnRequirements *string `form:"returnRequirements"` // 返投要求
	CompletionStatus *string `form:"completionStatus"` // 返投完成情况
	FundType *string `form:"fundType"` // 基金类型
	EstablishTime *time.Time `form:"establishTime"` // 成立时间
	DataSource *string `form:"dataSource"` // 数据来源
	IsBlacklist *int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	InvestmentField *string `form:"investmentField"` // 投资领域
	AvailableBalanceFund *string `form:"availableBalanceFund"` // 基金可投余额
	Tag *string `form:"tag"` // 行业数据标签
	UpdatedBy *int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *InvFundListRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// Handle 获取查询条件
func (r *InvFundListRequest) Handle() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
        if r.Id != nil {
            // 默认等于查询
            db = db.Where("id = ?", *r.Id)
        }
        if r.InvestId != nil {
            // 默认等于查询
            db = db.Where("invest_id = ?", *r.InvestId)
        }
        if r.FundName != nil {
            // 默认等于查询
            db = db.Where("fund_name = ?", *r.FundName)
        }
        if r.FundScale != nil {
            // 默认等于查询
            db = db.Where("fund_scale = ?", *r.FundScale)
        }
        if r.InvestIndustry != nil {
            // 默认等于查询
            db = db.Where("invest_industry = ?", *r.InvestIndustry)
        }
        if r.RoundPrefer != nil {
            // 默认等于查询
            db = db.Where("round_prefer = ?", *r.RoundPrefer)
        }
        if r.SingleInvestAmount != nil {
            // 默认等于查询
            db = db.Where("single_invest_amount = ?", *r.SingleInvestAmount)
        }
        if r.InvestMethod != nil {
            // 默认等于查询
            db = db.Where("invest_method = ?", *r.InvestMethod)
        }
        if r.InvestableBalance != nil {
            // 默认等于查询
            db = db.Where("investable_balance = ?", *r.InvestableBalance)
        }
        if r.ReturnRequirements != nil {
            // 默认等于查询
            db = db.Where("return_requirements = ?", *r.ReturnRequirements)
        }
        if r.CompletionStatus != nil {
            // 默认等于查询
            db = db.Where("completion_status = ?", *r.CompletionStatus)
        }
        if r.FundType != nil {
            // 默认等于查询
            db = db.Where("fund_type = ?", *r.FundType)
        }
        if r.EstablishTime != nil {
            // 默认等于查询
            db = db.Where("establish_time = ?", *r.EstablishTime)
        }
        if r.DataSource != nil {
            // 默认等于查询
            db = db.Where("data_source = ?", *r.DataSource)
        }
        if r.IsBlacklist != nil {
            // 默认等于查询
            db = db.Where("is_blacklist = ?", *r.IsBlacklist)
        }
        if r.InvestmentField != nil {
            // 默认等于查询
            db = db.Where("investment_field = ?", *r.InvestmentField)
        }
        if r.AvailableBalanceFund != nil {
            // 默认等于查询
            db = db.Where("available_balance_fund = ?", *r.AvailableBalanceFund)
        }
        if r.Tag != nil {
            // 默认等于查询
            db = db.Where("tag = ?", *r.Tag)
        }
        if r.UpdatedBy != nil {
            // 默认等于查询
            db = db.Where("updated_by = ?", *r.UpdatedBy)
        }
		return db
	}
}

// InvFundCreateRequest 创建res_invest_fund请求参数
type InvFundCreateRequest struct {
	models.Validator
	InvestId int64 `form:"investId"` // 投资机构-数据主键id
	FundName string `form:"fundName" validate:"required" message:"基金名称不能为空"` // 基金名称
	FundScale string `form:"fundScale" validate:"required" message:"基金规模不能为空"` // 基金规模
	InvestIndustry string `form:"investIndustry" validate:"required" message:"投资行业不能为空"` // 投资行业
	RoundPrefer string `form:"roundPrefer" validate:"required" message:"投资阶段偏好不能为空"` // 投资阶段偏好
	SingleInvestAmount string `form:"singleInvestAmount" validate:"required" message:"单笔投资金额不能为空"` // 单笔投资金额
	InvestMethod string `form:"investMethod" validate:"required" message:"投资方式不能为空"` // 投资方式
	InvestableBalance float64 `form:"investableBalance"` // 可投余额
	ReturnRequirements string `form:"returnRequirements" validate:"required" message:"返投要求不能为空"` // 返投要求
	CompletionStatus string `form:"completionStatus" validate:"required" message:"返投完成情况不能为空"` // 返投完成情况
	FundType string `form:"fundType" validate:"required" message:"基金类型不能为空"` // 基金类型
	EstablishTime *time.Time `form:"establishTime" validate:"required" message:"成立时间不能为空"` // 成立时间
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	InvestmentField string `form:"investmentField" validate:"required" message:"投资领域不能为空"` // 投资领域
	AvailableBalanceFund string `form:"availableBalanceFund" validate:"required" message:"基金可投余额不能为空"` // 基金可投余额
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *InvFundCreateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// InvFundUpdateRequest 更新res_invest_fund请求参数
type InvFundUpdateRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
	InvestId int64 `form:"investId"` // 投资机构-数据主键id
	FundName string `form:"fundName" validate:"required" message:"基金名称不能为空"` // 基金名称
	FundScale string `form:"fundScale" validate:"required" message:"基金规模不能为空"` // 基金规模
	InvestIndustry string `form:"investIndustry" validate:"required" message:"投资行业不能为空"` // 投资行业
	RoundPrefer string `form:"roundPrefer" validate:"required" message:"投资阶段偏好不能为空"` // 投资阶段偏好
	SingleInvestAmount string `form:"singleInvestAmount" validate:"required" message:"单笔投资金额不能为空"` // 单笔投资金额
	InvestMethod string `form:"investMethod" validate:"required" message:"投资方式不能为空"` // 投资方式
	InvestableBalance float64 `form:"investableBalance"` // 可投余额
	ReturnRequirements string `form:"returnRequirements" validate:"required" message:"返投要求不能为空"` // 返投要求
	CompletionStatus string `form:"completionStatus" validate:"required" message:"返投完成情况不能为空"` // 返投完成情况
	FundType string `form:"fundType" validate:"required" message:"基金类型不能为空"` // 基金类型
	EstablishTime *time.Time `form:"establishTime" validate:"required" message:"成立时间不能为空"` // 成立时间
	DataSource string `form:"dataSource" validate:"required" message:"数据来源不能为空"` // 数据来源
	IsBlacklist int `form:"isBlacklist"` // 是否负面清单: 0否 1是
	InvestmentField string `form:"investmentField" validate:"required" message:"投资领域不能为空"` // 投资领域
	AvailableBalanceFund string `form:"availableBalanceFund" validate:"required" message:"基金可投余额不能为空"` // 基金可投余额
	Tag string `form:"tag" validate:"required" message:"行业数据标签不能为空"` // 行业数据标签
	UpdatedBy int64 `form:"updatedBy"` // 修改人
}

// Validate 验证请求参数
func (r *InvFundUpdateRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// InvFundDeleteRequest 删除res_invest_fund请求参数
type InvFundDeleteRequest struct {
	models.Validator
	Id int64 `form:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *InvFundDeleteRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}

// InvFundGetByIDRequest 根据ID获取res_invest_fund请求参数
type InvFundGetByIDRequest struct {
	models.Validator
	Id int64 `uri:"id" validate:"required" message:"主键不能为空"` // 主键
}

// Validate 验证请求参数
func (r *InvFundGetByIDRequest) Validate(c *gin.Context) error {
	return r.Validator.Check(c, r)
}