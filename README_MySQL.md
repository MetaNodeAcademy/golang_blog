# MySQL数据库配置说明

## 数据库切换完成

项目已成功从SQLite切换到MySQL数据库。

## 配置文件

### 1. 环境变量配置
- `.env` - 实际环境变量文件
- `.env.example` - 环境变量模板文件

### 2. 环境变量说明
```
DB_HOST=localhost          # MySQL服务器地址
DB_PORT=3306              # MySQL端口
DB_USER=root              # MySQL用户名
DB_PASSWORD=              # MySQL密码（请根据实际情况设置）
DB_NAME=golang_blog       # 数据库名称
JWT_SECRET=your_jwt_secret_key_here  # JWT密钥
PORT=8080                 # 服务器端口
GIN_MODE=debug           # Gin运行模式
```

## MySQL设置步骤

### 1. 安装MySQL
```bash
# 使用Homebrew安装
brew install mysql

# 启动MySQL服务
brew services start mysql
```

### 2. 创建数据库
```bash
# 连接MySQL（根据你的密码设置调整）
mysql -u root -p

# 创建数据库
CREATE DATABASE golang_blog CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

# 退出MySQL
exit;
```

### 3. 配置密码
如果MySQL root用户没有密码，请保持`.env`文件中的`DB_PASSWORD`为空。
如果有密码，请在`.env`文件中设置正确的密码。

### 4. 测试连接
```bash
# 启动应用程序
go run cmd/main.go
```

## 注意事项

1. **密码配置**：请根据你的MySQL实际配置设置密码
2. **数据库创建**：应用启动时会自动创建表结构
3. **字符集**：使用utf8mb4字符集支持完整的Unicode字符
4. **安全性**：生产环境请设置强密码并配置适当的用户权限

## 故障排除

### 连接被拒绝错误
如果遇到"Access denied"错误：
1. 检查MySQL服务是否运行：`brew services list | grep mysql`
2. 确认用户名和密码是否正确
3. 尝试直接连接MySQL：`mysql -u root -p`

### 数据库不存在错误
如果遇到数据库不存在错误：
1. 手动创建数据库：`CREATE DATABASE golang_blog;`
2. 确认数据库名称在`.env`文件中配置正确