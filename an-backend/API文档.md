# API 接口文档

## 基础信息
- **服务地址**: `http://localhost:8086`
- **框架**: Gin (Go)
- **数据格式**: JSON

## 认证说明
- 需要认证的接口需在请求头中携带 `Authorization` 字段
- 格式: `Authorization: <token>`
- Token 通过登录接口获取，有效期 24 小时

---

## 1. 用户认证模块

### 1.1 邮箱登录
- **接口**: `POST /api/login`
- **认证**: 否

**请求参数**:
```json
{
  "email": "user@example.com",
  "password": "123456"
}
```

**成功响应** (200):
```json
{
  "token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9..."
}
```

**失败响应** (400):
```json
{
  "error": "password is incorrect"
}
```

---

### 1.2 发送验证码
- **接口**: `POST /api/send/code`
- **认证**: 否

**请求参数**:
```json
{
  "email": "user@example.com"
}
```

**成功响应** (200):
```json
{
  "message": "发送成功"
}
```

**失败响应** (400):
```json
{
  "error": "错误信息"
}
```

---

### 1.3 邮箱注册
- **接口**: `POST /api/register`
- **认证**: 否

**请求参数**:
```json
{
  "email": "user@example.com",
  "password": "123456",
  "code": "123456"
}
```

**成功响应** (200):
```json
{
  "message": "注册成功"
}
```

**失败响应** (400):
```json
{
  "error": "验证码错误"
}
```

---

## 2. 文章模块

### 2.1 添加文章
- **接口**: `POST /api/article/add`
- **认证**: 是 ✓

**请求参数**:
```json
{
  "title": "文章标题",
  "cover": "封面图片URL",
  "author": "作者名称",
  "content": "文章内容",
  "summary": "文章摘要",
  "category": "分类",
  "tags": "标签1,标签2"
}
```

**成功响应** (200):
```json
{
  "message": "文章创建成功"
}
```

**失败响应** (400/500):
```json
{
  "error": "错误信息"
}
```

---

### 2.2 获取单篇文章
- **接口**: `GET /api/article/one/:id`
- **认证**: 是 ✓

**路径参数**:
- `id`: 文章ID

**成功响应** (200):
```json
{
  "id": 1,
  "id_user": 1,
  "title": "文章标题",
  "cover": "封面URL",
  "author": "作者",
  "content": "文章内容",
  "summary": "摘要",
  "category": "分类",
  "tags": "标签",
  "likes": 0,
  "favorites": 0,
  "views": 0,
  "status": "published",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

**失败响应** (404):
```json
{
  "error": "错误信息"
}
```

---

### 2.3 获取所有文章
- **接口**: `GET /api/article/all`
- **认证**: 是 ✓

**成功响应** (200):
```json
[
  {
    "id": 1,
    "id_user": 1,
    "title": "文章标题",
    "cover": "封面URL",
    "author": "作者",
    "content": "文章内容",
    "summary": "摘要",
    "category": "分类",
    "tags": "标签",
    "likes": 0,
    "favorites": 0,
    "views": 0,
    "status": "published",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

---

## 3. AI 聊天模块

### 3.1 AI 对话
- **接口**: `POST /api/ai/chat`
- **认证**: 是 ✓

**请求参数**:
```json
{
  "user": "用户消息内容",
  "system": "系统提示词（可选）",
  "session_id": 1,
  "assistant": [
    {
      "UserContent": "历史用户消息",
      "AssistantContent": "历史AI回复"
    }
  ]
}
```

**成功响应** (200):
```json
{
  "content": "AI回复内容"
}
```

---

### 3.2 创建新会话
- **接口**: `POST /api/ai/new/session`
- **认证**: 是 ✓

**请求参数**:
```json
{
  "title": "会话标题（可选，默认：新的对话）"
}
```

**成功响应** (200):
```json
{
  "session_id": 1,
  "content": "创建成功"
}
```

---

### 3.3 获取会话列表
- **接口**: `GET /api/ai/sessions`
- **认证**: 是 ✓

**成功响应** (200):
```json
[
  {
    "session_id": 1,
    "user_id": 1,
    "title": "会话标题",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z",
    "expires_at": null
  }
]
```

---

### 3.4 删除会话
- **接口**: `DELETE /api/ai/session`
- **认证**: 是 ✓

**请求参数**:
```json
{
  "session_id": 1
}
```

**成功响应** (200):
```json
{
  "content": "删除成功"
}
```

---

### 3.5 更新会话
- **接口**: `POST /api/ai/session`
- **认证**: 是 ✓

**请求参数**:
```json
{
  "session_id": 1,
  "title": "新的会话标题"
}
```

**成功响应** (200):
```json
{
  "content": "更新成功"
}
```

---

### 3.6 获取会话对话记录
- **接口**: `GET /api/ai/session/conversation`
- **认证**: 是 ✓

**请求参数**:
```json
{
  "session_id": 1
}
```

**成功响应** (200):
```json
[
  {
    "id": 1,
    "session_id": 1,
    "role": "user",
    "content": "用户消息",
    "model": "qwen-plus",
    "created_at": "2024-01-01T00:00:00Z"
  },
  {
    "id": 2,
    "session_id": 1,
    "role": "assistant",
    "content": "AI回复",
    "model": "qwen-plus",
    "created_at": "2024-01-01T00:00:00Z"
  }
]
```

---

## 错误码说明

| 状态码 | 说明 |
|--------|------|
| 200 | 请求成功 |
| 400 | 请求参数错误 |
| 401 | 未授权（Token无效或过期） |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |

---

## 注意事项

1. 所有请求和响应的 Content-Type 均为 `application/json`
2. 需要认证的接口必须在请求头中携带有效的 Token
3. Token 在登录成功后获取，有效期为 24 小时
4. 验证码有效期由 Redis 配置决定
5. AI 对话使用的默认模型为 `qwen-plus`

