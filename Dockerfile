FROM golang:1.21 AS build

WORKDIR /workdir

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /app

FROM build AS test

RUN go test -v ./...

# 最新版本：https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/base-debian12 AS release

# 这里需要重新设置工作目录
WORKDIR /

COPY --from=build /app /app
COPY --from=build /workdir/config.toml /config.toml

EXPOSE 8080

# 权限
USER nonroot:nonroot

CMD ["/app"]