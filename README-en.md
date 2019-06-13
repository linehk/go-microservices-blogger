# gin-blog

[![Build Status](https://travis-ci.org/linehk/gin-blog.svg?branch=master)](https://travis-ci.org/linehk/gin-blog)
[![codecov](https://codecov.io/gh/linehk/gin-blog/branch/master/graph/badge.svg)](https://codecov.io/gh/linehk/gin-blog)
[![Go Report Card](https://goreportcard.com/badge/github.com/linehk/gin-blog)](https://goreportcard.com/report/github.com/linehk/gin-blog)

English | [简体中文](./README.md "简体中文")

gin-blog is a simple blog RESTful API example that uses MySQL as the database.

## Installation

```bash
git clone https://github.com/linehk/gin-blog.git
```

Or:

```bash
go get -u github.com/linehk/gin-blog
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
  1. get article: GET `http://localhost:8888/api/v1/articles/1`
  2. get all article: GET `http://localhost:8888/api/v1/articles`
  3. new article: POST `http://localhost:8888/api/v1/articles`, form allow to use `form-data` or `x-www-form-urlencoded` format
  4. edit article: PUT `http://localhost:8888/api/v1/articles/1`
  5. delete article: DELETE `http://localhost:8888/api/v1/articles/1`

* Tag
  1. get tag: GET `http://localhost:8888/api/v1/tags/1`
  2. get all tag: GET `http://localhost:8888/api/v1/tags`
  3. new tag: POST `http://localhost:8888/api/v1/tags`
  4. edit tag: PUT `http://localhost:8888/api/v1/tags/1`
  5. delete tag: DELETE `http://localhost:8888/api/v1/tags/1`

## Contributing

If you feel that there is something to improve this project, please feel free to launch Pull Request.

For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT License](./LICENSE "MIT License")