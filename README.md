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
