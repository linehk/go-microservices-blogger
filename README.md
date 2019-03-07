# 项目描述

gin-blog 是一个简单的博客 RESTful API 示例，使用 MySQL 作为数据库。

技术栈：

1. Go
2. Gin
3. GORM

第三方包：

1. go-ini 读取 `.ini` 格式的配置文件
2. validation 验证表单

测试接口：

1. 文章
    1. 获取指定文章
    GET `http://127.0.0.1:8888/api/v1/articles/1`
    2. 获取全部文章
    GET `http://127.0.0.1:8888/api/v1/articles`
    3. 新增文章
    POST `http://127.0.0.1:8888/api/v1/articles`
    表单可用 `form-data` 或 `x-www-form-urlencoded` 形式。
    4. 修改文章
    PUT `http://127.0.0.1:8888/api/v1/articles/1`
    5. 删除文章
    DELETE `http://127.0.0.1:8888/api/v1/articles/1`

2. 标签
    1. 获取指定标签
    GET `http://127.0.0.1:8888/api/v1/tags/1`
    2. 获取全部标签
    GET `http://127.0.0.1:8888/api/v1/tags`
    3. 新增标签
    POST `http://127.0.0.1:8888/api/v1/tags`
    4. 修改标签
    PUT `http://127.0.0.1:8888/api/v1/tags/1`
    5. 删除标签
    DELETE `http://127.0.0.1:8888/api/v1/tags/1`