# 项目结构

```
blog-backend/
├── cmd/
│ └── server/
│ └── main.go # 程序入口，初始化配置、数据库、路由和启动 HTTP 服务
│
├── internal/ # 核心业务代码（项目私有，外部无法 import）
│ ├── config/
│ │ ├── config.go # 配置加载逻辑（yaml + .env）
│ │ └── config.yaml # 应用和数据库配置
│ │
│ ├── database/
│ │ ├── database.go # 初始化数据库连接，Ping 测试
│ │ └── postgres.go # PostgreSQL 连接封装
│ │
│ ├── router/
│ │ └── router.go # Gin 路由注册，绑定中间件和 handler
│ │
│ ├── api/ # API 路由分组（可支持版本管理）
│ │ └── v1/
│ │  ├── post.go # post 模块路由注册
│ │  └── auth.go # auth 模块路由注册
│ │
│ ├── handler/ # HTTP 层，处理请求参数、调用 service 并返回 JSON
│ │ ├── post_handler.go
│ │ └── auth_handler.go
│ │
│ ├── service/ # 业务逻辑层，封装业务规则，不直接操作数据库
│ │ ├── post_service.go
│ │ └── auth_service.go
│ │
│ ├── repository/ # 数据访问层，封装所有数据库操作
│ │ ├── post_repo.go
│ │ └── user_repo.go
│ │
│ ├── model/ # 数据模型（与数据库表结构对应）
│ │ ├── post.go
│ │ └── user.go
│ │
│ ├── middleware/ # Gin 中间件，如鉴权、日志、跨域
│ │ └── auth.go
│ │
│ └── pkg/ # 可复用工具模块
│  ├── response/ # HTTP 响应封装
│  │ └── response.go
│  └── jwt/ # JWT 工具封装
│     └── jwt.go
│
├── migrations/ # 数据库迁移文件
│ └── 001_init.up.sql # 初始建表 SQL
│
├── .env # 敏感配置，如数据库账号密码（不提交 git）
├── go.mod
├── go.sum
└── README.md
```

## 结构说明
