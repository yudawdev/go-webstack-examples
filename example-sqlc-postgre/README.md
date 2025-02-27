# SQLC PostgreSQL 示例项目

这是一个使用 SQLC 和 PostgreSQL 的示例项目，展示了如何使用 SQLC 生成类型安全的 Go 代码来操作 PostgreSQL 数据库。

## 前置要求

- Go 1.24
- Docker
- sqlc

## 安装 SQLC

[SQLC Official Web](https://docs.sqlc.dev/)

可以通过以下方式安装 SQLC：

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

## 设置 PostgreSQL

使用 Docker 启动一个本地 PostgreSQL 实例， 参考启动参数如下：

```bash
docker run --name local-postgres \
-e POSTGRES_PASSWORD=root \
-e POSTGRES_USER=root \
-e POSTGRES_DB=dev \
-p 5432:5432 \
-v ~/docker-mount/postgres-data:/var/lib/postgresql/data \
-d postgres
```

数据库连接信息：
- 主机：localhost
- 端口：5432
- 用户名：root
- 密码：root
- 数据库名：dev

## 项目结构

```
.
├── database
│   ├── query.sql    # SQL 查询语句
│   └── authors.sql   # 数据库表结构
├── db
│   ├── db.go        # 生成的数据库接口
│   ├── models.go    # 生成的数据模型
│   └── query.sql.go # 生成的查询方法
└── sqlc.yaml        # SQLC 配置文件
```

## 生成代码

在项目根目录执行以下命令生成 Go 代码：

```bash
sqlc generate

// 如果 sqlc.yaml 不在本项目目录下，请使用命令
sqlc generate -f xxxx
```

## 数据库操作

本项目实现了以下数据库操作：

- `GetAuthor`: 通过 ID 获取作者信息
- `ListAuthors`: 获取所有作者列表
- `CreateAuthor`: 创建新作者
- `UpdateAuthor`: 更新作者信息
- `DeleteAuthor`: 删除作者
- `UpdateAuthorReturnRecord`: 更新作者并返回更新后的记录

## 依赖

项目使用以下主要依赖：

```go
require (
	github.com/jackc/pgx/v5 v5.7.2
)
```

## 配置说明

`sqlc.yaml` 配置文件说明：

```yaml
version: "2"
sql:
  - engine: "postgresql"    # 使用 PostgreSQL 引擎
    queries: "database/query.sql"    # SQL 查询文件位置
    schema: "database/authors.sql"    # 数据库架构文件位置
    gen:
      go:
        package: "db"    # 生成的 Go 包名
        out: "db"        # 输出目录
        sql_package: "pgx/v5"    # 使用的 SQL 驱动包
