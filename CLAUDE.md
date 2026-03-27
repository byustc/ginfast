# CLAUDE.md

本文档为 Claude Code (claude.ai/code) 提供在此代码仓库中工作的指导。

## 项目概述

GinFast 是一个基于 Gin 框架的多租户 RBAC（基于角色的访问控制）后端脚手架。提供 JWT 认证、Casbin 权限管理、代码生成、插件系统和定时任务等功能。

## 构建与测试命令

```bash
# 安装依赖
go mod tidy

# 构建
go build -o gin-fast .

# 运行
go run main.go

# 生成 Swagger 文档 (Windows)
scripts\swagger.bat

# 生成 Swagger 文档 (Linux/Mac)
./scripts/swagger.sh

# 运行所有测试
go test ./...

# 运行指定包测试
go test ./app/utils/cachehelper/...

# 运行测试并显示覆盖率
go test -cover ./...

# 运行测试并输出详细日志
go test -v ./app/utils/casbinhelper/...
```

## 架构

### 分层结构
- **Controllers** (`app/controllers/`): HTTP 请求处理器，继承 `controllers.Common`
- **Services** (`app/service/`): 业务逻辑层
- **Models** (`app/models/`): 数据模型，使用 GORM 标签，继承 `models.BaseModel`
- **Routes** (`app/routes/routes.go`): Gin 路由注册
- **Middleware** (`app/middleware/`): JWT 认证、Casbin、CORS、验证码、操作日志
- **Bootstrap** (`bootstrap/init.go`): 应用初始化（数据库、Casbin、缓存、调度器）

### 关键全局变量 (`app/global/app/variable.go`)
- `app.DB()` - 获取数据库连接（mysql/sqlserver/postgresql）
- `app.ZapLog` - Zap 日志器
- `app.CasbinV2` - Casbin 权限执行器
- `app.Cache` - Redis 或内存缓存
- `app.TokenService` - JWT 令牌管理
- `app.JobScheduler` - 基于 Cron 的任务调度器

### 多租户数据隔离
- 所有需要租户隔离的模型必须包含 `TenantID uint` 字段
- GORM 钩子在创建/更新时自动注入 TenantID
- `tenanthelper.TenantScope(c)` 获取当前租户的数据范围
- JWT 中间件提取并注入租户信息

### Casbin 权限模型
- 策略格式: `sub, obj, act, dom` (主体、对象、操作、域/租户)
- 角色定义: `g = _, _, _` (支持带域的角色继承)
- 自动从数据库同步策略（默认 120 秒，可配置）

## 配置

主配置文件: `config/config.yml` (使用 Viper)
- 数据库连接: `gormv2.mysql`, `gormv2.sqlserver`, `gormv2.postgresql`
- JWT 配置: `token.jwttokensignkey`, `token.jwttokenexpire`
- Casbin 模型: `casbin.modelconfig`
- 调度器: `scheduler.log.dir`, `scheduler.job_results_buffer_size`

### 单元测试模式

测试文件放在同包目录下，命名 `*_test.go`。使用 `testify/assert` 进行断言：

```go
func TestMemoryHelper_SetAndGet(t *testing.T) {
    cache := NewMemoryHelper()
    defer cache.Close()
    ctx := context.Background()

    // 表驱动测试
    tests := []struct {
        name       string
        key        string
        value      string
        expiration time.Duration
        wantErr    bool
    }{
        {name: "正常设置和获取", key: "test_key", value: "test_value", expiration: 5 * time.Second, wantErr: false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := cache.Set(ctx, tt.key, tt.value, tt.expiration)
            assert.NoError(t, err)

            value, err := cache.Get(ctx, tt.key)
            assert.NoError(t, err)
            assert.Equal(t, tt.value, value)
        })
    }
}

// Benchmark 性能测试
func BenchmarkMemoryHelper_Set(b *testing.B) {
    cache := NewMemoryHelper()
    defer cache.Close()
    ctx := context.Background()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        cache.Set(ctx, "key", "value", time.Minute)
    }
}
```

## 技能文档

框架专用技能位于 `ginfast-skill/` 目录:

| 优先级 | 技能 | 用途 |
|--------|------|------|
| 1 | `core/model-skill.md` | 模型定义、GORM 标签、CRUD |
| 2 | `core/service-skill.md` | 服务层模式 |
| 3 | `core/controller-skill.md` | 控制器结构、Swagger 注解 |
| 4 | `core/routes-skill.md` | 路由注册 |
| 5 | `features/middleware-skill.md` | 中间件开发 |
| 6 | `features/validate-skill.md` | 参数验证 |
| 7 | `features/response-skill.md` | API 响应格式 |
| 8 | `features/casbin-skill.md` | RBAC 权限配置 |
| 9 | `features/tenant-skill.md` | 多租户隔离 |
| 10 | `features/scheduler-skill.md` | Cron 任务开发 |

## 开发模式

### 创建新功能
1. 在 `app/models/` 定义模型（继承 `BaseModel`，添加 `TenantID`）
2. 在 `app/models/*param.go` 创建参数验证器
3. 在 `app/service/` 实现服务层
4. 在 `app/controllers/` 创建控制器
5. 在 `app/routes/routes.go` 注册路由
6. 配置 Casbin 策略以控制 API 权限

### 模型定义
```go
type SysUser struct {
    BaseModel
    TenantID    uint   `gorm:"column:tenant_id;default:0"`
    Username    string `gorm:"size:50"`
    Password    string `gorm:"size:255"`
    // ...
}
```

### 控制器结构
```go
type UserController struct {
    controllers.Common  // 继承以使用 Success/Fail 响应
}
// @Summary 用户列表
// @Tags 用户管理
// @Router /users/list [get]
func (uc *UserController) List(c *gin.Context) { ... }
```

## 数据库

通过 GORM 驱动支持 MySQL、PostgreSQL、SQL Server。默认: MySQL。
- 字符集: utf8mb4
- 命名约定: snake_case
- 连接池在 `config/config.yml` 中配置

## Swagger API 文档

调试模式下访问 `http://localhost:8080/swagger/index.html`
