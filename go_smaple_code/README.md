# Go Sample Code

Go 语言示例代码集合，提供实用的代码示例和可复用的工具包。

## 📦 项目结构

```
go_sample_code/
├── cmd/
│   └── server/         # Fiber Web 服务
├── internal/
│   ├── handler/        # HTTP 处理器
│   └── middleware/     # 中间件
├── pkg/
│   ├── filetree/       # 文件树 N 叉树实现
│   └── logger/         # 通用日志工具包
└── README.md
```

## 📚 模块说明

### filetree - 文件树数据结构

基于 N 叉树的文件/目录结构实现，支持序列化存储。

[查看详细文档](./pkg/filetree/README.md)

### logger - 日志工具包

基于 zap 的高性能日志库，支持 OpenTelemetry 集成。

[查看详细文档](./pkg/logger/README.md)

## 🚀 快速开始

```bash
# 运行服务
go run cmd/server/main.go

# 运行测试
go test ./...
```

## 🌐 API 端点

服务启动后访问 `http://localhost:8080`

### 健康检查

```bash
curl http://localhost:8080/api/health
```

### 文件树操作

```bash
# 添加目录
curl -X POST http://localhost:8080/api/filetree/dir \
  -H "Content-Type: application/json" \
  -d '{"parent_path": "/", "dir_name": "docs"}'

# 添加文件
curl -X POST http://localhost:8080/api/filetree/file \
  -H "Content-Type: application/json" \
  -d '{"dir_path": "/docs", "file_name": "readme.md", "file_id": 1001}'

# 获取完整树结构
curl http://localhost:8080/api/filetree/tree

# 获取所有文件
curl http://localhost:8080/api/filetree/files

# 重命名节点
curl -X PUT http://localhost:8080/api/filetree/rename \
  -H "Content-Type: application/json" \
  -d '{"old_path": "/docs/readme.md", "new_name": "README.md"}'

# 删除节点
curl -X DELETE http://localhost:8080/api/filetree/node \
  -H "Content-Type: application/json" \
  -d '{"path": "/docs/README.md"}'
```

## 📝 开发指南

每个模块都包含：
- 完整的实现代码
- 单元测试
- 使用文档

欢迎贡献新的示例和工具包！