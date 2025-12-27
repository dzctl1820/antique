# Antique Project

这是一个包含前后端的全栈项目，旨在构建一个关于古董/文物的展示与交流平台。项目集成了 3D 文物展示、AI 智能助手对话以及社区交流等功能。

## 📂 目录结构

- `an-backend/`: 后端服务项目 (基于 Go 语言)
- `antique/`: 前端应用项目 (基于 Vue 3)

---

## 🖥️ 后端 (an-backend)

后端服务提供 RESTful API，处理用户认证、数据持久化以及与 AI 模型的交互。

### 技术栈
- **语言**: Go 1.25.3
- **Web 框架**: [Gin](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io/)
- **数据库**: MySQL, Redis
- **认证**: JWT (JSON Web Token)
- **AI 服务**: 阿里云 DashScope (通义千问)

### 🚀 快速开始

#### 1. 环境准备
确保本地已安装以下服务：
- Go (1.25+)
- MySQL (默认端口 3306)
- Redis (默认端口 6379)

#### 2. 配置说明
**环境变量**:
在 `an-backend` 目录下创建 `.env` 文件，并配置以下必要的环境变量：

```env
# JWT 签名密钥（必须设置，建议长度 > 32 字符）
JWT_SECRET=your_secure_random_string

# 阿里云 DashScope API Key (用于 AI 对话功能)
DASHSCOPE_API_KEY=your_dashscope_api_key
```

**其他配置 (代码修改)**:
目前部分配置项在代码中硬编码，请根据您的本地环境修改相应文件：
- **MySQL 连接**: 修改 `an-backend/database/db.go` 中的 DSN 字符串。
- **Redis 连接**: 修改 `an-backend/database/redis.go` 中的地址配置。
- **邮件服务 (SMTP)**: 修改 `an-backend/internal/handler/login.go` 中的发件人邮箱和授权码。

#### 3. 运行服务
```bash
cd an-backend

# 安装依赖
go mod tidy

# 启动服务
go run cmd/main.go
```
服务默认运行在 `:8086` 端口。

### 📖 API 文档
详细的接口定义、参数说明及示例请参考 [an-backend/API文档.md](an-backend/API文档.md)。

---

## 🎨 前端 (antique)

前端应用构建了一个交互式的 Web 界面，利用 Three.js 实现 3D 视觉效果。

### 技术栈
- **框架**: Vue 3
- **构建工具**: Vue CLI / Vite
- **HTTP 客户端**: Axios
- **3D 引擎**: Three.js
- **动画库**: GSAP

### 🚀 快速开始

#### 1. 安装依赖
```bash
cd antique
yarn install
```

#### 2. 开发模式运行
```bash
yarn serve
```
启动后，访问终端输出的本地地址（通常为 `http://localhost:8080`）。

#### 3. 构建生产版本
```bash
yarn build
```

#### 4. 代码检查
```bash
yarn lint
```

---

## ✨ 功能特性

- **🔐 用户系统**: 支持邮箱注册、登录，集成验证码发送与验证功能。
- **🤖 AI 智能助手**: 集成大语言模型，支持流式对话与上下文记忆（基于 Session）。
- **🏺 3D 文物展示**: 使用 Three.js 渲染高精度文物模型，支持交互式查看。
- **📝 内容管理**: 文章发布与浏览，支持富文本内容。
- **💬 社区互动**: 提供用户交流互动的社区板块。
