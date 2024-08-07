FROM golang:1.20 as builder
WORKDIR app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on GOPROXY=https://goproxy.cn,direct \
    && go build -o laya-template cmd/main.go

FROM debian:bookworm-slim
COPY --from=builder /app/laya-template /var/www/code/app/layatp
COPY --from=builder /app/config /var/www/code/app/config

WORKDIR /var/www/code/app
CMD ["./laya-template"]
