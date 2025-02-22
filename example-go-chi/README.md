# Go-Chi Demo

这是一个使用 [go-chi](https://go-chi.io/#/README) 框架构建的简单 REST API 示例项目。Go-chi 是一个轻量级、灵活的路由器和中间件框架，用于构建 Go HTTP 服务。

- go-chi 官网：https://go-chi.io/#/README
- GitHub 仓库：https://github.com/go-chi/chi

## 项目启动

```bash
go run main.go
```

服务将在 `http://localhost:3000` 启动

## API 测试

你可以使用以下 curl 命令测试 API：

1. 测试首页接口
```bash
curl http://localhost:3000/
```

2. 获取指定 ID 的文章
```bash
curl http://localhost:3000/article/1
```

3. 获取文章列表
```bash
curl -X POST http://localhost:3000/articles
```
