# Golang Blog API

一个使用Go语言、Gin框架和GORM库开发的个人博客系统后端API。

## 功能特性

- 用户注册和登录
- JWT认证和授权
- 文章的CRUD操作
- 评论功能
- 统一的错误处理
- 日志记录
- 数据库自动迁移

## 技术栈

- **Go 1.24.2**
- **Gin** - Web框架
- **GORM** - ORM库
- **SQLite** - 数据库
- **JWT** - 身份认证
- **Logrus** - 日志库
- **bcrypt** - 密码加密

## 项目结构

```
golang_blog/
├── cmd/
│   └── main.go          # 程序入口
├── config/
│   └── database.go      # 数据库配置
├── controllers/
│   ├── auth.go          # 认证控制器
│   ├── post.go          # 文章控制器
│   └── comment.go       # 评论控制器
├── middleware/
│   ├── auth.go          # JWT认证中间件
│   └── logger.go        # 日志中间件
├── models/
│   ├── user.go          # 用户模型
│   ├── post.go          # 文章模型
│   └── comment.go       # 评论模型
├── routes/
│   └── routes.go        # 路由配置
├── utils/
│   ├── jwt.go           # JWT工具
│   └── response.go      # 响应工具
├── go.mod
├── go.sum
└── README.md
```

## 数据库设计

### Users表
- id (主键)
- username (用户名，唯一)
- email (邮箱，唯一)
- password (加密密码)
- created_at, updated_at, deleted_at

### Posts表
- id (主键)
- title (标题)
- content (内容)
- user_id (外键，关联users表)
- created_at, updated_at, deleted_at

### Comments表
- id (主键)
- content (内容)
- user_id (外键，关联users表)
- post_id (外键，关联posts表)
- created_at, updated_at, deleted_at

## API接口

### 认证接口
- `POST /api/v1/auth/register` - 用户注册
- `POST /api/v1/auth/login` - 用户登录
- `GET /api/v1/profile` - 获取用户信息 (需要认证)

### 文章接口
- `GET /api/v1/posts` - 获取文章列表 (公开)
- `GET /api/v1/posts/:id` - 获取文章详情 (公开)
- `POST /api/v1/posts` - 创建文章 (需要认证)
- `PUT /api/v1/posts/:id` - 更新文章 (需要认证，仅作者)
- `DELETE /api/v1/posts/:id` - 删除文章 (需要认证，仅作者)

### 评论接口
- `GET /api/v1/comments/post/:post_id` - 获取文章评论 (公开)
- `POST /api/v1/posts/:post_id/comments` - 创建评论 (需要认证)

### 其他接口
- `GET /health` - 健康检查

## 运行项目

1. 确保已安装Go 1.21+
2. 克隆项目到本地
3. 安装依赖：
   ```bash
   go mod tidy
   ```
4. 运行项目：
   ```bash
   go run cmd/main.go
   ```
5. 服务器将在 `http://localhost:8080` 启动

## 使用示例

### 用户注册
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }'
```

### 用户登录
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }'
```

### 创建文章
```bash
curl -X POST http://localhost:8080/api/v1/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "我的第一篇文章",
    "content": "这是文章内容..."
  }'
```

## 注意事项

- JWT密钥在生产环境中应该从环境变量读取
- 所有密码都会使用bcrypt进行加密存储
- API返回统一的JSON格式响应