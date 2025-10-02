# prompt-share-backend

## 1. 生成 Swagger docs

在项目根目录运行（需先安装 swag）：

```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g cmd/main.go
```

会生成 docs 文件夹，确保 api/router.go 已引入 _ "prompt-share-backend/docs"。

访问文档：

http://localhost:8080/swagger/index.html

## 2. 运行项目

保存所有文件到对应路径（确保 module 与导入路径一致——这里使用 prompt-share-backend）。

执行 go mod tidy（会生成 go.sum）。

执行 go run ./cmd 或 go run cmd/main.go。
若要 Swagger：先 go install github.com/swaggo/swag/cmd/swag@latest，再在项目根运行 swag init -g cmd/main.go，然后重新启动服务。

### 3. Format

```bash
# 安装 goimports（推荐）
go install golang.org/x/tools/cmd/goimports@latest

# 或安装 gofumpt（更严格的格式化工具）
go install mvdan.cc/gofumpt@latest

# 安装 golangci-lint（综合检查工具）
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

```bash
# 格式化单个文件
goimports -w service/prompt_service.go

# 格式化整个项目
goimports -w .

# 或使用 gofumpt
gofumpt -w .
```