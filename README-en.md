# gin-blog

![Travis (.org)](https://img.shields.io/travis/linehk/gin-blog)
[![codecov](https://codecov.io/gh/linehk/gin-blog/branch/master/graph/badge.svg)](https://codecov.io/gh/linehk/gin-blog)
[![Go Report Card](https://goreportcard.com/badge/github.com/linehk/gin-blog)](https://goreportcard.com/report/github.com/linehk/gin-blog)

English | [简体中文](./README.md "简体中文")

gin-blog is a simple blog RESTful API example that uses MySQL as the database.

## Installation

```bash
git clone https://github.com/linehk/gin-blog.git
```

Then, build it:

```bash
go build -o gin-blog
```

And, run it:

```bash
./gin-blog
```

## Usages

* Article
  * get article: GET `http://localhost:8888/api/v1/articles/1`
  * get all article: GET `http://localhost:8888/api/v1/articles`
  * new article: POST `http://localhost:8888/api/v1/articles`, form allow to use `form-data` or `x-www-form-urlencoded` format
  * edit article: PUT `http://localhost:8888/api/v1/articles/1`
  * delete article: DELETE `http://localhost:8888/api/v1/articles/1`

* Tag
  * get tag: GET `http://localhost:8888/api/v1/tags/1`
  * get all tag: GET `http://localhost:8888/api/v1/tags`
  * new tag: POST `http://localhost:8888/api/v1/tags`
  * edit tag: PUT `http://localhost:8888/api/v1/tags/1`
  * delete tag: DELETE `http://localhost:8888/api/v1/tags/1`

## Contributing

If you feel that there is something to improve this project, please feel free to launch Pull Request.

For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT License](./LICENSE "MIT License")
