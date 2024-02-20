FROM golang:1.22.0 AS build

WORKDIR /workdir

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -o /app ./service/user/rpc/user.go

FROM gcr.io/distroless/base-debian12 AS release

WORKDIR /

COPY --from=build /app /app
COPY --from=build /workdir/service/user/rpc/etc/user.yaml /config.yaml

EXPOSE 10000

USER nonroot:nonroot

CMD ["/app", "-f", "/config.yaml"]