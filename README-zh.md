# gin-blog

[![codecov](https://codecov.io/gh/linehk/gin-blog/branch/master/graph/badge.svg)](https://codecov.io/gh/linehk/gin-blog)
[![Go Report Card](https://goreportcard.com/badge/github.com/linehk/gin-blog)](https://goreportcard.com/report/github.com/linehk/gin-blog)

[English](./README-en.md "English") | 简体中文

gin-blog 是一个简单的博客 RESTful API 示例，使用 MySQL 作为数据库。

## 安装

```bash
git clone https://github.com/linehk/gin-blog.git
```

然后进行编译：

```bash
go build -o gin-blog
```

再运行：

```bash
./gin-blog
```

## 使用

* 文章
  * 获取指定文章：GET `http://localhost:8888/api/v1/articles/1`
  * 获取全部文章：GET `http://localhost:8888/api/v1/articles`
  * 新增文章：POST `http://localhost:8888/api/v1/articles`，表单可用 `form-data` 或 `x-www-form-urlencoded` 形式
  * 修改文章：PUT `http://localhost:8888/api/v1/articles/1`
  * 删除文章：DELETE `http://localhost:8888/api/v1/articles/1`
* 标签
  * 获取指定标签：GET `http://localhost:8888/api/v1/tags/1`
  * 获取全部标签：GET `http://localhost:8888/api/v1/tags`
  * 新增标签：POST `http://localhost:8888/api/v1/tags`
  * 修改标签：PUT `http://localhost:8888/api/v1/tags/1`
  * 删除标签：DELETE `http://localhost:8888/api/v1/tags/1`

## 参与贡献

如果您觉得这个项目有什么需要改进的地方，欢迎发起 Pull Request。

如果有重大变化，请先打开一个 Issue，讨论您想要改变的内容。

## 开源许可证

[MIT License](./LICENSE "MIT License")
